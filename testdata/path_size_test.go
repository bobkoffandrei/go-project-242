package code

import (
	"github.com/bobkoffandrei/go-project-242/code"
	"testing"
)

func TestGetPathSize_FileEmpty(t *testing.T) {
	const trueSize = "0B"
	value, err := code.GetPathSize("./testdir/file_empty", false, false, false)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}

func TestGetPathSize_File(t *testing.T) {
	const trueSize = "425715B"
	value, err := code.GetPathSize("./testdir/file_with_data", false, false, false)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}

func TestGetPathSize_Dir(t *testing.T) {
	const trueSize = "425715B"
	value, err := code.GetPathSize("./testdir", false, false, false)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}

func TestGetPathSize_Dir_Rek(t *testing.T) {
	const trueSize = "1277145B"
	value, err := code.GetPathSize("./testdir", true, false, false)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}

func TestGetPathSize_Dir_Rek_Human(t *testing.T) {
	const trueSize = "1.22MB"
	value, err := code.GetPathSize("./testdir", true, true, false)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}

func TestGetPathSize_Dir_Hidden(t *testing.T) {
	const trueSize = "425715B"
	value, err := code.GetPathSize("./testdir", false, false, true)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}

func TestGetPathSize_All(t *testing.T) {
	const trueSize = "1.62MB"
	value, err := code.GetPathSize("./testdir", true, true, true)
		if err != nil {
			t.Fatal(err)
		}

		if trueSize != value {
			t.Errorf("got %s, want %s", value, trueSize)
		}
}