package advent11

import (
	"os"
	"strconv"
	"strings"
	// "slices"
)

// func CountStones(pathToFile string) (int, error) {
// 	bytesData, err := os.ReadFile(pathToFile)
// 	if err != nil {
// 		return 0, err
// 	}

// 	stringData := string(bytesData)
// 	input := strings.Split(stringData, " ")
	
// 	for blinks := 0 ; blinks < 25 ; blinks++ {

// 		nextInput := []string{}

// 		for _,v := range input {
// 			stoneNumber, err := strconv.Atoi(v)
// 			if err != nil {
// 				return 0, err
// 			}
	
		
// 			if stoneNumber == 0 {
// 				nextInput = append(nextInput, "1")
// 			} else if len(v) % 2 == 0 {
// 				left,right, err := SplitNumber(v)
// 				if err != nil {
// 					return 0, err
// 				}
// 				nextInput = append(nextInput, left)
// 				nextInput = append(nextInput, right)
// 			} else {
// 				toInsert := strconv.Itoa(stoneNumber * 2024)
// 				nextInput = append(nextInput, toInsert)
// 			}
			
// 		}
// 		input = nextInput
// 		fmt.Printf("%+v", input)
// 	}
// 	return len(input), nil
// }

func SplitNumber(original string)(string, string, error) {

	length := len(original)
	left := original[0 : length/2]
	right := original[length/2 : length]

	numRight, err := strconv.Atoi(right)
	if err != nil {
		return "","", err
	}

	right = strconv.Itoa(numRight)

	return left, right, nil
}


// func CountStonesPt2(pathToFile string) (int, error) {
// 	bytesData, err := os.ReadFile(pathToFile)
// 	if err != nil {
// 		return 0, err
// 	}

// 	stringData := string(bytesData)
// 	input := strings.Split(stringData, " ")
// 	cumulativeLen := 0

// 	for _,v := range input {
		
// 		nextInput := []string{v}
// 		ch := make(chan int)
// 		go func () {
// 			for blinks := 0 ; blinks < 75 ; blinks++ {
			
// 				currValueExpansion := []string{}
				
// 				for _, val := range nextInput {
					
					
// 					stoneNumber, err := strconv.Atoi(val)
// 					if err != nil {
// 						ch <- -1
// 					}
			
				
// 					if stoneNumber == 0 {
// 						currValueExpansion = append(currValueExpansion, "1")
// 					} else if len(val) % 2 == 0 {
// 						left,right, err := SplitNumber(val)
// 						if err != nil {
// 							ch <- -1
// 						}
// 						currValueExpansion = append(currValueExpansion, left)
// 						currValueExpansion = append(currValueExpansion, right)
// 					} else {
// 						toInsert := strconv.Itoa(stoneNumber * 2024)
// 						currValueExpansion = append(currValueExpansion, toInsert)
// 					}
// 				}
// 				nextInput = currValueExpansion
// 			}
// 			ch <- len(nextInput)
// 		}()
// 		cumulativeLen += <- ch
// 	}

// 	return cumulativeLen, nil
// }



func CountStones(pathToFile string) (int, error) {
	bytesData, err := os.ReadFile(pathToFile)
	if err != nil {
		return 0, err
	}

	stringData := string(bytesData)
	input := strings.Split(stringData, " ")
	m := map[int]int{}
	for _,v := range input {
		n,_ := strconv.Atoi(v)
		m[n]++
	}

	for blinks := 0 ; blinks < 75 ; blinks++ {
		
		tmpMap := map[int]int{}

		for k := range m {
					
			if k == 0 {
				tmpMap[1] += m[k]
			} else if len(strconv.Itoa(k)) % 2 == 0 {
				left,right, err := SplitNumber(strconv.Itoa(k))
				if err != nil {
					return 0, err
				}
				leftNum,_ := strconv.Atoi(left)
				rightNum,_ := strconv.Atoi(right)
				tmpMap[leftNum] += m[k]
				tmpMap[rightNum] += m[k]

			} else {
				tmpMap[k*2024] += m[k]
			}
			
		}
		m = tmpMap
	}

	ret := 0
	for _,val := range m {
		ret +=  val
	}
	return ret, nil
}