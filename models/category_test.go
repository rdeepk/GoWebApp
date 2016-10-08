package models

import (
	"testing"
)

func Test_ReturnsNonEmptyCategorySlice(t *testing.T) {
	categories := GetCategories()
	if len(categories) == 0 {
		t.Log("Empty slice returned")
		t.Fail()
	}
}
