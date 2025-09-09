package advent1

import (
	"unicode"
	// "io"
	"math"
	"os"
	"strconv"
	"slices"
)




func PrepareData(pathToFile string) ([]int, []int, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, nil, err
	}

	runeData := []rune(string(data))
	var toAppend []rune
	var numbers1 []int
	var numbers2 []int

	for _,v := range runeData {
		if unicode.IsSpace(v) != true {
			toAppend = append(toAppend, v)
			continue
		}
		if len(toAppend) == 0 {
			continue
		}
		
		if v == '\t' {
			numbers1, toAppend, err = AppendNumber(numbers1, toAppend)
			if err != nil {
				return nil, nil, err
			}
		} 

		if v == '\n' {
			numbers2, toAppend, err = AppendNumber(numbers2, toAppend)
		} else {
			numbers1, toAppend, err = AppendNumber(numbers1, toAppend)
		} 
		if err != nil {
			return nil, nil, err
		}
	}
	return numbers1, numbers2, nil
}



func AppendNumber(slice []int, toAppend []rune) ([]int, []rune, error) {
	numToInsert, err := strconv.Atoi(string(toAppend))
	if err != nil {
		return slice, toAppend, err
	}

	slice = append(slice, numToInsert)
	toAppend = toAppend[:0]	

	return slice, toAppend, nil
}


func TotalDistance(numbers1, numbers2 []int) int {


	slices.Sort(numbers1)
	slices.Sort(numbers2)

	sum := 0
	for i := 0 ; i < len(numbers1) ; i++ {
		sum += int(math.Abs(float64(numbers1[i] - numbers2[i])))
	}

	return sum
}