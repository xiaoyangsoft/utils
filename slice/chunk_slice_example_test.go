//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:44 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

import "fmt"

func ExampleChunkSlice() {
	slice := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}
	res := ChunkSlice(slice, 3)

	fmt.Printf("%+v", res)

	// Output:
	// [[a b c] [d e f] [g h i] [j k]]
}
