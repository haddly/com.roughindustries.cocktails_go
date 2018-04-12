// Copyright 2017 Rough Industries LLC. All rights reserved.
//controller/i18n.go:
package www

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"strings"
)

var i18ns map[string]I18n

type I18n struct {
	Site    string
	Locales map[string]Locale
}

type Locale struct {
	Locale map[string]Translation
}

type Translation struct {
	Translation string `json:"translation"`
}

func GetI18nMap() map[string]I18n {
	if len(i18ns) == 0 {
		i18n_list := viper.Get("i18n")
		ilist := i18n_list.([]interface{})
		i18ns = make(map[string]I18n)
		for _, val := range ilist {
			tmp := val.(map[string]interface{})
			var i18nItem I18n
			if tmp["site"] != nil {
				if tmp["site"] != nil {
					i18nItem.Site = tmp["site"].(string)
				} else {
					i18nItem.Site = ""
				}
				if tmp["I18n"] != nil {
					locales := tmp["I18n"].([]interface{})
					i18nItem.Locales = make(map[string]Locale)
					for _, locale := range locales {
						raw, err := ioutil.ReadFile("./view/webcontent/" + i18nItem.Site + "/i18n/" + strings.ToLower(locale.(string)) + ".json")
						if err != nil {
							log.Errorln(err.Error())
							break
						}
						var t map[string]Translation
						json.Unmarshal(raw, &t)
						i18nItem.Locales[locale.(string)] = Locale{
							Locale: make(map[string]Translation),
						}
						for k, v := range t {
							i18nItem.Locales[locale.(string)].Locale[k] = v
						}
					}
				}
				i18ns[i18nItem.Site] = i18nItem
			}
		}
	}
	// log.Infoln(i18ns)

	return i18ns
}
