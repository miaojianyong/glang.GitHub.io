## 数组
描述：即输入一个整数n，生成一个n*n的二维数组，其元素是 1-n*n，且为螺旋顺序
输入：n = 3
输出：[[1,2,3],[8,9,4],[7,6,5]]
即：先上行 左到右、再左列 上到下、在下行 右到左、再右列 下到上 如此螺旋顺序

func generateMatrix(n int) [][]int {
	top, bottom := 0, n-1
	left, right := 0, n-1
	num := 1                   // 填充数 从1开始
	matrix := make([][]int, n) // 二维矩阵
	for i := 0; i < n; i++ {   // 初始化
		matrix[i] = make([]int, n)
	}
	for num <= n*n { // 填充的数 小于等于 n的平方
		for i := left; i <= right; i++ { // 左到右
			matrix[top][i] = num // 填充上行
			num++
		}
		top++                            // 从上 移动到 下一行
		for i := top; i <= bottom; i++ { // 上到下
			matrix[i][right] = num // 填充右列
			num++
		}
		right--                          // 从右 移动到 左一列
		for i := right; i >= left; i-- { // 右到左 倒序
			matrix[bottom][i] = num // 填充下行
			num++
		}
		bottom--                         // 从下 移动到 上一行
		for i := bottom; i >= top; i-- { // 下到上 倒序
			matrix[i][left] = num // 填充左列
			num++
		}
		left++ // 从左 移动到 右一列 继续上述循环
	}
	return matrix
}

func main() {
	fmt.Println(generateMatrix(3)) // [[1 2 3] [8 9 4] [7 6 5]]
}
