package main

type Foo struct {
	Field1 string
	Field2 int
}

func MakeFoo(f *Foo) error {
	/*
		if pointer passed is nil dereferencing 
		The statement below will yield a panic because it's equivalent to *f.Field (nil derefrencing)
	*/
	f.Field1 = "val" 
	f.Field2 = 20
	return nil
}

func main() {
	var f *Foo
	_ = MakeFoo(f)
}
