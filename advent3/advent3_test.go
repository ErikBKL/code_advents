package advent3

import (
	"fmt"
	"testing"
)

func TestFindTargetInSrc(t *testing.T) {

	buffer := []rune("0123mul78mul")

	t.Run("first mul", func(t *testing.T){
		got, err := FindTargetInSrc([]rune("mul"), buffer, 0)
		if err != nil {
			fmt.Println(err.Error())
		}
		want := 4
	
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("second mul", func(t *testing.T){
		got, err := FindTargetInSrc([]rune("mul"), buffer,  5)
		if err != nil {
			fmt.Println(err.Error())
		}
		want := 9
	
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestFindComma(t *testing.T) {
	got, err := FindComma([]rune("aaa,aaa"),0)
	if err != nil {
		fmt.Println(err.Error())
	}

	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSumOfAllMul (t *testing.T) {
	got, err := SumOfAllMul("./test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	want := 10000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}