/*
 * Copyright (c) 2021 BlueStorm
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFINGEMENT IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package datatable

import (
	"testing"
)

func GetTestDataTable() *DataTable {
	//测试数据 test case
	var rows = []map[string]interface{}{
		{"id": 5, "code": "BJS", "name": "CN", "money": 1.23},
		{"id": 2, "code": "BJS", "name": "CN", "money": 2.21},
		{"id": 3, "code": "SHA", "name": "CN", "money": 1.26},
		{"id": 4, "code": "NYC", "name": "US", "money": 3.99},
		{"id": 7, "code": "MEL", "name": "US", "money": 3.99},
		{"id": 1, "code": "", "name": "CN", "money": 2.99},
	}
	for i := 10; i < 9999; i++ {
		rows = append(rows, map[string]interface{}{"id": i, "code": "BJS", "name": "CN", "money": 1.23})
	}
	return New(rows)
}

func TestWhere(t *testing.T) {
	dt := GetTestDataTable()
	whereDT1 := dt.Where("code==''").OrderBy("id")
	t.Log(whereDT1)
}

func TestDataTable(t *testing.T) {
	dt := GetTestDataTable()
	whereDT1 := dt.Where("id=99 or id >= 999").OrderBy("id") //[id asc , name desc]...
	t.Log(whereDT1)
	//Group By 分组
	groupDT := dt.GroupBy("name")
	for _, row := range groupDT.Rows {
		name := row["name"].(string)
		//$GroupCount$ 分组的数量 Number of groups
		t.Log(name, row["$GroupCount$"])
		//Where
		whereDT := dt.Where("name='" + name + "' and (code='BJS' or code='SHA')").OrderBy("id") //[id asc , name desc]...
		t.Log(whereDT)
	}
	//使用模糊搜索 Use fuzzy search
	likeDT := dt.Like("name='CN%' and money=1.2%").OrderBy("id desc")
	t.Log(likeDT)
	//使用正则表达式 Use regular expressions
	findDT := dt.Find("code='[A-Z]{3}'").OrderBy("id desc")
	t.Log(findDT)
}

func Benchmark_DataTable(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			dt := GetTestDataTable()
			dt.Where("name='CN' or id=9999")
		}
	})
}
