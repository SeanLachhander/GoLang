package main

import(
	"fmt"
)

 var a string = "Stored in the variable a"
 var b, c string = "stored in b", "stored in c"
 var d string

func main(){
	a_1 := 10
	b_1 := "Golang"
	c_1 := 4.17
	d_1 := true

	fmt.Printf("Valeu: %v :: Type: %T \n", a_1,a_1)
	fmt.Printf("%v :: Type: %T \n", b_1, b_1)
	fmt.Printf("%v :: Type: %T \n", c_1, c_1)
	fmt.Printf("%v :: Type: %T \n", d_1, d_1)


	d = "stored in d"
	var e int = 42
	f := 43
	g := "stored in g"
	h, i := "stored in h", "stored in i"
	j, k, l, m := 44.7, true, false, 'm'
	n := "n"
	o := 'o'

	fmt.Println("e : ", e)
	fmt.Println("f : ", f)
	fmt.Println("g : ", g)
	fmt.Println("h : ", h)
	fmt.Println("i : ", i)
	fmt.Println("j : ", j)
	fmt.Println("k : ", k)
	fmt.Println("l : ", l)
	fmt.Println("m : ", m)
	fmt.Println("n : ", n)
	fmt.Println("o : ", o)

}