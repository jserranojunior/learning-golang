package main
import "fmt"

type person struct{
	name string
	age int
}

//func newPerson(name string, age int) *person{
//	p := person{name: name, age: age}
//	return &p
//}

func main(){
	newBoss := person{"jorge", 15}
	fmt.Println(newBoss.sun())
}

func (p person) sun() int{
		return p.age * 2
}