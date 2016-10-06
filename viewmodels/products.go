package viewmodels

import ()

type Products struct {
	Title    string
	Active   string
	Products []Product
}

type Product struct {
	ImageUrl         string
	Name             string
	PricePerLiter    float32
	DescriptionShort string
	DescriptionLong  string
	PricePer10Liter  float32
	Origin           string
	IsOrganic        bool
	Id               int
}

type ProductVM struct {
	Title   string
	Active  string
	Product Product
}

func GetProduct(id int) ProductVM {
	var result ProductVM
	var product Product
	productList := getProductList()
	for _, p := range productList {
		if id == p.Id {
			product = p
			break
		}
	}

	result.Active = "shop"
	result.Title = "Lemonade Stand Society - " + product.Name
	result.Product = product
	return result
}

func GetProducts(id int) Products {
	var shopName string
	result := Products{
		Active: "shop",
	}

	switch id {
	case 1:
		shopName = "Juice"
	case 2:
		shopName = "Supply"
	case 3:
		shopName = "Advertising"
	}

	result.Title = "Lemonade Stand Society - " + shopName + "shop"

	if id == 1 {
		result.Products = getProductList()
	}

	return result

}

func getProductList() []Product {
	lemonJuice := MakeLemonJuiceProduct()
	appleJuice := MakeAppleJuiceProduct()
	watermelonJuice := MakeWatermelonJuiceProduct()
	kiwiJuice := MakeKiwiJuiceProduct()
	mangosteenJuice := MakeMangosteenJuiceProduct()
	orangeJuice := MakeOrangeJuiceProduct()
	pineappleJuice := MakePineappleJuiceProduct()
	strawberryJuice := MakeStrawberryJuiceProduct()

	result := []Product{
		lemonJuice,
		appleJuice,
		watermelonJuice,
		kiwiJuice,
		mangosteenJuice,
		orangeJuice,
		pineappleJuice,
		strawberryJuice,
	}
	return result
}

func MakeLemonJuiceProduct() Product {
	result := Product{
		ImageUrl:         "lemon.png",
		Name:             "Lemon Juice",
		PricePerLiter:    1.09,
		DescriptionShort: "Made from fresh, organic California lemons.",
		DescriptionLong:  `Made from premium, organic Meyer lemons. These fruit are left on the tree until they reach the peak of ripeness and then juiced within 8 hours of being picked.<br/>Lemonade made from our premium juice is sure to make your stand the most popular in the neighborhood.`,
		PricePer10Liter:  1.04,
		Origin:           "California",
		IsOrganic:        true,
		Id:               1,
	}
	return result
}

func MakeAppleJuiceProduct() Product {
	result := Product{
		ImageUrl:         "apple.png",
		Name:             "Apple Juice",
		PricePerLiter:    0.89,
		PricePer10Liter:  0.85,
		DescriptionShort: "The perfect blend of Washington apples.",
		DescriptionLong:  "The perfect blend of Washington apples.",
		Origin:           "Ohio",
		IsOrganic:        true,
		Id:               2,
	}
	return result
}

func MakeWatermelonJuiceProduct() Product {
	result := Product{
		ImageUrl:         "watermelon.png",
		Name:             "Watermelon Juice",
		PricePerLiter:    3.99,
		PricePer10Liter:  3.84,
		DescriptionShort: "From sun-drenched fields in Florida.",
		DescriptionLong:  "From sun-drenched fields in Florida.",
		Origin:           "Florida",
		IsOrganic:        true,
		Id:               3,
	}
	return result
}

func MakeKiwiJuiceProduct() Product {
	result := Product{
		ImageUrl:         "kiwi.png",
		Name:             "Kiwi Juice",
		PricePerLiter:    5.99,
		PricePer10Liter:  5.54,
		DescriptionShort: "California sunshine and rain distilled into sweet goodness",
		DescriptionLong:  "California sunshine and rain distilled into sweet goodness",
		Origin:           "California",
		IsOrganic:        false,
		Id:               4,
	}
	return result
}

func MakeMangosteenJuiceProduct() Product {
	result := Product{
		ImageUrl:         "mangosteen.png",
		Name:             "Mangosteen Juice",
		PricePerLiter:    6.87,
		PricePer10Liter:  6.79,
		DescriptionShort: "Our latest taste sensation, imported directly from Hawaii",
		DescriptionLong:  "Our latest taste sensation, imported directly from Hawaii",
		Origin:           "Hawaii",
		IsOrganic:        false,
		Id:               5,
	}
	return result
}

func MakeOrangeJuiceProduct() Product {
	result := Product{
		ImageUrl:         "orange.png",
		Name:             "Orange Juice",
		PricePerLiter:    1.67,
		PricePer10Liter:  1.63,
		DescriptionShort: "Fresh squeezed from Florida's best oranges.",
		DescriptionLong:  "Fresh squeezed from Florida's best oranges.",
		Origin:           "florida",
		IsOrganic:        false,
		Id:               6,
	}
	return result
}

func MakePineappleJuiceProduct() Product {
	result := Product{
		ImageUrl:         "pineapple.png",
		Name:             "Pineapple Juice",
		PricePerLiter:    2.48,
		PricePer10Liter:  2.27,
		DescriptionShort: "An exotic and refreshing offering. Straight from Hawaii.",
		DescriptionLong:  "An exotic and refreshing offering. Straight from Hawaii.",
		Origin:           "Hawaii",
		IsOrganic:        false,
		Id:               7,
	}
	return result
}

func MakeStrawberryJuiceProduct() Product {
	result := Product{
		ImageUrl:         "strawberry.png",
		Name:             "Strawberry Juice",
		PricePerLiter:    4.36,
		PricePer10Liter:  4.27,
		DescriptionShort: "The perfect balance of sweet and tart.",
		DescriptionLong:  "The perfect balance of sweet and tart.",
		Origin:           "California",
		IsOrganic:        false,
		Id:               8,
	}
	return result
}
