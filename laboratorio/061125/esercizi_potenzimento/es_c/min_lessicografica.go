package main
import "fmt"
import "io"


func main() {
  var parola, prima string;
  var counter, index int = 1, -1;
    
  for {
    _,err := fmt.Scan(&parola);
    if err == io.EOF {
      break;
    }
    if prima == "" || prima > parola {
      prima = parola;
      index = counter;
    }
    counter++;
  }
  fmt.Printf("La prima parola in ordine lessicografico è %s e si trova in posizione %d\n",prima,index);
}
