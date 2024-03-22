package main

import (
	"sort"
	"strings"
)

func GetCharCount(str string) (map[string]int, []string) {
	arr := strings.Split(strings.ToLower(str), "")
	// Count number of chars
	result := make(map[string]int)
	// List of chars
	list := make([]string, 0)
	for _, c := range arr {
		if c == "'" || c == "," || c == "." || c == " " || c == "-" || c == "_" ||
			c == "!" || c == "?" || c == ":" || c == ";" || c == "(" || c == ")" {
			continue
		}
		if _, ok := result[c]; ok {
			result[c]++
			continue
		}
		list = append(list, c)
		result[c] = 1
	}

	return result, list
}

func FindTopDuplicates1(str string) []string {

	acc, list := GetCharCount(str)

	sort.Slice(list, func(i, j int) bool {
		return acc[list[i]] > acc[list[j]]
	})
	return list[0:2]
}

func FindTopDuplicates2(str string) (result []string) {
	acc, _ := GetCharCount(str)
	top1 := ""
	top2 := ""
	for k, v := range acc {
		if top1 == "" {
			top1 = k
			continue
		}
		if top2 == "" && acc[top1] < v {
			top2 = top1
			top1 = k
			continue
		}
		if top2 == "" && acc[top1] >= v {
			top2 = k
			continue
		}

		if v > acc[top1] {
			top2 = top1
			top1 = k
			continue
		}

		if v > acc[top2] && v < acc[top1] {
			top2 = k
			continue
		}

	}

	return []string{top1, top2}
}
