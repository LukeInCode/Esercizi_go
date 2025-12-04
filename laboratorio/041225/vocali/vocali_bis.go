package main
import "fmt"
import "unicode"
import "bufio"
import "os"
import "strings"
import "slices"

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

  var keys []rune;
  for k,_ := range vowels {
    keys = append(keys, k);
  }
  slices.Sort(keys);
  
  for _, key := range keys {
    fmt.Printf("%s : %d\n", string(key), vowels[key]);
  }
}
