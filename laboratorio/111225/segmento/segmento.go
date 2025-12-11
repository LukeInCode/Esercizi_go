package main
import "fmt"

type Segmento struct {
  estremi, interno byte;
  orizzontale bool;
  lunghezza int;
}

func (s Segmento) String() string {
  if s.lunghezza < 2 {
    return "";
  }
  var res string;
  for i := 0; i < s.lunghezza; i++ {
    if i == 0 || i == s.lunghezza -1 {
      res += string(s.estremi);
    } else {
      res += string(s.interno);
    }
    if !s.orizzontale && i < s.lunghezza - 1 {
      res += "\n";
    }
  }
  return res;
}

func allunga(s *Segmento, n int) {
  if n < 0 {
      s.lunghezza += n;
  } else {
      s.lunghezza = n;
  }
}


func main() {
  var estremi, interno rune;
  var orientamento bool;
  var lung int;

  fmt.Scanf("%c %c",&estremi, &interno);
  fmt.Scan(&orientamento, &lung);

  var segmento Segmento = Segmento{byte(estremi), byte(interno), orientamento, lung};
  fmt.Println(segmento);
  
  allunga(&segmento, segmento.lunghezza + 3);
  fmt.Println(segmento);

  segmento.orizzontale = false;
  fmt.Println(segmento);
}
