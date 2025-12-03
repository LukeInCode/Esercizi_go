package main
import "fmt"
import "bufio"
import "math"
import "os"
import "strings"
import "strconv"

func main() {
  scanner := bufio.NewScanner(os.Stdin);

  for !scanner.Scan() {
  }
  
  var polInput []string = strings.Fields(scanner.Text());
  var solution int = 0;
  x,_ := strconv.Atoi(polInput[ len(polInput) - 1 ]);
  
  for i,v := range polInput[:len(polInput) - 1] {
    n,_ := strconv.Atoi(v);
    var tmp int = 1;
    if i > 0 {
      tmp = int(math.Pow(float64(x), float64(i)));
    }
    solution += n * tmp;
  }

  fmt.Println(solution);
}
