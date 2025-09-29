package advent2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func AmountSafeReports(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := 0
	for scanner.Scan() {
		lineSlice := strings.Fields(scanner.Text())
		if len(lineSlice) == 0 {
			continue
		}

		numSlice, err := StrSliceToInts(lineSlice)
		if err != nil {
			return 0, err
		}

		if isSafe(numSlice) {
			ret++
		}
	}

	return ret, scanner.Err()
}

func StrSliceToInts(lineSlice []string) ([]int, error) {
	numSlice := make([]int, 0, len(lineSlice))

	for i := range lineSlice {
		toAppend, err := GetIntAtIdx(lineSlice, i)
		if err != nil {
			return nil, err
		}

		numSlice = append(numSlice, toAppend)
	}

	return numSlice, nil
}

func GetIntAtIdx(lineSlice []string, idx int) (int, error) {
	ret, err := strconv.Atoi(lineSlice[idx])
	if err != nil {
		return 0, err
	}

	return ret, nil
}

// isSafe validates a report per Day 2 Part 1 rules.
func isSafe(nums []int) bool {

	increasing := nums[1] > nums[0]

	for i := 1; i < len(nums); i++ {
		diff := nums[i] - nums[i-1]
		if diff == 0 {
			return false
		}
		if diff > 0 && !increasing {
			return false
		}
		if diff < 0 && increasing {
			return false
		}
		if diff > 3 || diff < -3 {
			return false
		}
	}
	return true
}

// AmountSafeReportsWithDampener counts safe reports allowing removal of one level.
func AmountSafeReportsWithDampener(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	ret := 0
	for scanner.Scan() {
		lineSlice := strings.Fields(scanner.Text())
		if len(lineSlice) == 0 {
			continue
		}

		numSlice, err := StrSliceToInts(lineSlice)
		if err != nil {
			return 0, err
		}

		if isSafe(numSlice) || isSafeWithOneRemoval(numSlice) {
			ret++
		}
	}

	return ret, scanner.Err()
}

func isSafeWithOneRemoval(nums []int) bool {
	for i := 0; i < len(nums); i++ {
		if isSafe(skipIndex(nums, i)) {
			return true
		}
	}
	return false
}

func skipIndex(nums []int, skip int) []int {
	// res := make([]int, 0, len(nums)-1)
	res := []int{}
	for i := 0; i < len(nums); i++ {
		if i == skip {
			continue
		}
		res = append(res, nums[i])
	}
	return res
}
