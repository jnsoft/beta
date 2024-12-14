package main

import (
	"fmt"
)

func eq(c1, c2 rune) bool {
	if c2 == '.' {
		return true
	} else if c1 == c2 {
		return true
	} else {
		return false
	}

}

func isMatch(s string, p string) bool {
	p_ix := 0
	var c1, c2 rune

	for i := 0; i < len(s); {
		if p_ix >= len(p) {
			return false
		}
		c1 = rune(s[i])
		c2 = rune(p[p_ix])

		if p_ix < len(p)-1 && rune(p[p_ix+1]) == '*' {
			p_ix++
		} else if c2 == '*' {
			c2 = rune(p[p_ix-1])
			for eq(c1, c2) {
				if i < len(s)-1 {
					i++
					c1 = rune(s[i])
				} else {
					if p_ix == len(p)-1 || eq(c1, rune(p[p_ix+1])) {
						return true
					} else {
						i++
						p_ix++
					}
				}
			}
			p_ix++
		} else if !eq(c1, c2) {
			return false
		} else {
			i++
			if i < len(s) {
				p_ix++
			}
		}
	}
	if p_ix >= len(p)-1 {
		return true
	} else {
		return false
	}
}

func main() {
	s := "aa"
	p := ".*"
	test := isMatch(s, p)
	AssertTrue(test)

	s = "mississippi"
	p = "mis*is*ip*."

	test = isMatch(s, p)
	AssertTrue(test)

	s = "ab"
	p = ".*c"
	test = isMatch(s, p)
	AssertFalse(test)

	s = "aaa"
	p = "aaaa"
	test = isMatch(s, p)
	AssertFalse(test)

	s = "aaa"
	p = "a*a"
	test = isMatch(s, p)
	AssertTrue(test)

	s = "aaa"
	p = "ab*a*c*a"
	test = isMatch(s, p)
	AssertTrue(test)

}

func AssertTrue(ex bool) {
	if ex {
		fmt.Println("OK")
	} else {
		panic("Error!")
	}
}

func AssertFalse(ex bool) {
	if !ex {
		fmt.Println("OK")
	} else {
		panic("Error!")
	}
}
