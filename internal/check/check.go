package check

import "fmt"

func ParseCorrect(s string, h string) (string, error) {
	var tmp string
	rs := []rune(s)
	fmt.Println(len(rs))
	if len(rs) != 5 || len(h) != 5 {
		return "", fmt.Errorf("poke name or hits not match len")
	}
	for i, r := range rs {
		if string(h[i]) == "." {
			tmp += "_"
			continue
		}
		if string(h[i]) == "o" {
			tmp += string(r)
		} else {
			tmp += "_"
		}
	}
	return tmp, nil
}

func ParsePartial(s string, h string) ([]string, error) {
	var tmp []string
	rs := []rune(s)
	if len(rs) != 5 && len(h) != 5 {
		return []string{}, fmt.Errorf("poke name or hits not match len")
	}
	for i, r := range rs {
		if string(string(r)) == "." {
			continue
		}
		if string(h[i]) == "x" {
			tmp = append(tmp, string(r))
		}
	}
	return tmp, nil
}
