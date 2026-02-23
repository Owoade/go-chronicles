package main

func main() {
	if err := ToTar("sample"); err != nil {
		panic(err)
	}

	if err := FromTar("tar-1771852301792.tar"); err != nil {
		panic(err)
	}
}
