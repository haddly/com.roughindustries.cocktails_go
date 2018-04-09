// Copyright 2017 Rough Industries LLC. All rights reserved.
//model/search.go:package model
package model

import ()

//DATA STRUCTURES
//Search data structure
type Search struct {
	Errors                 map[string]string
	Include_Ingredients    []int
	Include_NonIngredients []int
	Include_Metas          []int
	Exclude_Ingredients    []int
	Exclude_NonIngredients []int
	Exclude_Metas          []int
	RatingMin              int
	RatingMax              int
	Keywords               string
}

//ENUMERATIONS
