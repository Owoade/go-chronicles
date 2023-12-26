package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Product struct {
	name     string
	in_stock int
}

func purchase_random_product(prods []Product, mux *sync.Mutex, wg *sync.WaitGroup) {

	mux.Lock()

	defer mux.Unlock()

	random_product_index := rand.Intn(len(prods))

	prods[random_product_index].in_stock--

	wg.Done()

}

func main() {

	products := []Product{
		{name: "PS5", in_stock: 20},
		{name: "Nitendo switch", in_stock: 30},
		{name: "Iphone 15", in_stock: 30},
	}

	unserious_customers := 30

	var (
		wg  sync.WaitGroup
		mux sync.Mutex
	)

	wg.Add(unserious_customers)

	for i := 1; i <= unserious_customers; i++ {

		go purchase_random_product(products, &mux, &wg)

	}

	wg.Wait()

	fmt.Println("Done", products)

}
