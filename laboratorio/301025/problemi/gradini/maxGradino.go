package main
import "fmt"
import "unicode"
import "time"

func estraiGradino(s string) (string, int) {
  var gradino string = "";
  var group, lastIndex int = 0, -1;
  var previous rune;
  
  for index,r := range s {
    newGroup := r != previous;

    if group == 1 {
      lastIndex = index;
    }
    
    if group == 2 && newGroup {
      break;
    }
    gradino += string(r);
    if newGroup {
      group++;
    }
    previous = r;
  }
  return gradino, lastIndex;
}

func main() {
  var seq string;
  var ch, previous rune;

  for {
    fmt.Scanf("%c", &ch);
    if !unicode.IsDigit(ch) {
      continue;
    }
    seq += string(ch);
    if (ch - '0') < (previous - '0') {
      break;
    }
    previous = ch;
  }

  var index int;
  var gr, max string;
  var cpSeq string = seq;
  
  for {
    gr,index = estraiGradino(cpSeq);
    if index >= len(seq) || index < 0 {
      break;
    }
    cpSeq = cpSeq[index:];

    if len(gr) > len(max) {
      max = gr;
    }
    time.Sleep(100 * time.Millisecond)
  }
  fmt.Println("Il gradino massimo è:",max);
}
