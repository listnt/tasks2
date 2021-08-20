package mymodule

import (
	"sort"
	"strings"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func Annagrams(st *[]string) *map[string][]string {
	res := make(map[string][]string)
	keys := make(map[string]string)
	for i := 0; i < len(*st); i++ {
		k := sortString(strings.ToLower((*st)[i]))
		if keys[k] == "" {
			keys[k] = strings.ToLower((*st)[i])
		}
		res[keys[k]] = append(res[keys[k]], strings.ToLower((*st)[i]))
	}
	for k, v := range res {
		if len(res[k]) == 1 {
			delete(res, k)
			continue
		}
		sort.Strings(v)
		n := 1
		for i := 1; i < len(v); i++ {
			if v[i] != v[n-1] {
				v[n] = v[i]
				n++
			}
		}
		res[k] = v[:n]
	}
	return &res
}
