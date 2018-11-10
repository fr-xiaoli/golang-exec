package main

import "fmt"

func main() {
	a := 42
	b := a
	c := &a

	fmt.Printf("====== Values ====== \n a = %v\n b = %v\n c = %v\n *c = %v\n",
		a, b, c, *c)
	fmt.Printf("====== Types ======= \n a: %T\n b: %T\n c: %T\n *c: %T\n",
		a, b, c, *c)

	b = 43
	fmt.Printf("b = 43, now a = %v\n", a)
	*c = 44
	fmt.Printf("*c = 44, now a = %v, b = %v, c = %v\n", a, b, c)
	fmt.Println("==============")
	d := c
	*d = 45
	fmt.Printf("d := c, *d = 45\n now a = %v, b = %v\n c = %v, *c = %v, d = %v\n",
		a, b, c, *c, d)

	e := "I am a string"

	// You can't do this since "cannot use &e (type *string) as type *int in assignment"
	// c = &e
	f := &e
	fmt.Printf("\nWhat's the difference between a *int and a *string?: c(*int): %v, f(*string): %v\n", c, f)
}
