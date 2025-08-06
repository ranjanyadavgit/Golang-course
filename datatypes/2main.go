import "fmt"

func main() {

	a := 45.65
	b := 65.25

	c := b - a

	fmt.Printf("A = %f\n", a)
	fmt.Printf("B = %f\n", b)

	fmt.Printf("%f - %f = %f\n", b, a, c)
	fmt.Printf("type of c : %T", c)
}
