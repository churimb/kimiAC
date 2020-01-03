package ac

// by cr

type AcNode struct {
	fail      *AcNode
	isend     bool
	next      map[rune]*AcNode
	name      rune
	signature []string
}

type Ac struct {
	root *AcNode
}

func NewNode() *AcNode {
	return &AcNode{
		fail:      nil,
		isend:     false,
		next:      map[rune]*AcNode{},
		name:      '*',
		signature: []string{},
	}
}

func NewAhoCorasick() *Ac {
	root_node := Ac{root: NewNode()}
	root_node.root.signature = []string{"root"}
	return &root_node
}

func (ac Ac) BatchInsert(items []map[string]string) {
	for _, item := range items {
		ac.Insert(item["keyword"], item["signature"])
	}
}

func (ac Ac) Insert(s string, signature string) {
	chars := []rune(s)
	iter := ac.root
	for _, c := range chars {
		if _, existed := iter.next[c]; !existed {
			iter.next[c] = NewNode()
			iter.next[c].name = c
			//iter.next[c].signature = signature
		}
		iter = iter.next[c]
	}
	if iter.signature == nil {
		iter.signature = []string{signature}
	} else {
		iter.signature = append(iter.signature, signature)
	}
	if iter != ac.root {
		iter.isend = true
	}
}

func (ac Ac) Build() {
	tmplist := []*AcNode{} // 临时存放节点

	tmplist = append(tmplist, ac.root) // 从根节点开始

	for len(tmplist) != 0 {
		parrent := tmplist[0]
		tmplist = tmplist[1:]
		//fmt.Printf("parrent %c\n", parrent.name)

		for char, child := range parrent.next { // 从root开始遍历所有节点， 首先获得root的子节点们
			//fmt.Printf("child node %c\n", char)

			// 设置fail方向
			if parrent == ac.root { // 如果是root节点，fail指向root自身
				child.fail = ac.root
			} else { // 不是root节点
				failAcNode := parrent.fail
				for failAcNode != nil {
					if _, ok := failAcNode.next[char]; ok {
						child.fail = parrent.fail.next[char]
						//fmt.Printf("direct %c fail to %c\n", child.name, parrent.fail.next[char].name)
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil { // 最终的fail均指向root节点
					child.fail = ac.root
					//fmt.Printf("direct %c fail to root\n", child.name)
				}
			}

			// 通过tmplist实现遍历子节点
			tmplist = append(tmplist, child)
		}
	}
}

func (ac Ac) Match(content string) (results []string) {
	chars := []rune(content)
	iter := ac.root // 从根节点开始

	//var start, end int

	for _, c := range chars { // 逐个词开始比较

		_, ok := iter.next[c]        // 是否存在
		for !ok && iter != ac.root { // 不存在且非根节点
			iter = iter.fail // 转到fail节点
		}

		if _, ok = iter.next[c]; ok { // 存在
			if iter == ac.root { // 如果是根节点
				//start = i
			}

			iter = iter.next[c] // 转到接下去的节点
			if iter.isend {
				//end = i
				//results = append(results, string([]rune(content)[start:end+1]))  // 保存检出结果
				for _, sign := range iter.signature {
					results = append(results, sign)
				}
			}
		}
	}
	return
}
