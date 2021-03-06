package clear

import (
	"math/rand"
	"sort"
	"strings"
)

type SamePart struct {
	// same part text
	Text string
	// how much rate allTxt has the same txt
	Rate float64
}

// a same Text detector to discriminate same Text in different page
// allTxt: all text you want to compare
// minLen: the min len keyword you want discriminate, better greater than 10
// minRate: the mini rate to judge this part is same part, suggest greater than 0.3
// return: []*SamePart this result will sort by text len desc, so you can clear the same part in asc order
func DiscriminateSiteTextSamePart(allTxt []string, minLen int, minRate float64) []*SamePart {
	all := len(allTxt)
	if all <= 3 {
		return nil
	}
	if all > 1000 {
		all = 1000
	}
	if len(allTxt) == 0 {
		return nil
	}
	c := 0
	var innerSamePartStrings []string
	//allSpend := int64(0)
	maxCount := all / 10
	if maxCount < 5 {
		maxCount = 5
	}
	for {
		c++
		first := rand.Intn(len(allTxt))
		second := rand.Intn(len(allTxt))
		if first == second {
			c--
			continue
		}
		//timeS := time.Now().UnixNano()/10000000
		cps := findSamePartFromTwoString(allTxt[first], allTxt[second], minLen)
		//allSpend+=time.Now().UnixNano()/10000000 - timeS
		//fmt.Println(c,allSpend)
		for _, cp := range cps {
			same := false
			for _, part := range innerSamePartStrings {
				if part == cp {
					same = true
					break
				}
			}
			if !same {
				innerSamePartStrings = append(innerSamePartStrings, cp)
			}
		}

		if c == maxCount {
			break
		}
	}
	var sps []*SamePart
	for _, part := range innerSamePartStrings {
		c := 0
		for _, s := range allTxt {
			if strings.Contains(s, part) {
				c++
			}
		}
		sp := &SamePart{
			Text: part,
			Rate: float64(c) / float64(len(allTxt)),
		}
		if sp.Rate > minRate {
			sps = append(sps, sp)
		}
	}
	sort.Slice(sps, func(i, j int) bool {
		return len(sps[i].Text) > len(sps[j].Text)
	})
	return sps

}

// This function can find the same part between two strings.
// warning: this function is low efficient, don't drink it too much
func findSamePartFromTwoString(first string, second string, sameMinLen int) []string {
	if len(first) == 0 || len(second) == 0 {
		return nil
	}
	firstParts := strings.Split(first, "")
	firstLen := len(firstParts)
	var commParts []string
	startIndex := 0
	commStr := ""
	commStartIndex := -1
	subLen := sameMinLen
	for startIndex <= firstLen-sameMinLen {
		sub := strings.Join(firstParts[startIndex:startIndex+subLen], "")
		if strings.Contains(second, sub) {
			commStr = sub
			if commStartIndex == -1 {
				commStartIndex = startIndex
			}
			subLen++
		} else {
			if commStartIndex != -1 {
				commStr = strings.TrimSpace(commStr)
				if len(strings.Split(commStr, "")) >= sameMinLen {
					commParts = append(commParts, commStr)
				}
				commStr = ""
				commStartIndex = -1
				startIndex += subLen - 1
				subLen = sameMinLen
			} else {
				startIndex++
			}

		}
		if startIndex+subLen > firstLen {
			if commStartIndex != -1 {
				commStr = strings.TrimSpace(commStr)
				if len(strings.Split(commStr, "")) >= sameMinLen {
					commParts = append(commParts, commStr)
				}
			}
			break
		}

	}

	return commParts
}
