package main
import "fmt"

type Stack struct {
  Stack []float64;
  Head float64;
}

func (s *Stack) Push(value float64) {
  s.Stack = append(s.Stack, value);
  s.Head = value;
}

func (s *Stack) Pop() float64 {
  if len(s.Stack) == 1 {
    return_value := s.Stack[len(s.Stack) - 1];
    s.Stack = make([]float64, 0);
    return return_value;
  }
  if len(s.Stack) > 1 {
    return_value := s.Stack[len(s.Stack) - 1];
    s.Stack = s.Stack[:len(s.Stack) - 1];
    s.Head = s.Stack[len(s.Stack) - 1]
    return return_value;
  }
  return -1;
}

func (s *Stack) Peek() float64 {
  return s.Head;
}

func (s *Stack) Empty() bool {
  if len(s.Stack) == 0 {
    return true;
  }
  return false;
}

var s Stack;

func main() {
  var choice string;
  var val float64;
  mainLoop: for {
    fmt.Println("Operazione (push/pop/top/empty/quit)?");
    fmt.Scanf("%s", &choice);
    switch choice {
      case "push":
        fmt.Println("valore?");
        fmt.Scanf("%f", &val);
        s.Push(val);
        fmt.Println(s.stack);
      case "pop":
        previous_head := s.Pop();
        fmt.Println("testa", previous_head);
        fmt.Println(s.stack);
      case "top":
        fmt.Println("testa", s.Peek());
        fmt.Println(s.stack);
      case "empty":
        fmt.Println(s.Empty());
      case "quit":
        break mainLoop;
    }
  }
}
