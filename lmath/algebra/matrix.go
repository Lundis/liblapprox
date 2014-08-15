package algebra

type Matrix [][]float64

func NewMatrix(rows, cols int) Matrix {
	matrix := make([][]float64, rows)
	for i := range matrix {
		matrix[i] = make([]float64, cols)
	}
	return matrix
}


func (self Matrix) Solve() {
	if len(self) != len(self[0]) - 1 {
		panic("Matrix.Solve(): rows must equal columns-1")
	}
	for col := 0; col < len(self[0]) - 1; col++ {
		self.ensureDiagonal1(col)
		self.reduceAllBut(col)
	}
}

// Make sure that the leading number of row col is 1, by dividing it with itself.
// If it is zero, swap it with one further down, then do it again
func (self Matrix) ensureDiagonal1(col int) {
	if self[col][col] != 0 {
		div := self[col][col]
		for i := col; i < len(self[col]); i++ {
			self[col][i] /= div
		}
	} else {
		for i := col; i < len(self); i++ {
			if self[i][col] != 0 {
				self.swapRows(i, col)
				self.ensureDiagonal1(col)
			}
		}
	}
}

func (self Matrix) swapRows(r1, r2 int) {
	tmp := self[r1]
	self[r1] = self[r2]
	self[r2] = tmp
}

// subtracts rows with each other so that everything in column col but matrix[col][col] is 0
func (self Matrix) reduceAllBut(col int) {
	for row := 0; row < len(self); row++ {
		if row == col {
			continue
		}
		// calculate the multiplier
		mult := self[row][col]
		self[row][col] = 0
		for i := col + 1; i < len(self[row]); i++ {
			self[row][i] -= mult*self[col][i]
		}
	}
}