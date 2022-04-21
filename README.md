# datatable


``` golang
import (
    "fmt"
    "github.com/BlueStorm001/datatable"
)
```

``` golang
func main() {
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
    
    dt := datatable.New(rows)

    // Where 条件匹配 (a=1 and b=2) or (c=2 and d=3) Condition match
    table := dt.Where("Text='CN' and (code='BJS' or code='SHA')").OrderBy("id") 
    for i, row := range table.Rows {
        fmt.Println(i,row)
    }
    
    // 使用模糊搜索 
    // Use fuzzy search
    table = dt.Like("name='CN%' and money=1.2%").OrderBy("id desc")
    
    // 使用正则表达式 
    // Use regular expressions
    table = dt.Find("code='[A-Z]{3}'").OrderBy("id desc")
    
    // 分组
    // Group
    table = dt.GroupBy("name")
    for i, row := range table.Rows {
        name := row["name"].(string)
        newTable := dt.Where("name='" + name + "' and (code='BJS' or code='SHA')").OrderBy("id") //[id asc , name desc]...
        fmt.Println(newTable)
    }
   
}
```