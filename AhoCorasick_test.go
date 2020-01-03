package ac_test

// by cr

import (
	"../ac"
	"fmt"
	"testing"
)


func TestAc(t *testing.T) {
	//fmt.Println("hello!")

	acm := ac.NewAhoCorasick()
	acm.Insert("你真的很好", "{\"descript\": \"good people\"}")
	acm.Insert("你真的很坏", "{\"descript\": \"bad people\"}")
	acm.Insert("很自信", "{\"descript\": \"confidence\"}")
	acm.BatchInsert([]map[string]string{
		{"keyword":"说", "signature":"{\"descript\": \"action speek\"}"},
		{"keyword":"老师", "signature":"{\"descript\": \"relative teacher\"}"},
	})

	acm.Build()

	res := acm.Match("老实说，你很自信，我也是这么认为的，你真的很好，给你小红花。")

	for _, result := range res {
		fmt.Println(result)
		//var json_data interface{}
		//fmt.Println(result)
		//err := json.Unmarshal([]byte(result), &json_data)
		//if err != nil  {
		//	fmt.Println(err)
		//}
		//buf, _ := json.Marshal(json_data)
		//fmt.Println(string(buf))
	}
}

