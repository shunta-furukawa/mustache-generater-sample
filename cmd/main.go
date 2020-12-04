package main

import (
	"fmt"
	"strings"

	"github.com/shunta-furukawa/mustache-generater-sample/app/gift"
)

func main() {

	// サンタクロースの袋の中
	sack := make([]Gift, 0)

	// サンタクロースの袋の中身を詰める
	sack = append(sack, gift.NewSportsCar())

	// 子供たち
	kids := []Kid{
		kid.Taro,
	}

	// 子供たちへのプレゼント を 表示
	giftlist = make([]string, 0)
	for _, gift := range sack {
		giftlist = append(giftlist, gift.Name())
	}

	fmt.Printf("====【:*･ﾟ☆†　Merry Ⅹ’mas　†.｡.:*･ﾟ】===	\nよういした プレゼント : \n%s\n\n", strings.Join(giftlist, "\n"))

	for _, kid = range kids {
		fmt.Printf("==== %s ====\nほしいもの : \n%s\nもらえるおもちゃ: \n%s \n",
			kid.Name(),
			kid.Wishlist(),
			kid.CanGet(sack),
		)
	}

}
