package main
import "fmt"
import "strings"

func cifra_parola(parola string, k int) (cifrato string) {
  cifrato = "";
  alfabeto := "abcdefghijklmnopqrstuvwxyz";
  for _,r := range parola {
    cifrato += string(alfabeto[(strings.Index(alfabeto, string(r)) + k) % len(alfabeto)]);
  }
  return;
}

func main() {
  var k int;
  var parola string;

  fmt.Print("Inserire la chiave e la parola da cifrare (tutta in minuscolo): ");
  _,err := fmt.Scan(&k, &parola);

  var check bool = true;
  for _,r := range parola {
    if r < 97 || r > 122 {
      check = false;
      break;
    }
  }
  if err != nil || k < 0 || !check{
    fmt.Println("parametri errati");
  }else {
    fmt.Println(cifra_parola(parola,k));
  }
}
