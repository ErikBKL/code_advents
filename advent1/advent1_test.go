package advent1

import (
	"fmt"
	"testing"
)

func TestTotalDistance(t *testing.T) {

	got1, got2, err := PrepareData("./lists.txt")
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Printf("TotalDistance: %d", TotalDistance(got1, got2))
	fmt.Printf("SimilarityScore: %d", SimilarityScore(got1, got2))
}
