package gift

import "fmt"

// SportsCar represents SportCar.
type SportsCar struct {
	Name     string
	Category string
	Color    string
	Gender   string
}

// NewSportsCar returns new SportCar.
func NewSportsCar() SportsCar {
	return SportsCar{
		Name:     "スポーツカー",
		Category: "のりもの",
		Color:    "あか",
		Gender:   "おとこのこ",
	}
}

// Display returns spec of SportCar.
func (g SportsCar) Display() string {
	return fmt.Sprintf(`%s【%s|%s|%s向け】`,
		g.Name,
		g.Category,
		g.Color,
		g.Gender,
	)
}

// GetName returns its name.
func (g SportsCar) GetName() string {
	return g.Name
}

// GetCategory returns its category.
func (g SportsCar) GetCategory() string {
	return g.Category
}

// GetGender returns its gender.
func (g SportsCar) GetGender() string {
	return g.Gender
}

// GetColor returns its color.
func (g SportsCar) GetColor() string {
	return g.Color
}
