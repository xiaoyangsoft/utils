//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:45 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

import "fmt"

func ExampleColumnSlice() {
	type Person struct {
		Name string
		Age  int
	}
	people := []Person{
		{"Alice", 18},
		{"Bob", 20},
		{"Charlie", 22},
	}

	res := ColumnSlice(people, "Age")

	fmt.Printf("%+v", res)

	// Output:
	// [18 20 22]
}
