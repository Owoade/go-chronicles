package main

import "time"

func main() {
	now := time.Now()
	tenMinsAgo := now.Add(-10 * time.Minute)
	println(tenMinsAgo.Format(time.RFC3339)) // prints iso format
	

	println(time.Now().Add(10*time.Minute).Format(time.RFC3339), time.Now().Format(time.RFC3339))

	println(now.Before((tenMinsAgo)), now.After(tenMinsAgo))

	println(now.Location().String())
}
