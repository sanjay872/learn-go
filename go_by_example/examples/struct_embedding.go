package examples

import "fmt"

type base struct {
	num int
}

type container struct {
	base
	name string
}

func (b base) describeBase() string {
	return fmt.Sprintf("The base value is %v", b.num)
}

func (c container) describeContainer() string {
	return fmt.Sprintf("The container has base with val of %v and name as %s", c.base.num, c.name)
}

func StructEmbedding() {

	c := container{base: base{num: 12}, name: "sanjay"}

	fmt.Println(c)

	fmt.Println(c.base.num)
	fmt.Println(c.name)

	type describer interface {
		describeBase() string
		describeContainer() string
	}

	var d describer = c

	fmt.Println(d.describeBase())
	fmt.Println(d.describeContainer())
}
