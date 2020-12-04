package main

import (
	"errors"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"

	"./app/model"
	"github.com/cbroglie/mustache"
	"gopkg.in/yaml.v2"
)

// Renderer represents renderer
type Renderer struct {
	Tmpl    *mustache.Template
	Path    string
	Postfix string
}

var (

	// mustache拡張子
	// ベース
	basePath = "./"

	// データ
	inputBasePath  = basePath + "variables/"
	kidsInputPath  = inputBasePath + "kids.yml"
	giftsInputPath = inputBasePath + "gifts.yml"

	// テンプレ
	templateBasePath = basePath + "templates/"

	kidTemplatePath       = templateBasePath + "kid.mustache"
	giftTemplatePath      = templateBasePath + "gift.mustache"
	christmasTemplatePath = templateBasePath + "christmas.mustache"

	// 出力
	//outputBasePath                  = basePath + "out/"
	outputBasePath = "../app/"

	kidOutputPath  = outputBasePath + "kid/"
	giftOutputPath = outputBasePath + "gift/"

	// main 出力用
	christmasOutputPath = "../cmd/"
)

//go:generate go run main.go
func main() {

	// Kid 生成
	kidBuf, err := ioutil.ReadFile(kidsInputPath)
	if !errors.Is(err, nil) {
		panic(err)
	}
	kidContainer := model.KidContainer{}
	kidContainer.Kids = make([]model.Kid, 0)
	err = yaml.Unmarshal(kidBuf, &kidContainer)
	if !errors.Is(err, nil) {
		panic(err)
	}
	generateKids(kidContainer.Kids)

	// Gift 生成
	giftBuf, err := ioutil.ReadFile(giftsInputPath)
	if !errors.Is(err, nil) {
		panic(err)
	}
	giftContainer := model.GiftContainer{}
	giftContainer.Gifts = make([]model.Gift, 0)
	err = yaml.Unmarshal(giftBuf, &giftContainer)
	if !errors.Is(err, nil) {
		panic(err)
	}
	generateGifts(giftContainer.Gifts)

	// main.go の生成
	christmas := model.Christmas{}
	christmas.Kids = kidContainer.Kids
	christmas.Gifts = giftContainer.Gifts

	generateChristmas(christmas)
}

func generateKids(kids []model.Kid) {

	// テンプレートの読み込み
	kidTemplate, err := mustache.ParseFile(kidTemplatePath)
	if !errors.Is(err, nil) {
		panic(err)
	}
	// Kids からの 書き出し
	for _, k := range kids {
		// テンプレートが増えた時に、ここの要素を増やすと拡張できる。
		for _, r := range []Renderer{
			{
				Tmpl: kidTemplate,
				Path: kidOutputPath,
			},
		} {
			outputFile(r, k.Name, k)
		}
	}

}

func generateGifts(gifts []model.Gift) {

	// テンプレートの読み込み
	giftTemplate, err := mustache.ParseFile(giftTemplatePath)
	if !errors.Is(err, nil) {
		panic(err)
	}
	// Gifts からの 書き出し
	for _, p := range gifts {
		// テンプレートが増えた時に、ここの要素を増やすと拡張できる。
		for _, r := range []Renderer{
			{
				Tmpl: giftTemplate,
				Path: giftOutputPath,
			},
		} {
			outputFile(r, p.Name, p)
		}
	}

}

func generateChristmas(christmas model.Christmas) {

	// テンプレートの読み込み
	christmasTemplate, err := mustache.ParseFile(christmasTemplatePath)
	if !errors.Is(err, nil) {
		panic(err)
	}

	r := Renderer{
		Tmpl: christmasTemplate,
		Path: christmasOutputPath,
	}
	outputFile(r, "main", christmas)

}

func outputFile(r Renderer, name string, data interface{}) {
	output, err := r.Tmpl.Render(data)
	if !errors.Is(err, nil) {
		panic(err)
	}

	outputBytes, err := format.Source([]byte(output))
	if !errors.Is(err, nil) {
		panic(err)
	}
	// outputBytes := []byte(output)

	// ディレクトリを作成、存在する場合は無視する。
	_ = os.MkdirAll(r.Path, 0755)

	// []byte をファイルに上書きしています。
	filename := r.Path + strings.ToLower(name) + r.Postfix + ".go"
	err = ioutil.WriteFile(filename, outputBytes, 0755)
	if err != nil {
		panic(err)
	}

	fmt.Printf("mustache: generate %s\n", filename)

}
