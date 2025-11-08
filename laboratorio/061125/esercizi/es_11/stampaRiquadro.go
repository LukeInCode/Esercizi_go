package main
import "fmt"
import "strings"

func horizLine(ch byte, n int)  string {
  return " " + strings.Repeat(string(ch),n) + " ";
}

func main() {
  var parola, max string;
  
  for {
    fmt.Scan(&parola);
    if parola == "stop" {
      break;
    }
    if len(parola) > len(max) {
      max = parola
    }
  }
  var ch byte = byte('-');
  fmt.Println(horizLine(ch,len(max) + 1));
  fmt.Printf("| %s |\n", max);
  fmt.Println(horizLine(ch,len(max) + 1));
}
