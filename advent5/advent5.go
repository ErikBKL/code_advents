package advent5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"slices"
)
func SumMidPageInvalidUpdates(pathToFile string)(int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m, err := CreateMap (scanner)
	if err != nil {
		return 0, err
	}
	ret := 0
	for {
		update, isFinish := GetUpdate(scanner)
		if isFinish {
			break
		}

		if ! IsUpdateValid(m, update) {
			SortInvalidUpdate(m, update)
			res, err := ValueMidNumber(update)
			if err != nil {
				return 0, err
			}

			ret += res
		}
	}

	return ret, nil
}

func SortInvalidUpdate(m map[string][]string, update []string) {
	for i := 1 ; i < len(update) ; i++ {
		curr := update[i]
		leftOfCurr := update[0 : i]
		
		for j := 0 ; j < len(leftOfCurr) ; j++ {
			if slices.Contains(m[curr], leftOfCurr[j]) {
				update[i], update[j] = update[j], update[i]
			}
		}	
	}
}


func SumMidPageValidUpdates(pathToFile string) (int, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	m, err := CreateMap (scanner)
	if err != nil {
		return 0, err
	}

	ret := 0
	for {
		res := 0
		update, isFinish := GetUpdate(scanner)
		if isFinish {
			break
		}
		
		if ! IsUpdateValid(m, update) {
			continue
		}

		res, err := ValueMidNumber(update)
		if err != nil {
			return 0, err
		}

		ret += res
	}

	return ret, nil
}

func ValueMidNumber(update []string) (int, error) {
	midIdx := len(update) / 2

	ret, err := strconv.Atoi(update[midIdx])
	if err != nil {
		return 0, err
	}
	return ret, nil
}


func IsUpdateValid(m map[string][]string, update []string) bool {

	// foreach number in updates from end to start, 
	for i := len(update)-1 ; i >= 0 ; i-- {
		// look up element in map.
		toInspect := update[i]
		rules, isExist := m[toInspect]
		if ! isExist {
			continue
		}
		
		// if any of the numbers in his rule show up in updates before him - non valid update
		if IsNumberOutOfPlace(rules, update[0 : i + 1]) {
			return false
		}
	}
	return true
}

func IsNumberOutOfPlace(rules []string, update []string ) bool {
	for _,v := range rules {
		if slices.Contains(update, v) {
			return true
		}
	}
	return false
}

func GetUpdate(scanner *bufio.Scanner) ([]string, bool) {

	if scanner.Scan() {
		line := scanner.Text()
		return strings.Split(line, ","), false
	}
	return nil, true
}

func CreateMap(scanner *bufio.Scanner) (map[string][]string, error) {
	m := map[string][]string{}

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 3 {
			break
		}

		runes := []rune(line)

		key, value := GetNumbers(runes)

		m[key] = append(m[key], value)
	}

	return m, nil
}

func GetNumbers(runes []rune) (string, string) {

	retKey := []rune{}
	retValue := []rune{}

	for i,v := range runes {
		if v == '|' {
			retKey = runes[0 : i]
			retValue = runes[i + 1 : ]
			break
		}
	}

	return string(retKey), string(retValue)
}