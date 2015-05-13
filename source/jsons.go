package source
import (
	"encoding/json"
	"fmt"
)

type Response1 struct {
	Page int
	Fruits []string
}

type Response2 struct {
	Page int           `json:"page"`
	Fruits []string    `json:"fruits"`
}

func Jsons() {

	bloB, _ := json.Marshal(true)
	fmt.Println(string(bloB))

	intB, _ := json.Marshal(1)
	fmt.Println(string(intB))

	fltB, _ := json.Marshal(2.34)
	fmt.Println(string(fltB))

	strB, _ := json.Marshal("gopher")
	fmt.Println(string(strB))

	slcD := []string{"apple", "peach", "pear"}
	slcB, _ := json.Marshal(slcD)
	fmt.Println(string(slcB))

	mapD := map[string]int{"apple":5, "lettuce":7}
	mapB, _ := json.Marshal(mapD)
	fmt.Println(string(mapB))

	res1D := &Response1{Page:1, Fruits:[]string{"apple", "peach", "pear"}}
	res1B, _ := json.Marshal(res1D)
	fmt.Println(string(res1B))

	res2D := &Response2{Page:1, Fruits:[]string{"apple", "peach", "pear"}}
	res2B, _ := json.Marshal(res2D)
	fmt.Println(string(res2B))

	byt := []byte(`{"num":6.13,"strs":["a","b"]}`)
	var data map[string]interface{}

	if err := json.Unmarshal(byt, &data); err!=nil {
		panic(err)
	}
	fmt.Println(data)

	num := data["num"].(float64)
	fmt.Println(num)

	strs := data["strs"].([]interface{})
	str1 := strs[0].(string)
	fmt.Println(str1)

	str := `{"page":1,"fruits":["apple","peach"]}`
	res := &Response2{}
	json.Unmarshal([]byte(str), res)
	fmt.Println(res)
	fmt.Println(res.Fruits[0])

}
