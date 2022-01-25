## 回溯_棋盘
n 皇后问题 研究的是如何将 n 个皇后放置在 n×n 的棋盘上，并且使皇后彼此之间不能相互攻击。
给你一个整数 n ，返回所有不同的 n 皇后问题 的解决方案

  . Q . .   -->   . . Q .
  . . . Q   -->   Q . . .
  Q . . .   -->   . . . Q
  . . Q .   -->   . Q . .

输入：n = 4
输出：[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]
解释：如上图所示，4 皇后问题存在两个不同的解法。
即Q所在的 行、列、左上、右上、左下、右下 的位置只能有1个Q

var res [][]string // 全局变量 结果集
func solveNQueens(n int) [][]string {
	res = [][]string{}
	board := make([][]string, n)
	for i := 0; i < n; i++ { // 初始化 每行
		board[i] = make([]string, n)
	}
	for i := 0; i < n; i++ { // 给二维棋盘 每个元素都赋值为 "."
		for j := 0; j < n; j++ {
			board[i][j] = "."
		}
	}
	backTrack(board, 0)
	return res
}
func backTrack(board [][]string, row int) {
	size := len(board)
	if row == size { // 终止条件
		temp := make([]string, size)
		for i := 0; i < size; i++ {
			temp[i] = strings.Join(board[i], "") // 把board第i行数组拼接后给temp[i]
		}
		res = append(res, temp)
		return
	}
	for col := 0; col < size; col++ { // 遍历当前行的 每个元素
		if !isValid(board, row, col) { // 不合法就跳过
			continue
		}
		board[row][col] = "Q"   // 放置皇后
		backTrack(board, row+1) // 递归下一行
		board[row][col] = "."   // 回溯 把皇后撤销
	}
}
func isValid(board [][]string, row, col int) (res bool) { // 检测是否合法
	n := len(board)
	for i := 0; i < row; i++ { // 检查每行
		if board[i][col] == "Q" { // 如该行有元素为皇后 返回false
			return false
		}
	}
	for i := 0; i < n; i++ { // 检查每列
		if board[row][i] == "Q" { // 如该列有元素为皇后 返回false
			return false
		}
	}
	// 检测45度角
	for i, j := row, col; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	// 检测135度角
	for i, j := row, col; i >= 0 && j < n; i, j = i-1, j+1 {
		if board[i][j] == "Q" {
			return false
		}
	}
	return true
}
