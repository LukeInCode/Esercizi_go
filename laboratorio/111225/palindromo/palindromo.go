package main
import "fmt"
import "os"

func isPalindrome(s string) bool {
	r := []rune(s);
	if len(r) <= 1 {
		return true;
	}
	if r[0] != r[len(r) - 1] {
		return false;
	}
	return isPalindrome(string(r[1 : len(r)-1]));
}

func main() {
  s := os.Args[1];
  if (isPalindrome(s) == true) {
    fmt.Printf("%s è palindroma\n", s);
  } else {
    fmt.Printf("%s non è palindroma\n", s);
  }
}
