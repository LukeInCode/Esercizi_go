package main
import "fmt"
import "unicode"


func hasUpper(s string) bool {
  var hasUpper bool = false;
  for _,r := range s {
    if unicode.IsUpper(r) {
      hasUpper = true;
      break;
    }
  }
  return hasUpper;
}

func firstUpper(s string) int {
  var pos int = -1;
  for i,r := range s {
    if unicode.IsUpper(r) {
      pos = i;
      break;
    }
  }
  return pos;
}

func lastUpper(s string) int {
  var pos int = -1;
  for i,r := range s {
    if unicode.IsUpper(r) {
      pos = i;
    }
  }
  return pos;
}

func countDigitsLettersOthers(s string) (int,int,int) {
  var digits,letters,others int;
  
  for _,r := range s {
    switch {
        case unicode.IsDigit(r):
          digits++;
        case unicode.IsLetter(r):
          letters++;
        default:
          others++;
    }
  }
  return digits,letters,others;
}

func isPalindrome(s string) bool {
  var backward string = "";
  for i := len(s) - 1; i >= 0; i-- {
    backward += string(s[i]);
  }
  return s == backward;
}

func main() {
  var s string;
  fmt.Scan(&s);

  var upper bool = hasUpper(s);
  if upper {
    fmt.Println("ha maiuscole");
  }else {
    fmt.Println("non ha maiuscole");
  }
  
  fmt.Println("prima maiuscola in posizione",firstUpper(s));
  fmt.Println("ultima maiuscola in posizione",lastUpper(s));

  var digits,letters,others int = countDigitsLettersOthers(s);
  fmt.Printf("cifre, lettere, altro: %d %d %d\n",digits,letters,others);
  
  var palindrome bool = isPalindrome(s);
  if palindrome {
    fmt.Println("palindroma");
  }else {
    fmt.Println("non palindroma");
  }
}
