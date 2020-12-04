package gift

// Gift represents gift
type Gift interface {
	// Display returns name of the gift.
	Display() string
	// GetName returns its name.
	GetName() string
	// GetCategory returns its category.
	GetCategory() string
	// GetGender returns its gender.
	GetGender() string
	// GetColor returns its color.
	GetColor() string
}
