package s
import "fmt"

func init() {
	fmt.Println("init in sandbox.go")
}
func S() int64 {
	fmt.Println("calling s() in sandbox.go")
	return 1
}