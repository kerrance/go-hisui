package app

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"strconv"
	"strings"
)

func ConvertStringToTitleCase(stringToConvert string) string {
	return cases.Title(language.Und, cases.NoLower).String(stringToConvert)
}

func ConvertStringToKebabCase(stringToConvert string) string {
	return strings.ToLower(ReplaceHyphensWithSpaces(stringToConvert))
}

func ReplaceHyphensWithSpaces(stringToConvert string) string {
	return strings.ReplaceAll(stringToConvert, " ", "-")
}

func ReplaceSpacesWithHyphens(stringToConvert string) string {
	return strings.ReplaceAll(stringToConvert, "-", " ")
}

func ConvertDecimetersToMeters(height int) string {
	if height >= 10 {
		var heightAsFloat = float64(height)
		var adjustedHeight = heightAsFloat / float64(10)
		return fmt.Sprintf("%g", adjustedHeight) + "m"
	} else {
		return "0." + strconv.Itoa(height) + "m"
	}
}

func ConvertHectogramsToKilograms(weight int) string {
	if weight >= 10 {
		var weightAsFloat = float64(weight)
		var adjustedWeight = weightAsFloat / float64(10)
		return fmt.Sprintf("%g", adjustedWeight) + "kg"
	} else {
		var weightAsFloat = float64(weight)
		var adjustedWeight = weightAsFloat * 2.2
		var weightAsPounds = adjustedWeight / float64(10)
		return fmt.Sprintf("%.1f", weightAsPounds) + "lbs"
	}
}
