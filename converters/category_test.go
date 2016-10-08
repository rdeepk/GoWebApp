package converters

import (
	"models"
	"testing"
)

func Test_ConvertsCategoryToViewModel(t *testing.T) {
	category := models.Category{}
	category.SetImageUrl("Image Url")
	category.SetDescription("this is decsription")
	category.SetId(43)
	category.SetIsOrientRight(true)
	category.SetTitle("the title")

	result := ConvertCategoryToViewModel(category)

	if result.ImageUrl != category.ImageUrl() {
		t.Log("Image Url not converted properly")
		t.Fail()
	}

	if result.Title != category.Title() {
		t.Log("Title not converted properly")
		t.Fail()
	}

	if result.Description != category.Description() {
		t.Log("Description not converted properly")
		t.Fail()
	}

	if result.Id != category.Id() {
		t.Log("Id not converted properly")
		t.Fail()
	}

	if result.IsOrientRight != category.IsOrientRight() {
		t.Log("IsOrientRight not converted properly")
		t.Fail()
	}

}
