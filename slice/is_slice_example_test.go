//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:47 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

import "fmt"

func ExampleIsSlice() {
	slice1 := []int{1, 2, 3}
	slice2 := make(chan int)

	res1 := IsSlice(slice1)
	res2 := IsSlice(slice2)

	fmt.Println(res1)
	fmt.Println(res2)

	// Output:
	// true
	// false
}
