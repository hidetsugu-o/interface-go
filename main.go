package main

func main() {
	// NewではInterfaceが帰ってくる
	client := New()

	client.Get()
	client.Post()
}
