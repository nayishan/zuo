package tries

type Node struct {
	Pass int
	End  int
	Next [26]*Node
}

func (t *Node) Insert(word string) {
	b := []byte(word)
	N := len(b)
	node := t

	node.Pass++
	for i := 0; i < N; i++ {
		index := b[i] - 'a'
		if node.Next[index] == nil {
			node.Next[index] = &Node{
				Pass: 0,
				End:  0,
				Next: [26]*Node{},
			}
		}
		node = node.Next[index]
		node.Pass++
	}
	node.End++
}
func (t *Node) Search(word string) int {
	if word == "" {
		return 0
	}
	b := []byte(word)
	N := len(b)
	node := t
	for i := 0; i < N; i++ {
		index := b[i] - 'a'
		if node.Next[index] == nil {
			return 0
		}
		node = node.Next[index]
	}
	return node.End
}
func (t *Node) prefixNum(word string) int {
	if word == "" {
		return 0
	}
	b := []byte(word)
	N := len(b)
	node := t
	for i := 0; i < N; i++ {
		index := b[i] - 'a'
		if node.Next[index] == nil {
			return 0
		}
		node = node.Next[index]
	}
	return node.Pass
}

func (t *Node) Delete(word string) {
	if word == "" {
		return
	}
	if t.Search(word) != 0 {
		b := []byte(word)
		N := len(b)
		node := t
		node.Pass--
		for i := 0; i < N; i++ {
			index := b[i] - 'a'
			if node.Next[index].Pass == 0 {
				node.Next[index] = nil
				return
			}
			node = node.Next[index]
			node.Pass--
		}
		node.End--
	}
}
