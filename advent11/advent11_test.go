package advent11

import(
	"testing"
)

func TestCountStones(t *testing.T){
	got,_ := CountStones("./test.txt") 
	want := 3

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestSplitNumber(t *testing.T) {
	var table = []struct{
		input	string
		left	string
		right	string
	}{
		{input: "1122", left: "11", right: "22"},
		{input: "132227", left: "132", right: "227"},
		{input: "1000", left: "10", right: "0"},

	}

	for _,testcase := range table {
		left, right, err := SplitNumber(testcase.input)

		if err != nil {
			t.Fatalf("unexpected error")
		}
		if left != testcase.left || right != testcase.right {
			t.Errorf("got left %v and right %v, want left %v and right %v", left, right, testcase.left, testcase.right)
		}
	}
}