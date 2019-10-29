package add

func Add(a int, b int , c chan int) {
	sum := a + b
	c <- sum
}
