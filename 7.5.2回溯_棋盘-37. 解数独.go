## 回溯_棋盘

编写一个程序，通过填充空格来解决数独问题。
数独的解法需 遵循如下规则：
数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次

package main

func solveSudoku(board [][]byte) {
	var backTrack func() bool
	backTrack = func() bool {
		for i := 0; i < 9; i++ { // 遍历二维数组
			for j := 0; j < 9; j++ {
				if board[i][j] != '.' { // 如当前元素已填数字就跳过
					continue
				}
				// 否则 尝试填1-9
				for k := '1'; k <= '9'; k++ {
					// 简单该数字 是否可以填
					if isvalid(i, j, byte(k), board) {
						board[i][j] = byte(k)
						if backTrack() { // 递归
							return true
						}
						board[i][j] = '.' // 回溯
					}
				}
				return false // 如上述的1-9都无法填返回 false
			}
		}
		return true // 二维数组遍历完了 说明都填上数了 返回true
	}
	backTrack()
}

func isvalid(row, col int, k byte, board [][]byte) bool {
	for i := 0; i < 9; i++ { // 检测行
		if board[row][i] == k { // 如该行有元素 等于k 返回false
			return false
		}
	}
	for i := 0; i < 9; i++ { // 检测列
		if board[i][col] == k { // 如该列有元素 等于k 返回false
			return false
		}
	}
	// 检查 3*3区域 是否有元素 等于k
	startRow := row / 3 * 3
	startCol := col / 3 * 3
	for i := startRow; i < startRow+3; i++ { // 每3行
		for j := startCol; j < startCol+3; j++ { // 每3列
			if board[i][j] == k { // 如该区域有元素 等于k 返回false
				return false
			}
		}
	}
	return true // 否则返回true
}
