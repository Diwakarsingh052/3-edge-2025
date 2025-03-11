package main

func main() {
	add
	sub
	mult

	go Add(1, 2)
	go Sub(1, 2)
	go Mult(1, 2)
	go CollectResult()
}

func Add(a, b int) int {
	// send the result to the channel
}

func Sub(a, b int) int {}

func Mult(a, b int) int {

}
func CollectResult() int {

	//recv all the values
}
