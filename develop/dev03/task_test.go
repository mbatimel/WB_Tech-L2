package main

import (
	"reflect"
	"testing"
)

func TestSortText(t *testing.T) {
	sf := NewSortFile("text.txt", 1, false, false, false, false, false, false, false)
	err := sf.TextFromFile()
	if err != nil {
		t.Errorf("Error reading file: %v", err)
	}

	sorted, err := sf.SortText()
	if err != nil {
		t.Errorf("Error sorting text: %v", err)
	}

	expected := []string{"            JOP 5 orange", "apple 10 red", "apple 10 red", "banana 8 yellow", "grape 15 purple"}
	if !reflect.DeepEqual(sorted, expected) {
		t.Errorf("SortText() = %v, want %v", sorted, expected)
	}
}

func TestMakeUnique(t *testing.T) {
	strs := []string{"apple", "banana", "cherry", "banana", "date", "apple"}
	unique := MakeUnique(strs)
	expected := []string{"apple", "banana", "cherry", "date"}
	if !reflect.DeepEqual(unique, expected) {
		t.Errorf("MakeUnique() = %v, want %v", unique, expected)
	}
}

func TestExtractNumericSuffix(t *testing.T) {
	num, suff := extractNumericSuffix("test123")
	if num != 123 || suff != "test" {
		t.Errorf("ExtractNumericSuffix() = %d, %s, want 123, test", num, suff)
	}

	num, suff = extractNumericSuffix("123test")
	if num != 123 || suff != "test" {
		t.Errorf("ExtractNumericSuffix() = %d, %s, want 123, test", num, suff)
	}
}

func TestIsSorted(t *testing.T) {
	sorted := []string{"apple", "banana", "cherry", "date"}
	if !isSorted(sorted) {
		t.Errorf("IsSorted() = false, want true")
	}

	unsorted := []string{"banana", "apple", "date", "cherry"}
	if isSorted(unsorted) {
		t.Errorf("IsSorted() = true, want false")
	}
}
