
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


func NamedZeros(rowNames, colNames map[string] int) *DenseNamedMatrix {
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


func (self *DenseNamedMatrix) RowNames() []string {
	out := make([]string, len(self.rowNames))
	i := 0
	for k, _ := range(self.rowNames) {
		out[i] = k
		i++
	}
	return out
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