// Copyright (c) 2021 BlueStorm
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFINGEMENT IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package datatable

func New(rows []map[string]interface{}) *DataTable {
	dt := &DataTable{}
	dt.Count = len(rows)
	if dt.Count > 0 {
		for k := range rows[0] {
			dt.Columns = append(dt.Columns, &Column{Name: k})
		}
	}
	dt.Rows = rows
	return dt
}

type findKind uint

const (
	normal findKind = iota
	likeMode
	regXMode
)

type Column struct {
	Name   string
	Type   string
	Length int64
}

type Field struct {
	Tag string
	Val interface{}
}

type DataTable struct {
	Name    string
	Columns []*Column
	Rows    []map[string]interface{}
	Count   int
	mode    findKind
}

type DataSet struct {
	Tables []*DataTable
}

type UseMode int

const (
	Not UseMode = iota
	Get
	Add
	Set
	Del
	Count
)
