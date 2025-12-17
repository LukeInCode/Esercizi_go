package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func getAllIntegers(s string) []int {
  fields := strings.Fields(s);
  digits := make([]int, len(fields));

  for _,v := range fields {
    intValue,_ := strconv.Atoi(v);
    digits = append(digits, intValue);
  }
  return digits;
}


func recursiveMax(list []int) int {
  if len(list) == 1 {
    return list[0];
  }
  max := recursiveMax(list[1:]);
  if max > list[0] {
    return max;
  }
  return list[0];
}


func main() {
  scanner := bufio.NewScanner(os.Stdin);
  for scanner.Scan() {
    fmt.Println(recursiveMax(getAllIntegers(scanner.Text())));
  }
}
