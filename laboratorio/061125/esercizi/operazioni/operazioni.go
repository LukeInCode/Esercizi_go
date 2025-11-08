package main
import "fmt"

func operazioniV1(n1, n2 int) (int, int, int) {
  return n1 + n2, n1 * n2, n1 - n2;
}

func operazioniV2(n1, n2 int) (somma int, prodotto int, differenza int) {
  somma = n1 + n2;
  prodotto = n1 * n2;
  differenza = n1 - n2;

  return;
}

func main() {
  var n1, n2 int;
  fmt.Scan(&n1,&n2);

  fmt.Println(operazioniV1(n1,n2));
  fmt.Println(operazioniV2(n1,n2));
}
