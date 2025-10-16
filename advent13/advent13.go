package advent13

import (
	"bufio"
	"os"
	"strings"
	"strconv"
	"reflect"
)
type Point struct{ 
	X	int 
	Y 	int 
}

type Challange struct {
	A     	Point
	B     	Point
	Prize 	Point
	Moves	[2]int
}


const (
	APRICE = 3
	BPRICE = 1
	OFFSET = 10000000000000
) 
 

func Min( a, b int) int {
	if a < b {
		return a
	}

	return b
}
func LowestPriceForMaxPrizes(pathToFile string) (int,error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}

	challanges, err := ReadFile(file)
	if err != nil {
		return 0, err
	}

	ret := 0
	for _, ch := range challanges {
		m := MarkAllPoints(ch)

		prices, ok := m[ch.Prize]
		if ok {
			ret += Min(prices[0], prices[1])
		} 
			
	}

	return ret, nil
}


func MarkAllPoints(ch Challange) map[Point][]int {
	m := map[Point][]int{}
	maxMoves := 100
	
	for a := 1 ; a <= maxMoves ; a++ {
		for b := 0 ; b <= maxMoves ; b++ {
			pointToInsert := Point{
				X: ch.A.X * a + ch.B.X * (b), 
				Y: ch.A.Y * a + ch.B.Y * (b),
			}
	
			m[pointToInsert] = append(m[pointToInsert], a * APRICE + b * BPRICE)

			pointToInsert = Point{
				X: ch.B.X * a + ch.A.X * (b), 
				Y: ch.B.Y * a + ch.A.Y * (b),
			}

			m[pointToInsert] = append(m[pointToInsert], a * BPRICE + b * APRICE)

		}
	}

	return m
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

	if !reflect.DeepEqual(ch, emptyChallange) {
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

















func ReadFilePt2(file *os.File ) ([]Challange, error) {
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
			ch.Prize.X = x + OFFSET
			ch.Prize.Y = y + OFFSET
		default:
			challange = append(challange, ch)
			ch = emptyChallange
		}
	}

	
	if !reflect.DeepEqual(ch, emptyChallange) {
		challange = append(challange, ch)
	}

	return challange, nil
}



func LowestPriceForMaxPrizesPt2(pathToFile string) (int,error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return 0, err
	}

	challanges, err := ReadFilePt2(file)
	if err != nil {
		return 0, err
	}

	ret := 0
	for _, ch := range challanges {
		prices := []int{}
		prices = MinPriceForChallange(ch, ch.Prize, Point{0,0}, prices)
		if len(prices) > 1 {
			ret += Min(prices[0], prices[1])
		}
	}

	return ret, nil
}

func MinPriceForChallange(ch Challange, target, current Point, prices []int) []int {
	buttons := []Point{ch.A, ch.B}

	for i,button := range buttons {

		if current.Y > target.Y || current.X > target.X {
			return prices
		} 
	
		if current == target {
			prices = append(prices, APRICE * ch.Moves[0] + BPRICE * ch.Moves[1])
			return prices
		}
		
		ch.Moves[i]++
		prices = MinPriceForChallange(ch, ch.Prize, Point{X: current.X + button.X, Y: current.Y + button.Y}, prices)
		ch.Moves[i]--

	}
	
	return prices
}