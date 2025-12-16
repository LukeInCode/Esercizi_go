package main
import "fmt"
import "bufio"
import "os"

func reverseString(s string) string {
  if len(s) == 0 {
    return "";
  }
  slice := []rune(s);
  if len(s) == 1 {
    return string(slice[0]);
  }
  return string(slice[len(slice) - 1]) + reverseString(string(slice[:len(slice) - 1]));
}

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  for scanner.Scan() {
      fmt.Println(reverseString(scanner.Text()));
  }
}
