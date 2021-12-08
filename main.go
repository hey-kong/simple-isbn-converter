package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 将 ISBN-10 转换为 ISBN-13
func ConvertToISBN13(isbn10 string) (string, error) {
	if len(isbn10) != 10 {
		return "", fmt.Errorf("ISBN with length 10 is required. Given: %s", isbn10)
	}

	first12Digits := "978" + string(isbn10[:9])
	checkDigitStr, err := calCheckDigitISBN13(first12Digits)
	if err != nil {
		return "", fmt.Errorf("Failed to calculate check digit. Error: %v", err)
	}
	return first12Digits + checkDigitStr, nil
}

// 计算 ISBN-13 的校验位
func calCheckDigitISBN13(isbn13 string) (string, error) {
	multipliers := []int{1, 3}
	sum := 0
	for idx, char := range isbn13[:12] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", fmt.Errorf("Failed to convert char to int: %s", string(char))
		}
		sum += digit * multipliers[idx%2]
	}

	checkDigit := (10 - sum%10) % 10
	return strconv.Itoa(checkDigit), nil
}

// 将 ISBN-13 转换为 ISBN-10
func ConvertToISBN10(isbn13 string) (string, error) {
	if len(isbn13) != 13 {
		return "", fmt.Errorf("ISBN with length 13 is required. Given: %s", isbn13)
	}
	if !strings.HasPrefix(isbn13, "978") {
		return "", fmt.Errorf("Given ISBN-13 is not convertible to ISBN10: %s", isbn13)
	}

	checkDigitStr, err := calCheckDigitISBN10(string(isbn13[3:12]))
	if err != nil {
		return "", fmt.Errorf("Failed to calculate check digit. Error: %v", err)
	}
	return string(isbn13[3:12]) + checkDigitStr, nil
}

// 计算 ISBN-10 的校验位
func calCheckDigitISBN10(isbn10 string) (string, error) {
	sum := 0
	for idx, char := range isbn10[:9] {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			return "", fmt.Errorf("Failed to convert char to int: %s", string(char))
		}
		sum += digit * (10 - idx)
	}

	checkDigit := (11 - sum%11) % 11

	if checkDigit == 10 {
		return "X", nil
	}
	return strconv.Itoa(checkDigit), nil
}

func main() {
	ISBN10 := "7506287641"
	expected := "9787506287647"
	res, _ := ConvertToISBN13(ISBN10)
	fmt.Printf("Convert to ISBN13, result: %s, expect: %s\n", res, expected)

	ISBN13 := "9787307047310"
	expected = "7307047314"
	res, _ = ConvertToISBN10(ISBN13)
	fmt.Printf("Convert to ISBN10, result: %s, expect: %s\n", res, expected)
}
