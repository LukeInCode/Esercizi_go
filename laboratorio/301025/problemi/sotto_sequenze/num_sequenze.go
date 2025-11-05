package main
import "fmt"
import "io"

func main() {
  var r rune;

  var counter int = 0;
  const ASCII_1 = 49;
  var subseq bool = false;

  for {
    _,err := fmt.Scanf("%c", &r);
    if r == (ASCII_1 + 1) || err == io.EOF {
        break;
    }
    if r == ASCII_1 {
      subseq = true;
    }
    if subseq && r == (ASCII_1 - 1) {
      counter++;
      subseq = false;
    }
  }
  fmt.Printf("Le sottosequenze di zeri (0) sono: %d\n", counter);
}
