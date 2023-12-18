//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 4:45 PM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package math

import "fmt"

func ExampleRand() {
	minInt := 1
	maxInt := 100
	res1 := Rand(minInt, maxInt)
	if res1 >= 1 && res1 < 100 {
		fmt.Println("ok")
	}

	minFloat := 1.0
	maxFloat := 10.0
	res2 := Rand(minFloat, maxFloat)
	if res2 >= 1.0 && res2 < 10.0 {
		fmt.Println("ok")
	}

	// Output:
	// ok
	// ok
}
