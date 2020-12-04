package model

type (
	// KidContainer wrap interfaces
	KidContainer struct {
		Kids []Kid `yaml:"kids"`
	}

	// Kid represents kid
	Kid struct {
		Name         string       `yaml:"name"`
		JName        string       `yaml:"jname"`
		Gender       string       `yaml:"gender"`
		Age          int          `yaml:"age"`
		Preferences  []Preference `yaml:"preference"`
		GenderBiased bool         `yaml:"genderBiased"`
	}

	Preference struct {
		// Attribute ほしいものの特徴軸
		Attribute string `yaml:"attribute"`
		// Attribute ほしいものの特徴
		Value string `yaml:"value"`
	}
)
