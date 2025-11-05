package main
import "fmt"
import "io"
import "strconv"
import "strings"

func main() {
  var n, previous, group int = 0, -1, 0;
  var gradino,max, lastGroup string = "", "", "";
  
  for {
    _,err := fmt.Scanf("%d", &n);
    if err == io.EOF || n < previous {
      break;
    }
    nString := strconv.Itoa(n);
    if !strings.ContainsAny(gradino, nString) {
      if group == 2 {
        if len(max) < len(gradino) {
          max = gradino;
        }
        group = 0;
        gradino = lastGroup;
        fmt.Println(gradino)
      }
      group++;
    }
    if group < 2 || n == previous {
      gradino += nString;
    }
    if strings.ContainsAny(lastGroup, nString) {
      lastGroup += nString;
    }else {
      lastGroup = nString;
    }
            fmt.Println(gradino)
    previous = n;
  }
  fmt.Printf("Il gradino massimo è %s, con una lunghezza di %d\n", max, len(max));
}
