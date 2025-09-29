package advent5

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

func TestUnits(t *testing.T) {

	file, _ := os.Open("./test.txt")
	defer file.Close()

	t.Run("TestCreateMap", func(t *testing.T) {
		scanner := bufio.NewScanner(file)
		m, err := CreateMap(scanner)
		if err != nil {
			t.Fatal("Test crashed unexpectedly")
		}

		fmt.Printf("Map is: %v", m)
	})

	t.Run("TestGetNumber", func(t *testing.T) {
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()
			if len(line) < 3 {
				break
			}

			runes := []rune(line)

			strKey, strValue := GetNumbers(runes)

			fmt.Printf("key: %s, value: %s\n", strKey, strValue)
		}
	})
}

func TestSumMidPageValidUpdates(t *testing.T) {
	got, err := SumMidPageValidUpdates("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	want := 1

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}

func TestSumMidPageInvalidUpdates(t *testing.T) {
	got, err := SumMidPageInvalidUpdates("./test.txt")

	if err != nil {
		t.Fatalf("Unexpected error")
	}

	want := 1

	if got != want {
		t.Errorf("got: %d, want: %d", got, want)
	}
}
