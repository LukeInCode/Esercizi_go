package main
import "fmt"

func main() {
  var anno int;

  fmt.Print("Inserire un anno da controllare: ");
  _,err := fmt.Scan(&anno);

  if err != nil || anno < 0 {
    fmt.Println("Valori Inseriti non validi!");
  }else {
    var bisestile bool = false;
    if anno % 100 == 0 {
      if anno % 400 == 0 {
        bisestile = true;
      }
    }else {
      if anno % 4 == 0 {
        bisestile = true;
      }
    }
    if bisestile {
      fmt.Println("L'anno",anno,"è bisestile");
    }else {
      fmt.Println("L'anno",anno,"non è bisestile");
    }
  }
}
