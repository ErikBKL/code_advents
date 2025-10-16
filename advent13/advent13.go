package advent13

import (
	"bufio"
	"os"
	"strings"
	"strconv"
)
type Point struct{ 
	X	int 
	Y 	int 
}

type Challange struct {
	A     Point
	B     Point
	Prize Point
}

const APRICE = 3
const BPRICE = 1

func LowestPriceForMaxPrizes(pathToFile string) (int,error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}

	challanges, err := ReadFile(file)
	if err != nil {
		return 0, err
	}

	for i, ch := range challanges {
		if ! IsPointReachable(ch) {
			continue
		}



	}


	return 0, nil
}

func ReadFile(file *os.File ) ([]Challange, error) {
	scanner := bufio.NewScanner(file)
	challange := []Challange{}

	// foreach line in file:
	ch := Challange{}
	emptyChallange := Challange{}
	for scanner.Scan() {
		// separate by : - discard left - keep right
		line := scanner.Text()
		switch {
		case strings.Contains(line, "Button A:"):
			x,y,err := ExtractNumbers(line, "+")
			if err != nil {
				return nil, err
			}
			ch.A.X = x
			ch.A.Y = y
		case strings.Contains(line, "Button B:"):
			x,y,err := ExtractNumbers(line, "+")
			if err != nil {
				return nil, err
			}
			ch.B.X = x
			ch.B.Y = y
		case strings.Contains(line, "Prize:"):
			x,y,err := ExtractNumbers(line, "=")
			if err != nil {
				return nil, err
			}
			ch.Prize.X = x
			ch.Prize.Y = y
		default:
			challange = append(challange, ch)
			ch = emptyChallange
		}
	}

	
	if ch != emptyChallange {
		challange = append(challange, ch)
	}

	return challange, nil
}

func ExtractNumbers(line, separator string) (int, int, error) {
		subStrings := strings.Split(line, ":")
		subStrings = subStrings[1:]
		subStrings[0] = strings.TrimSpace(subStrings[0])
		subStrings = strings.Split(subStrings[0], ",")
		subStrings[1] = strings.TrimSpace(subStrings[1])
		subStrings[0] = strings.Split(subStrings[0], separator)[1]
		subStrings[1] = strings.Split(subStrings[1], separator)[1]

		x, err := strconv.Atoi(subStrings[0])
		if err != nil {
			return -1,-1, err
		}
		y, err := strconv.Atoi(subStrings[1])
		if err != nil {
			return -1,-1, err
		}

		return x,y,nil

}

func IsPointReachable(ch Challange) bool {

	if ch.Prize.X % ch.A.
	return true
}