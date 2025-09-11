package advent3

import (
	// "bufio"
	"errors"
	"os"
	"slices"
	"strconv"
	"unicode"
)

func SumOfAllMul(pathToFile string) (int, error) {

	bytesData, err := os.ReadFile(pathToFile)
	if err != nil {
		return 0, err
	}

	runeData := []rune(string(bytesData))
	ret := 0
	for idxMul := 0; idxMul < len(runeData)-1; idxMul++ {

		idxMul, err = FindMul(runeData, idxMul)
		if err != nil {
			return ret, nil
		}

		idxLeftParenthesis, err := FindLeftParenthesis(runeData, idxMul+3)
		if idxLeftParenthesis != idxMul+3 {
			idxMul = idxMul + 1
			continue
		}

		idxComma, err := FindComma(runeData, idxLeftParenthesis+1)
		if err != nil {
			idxMul = idxMul + 1
			continue
		}
		if idxComma > idxLeftParenthesis+4 {
			idxMul = idxMul + 1
			continue
		}

		idxRightParenthesis, err := FindRightParenthesis(runeData, idxComma+1)
		if err != nil {
			idxMul = idxMul + 1
			continue
		}
		if idxRightParenthesis > idxComma+4 {
			idxMul = idxMul + 1
			continue
		}

		num1, err := ValidNumber(runeData, idxLeftParenthesis+1, idxComma)
		if err != nil {
			idxMul = idxMul + 1
			continue
		}
		num2, err := ValidNumber(runeData, idxComma+1, idxRightParenthesis)
		if err != nil {
			idxMul = idxMul + 1
			continue
		}

		ret += num1 * num2
		idxMul = idxRightParenthesis

	}

	return ret, nil
}

func ValidNumber(buffer []rune, start, end int) (int, error) {
	for i := start; i < end; i++ {
		if !unicode.IsDigit(buffer[i]) {
			return 0, errors.New("Invalid token")
		}
	}

	num1, err := strconv.Atoi(string(buffer[start:end]))
	if err != nil {
		return 0, err
	}

	return num1, nil
}

func FindTargetInSrc(target []rune, src []rune, startIdx int) (int, error) {
	// assuming no whitespace

	for i := startIdx; i < len(src)-(len(target)-1); i++ {
		toInspect := src[i : i+len(target)]

		if slices.Equal(toInspect, target) {
			return i, nil
		}
	}

	return 0, errors.New("Target not found")
}

func FindMul(src []rune, idxStart int) (int, error) {
	ret, err := FindTargetInSrc([]rune("mul"), src, idxStart)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func FindLeftParenthesis(src []rune, idxAfterMul int) (int, error) {
	ret, err := FindTargetInSrc([]rune("("), src, idxAfterMul)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func FindComma(src []rune, idxAfterParenthesis int) (int, error) {
	ret, err := FindTargetInSrc([]rune(","), src, idxAfterParenthesis)
	if err != nil {
		return 0, err
	}

	return ret, nil
}

func FindRightParenthesis(src []rune, idxAfterComma int) (int, error) {
	ret, err := FindTargetInSrc([]rune(")"), src, idxAfterComma)
	if err != nil {
		return 0, err
	}

	return ret, nil
}
