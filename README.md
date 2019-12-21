# 用法

学习用GO实现AC自动机（Aho-Corasick automation），支持添加自己的匹配信息。

使用例子如下：

```go
acm := NewAhoCorasick()
acm.insert([]byte("你真的很好"), []byte("{\"descript\": \"good people\"}"))
acm.insert([]byte("你真的很坏"), []byte("{\"descript\": \"bad people\"}"))
acm.insert([]byte("很自信"), []byte("{\"descript\": \"confidence\"}"))
acm.batchInsert([]map[string][]byte{{"keyword":[]byte("说"), "signature":[]byte("{\"descript\": \"action speek\"}")},
  {"keyword":[]byte("老师"), "signature":[]byte("{\"descript\": \"relative teacher\"}")}})

acm.Build()

res := acm.Match("老实说，你很自信，我也是这么认为的，你真的很好，给你小红花。")
for _, result := range res {
  //fmt.Println(string(result))
  var json_data interface{}
  //fmt.Println(result)
  err := json.Unmarshal(result, &json_data)
  if err != nil  {
    fmt.Println(err)
  }
  buf, _ := json.Marshal(json_data)
  fmt.Println(string(buf))
}

// {"descript":"action speek"}
// {"descript":"confidence"}
// {"descript":"good people"}

```