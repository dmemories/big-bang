package utils

import (
	"fmt"
	"strings"
)

func ToFirstUpper(input string) string {
	if len(input) == 0 {
		return input
	}
	return strings.ToUpper(input[:1]) + strings.ToLower(input[1:])
}

func ToFirstLower(input string) string {
	if len(input) > 0 {
		input = strings.ToLower(string(input[0])) + input[1:]
	}

	return input
}

func GetMethodFileName(requestMethod, subName string) string {
	return fmt.Sprintf("%s_%s", strings.ToLower(requestMethod), subName)
}

var irregularPlurals = map[string]string{
	"mouse":    "mice",
	"goose":    "geese",
	"child":    "children",
	"person":   "people",
	"tooth":    "teeth",
	"foot":     "feet",
	"ox":       "oxen",
	"cactus":   "cacti",
	"radius":   "radii",
	"analysis": "analyses",
	"thesis":   "theses",
	"octopus":  "octopi",
	"fungus":   "fungi",
	"datum":    "data",
	"medium":   "media",
	"index":    "indices",
}

func GetPluralName(baseMethodName string) string {
	if len(baseMethodName) == 0 {
		return ""
	}

	// Check if it's an irregular plural
	if plural, exists := irregularPlurals[baseMethodName]; exists {
		return plural
	}

	// คำลงท้ายที่ต้องเติม "es"
	esSuffixes := []string{"s", "sh", "ch", "x", "z", "o"}
	for _, suffix := range esSuffixes {
		if strings.HasSuffix(baseMethodName, suffix) {
			return baseMethodName + "es"
		}
	}

	// เปลี่ยน "y" เป็น "ies" ถ้าหน้า "y" เป็นพยัญชนะ
	if strings.HasSuffix(baseMethodName, "y") && len(baseMethodName) > 1 {
		lastCharBeforeY := baseMethodName[len(baseMethodName)-2]
		if !strings.ContainsRune("aeiou", rune(lastCharBeforeY)) {
			return baseMethodName[:len(baseMethodName)-1] + "ies"
		}
	}

	// เปลี่ยน "f" หรือ "fe" เป็น "ves"
	vesExceptions := map[string]bool{"roo": true, "che": true, "chie": true}
	if strings.HasSuffix(baseMethodName, "fe") {
		root := baseMethodName[:len(baseMethodName)-2]
		if !vesExceptions[root] {
			return root + "ves"
		}
	}
	if strings.HasSuffix(baseMethodName, "f") {
		root := baseMethodName[:len(baseMethodName)-1]
		if !vesExceptions[root] {
			return root + "ves"
		}
	}

	// คำทั่วไปเติม "s"
	return baseMethodName + "s"
}
