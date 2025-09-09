package advent1

import (
	"strconv"
	"slices"
	"math"
	"os"
)




func PrepareData(pathToFile string) ([]int, error) {
	data, err := os.ReadFile(pathToFile)
	if err != nil {
		return nil, err
	}

	runeData := []rune(string(data))
	var toAppend []rune
	var ret []int
	for _,v := range runeData {
		if v != '\n' {
			toAppend = append(toAppend, v)
			continue
		}
		
		ret, toAppend, err = AppendNumber(ret, toAppend)
		if err != nil {
			return nil, err
		}
	}

	if len(toAppend) > 0 {
		ret, toAppend, err = AppendNumber(ret, toAppend)
		if err != nil {
			return nil, err
		}
	}
	return ret, nil
}



func AppendNumber(ret []int, toAppend []rune) ([]int, []rune, error) {
	numToInsert, err := strconv.Atoi(string(toAppend))
	if err != nil {
		return ret, toAppend, err
	}

	ret = append(ret, numToInsert)
	toAppend = toAppend[:0]	

	return ret, toAppend, nil
}


// func PreapreData (pathToFile string) ([]int, []int, error) {
// 	// open file
// 	// read the whole file into a buffer
// 	data, err := os.ReadFile("./lists")
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	runeData := []rune(string(data))
// 	// iterate the buffer
// 	var runes []rune
// 	var numbers1 []int
// 	var numbers2 []int

// 	toAppend := 0

// 	for _,v := range runeData {

// 		// every time reach a tab append to slice1
// 		if unicode.IsSpace(v) && v == '\t' {
// 			toAppend, err = strconv.Atoi(string(runes))
// 			if err != nil {
// 				return nil, nil, err
// 			}
// 			numbers1 = append(numbers1, toAppend )
// 			runes = runes[:0]
// 			// every time reach a newline append to slice2
// 		} else if  unicode.IsSpace(v) && v == '\n'{
// 			toAppend, err = strconv.Atoi(string(runes))
// 			if err != nil {
// 				return nil, nil, err
// 			}
// 			numbers2 = append(numbers2,toAppend)
// 			runes = runes[:0]
// 		}
		
// 		runes = append(runes, v)
// 	}

// 	return numbers1, numbers2, nil
// }		

func TotalDistance(numbers1, numbers2 []int) int {


	slices.Sort(numbers1)
	slices.Sort(numbers2)

	sum := 0
	for i := 0 ; i < len(numbers1) ; i++ {
		sum += int(math.Abs(float64(numbers1[i] - numbers2[i])))
	}

	return sum
}