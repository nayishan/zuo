package tries

import "testing"

func TestTries(t *testing.T) {
	n := Node{}
	n.Insert("abc")
	n.Insert("abcd")
	t.Log(n.Search("abc"))
	t.Log(n.prefixNum("abc"))
	n.Delete("abcd")
	t.Log(n.Search("abc"))
	t.Log(n.prefixNum("abc"))
	n.Insert("abcd")
	t.Log(n.Search("abc"))
	t.Log(n.prefixNum("abc"))

}
