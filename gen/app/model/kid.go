package model

import "strings"

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
		Preferences  []Preference `yaml:"preferences"`
		GenderBiased bool         `yaml:"genderBiased"`
	}

	// Preference represents preference
	Preference struct {
		// Attribute ほしいものの特徴軸
		Attribute string `yaml:"attribute"`
		// Attribute ほしいものの特徴
		Value string `yaml:"value"`
	}
)

/* 構造体に引数を取らない関数を定義することで、
 * Mustache から 計算フィールドとして呼び出すことができる
 * うまく活用することで、Mustacheの記述を完結にすることができる。
 */

// Male returns whether kid is male or not.
func (k Kid) Male() bool {
	return k.Gender == "おとこのこ"
}

// Wishlist is generated sentence from preferences.
func (k Kid) Wishlist() string {
	var postfix = " がほしい"

	if len(k.Preferences) == 0 {
		postfix = "なんでもいい。"
	}

	if k.GenderBiased {
		postfix = postfix + "\nでも " + k.Gender + "向けじゃなきゃいやだ。"
	}

	colors := make([]string, 0)
	categories := make([]string, 0)
	genders := make([]string, 0)
	names := make([]string, 0)

	for _, p := range k.Preferences {
		switch p.Attribute {
		case "Color":
			colors = append(colors, p.Value)
		case "Gender":
			genders = append(genders, p.Value+"向け")
		case "Category":
			categories = append(categories, p.Value)
		case "Name":
			names = append(names, p.Value)
		}
	}

	ret := ""
	if len(names) > 0 {
		ret = ret + strings.Join(names, "/")
		return ret + postfix
	}

	if len(genders) > 0 {
		ret = ret + strings.Join(genders, "/")
		if len(colors) > 0 {
			ret = ret + " で "
		} else {
			ret = ret + " の "
		}

	}
	if len(colors) > 0 {
		ret = ret + strings.Join(colors, "/") + " の "
	}
	if len(categories) > 0 {
		ret = ret + strings.Join(categories, "/")
	} else {
		ret = ret + "もの"
	}
	return ret + postfix

}
