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

func Annagrams(strArray *[]string) *map[string][]string {
	annagrams := make(map[string][]string)
	keysToAnnagrams := make(map[string]string)
	for _, word := range *strArray {
		sortedWord := sortString(strings.ToLower(word))
		if keysToAnnagrams[sortedWord] == "" {
			keysToAnnagrams[sortedWord] = strings.ToLower(word)
		}
		annagrams[keysToAnnagrams[sortedWord]] = append(annagrams[keysToAnnagrams[sortedWord]], strings.ToLower(word))
	}
	for k, v := range annagrams {
		if len(annagrams[k]) == 1 {
			removeSingeElementKey(k, annagrams)
			continue
		}
		sort.Strings(v)
		annagrams[k] = removeDuplicates(v)
	}
	return &annagrams
}

func removeSingeElementKey(key string, mp map[string][]string) {
	if len(mp[key]) == 1 {
		delete(mp, key)
	}
}

func removeDuplicates(strArray []string) []string {
	n := 1
	for i := 1; i < len(strArray); i++ {
		if strArray[i] != strArray[n-1] {
			strArray[n] = strArray[i]
			n++
		}
	}
	return strArray[:n]
}
