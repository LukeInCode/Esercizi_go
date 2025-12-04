package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "slices"

func aggiornaConteggio(m map[int]int, riga string) {
  words := strings.Fields(riga);
  for _,v := range words {
    m[len(v)]++;
  }
}

func trovaChiaviMinMax(m map[int]int) (int, int) {
  var keys []int;

  for k,_ := range m {
    keys = append(keys, k);
  }
  slices.Sort(keys);
  return keys[0], keys[len(keys) - 1];
}

func main() {
  var lenMap map[int]int = map[int]int{};
  scanner := bufio.NewScanner(os.Stdin);
  
  for scanner.Scan() {
    aggiornaConteggio(lenMap, scanner.Text());
  }

  fmt.Println(lenMap);

  var min, max int = trovaChiaviMinMax(lenMap);
  for i := min; i <= max; i++ {
    fmt.Printf("%d;%d ", i, lenMap[i]);
  }
  fmt.Println("");
  
  for i := min; i <= max; i++ {
    _, ok := lenMap[i];
    if ok {
      fmt.Printf("%d|%d ", i, lenMap[i]);
    } 
  }
  fmt.Println();
}
