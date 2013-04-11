
package matrix

import (
	"io"
	"bufio"
	"bytes"
	"strings"
	"strconv"
	"errors"
	"fmt"
	"container/list"
)

func LoadTSV(input io.Reader) (out *DenseNamedMatrix, err error) {

	var (
		part []byte
		prefix bool
	)
	out = nil

	reader := bufio.NewReader(input)
	
	rowNames := make(map[string] int)
	colNames := make(map[string] int)
	
	head := true
	row_count := 0
	col_count := 0
	
	row_list := list.New()

	buffer := bytes.NewBuffer(make([]byte, 0))
	for {
		if part, prefix, err = reader.ReadLine(); err != nil {
			break
		}
		buffer.Write(part)
		if !prefix {
			var tmp []string
			line := buffer.String()
			tmp = strings.Split(line, "\t")
			if head {
				for i := 1; i < len(tmp); i++ {
					colNames[tmp[i]] = i-1
				}
				head = false
				col_count = len(colNames)
			} else {
				if _,ok :=rowNames[tmp[0]]; !ok {
					rowNames[tmp[0]] = row_count
					row_count += 1
					rvals := make([]float64, col_count)
					if (len(tmp)-1 != col_count) {
						err = errors.New( fmt.Sprintf("Col Mismatch %d %d", col_count, len(tmp)-1));
						return
					}
					for i := 1; i < len(tmp); i++ {
						rvals[i-1],_ = strconv.ParseFloat(tmp[i], 64)
					}		
					row_list.PushBack(rvals)		
				}
			}
			buffer.Reset()
		}
	}

	out = NamedZeros(rowNames, colNames)
	row_num := 0
	for e := row_list.Front(); e != nil; e = e.Next() {
		row := e.Value.([]float64)
		for col_num := 0; col_num < col_count; col_num++ {
			v := row[col_num]
			out.Set(row_num, col_num, v)
		}
		row_num++
	}
	return
}