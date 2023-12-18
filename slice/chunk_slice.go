//============================================================

// 作者: 杨大雷
// 日期: 2023/12/02 11:43 AM
// 版权: 济南晓杨信息科技有限公司 @Since 2022
//
//============================================================

package slice

// ChunkSlice 把slice分割为新的数组块
func ChunkSlice[T any](slice []T, size int) [][]T {
	var chunks [][]T
	for i := 0; i < len(slice); i += size {
		end := i + size
		if end > len(slice) {
			end = len(slice)
		}
		chunks = append(chunks, slice[i:end])
	}
	return chunks
}
