package main

import "fmt"

//Article of struct
type Article struct {
	Title  string
	Author string
}

//Post string
type Post struct {
	Title  string
	Author string
}

//Stringer return string
type Stringer interface {
	String() string
}

func (a Article) String() string {
	return fmt.Sprintf("The %q article %s, -- NOSSA ", a.Title, a.Author)
}

func (b Post) String() string {
	return fmt.Sprintf("Another %q article %s, -- NOSSA ", b.Title, b.Author)
}

// Print return string
func Print(s Stringer) {
	fmt.Println(s.String())
}

func main() {
	a := Article{
		Title:  "Undesande",
		Author: "Sammy",
	}
	b := Post{
		Title:  "Undesande",
		Author: "Sammy NEW",
	}
	Print(a)
	Print(b)
}

//Cria a struct
//Cria uma interface
//Cria um Metodo
//Cria uma função

//inicia a strutura
//inicia a função que chama o metodo passando a estrutura
