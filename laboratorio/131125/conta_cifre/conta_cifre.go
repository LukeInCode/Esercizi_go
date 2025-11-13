package main;
import "fmt";
import "unicode"

func contaCifre(s string, numCifre *[10]int) (haCifre bool) {
  for _,r := range s {
    if unicode.IsDigit(r) {
      n := r - '0';
      numCifre[n]++;
      haCifre = true;
    }
  }
  return;
}

func main() {
  var s string;
  var numCifre [10]int;
  var counter int;
  
  for {
    fmt.Scan(&s);
    if s == "stop" {
      break;
    }
    check := contaCifre(s,&numCifre);
    if check {
      counter++;
    }
  }
  fmt.Println(counter,"stringhe con cifre");
  fmt.Println("[0 1 2 3 4 5 6 7 8 9]");
  fmt.Println(numCifre);
}
