package gift

type Sportcar struct {
	Name     string
	Category string
	Color    string
}

func NewSportsCar() Gift {
	return Gift{
		Name:     "スポーツカー",
		Category: "のりもの",
		Color:    "あか",
	}
}
