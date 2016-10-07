package viewmodels

import ()

type Categories struct {
	Title      string
	Active     string
	Categories []Category
}

type Category struct {
	Id            int
	Title         string
	ImageUrl      string
	Description   string
	IsOrientRight bool
}

func GetCategories() Categories {
	result := Categories{
		Title:  "Lemonade Stand Society - Shop",
		Active: "shop",
	}

	return result
}
