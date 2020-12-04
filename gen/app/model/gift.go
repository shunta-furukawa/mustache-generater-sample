package model

type (
	// GiftContainer wrap interfaces
	GiftContainer struct {
		Gifts []Gift `yaml:"gifts"`
	}

	// Gift represents kid
	Gift struct {
		Name     string `yaml:"name"`
		JName    string `yaml:"jname"`
		Category string `yaml:"category"`
		Color    string `yaml:"color"`
		Gender   string `yaml:"gender"`
	}
)
