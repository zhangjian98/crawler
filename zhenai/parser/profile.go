package parser

import (
	"crawler/engine"
	"crawler/model"
)

//var allRe  = regexp.MustCompile(`<div class="m-btn purple" data-v-bff6f798>([\s\S]*)</div>`)
func ParseProfile(name string) engine.ParseResult {
	profile := model.Profile{}

	//age, err := strconv.Atoi(extractString(contents, allRe))
	//if err != nil {
	//	profile.Age = age
	//}
	//profile.Marriage = extractString(contents, marriageRe)
	age := "test"
	profile.Name = name
	profile.Age = 11
	profile.Gender = age
	profile.Height = 12
	profile.Weight = 113
	profile.Income = age
	profile.Marriage = age
	profile.Education = age
	profile.Occupation = age
	profile.Hokon = age
	profile.Xinzuo = age
	profile.House = age
	profile.Car = age

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}
	return result
}

//func extractString(contents []byte,  re *regexp.Regexp) string{
//	match := re.FindSubmatch(contents)
//
//	if len(match) > 2 {
//		return string(match[1])
//	}else{
//		return ""
//	}
//
//}
