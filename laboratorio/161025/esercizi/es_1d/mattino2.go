package main;
import "fmt";

func main() {
  var ora,minuto int;

  fmt.Print("Inserire un orario da verificare: ");
  _,err := fmt.Scan(&ora,&minuto);

  if err != nil {
    fmt.Println("Inserire solo interi!");
  }else {
    const MINUTI_ORE = 60;
    minutaggio := ora * MINUTI_ORE + minuto;
    if minutaggio >= 330 && minutaggio < 750 {
      fmt.Println("Mattino");   
    }
  }
}
