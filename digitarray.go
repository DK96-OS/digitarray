package digitarray

import (
    "fmt"
)

// Manages a slice of Digits.
type DigitArray []int8

// The Size of the Slice.
func (d DigitArray) Size() int {
    return len(d)
}

// Get a value at an index in the DigitArray.
func (d DigitArray) Get(index int) (int8, error) {
    if index < 0 || index >= len(d) {
        panic("Index out of bounds")
    }
    return d[index], nil
}

// String representation of DigitArray.
func (d DigitArray) String() string {
    result := ""
    for _, v := range d {
        result += fmt.Sprintf("%v", v)
    }
    return result

}

// Checks if the two DigitArrays are equal.
func (d DigitArray) Equal(other DigitArray) bool {
    if other == nil || len(d) != len(other) {
        return false
    }
    for i := range d {
        if d[i] != other[i] {
            return false
        }
    }
    return true
}

// Return a hash code for the DigitArray.
func (d DigitArray) Hash() int {
    h := 0
    for _, v := range d {
        h = h*31 + int(v)
    }
    return h
}

// Add two DigitArrays together.
func (d DigitArray) Plus(other DigitArray) DigitArray {
    var thisSize int = len(d)
    var otherSize int = len(other)
    newSize := max(thisSize, otherSize)
    result := make([]int8, newSize)
    if thisSize == newSize {
        if newSize == otherSize {
            for i, v := range d {
                result[i] = v + other[i]
            }
        } else {
            for i, v := range d {
                if i < otherSize {
                    result[i] = (v + other[i])
                } else {
                    result[i] = v
                }
            }
        }
    } else {
        for i, v := range other {
            if i < thisSize {
                result[i] = (v + d[i])
            } else {
                result[i] = v
            }
        }
    }
    for i := newSize - 1; i > 0; i-- {
        digitValue := result[i]
        if digitValue > 9 {
            result[i-1] = (result[i - 1] + digitValue / 10)
            result[i] = (digitValue % 10)
        }
    }
    return DigitArray(result)
}

// Check if the LeadDigit is overflowing.
func (d DigitArray) IsLeadDigitOverflowing() bool {
    return d[0] > 9 || d[0] < 0
}

// Obtain the Overflow Value from the LeadDigit.
func (d DigitArray) CollectOverflowFromLeadDigit() int8 {
    leadDigit := d[0]
    if leadDigit < 0 { // LeadDigit is Negative
        d[0] = (10 - ((-leadDigit) % 10))
        return (leadDigit / 10)
    }
    d[0] = (leadDigit % 10)
    return (leadDigit / 10)
}

// Subtract two DigitArrays.
func (d DigitArray) Minus(other DigitArray) DigitArray {
    var thisSize int = len(d)
    var otherSize int = len(other)
    newSize := max(thisSize, otherSize)
    result := make([]int8, newSize)
    // Copy this to result
    copy(result[:thisSize], d)
    for i := newSize - 1; i >= 0; i-- {
        var otherDigitValue int8
        if i < otherSize {
            otherDigitValue = other[i]
        } else {
            otherDigitValue = 0
        }
        if otherDigitValue < 0 || otherDigitValue > 9 {
            panic("Other DigitArray contains non-normalized digits.")
        }
        if otherDigitValue == 0 {
            continue
        }
        diff := result[i] - otherDigitValue
        if diff < 0 {
            borrowIndex := findBorrowableIndex(i-1, result)
            if borrowIndex > -1 {
                result[borrowIndex]--
                for j := borrowIndex + 1; j < i; j++ {
                    result[j] = 9
                }
                result[i] = diff + 10
            } else if i == 0 {
                result[i] = diff - 10
            } else {
                result[0] = -11
                for j := 1; j < i; j++ {
                    result[j] = 9
                }
                result[i] = diff + 10
            }
        } else {
            result[i] = diff
        }
    }
    return DigitArray(result)
}

// Find an index that can be borrowed from.
func (d DigitArray) FindBorrowableIndex(startIndex int) int {
    for i := startIndex; i >= 0; i-- {
        if d[i] > 0 {
            return i
        }
    }
    return -1
}

// Remove Trailing Zeros from the DigitArray.
func (d DigitArray) TrimTrailingZeros() DigitArray {
    var i int = len(d) - 1
    for ; i >= 0 && d[i] == int8(0); i-- {
    }
    return DigitArray(d[:i+1])
}

// Remove Leading Zeros from the DigitArray.
func (d DigitArray) TrimLeadingZeros() DigitArray {
    var i int = 0
    for ; i < len(d) && d[i] == int8(0); i++ {
    }
    return DigitArray(d[i:])
}

// Utility function to find the largest index below startIndex where value is greater than zero.
func findBorrowableIndex(startIndex int, digitArray []int8) int {
    for i := startIndex; i >= 0; i-- {
        if digitArray[i] > 0 {
            return i
        }
    }
    return -1
}
