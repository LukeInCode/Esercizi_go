package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

func getAllIntegers(s string) []int {
  fields := strings.Fields(s);
  digits := make([]int, len(fields));

  for i,v := range fields {
    intValue,_ := strconv.Atoi(v);
    digits[i] = intValue;
  }
  return digits;
}


func recursiveCount(el int, list []int) int {
  if len(list) == 1 {
    if list[0] == el {
      return 1;
    }
    return 0;
  }
  if list[0] == el {
    return 1 + recursiveCount(el, list[1:]); 
  }
  return 0 + recursiveCount(el, list[1:]); 
}

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  var target int;

  
  if scanner.Scan() {
    digits := getAllIntegers(scanner.Text());
    target = digits[0];
    fmt.Println(recursiveCount(target, digits));
  }
  
  for scanner.Scan() {
    fmt.Println(recursiveCount(target, getAllIntegers(scanner.Text())));
  }
}
