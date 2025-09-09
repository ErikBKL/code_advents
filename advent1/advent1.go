package advent1

import (
	"bufio"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)




func PrepareData(pathToFile string) (map[int]int, map[int]int, error) {
	// add a comment
	file, err := os.Open(pathToFile)
	defer file.Close()
	 
	if err != nil {
		return nil, nil, err
	}
	
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

func SimilarityScore(leftList map[int]int, rightList map[int]int) int {

	ret := 0
	for k,_ := range leftList {
		ret += k * rightList[k]
	}

	return ret
}