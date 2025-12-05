package main
import "fmt"
import "slices"

func extractAllDigits(n int) (digits []int) {
  for n > 0 {
    digits = append(digits, n % 10);
    n /= 10;
  }
  slices.Reverse(digits);
  return;
}

func createSet(numbers []int) (set map[int]string) {
  set = make(map[int]string);
  
  for _,v := range numbers {
    _,in := set[v];
    if !in {
      set[v] = ""; 
    }
  }
  return;
}


func main() {
  var n int;
  fmt.Print("un numero: ");
  fmt.Scanf("%d", &n);

  var numbers []int = extractAllDigits(n);
  set := createSet(numbers);
  
  for k,_ := range set {
    var value string;
    fmt.Printf("parola per %d ? ", k);
    fmt.Scanf("%s", &value);

    set[k] = value;
  }

  for i,n := range numbers {
    if i == len(numbers) - 1 {
      break;
    }
    fmt.Printf("%s - ", set[n]);
  }
  fmt.Println(set[numbers[len(numbers) - 1]]);
}
