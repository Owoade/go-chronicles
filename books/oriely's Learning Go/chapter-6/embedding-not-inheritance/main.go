package main

type Inner struct {
	Value int
}

type Outer struct {
	Value int
	Inner
}

func (i Inner) PrintValue() {
	println("Inner Value:", i.Value)
} 


func (o Outer) PrintValue() {
	println("Outer Value:", o.Value)
}

func (i Inner) GetPrintValue() {
	i.PrintValue()
}

func main(){
	
	inner := Inner{Value: 10}
	outer := Outer{Value: 20, Inner: inner}

	// Calling GetPrintValue on Inner
	// Though the underlying method is also defined on Outer, it will call Inner's PrintValue
	outer.GetPrintValue() // Calls Inner's PrintValue
}