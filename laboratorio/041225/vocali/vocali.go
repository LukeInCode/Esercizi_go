package main
import "fmt"
import "unicode"
import "bufio"
import "os"
import "strings"

func contaVocali(s string, vocali map[rune]int) {
  for _, r := range s {
    if strings.ContainsAny("aeiou", string(unicode.ToLower(r))) {
      vocali[r]++;
    }
  }
}

func main() {
  var vowels map[rune]int = make(map[rune]int);
  scanner := bufio.NewScanner(os.Stdin);

  for scanner.Scan() {
    contaVocali(scanner.Text(), vowels);
  }

  for k,v := range vowels {
    fmt.Printf("%s : %d\n", string(k), v);
  }
}



