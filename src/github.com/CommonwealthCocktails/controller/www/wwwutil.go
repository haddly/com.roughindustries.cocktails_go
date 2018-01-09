// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/www/util.go: Utility functions
package www

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/glog"
	"github.com/spf13/viper"
	"io"
	"math"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
)

//all alpha numeric ascii characters upper and lower case
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

//Generate a random sequence of length n characters from the alpha numeric
//ascii characters upper and lower case
func randSeq(n int) string {
	token, _ := GenerateRandomString(32)
	return token
}

//Helper function for producing a standard 404 page error when we through an
//panic
func Error404(w http.ResponseWriter, rec interface{}) {
	page := NewPage(nil, nil)
	page.View = "www"
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	glog.Infoln("Recovered %s:%d %s\n", file, line, f.Name())
	glog.Infoln(rec)
	page.RenderPageTemplate(w, nil, "404")
}

//Setup Page Only
func RenderSetupTemplate(w http.ResponseWriter, rec interface{}) {
	// CATCH ONLY HEADER START
	defer func() {
		// recover from panic if one occured. Set err to nil otherwise.
		if rec := recover(); rec != nil {
			Error404(w, rec)
		}
	}()
	// CATCH ONLY HEADER START
	page := NewSetupPage(nil, nil)
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	glog.Infoln("Setup Failed %s:%d %s\n", file, line, f.Name())
	glog.Infoln(rec)
	page.RenderSetupTemplate(w, nil, "setup")
}

//Validates the CSRF ID from the last page request.  True means the CSRf is good,
//and false means that the CSRF is not the one past to the previous page.
//This could indicate a CSRF attack.
func ValidateCSRF(r *http.Request, page *page) bool {
	r.ParseForm()
	if len(r.Form["CSRF"]) > 0 {
		if (r.Form["CSRF"][0] != page.UserSession.CSRF) || (decrypt([]byte(page.UserSession.CSRFKey), r.Form["CSRF"][0]) != page.UserSession.CSRFBase) {
			glog.Infoln(r.Form["CSRF"][0])
			glog.Infoln(page.UserSession.CSRF)
			glog.Infoln(decrypt([]byte(page.UserSession.CSRFKey), r.Form["CSRF"][0]))
			glog.Infoln(page.UserSession.CSRFBase)
			page.Messages["modifyFail"] = "Modification failed. You tried to navigate backwards and resubmit!"
			if r.Form["CSRF"][0] != page.UserSession.CSRF {
				glog.Errorln("ERROR: Incorrect CSRF, possible CSRF attack!")
			}
			return false
		}
	} else {
		glog.Errorln("ERROR: No CSRF ID provided, possible CSRF attack!")
	}
	return true
}

//Converts a float value to a vulgar fractional string i.e. .5 to ½
func FloatToVulgar(val float64) string {
	realPart := val
	integerPart := math.Floor(realPart)
	decimalPart := realPart - integerPart
	var intStringPart string
	if int(integerPart) == 0 {
		intStringPart = ""
	} else {
		intStringPart = strconv.Itoa(int(integerPart))
	}
	if decimalPart == 0.0 {
		return intStringPart
	} else if decimalPart <= 0.25 {
		return intStringPart + "¼"
	} else if decimalPart <= .5 {
		return intStringPart + "½"
	} else if decimalPart <= .75 {
		return intStringPart + "¾"
	}
	// if decimalPart == 0.0 {
	// 	return intStringPart
	// } else if decimalPart <= 0.125 {
	// 	return intStringPart + "⅛"
	// } else if decimalPart <= 0.25 {
	// 	return intStringPart + "¼"
	// } else if decimalPart <= 0.333 {
	// 	return intStringPart + "⅓"
	// } else if decimalPart <= 0.375 {
	// 	return intStringPart + "⅜"
	// } else if decimalPart <= .5 {
	// 	return intStringPart + "½"
	// } else if decimalPart <= .625 {
	// 	return intStringPart + "⅝"
	// } else if decimalPart <= .666 {
	// 	return intStringPart + "⅔"
	// } else if decimalPart <= .75 {
	// 	return intStringPart + "¾"
	// } else if decimalPart <= .875 {
	// 	return intStringPart + "⅞"
	// }
	return strconv.Itoa(int(math.Ceil(realPart)))
}

// encrypt string to base64 crypto using AES
func encrypt(key []byte, text string) string {
	//key := []byte(keyText)
	plaintext := []byte(text)

	block, err := aes.NewCipher(key)
	if err != nil {
		glog.Errorln(err)
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		panic(err)
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	// convert to base64
	return base64.URLEncoding.EncodeToString(ciphertext)
}

// decrypt from base64 to decrypted string
func decrypt(key []byte, cryptoText string) string {
	ciphertext, _ := base64.URLEncoding.DecodeString(cryptoText)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// The IV needs to be unique, but not secure. Therefore it's common to
	// include it at the beginning of the ciphertext.
	if len(ciphertext) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)

	// XORKeyStream can work in-place if the two arguments are the same.
	stream.XORKeyStream(ciphertext, ciphertext)

	return fmt.Sprintf("%s", ciphertext)
}

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}
	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}

func LoadFileToS3(filename string, bucket string) {
	aws_access_key_id := viper.GetString("aws_access_key_id")
	aws_secret_access_key := viper.GetString("aws_secret_access_key")
	token := ""
	creds := credentials.NewStaticCredentials(aws_access_key_id, aws_secret_access_key, token)

	_, err := creds.Get()
	if err != nil {
		glog.Infof("bad credentials: %s", err)
	}

	cfg := aws.NewConfig().WithRegion("us-east-1").WithCredentials(creds)

	svc := s3.New(session.New(), cfg)

	file, err := os.Open(filename)

	if err != nil {
		glog.Infof("err opening file: %s", err)
	}

	defer file.Close()

	fileInfo, _ := file.Stat()

	size := fileInfo.Size()

	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)

	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := filepath.Base(file.Name())

	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
		ACL:           aws.String("public-read"),
	}

	resp, err := svc.PutObject(params)
	if err != nil {
		glog.Infof("bad response: %s", err)
	}

	glog.Infof("response %s", awsutil.StringValue(resp))
}
