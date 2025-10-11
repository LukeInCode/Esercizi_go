package main
import "fmt"

func main() {
  var km_percorsi, litri_usati float64;
  fmt.Print("Inserire i kilometri percorsi e i litri di carburante utilizzati: ");
  _,err := fmt.Scan(&km_percorsi, &litri_usati);

  if(err != nil) {
    fmt.Println("I valori passati non rispettano i requisiti");
  }else {
    var consumo,resa float64;
    
    consumo = litri_usati / km_percorsi;
    resa = km_percorsi / litri_usati;

    fmt.Printf("Il consumo medio di carburante è %.2f e la resa del motore è %.2f\n",consumo,resa);
  }
}
