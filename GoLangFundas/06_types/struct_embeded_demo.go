package main

import "fmt"

type pen struct {
	nib float32
	inkColor string
}

func (p pen) display(){
	fmt.Println("Nib: ",p.nib,"\tInkColor: ",p.inkColor)
}

type marker struct {
	p pen
	isParminant bool
}

func (m marker) display(){
	fmt.Println("Nib: ",m.p.nib,"\tInkColor: ",m.p.inkColor,"\tIsParminent: ",m.isParminant)
}

func main(){
	p1 := pen{0.4,"blue"}
	m1 := marker{pen{0.32,"Green"},false}

	p1.display()
	m1.display()
}
