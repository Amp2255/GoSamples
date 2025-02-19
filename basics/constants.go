//Go supports constants of character, string, boolean, and numeric values.
package main
import(
	"fmt"
	"math"
)
	
const s string = "constant"
const n  =1
func main() {
	fmt.Println(s, math.Sin(n))
}