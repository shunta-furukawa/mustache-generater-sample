package kid

import "../gift"

// Kid represents kid.
type Kid interface {
	// Name returns name of the kid.
	Display() string
	// Wishlist returns the kid's wishlist.
	Wishlist() string
	// CanGet returns gift the kid can get.
	CanGet(sack []gift.Gift) string
}
