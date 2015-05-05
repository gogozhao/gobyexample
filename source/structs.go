package source
import "fmt"

type person struct {
	name string
	age int
}


type rect struct {
	width, height int
}

func (r *rect) area() int {
	return r.width*r.height
}

func (r rect) perim() int {
	return 2*r.width+2*r.height

}




func Structs() {

	fmt.Println(person{"Bob", 20})

	fmt.Println(person{name:"Alice", age:30})

	fmt.Println(person{name:"Fred"})

	fmt.Println(&person{name:"Ann", age:40})

	s := person{name:"Sean", age:50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age=51
	fmt.Println(sp.age)

	r := rect{width:10, height:5}
	fmt.Println("area:", r.area())
	fmt.Println("perim:", r.perim())

	rp := &r
	fmt.Println("area:", rp.area())
	fmt.Println("perim:", rp.perim())
}