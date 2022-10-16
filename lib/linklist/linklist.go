package linklist

type Fifo_Interface interface {
	Acl_push_back(data interface{})
	Acl_push_front(data interface{})
	Acl_pop_back() (interface{}, error)
	Acl_pop_front() (interface{}, error)
	Acl_head() (interface{}, error)
	Acl_tail() (interface{}, error)
	Acl_index(index int) (interface{}, error)
}

type acl_fifo_info struct {
	data interface{}
	prev *acl_fifo_info
	next *acl_fifo_info
}

func acl_fifo_info_new(data interface{}) *acl_fifo_info {
	return &acl_fifo_info{
		data: data,
		prev: nil,
		next: nil,
	}
}

type Acl_fifo struct {
	head *acl_fifo_info
	tail *acl_fifo_info
	cnt  int
}

func Acl_fifo_new() *Acl_fifo {
	return &Acl_fifo{
		head: nil,
		tail: nil,
		cnt:  0,
	}
}

func (fifo *Acl_fifo) Acl_push_back(data interface{}) {
	info := acl_fifo_info_new(data)
	if info == nil {
		return
	}

	if fifo.tail == nil {
		fifo.head = info
		fifo.tail = info
		info.prev = nil
		info.next = nil
	} else {
		fifo.tail.next = info
		info.prev = fifo.tail
		info.next = nil
		fifo.tail = info
	}

	fifo.cnt++

}
func (fifo *Acl_fifo) Acl_push_front(data interface{}) {
	info := acl_fifo_info_new(data)
	if info == nil {
		return
	}

	if fifo.head == nil {
		fifo.head = info
		fifo.tail = info
		info.prev = nil
		info.next = nil
	} else {
		info.next = fifo.head
		fifo.head.prev = info
		info.prev = nil
		fifo.head = info
	}

	fifo.cnt++
}
func (fifo *Acl_fifo) Acl_pop_back() interface{} {

	info := fifo.tail
	if fifo.tail.prev == nil {
		fifo.tail = nil
		fifo.head = nil
	} else {
		fifo.tail.prev.next = nil
		fifo.tail = fifo.tail.prev
	}

	fifo.cnt--

	return info.data
}
func (fifo *Acl_fifo) Acl_pop_front() interface{} {

	info := fifo.head
	if fifo.head.next == nil {
		fifo.head = nil
		fifo.tail = nil
	} else {
		fifo.head.next.prev = nil
		fifo.head = fifo.head.next
	}

	fifo.cnt--

	return info.data
}
func (fifo *Acl_fifo) Acl_size() int {
	return fifo.cnt
}
func (fifo *Acl_fifo) Acl_index(index int) interface{} {

	info := fifo.head
	for i := 0; i < fifo.cnt; i++ {
		if i == index {
			break
		}
		info = info.next
	}

	return info.data
}

func (fifo *Acl_fifo) Acl_head() interface{} {
	return fifo.head.data
}
func (fifo *Acl_fifo) Acl_tail() interface{} {
	return fifo.tail.data
}

// func main() {
// 	var size int
// 	var fifo *Acl_fifo
//
// 	fifo = Acl_fifo_new()
//
// 	size = fifo.acl_size()
// 	fmt.Print(size, " should be [0]\n")
//
// 	fmt.Printf("\n----------push_back[pop_back]-------------------\n")
// 	for i := 0; i < 10; i++ {
// 		fifo.acl_push_back(i)
// 		fmt.Printf("%d\t", i)
// 	}
// 	fmt.Println()
//
// 	head, _ := fifo.acl_head()
// 	tail, _ := fifo.acl_tail()
// 	fmt.Println(head.(int), tail.(int))
//
// 	size = fifo.acl_size()
// 	for i := 0; i < size; i++ {
// 		info, _ := fifo.acl_pop_back()
// 		fmt.Printf("%d\t", info.(int))
// 	}
//
// 	fmt.Printf("\n----------push_back[pop_front]-------------------\n")
//
// 	for i := 0; i < 10; i++ {
// 		fifo.acl_push_back(i)
// 		fmt.Printf("%d\t", i)
// 	}
// 	fmt.Println()
//
// 	head, _ = fifo.acl_head()
// 	tail, _ = fifo.acl_tail()
// 	fmt.Println(head.(int), tail.(int))
//
// 	size = fifo.acl_size()
// 	for i := 0; i < size; i++ {
// 		info, _ := fifo.acl_pop_front()
// 		fmt.Printf("%d\t", info.(int))
// 	}
// 	fmt.Printf("\n-----------------------------\n")
//
// }
