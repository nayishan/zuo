package liststack

import (
	"errors"
	"fmt"
)

type LinkStack struct {
	data int
	next *LinkStack
}

// func main() {
// 	//初始化一个链栈
// 	fmt.Println("---初始化链栈---")
// 	linkStack1 := initLinkStack()
//
// 	//入栈
// 	fmt.Println("---入栈---")
// 	for i := 1; i <= 5; i++ {
// 		linkStack1.push(i)
// 		fmt.Printf("第%v次出栈, 值为:%v\n", i, i)
// 	}
//
// 	//出栈
// 	fmt.Println("---出栈---")
// 	for i := 1; i <= 6; i++ {
// 		v, err := linkStack1.pop()
// 		if err != nil {
// 			fmt.Println(err)
// 			break
// 		}
// 		fmt.Printf("第%v次出栈, 值为:%v\n", i, v)
// 	}
// }

func initLinkStack() (linkStack *LinkStack) {
	//初始化链栈，新建链栈头节点
	linkStack = &LinkStack{
		data: -1,
		next: nil,
	}
	return linkStack
}

func (s *LinkStack) push(v int) {
	//链栈可先不考虑栈满，因为目前没有对栈做限制

	pushNode := &LinkStack{
		data: v,
		next: nil,
	}

	pushNode.next = s.next //入栈节点的next指向头节点的next节点
	s.next = pushNode      //头节点指向入栈节点
}

func (s *LinkStack) pop() (int, error) {
	var v int
	//判断栈是否为空
	if s.next == nil {
		return 0, errors.New("error: 栈为空")
	}

	tmpTop := s.next
	v = tmpTop.data //出栈，获取节点的元素值

	s.next = tmpTop.next //头节点指向原栈顶节点的下一个节点
	tmpTop.next = nil    //原栈顶节点指向nil

	return v, nil
}
