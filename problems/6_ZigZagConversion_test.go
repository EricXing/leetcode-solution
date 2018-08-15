package problems

import "testing"

func TestConvert(t *testing.T) {
	if "PAHNAPLSIIGYIR" == Convert("PAYPALISHIRING", 3) {
		t.Log("6. ZigZag Conversion: case 1 succeeded.")
	} else {
		t.Error("6. ZigZag Conversion: case 1 failed.")
	}
	if "PINALSIGYAHRPI" == Convert("PAYPALISHIRING", 4) {
		t.Log("6. ZigZag Conversion: case 2 succeeded.")
	} else {
		t.Error("6. ZigZag Conversion: case 2 failed.")
	}
}
