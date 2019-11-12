package main
import "fmt"

/*对于一个数n，求n的阶乘之和，即： 1！ + 2！ + 3！+…n! */
func main() {
	partion(4)
}

func partion(a int) {
	count := 0
	for i := 1; i <= a; i++ {
		tem := 1
		for j := 1; j <= i; j++ {
			tem *= j
		}
		count += tem
	}
	fmt.Println(count)
}
