// demos
package main

import (
	"fmt"
	"github.com/JonKerridge/pnp"
)

func main() {
	fmt.Printf("Concurrent and Parallel Systems in GO\n")
	fmt.Printf("Jon Kerridge Edinburgh Napier Univerity\n")
	fmt.Printf("email : j dot kerridge at napier.ac.uk\n")
	fmt.Printf(" code : https://github.com/JonKerridge\n\n")

	fmt.Printf("A Set of Demo Networks Showing the Operation of Package pnp\n\n")

	// this network generates a sequence of increasing numbers 0 .. 10
	fmt.Printf("Generate a sequence of increasing numbers 1 .. 10\n ")
	var n2c = make(chan int)
	var str1 = make(chan string)
	go pnp.Numbers(n2c, 1)
	go pnp.ConvertIntStr(n2c, str1)
	//	go pnp.Console(c2c, "List of Numbers", 20)  does not work as expected
	var v string
	var i int = 0
	for i < 10 {
		v = <-str1
		fmt.Printf("%v", v)
		i = i + 1
	}
	fmt.Printf("\n")
	fmt.Printf("Generate the running sum of a sequence of increasing numbers 0 .. 9\n ")
	n2i := make(chan int)
	i2p := make(chan int)
	str2 := make(chan string)
	go pnp.Numbers(n2i, 0)
	go pnp.Integrate(n2i, i2p)
	go pnp.ConvertIntStr(i2p, str2)
	i = 0
	for i < 10 {
		v = <-str2
		fmt.Printf("%v", v)
		i = i + 1
	}
	fmt.Printf("\n")
	fmt.Printf("Undo the effect of Integrate\n ")
	num2int := make(chan int)
	int2ri := make(chan int)
	ri2c := make(chan int)
	str4 := make(chan string)
	go pnp.Numbers(num2int, 1)
	go pnp.Integrate(num2int, int2ri)
	go pnp.ReverseIntegrate(int2ri, ri2c)
	go pnp.ConvertIntStr(ri2c, str4)
	i = 0
	for i < 10 {
		v = <-str4
		fmt.Printf("%v", v)
		i = i + 1
	}
	fmt.Printf("\n")
	fmt.Printf("Generate the squares of numbers\n ")
	nu2i := make(chan int)
	in2p := make(chan int)
	p2c := make(chan int)
	str3 := make(chan string)
	go pnp.Numbers(nu2i, 0)
	go pnp.Integrate(nu2i, in2p)
	go pnp.Pairs(in2p, p2c)
	go pnp.ConvertIntStr(p2c, str3)
	i = 0
	for i < 10 {
		v = <-str3
		fmt.Printf("%v", v)
		i = i + 1
	}
	fmt.Printf("\n")
	fmt.Printf("Another example: pnp.Example1(out) - how does it work?\n")
	oute := make(chan int)
	str5 := make(chan string)
	go pnp.Example1(oute)
	go pnp.ConvertIntStr(oute, str5)
	i = 0
	for i < 10 {
		v = <-str5
		fmt.Printf("%v", v)
		i = i + 1
	}
	fmt.Printf("\nFinished Example Processing")
	fmt.Printf("\n\n")
	fmt.Printf("Demonstration of Tabulating Many Outputs\n")
	numb2c1 := make(chan int)
	c12int := make(chan int)
	int2c2 := make(chan int)
	c22prs := make(chan int)
	prs2pref := make(chan int)
	str6 := make(chan string)
	var tabs = []chan int{make(chan int), make(chan int), make(chan int)}
	//	for i := range tabs {
	//		tabs[i] = make(chan int)
	//	}
	go pnp.Numbers(numb2c1, 0)
	go pnp.Copy2(numb2c1, c12int, tabs[0])
	go pnp.Integrate(c12int, int2c2)
	go pnp.Copy2(int2c2, c22prs, tabs[1])
	go pnp.Pairs(c22prs, prs2pref)
	go pnp.Prefix(prs2pref, tabs[2], 0)
	go pnp.Tabulate(tabs, str6)
	i = 0
	for i < 10 {
		v = <-str6
		fmt.Printf("%v", v)
		i = i + 1
	}
	fmt.Printf("\nFinished Example Processing Using Tabulate")

}
