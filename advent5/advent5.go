package advent5

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"slices"
)

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

	// iterate update from end to start:
	for i := len(update)-1 ; i >= 0 ; i-- {
		// look up element in map.
		toInspect := update[i]
		rules, isExist := m[toInspect]
		if ! isExist {
			continue
		}

		if IsRulesViolated(rules, update[0 : i + 1]) {
			return false
		}
	}
	return true
}

func IsRulesViolated(rules []string, update []string ) bool {
	
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