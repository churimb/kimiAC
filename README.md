# 用法

学习用GO实现AC自动机（Aho-Corasick automation），支持添加自己的匹配信息。

使用例子如下：

```go
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
}

// === RUN   TestAc
// {"descript": "action speek"}
// {"descript": "confidence"}
// {"descript": "good people"}
// --- PASS: TestAc (0.00s)
// PASS

```