package kid

// Kid represents kid.
type Kid interface {
	// Name returns name of the kid.
	func Name() string,
	// Wishlist returns the kid's wishlist.
	func Wishlist() string ,
	// CanGet returns gift the kid can get.
	func CanGet(sack []Gift) string,
}
