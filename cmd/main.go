package main

import (
	"fmt"
	"strings"

	"../app/gift"
	"../app/kid"
)

func main() {

	// サンタクロースの袋の中
	sack := make([]gift.Gift, 0)

	// サンタクロースの袋の中身を詰める
	sack = append(sack, gift.NewSportsCar())

	// 子供たち
	kids := make([]kid.Kid, 0)

	// サンタクロースの袋の中身を詰める
	kids = append(kids, kid.NewTaro())

	// 子供たちへのプレゼント を 表示
	giftlist := make([]string, 0)
	for _, gift := range sack {
		giftlist = append(giftlist, gift.Display())
	}

	fmt.Printf("=======================================\n===【:*･ﾟ☆†　Merry Ⅹ’mas　†.｡.:*･ﾟ】=== \n=======================================\n\nよういした プレゼント : \n  - %s\n\n", strings.Join(giftlist, "\n  - "))

	for _, kid := range kids {
		fmt.Printf("%s\nほしいもの : \n  %s\nもらえるおもちゃ: \n  %s \n",
			kid.Display(),
			kid.Wishlist(),
			kid.CanGet(sack),
		)
	}

	fmt.Printf("\n 良いお年を... Fin.\n")

}
