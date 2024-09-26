package digitarray

import (
	"testing"
)

func TestEquals_SameValues_ReturnsTrue(t *testing.T) {
	a1 := DigitArray{3, 4, 9}
	a2 := DigitArray{3, 4, 9}
	if !a1.Equal(a2) {
		t.Error("Expected Arrays to Be Equal")
	}
}

func TestEquals_DifferentValue1_ReturnsTrue(t *testing.T) {
	a1 := DigitArray{3, 4, 9}
	a2 := DigitArray{1, 4, 9}
	if a1.Equal(a2) {
		t.Error("Expected Arrays to Not Be Equal")
	}
}

func TestPlus_Zero(t *testing.T) {
    operand1 := DigitArray{5, 3}
    operand2 := DigitArray{0, }
    result := operand1.Plus(operand2)
    if len(result) != 2 || result[0] != 5 || result[1] != 3 {
        t.Errorf("Expected result %v but got %v", []byte{5, 3}, result)
    }
}

func TestPlus(t *testing.T) {
	operand1 := DigitArray{1, 2}
    operand2 := DigitArray{3, 4}
    result := operand1.Plus(operand2)
    if len(result) != 2 || result[0] != 4 || result[1] != 6 {
        t.Errorf("Expected result %v but got %v", []byte{4, 6}, result)
    }
}

func TestPlus_Carrying(t *testing.T) {
	operand1 := DigitArray{4, 5}
    operand2 := DigitArray{4, 5}
    result := operand1.Plus(operand2)
    if len(result) != 2 || result[0] != 9 || result[1] != 0 {
        t.Errorf("Expected result %v but got %v", []byte{9, 0}, result)
    }
}

func TestPlus_Overflow(t *testing.T) {
	operand1 := DigitArray{7, 5}
    operand2 := DigitArray{8, 5}
    result := operand1.Plus(operand2)
    if len(result) != 2 || result[0] != 16 || result[1] != 0 {
        t.Errorf("Expected result %v but got %v", []byte{16, 0}, result)
    }
}

func TestPlus_DifferentSizesA1B2(t *testing.T) {
	operand1 := DigitArray{7}
    operand2 := DigitArray{8, 5}
    result := operand1.Plus(operand2)
    if len(result) != 2 || result[0] != 15 || result[1] != 5 {
        t.Errorf("Expected result %v but got %v", []byte{15, 5}, result)
    }
}

func TestPlus_DifferentSizesA2B1(t *testing.T) {
	operand1 := DigitArray{6, 4}
    operand2 := DigitArray{5}
    result := operand1.Plus(operand2)
    if len(result) != 2 || result[0] != 11 || result[1] != 4 {
        t.Errorf("Expected result %v but got %v", []byte{11, 4}, result)
    }
}

func TestMinus_Zero(t *testing.T) {
    operand1 := DigitArray{2, 2}
    operand2 := DigitArray{0}
    result := operand1.Minus(operand2)
    if len(result) != 2 || result[0] != 2 || result[1] != 2 {
        t.Errorf("Expected result %v but got %v", "22", result)
    }
}

func TestMinus(t *testing.T) {
    operand1 := DigitArray{5, 3}
    operand2 := DigitArray{2, 2}
    result := operand1.Minus(operand2)
    if len(result) != 2 || result[0] != 3 || result[1] != 1 {
        t.Errorf("Expected result %v but got %v", "31", result)
    }
}

func TestMinus_Zero4_ReturnsLength4WithTrailingZero(t *testing.T) {
    operand1 := DigitArray{3, 9, 4}
    operand2 := DigitArray{0, 0, 0, 0}
    result := operand1.Minus(operand2)
    if len(result) != 4 ||
		result[0] != 3 ||
		result[1] != 9 ||
		result[2] != 4 ||
		result[3] != 0 {
        t.Errorf("Expected result %v but got %v", []byte{3, 9, 4, 0}, result)
    }
}

func TestTrimTrailingZeros_NoZeros(t *testing.T) {
	inputArray := DigitArray{2, 5}
	result := inputArray.TrimTrailingZeros()
	if len(result) != 2 || result[0] != 2 || result[1] != 5 {
		t.Errorf("Expected result %v but got %v", []byte{2, 5}, result)
	}
}

func TestTrimTrailingZeros_OneZero(t *testing.T) {
	inputArray := DigitArray{2, 5, 0}
	result := inputArray.TrimTrailingZeros()
	if len(result) != 2 || result[0] != 2 || result[1] != 5 {
		t.Errorf("Expected result %v but got %v", []byte{2, 5}, result)
	}
}

func TestTrimTrailingZeros_TwoZeros(t *testing.T) {
	inputArray := DigitArray{0, 1, 0, 0}
	result := inputArray.TrimTrailingZeros()
	if len(result) != 2 || result[0] != 0 || result[1] != 1 {
		t.Errorf("Expected result %v but got %v", []byte{0, 1}, result)
	}
}

func TestTrimLeadingZeros_NoZeros(t *testing.T) {
	inputArray := DigitArray{2, 5}
	result := inputArray.TrimLeadingZeros()
	if len(result) != 2 || result[0] != 2 || result[1] != 5 {
		t.Errorf("Expected result %v but got %v", []byte{2, 5}, result)
	}
}

func TestTrimLeadingZeros_OneZero(t *testing.T) {
	inputArray := DigitArray{0, 2, 5}
	result := inputArray.TrimLeadingZeros()
	if len(result) != 2 || result[0] != 2 || result[1] != 5 {
		t.Errorf("Expected result %v but got %v", []byte{2, 5}, result)
	}
}

func TestTrimLeadingZeros_TwoZeros(t *testing.T) {
	inputArray := DigitArray{0, 0, 1, 0}
	result := inputArray.TrimLeadingZeros()
	if len(result) != 2 || result[0] != 1 || result[1] != 0 {
		t.Errorf("Expected result %v but got %v", []byte{1, 0}, result)
	}
}

func TestString_(t *testing.T) {
	inputArray := DigitArray{2, 1, 5}
	result := inputArray.String()
	if len(result) != 3 || result != "215" {
		t.Errorf("Expected result %v but got %v", "215", result)
	}
}

func TestCollectOverflowFromLeadDigit_InitialCondition_ReturnsFalse(t *testing.T) {
	inputArray := DigitArray{3, 9, 4}
	if inputArray.IsLeadDigitOverflowing() {
		t.Error("Expected No Overflow")
	}
	result := inputArray.CollectOverflowFromLeadDigit()
	if result != 0 {
		t.Errorf("Expected 0, but got %v", result)
	}
}

func TestCollectOverflowFromLeadDigit_Add01_ReturnsFalse(t *testing.T) {
	inputArray := DigitArray{3, 9, 4}
	operand2 := DigitArray{0, 1, 0}
	sum := inputArray.Plus(operand2)
	if sum.IsLeadDigitOverflowing() {
		t.Error("Expected No Overflow")
	}
	result := sum.CollectOverflowFromLeadDigit()
	if result != 0 {
		t.Errorf("Expected 0, but got %v", result)
	}
}

func TestCollectOverflowFromLeadDigit_Add99_ReturnsTrue(t *testing.T) {
	inputArray := DigitArray{3, 9, 4}
	operand2 := DigitArray{9, 9}
	sum := inputArray.Plus(operand2)
	if !sum.IsLeadDigitOverflowing() {
		t.Error("Expected Lead Digit Overflow")
	}
	result := sum.CollectOverflowFromLeadDigit()
	if result != 1 {
		t.Errorf("Expected 1, but got %v", result)
	}
}

func TestCollectOverflowFromLeadDigit_Minus01_ReturnsFalse(t *testing.T) {
	inputArray := DigitArray{3, 9, 4}
	operand2 := DigitArray{0, 1, 0}
	sum := inputArray.Minus(operand2)
	if sum.IsLeadDigitOverflowing() {
		t.Error("Expected No Overflow")
	}
	result := sum.CollectOverflowFromLeadDigit()
	if result != 0 {
		t.Errorf("Expected 0, but got %v", result)
	}
}

func TestCollectOverflowFromLeadDigit_Minus99_ReturnsTrue(t *testing.T) {
	inputArray := DigitArray{3, 9, 4}
	operand2 := DigitArray{9, 9}
	sum := inputArray.Minus(operand2)
	if !sum.IsLeadDigitOverflowing() {
		t.Error("Expected Lead Digit Overflow")
	}
	result := sum.CollectOverflowFromLeadDigit()
	if result != -1 {
		t.Errorf("Expected -1, but got %v", result)
	}
}
