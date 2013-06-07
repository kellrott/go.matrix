
package matrix

import (

)

/*
A matrix backed by a flat array of all elements.
*/
type DenseNamedMatrix struct {
	DenseMatrix
	namedMatrix
}

func NamedZeros(rowNames, colNames []string) *DenseNamedMatrix {
	rows := len(rowNames)
	cols := len(colNames)
	A := new(DenseNamedMatrix)
	A.elements = make([]float64, rows*cols)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.rowNames = make(map[string] int)
	A.colNames = make(map[string] int)
	for i := 0; i < rows; i++ {
		A.rowNames[rowNames[i]] = i
	}
	for i := 0; i < cols; i++ {
		A.colNames[colNames[i]] = i
	}
	return A
}

func NamedZerosFromMap(rowNames, colNames map[string] int) *DenseNamedMatrix {
	rows := len(rowNames)
	cols := len(colNames)
	A := new(DenseNamedMatrix)
	A.elements = make([]float64, rows*cols)
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.rowNames = rowNames
	A.colNames = colNames
	return A
}

func NamedFromDense(matrix *DenseMatrix, rowNames, colNames map[string] int) *DenseNamedMatrix {
	rows := len(rowNames)
	cols := len(colNames)
	A := new(DenseNamedMatrix)
	A.elements = matrix.elements
	A.rows = rows
	A.cols = cols
	A.step = cols
	A.rowNames = rowNames
	A.colNames = colNames
	return A
}

func (self *DenseNamedMatrix) RowNames() []string {
	out := make([]string, len(self.rowNames))
	i := 0
	for k, _ := range(self.rowNames) {
		out[i] = k
		i++
	}
	return out
}

func (self *DenseNamedMatrix) RowMap() map[string] int {
	return self.rowNames
}

func (self *DenseNamedMatrix) ColMap() map[string] int {
	return self.colNames
}


func (self *DenseNamedMatrix) ColNames() []string {
	out := make([]string, len(self.colNames))
	i := 0
	for k, _ := range(self.colNames) {
		out[i] = k
		i++
	}
	return out
}



func (A *DenseNamedMatrix) GetNamedColVector(col string) *DenseNamedMatrix {
	return NamedFromDense(A.GetMatrix(0, A.colNames[col], A.rows, 1), A.rowNames, map[string] int { col : 0 })
}

func (A *DenseNamedMatrix) GetNamedRowVector(row string) *DenseNamedMatrix {
	return NamedFromDense(A.GetMatrix(A.rowNames[row], 0, 1, A.cols), map[string] int {row:0}, A.colNames)
}


func (A *DenseNamedMatrix) GetNamedRowSlices(rows []string) *DenseNamedMatrix {
	rowmap := make(map[string] int)
	for i := range rows {
		rowmap[rows[i]] = i
	}
	out := NamedZerosFromMap(rowmap, A.colNames)
	for name, i := range rowmap {
		orig_i := A.rowNames[name]
		for j := 0; j < A.cols; j++ {
			out.elements[ i * out.step + j ] = A.elements[ orig_i * A.step + j]
		}
	}
	return out
}