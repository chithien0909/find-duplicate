package main

import (
	"fmt"
	"sort"
	"strings"
)

func FindTopDuplicates1(str string) (list []string) {
	arr := strings.Split(strings.ToLower(str), "")
	acc := make(map[string]int)
	list = make([]string, 2)
	for _, c := range arr {
		if c == "'" || c == "," || c == "." || c == " " || c == "-" || c == "_" ||
			c == "!" || c == "?" || c == ":" || c == ";" || c == "(" || c == ")" {
			continue
		}
		if _, ok := acc[c]; ok {
			acc[c]++
			continue
		}
		list = append(list, c)
		acc[c] = 1
	}
	sort.Slice(list, func(i, j int) bool {
		return acc[list[i]] > acc[list[j]]
	})
	return list[0:2]
}

func FindTopDuplicates2(str string) (result []string) {
	arr := strings.Split(strings.ToLower(str), "")
	acc := make(map[string]int)
	list := make([]string, 0)
	for _, c := range arr {
		if c == "'" || c == "," || c == "." || c == " " || c == "-" || c == "_" ||
			c == "!" || c == "?" || c == ":" || c == ";" || c == "(" || c == ")" {
			continue
		}
		if _, ok := acc[c]; ok {
			acc[c]++
			continue
		}
		list = append(list, c)
		acc[c] = 1
	}
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

	fmt.Println(acc)

	return []string{top1, top2}
}
