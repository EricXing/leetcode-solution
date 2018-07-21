package problems

import "testing"

func TestMySqrt(t *testing.T) {
	if 2 == MySqrt(4){
		t.Log("Sqrt(x): case 1 succeed.")
	} else {
		t.Error("Sqrt(x): case 1 failed.")
	}
	if 2 == MySqrt(8){
		t.Log("Sqrt(x): case 2 succeed.")
	} else {
		t.Error("Sqrt(x): case 2 failed.")
	}
	if 1 == MySqrt(1){
		t.Log("Sqrt(x): case 3 succeed.")
	} else {
		t.Error("Sqrt(x): case 3 failed.")
	}
}
