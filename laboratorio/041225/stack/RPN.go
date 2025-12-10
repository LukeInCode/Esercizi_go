package main
import "fmt"
import "strconv"
import "strings"

func main() {
  var input string;
  var terms Stack;
  var op string = "+-*/";
  
  for {
    fmt.Print("Next? (+, -, *, /, q o un numero) ");
    fmt.Scanf("%s", &input);
    if input == "q" || input == "Q" {
      break;
    }
    term, err := strconv.ParseFloat(input, 64);
    if err == nil {
        terms.Push(term);
    } else {
      if strings.ContainsAny(input, op) {
        if len(terms.Stack) < 2 {
          fmt.Println("not enough data");
        } else {
          var t1, t2 float64 = terms.Pop(), terms.Pop();
          switch input {
            case "+":
              terms.Push(t1 + t2);
            case "-":
              terms.Push(t1 - t2);
            case "*":
              terms.Push(t1 * t2);
            case "/":
              terms.Push(t1 / t2);
          }
        }
      }
    }
    fmt.Println(terms.Stack);
  }
}
