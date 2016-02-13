package common

import (
	"log"
	"os"
	"strings"
)

//replace substrings in string by search=>replace slices
func StringReplaceBySlice(subject string, search, replace []string) string {
	if len(search) != len(replace) {
		log.Printf("Unable to replace in string by slice, length of search and replace slices are not equals!"+
			" string: %s, search: %v, replace: %v  ", subject, search, replace)

		return subject
	}

	replaceData := make([]string, 0)
	for key, srch := range search {
		replaceData = append(replaceData, srch)
		replaceData = append(replaceData, replace[key])
	}

	replacer := strings.NewReplacer(replaceData...)
	subject = replacer.Replace(subject)

	return subject
}

//util,check exists in slice
func CheckInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

//util, check slices intersects
func CheckIntersect(original, extras []string) bool {
	for _, i := range original {
		for _, x := range extras {
			if i == x {
				return true
			}
		}
	}

	return false
}

//util, remove string from slice
func RemoveStringFromStringSlice(str string, elements []string) []string {
	result := []string{}
	for _, el := range elements {
		if str != el {
			result = append(result, el)
		}
	}
	return result
}

//check file exists
func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func SliceCountValues(input []string) map[string]int {
	result := make(map[string]int)
	for _, val := range input {
		if _, ok := result[val]; ok {
			result[val]++
		} else {
			result[val] = 1
		}
	}

	return result
}
