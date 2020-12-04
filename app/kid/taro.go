package kid

import (
	"strings"

	"../gift"
)

// Taro represents Taro.
type Taro struct {
	Name   string
	Gender string
	Age    int
}

// NewTaro returns instance of Kids Impl as Taro
func NewTaro() Taro {
	return Taro{
		Name:   "たろう",
		Gender: "おとこのこ",
		Age:    4,
	}
}

// Display returns name of the kid.
func (k Taro) Display() string {
	return "☆★☆★ たろうくん (4) ★☆★☆"
}

// Wishlist returns the kid's wishlist.
func (k Taro) Wishlist() string {
	return "あか" + "の" + "のりもの" + "がほしい"
}

// CanGet returns gift the kid can get.
func (k Taro) CanGet(sack []gift.Gift) string {
	gds := make([]string, 0)

	for _, gift := range sack {
		if gift.GetColor() == "あか" && gift.GetCategory() == "のりもの" {
			gds = append(gds, gift.Display())
		}
	}

	if len(gds) == 0 {
		return "ほしいものがみつからなかった..."
	}
	return strings.Join(gds, "\n  ")
}
