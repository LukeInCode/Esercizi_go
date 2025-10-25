package main
import "fmt"

func determine_seed(val int) int {
  const EXCHANGE = 13;
  
  return val / EXCHANGE;
}

func determine_value(val int) int {
  const EXCHANGE = 13;
  
  return val % EXCHANGE;
}


func main() {
  var value int;

  fmt.Print("Inserire una carta del mazzo [0-51]: "); _,err := fmt.Scan(&value);

  if err != nil || value < 0 || value > 51{
    fmt.Println("Input non valido");
  }else {
    var seed rune;
    var val string;
    
    switch determine_seed(value) {
      case 0:
        seed = '\U00002665';
      case 1:
        seed = '\U00002666';
      case 2:
        seed = '\U00002667';
      case 3:
        seed = '\U00002660';
      default:
        fmt.Println("errore");
    }
    
    fmt.Printf("La carta numero %d corrisponde al ",value);
    switch determine_value(value) {
      case 0:
        val = "A";
      case 10:
        val = "J";
      case 11:
        val = "q";
      case 12:
        val = "K";
      default:
        fmt.Print(value);
    }
    fmt.Printf("%s di %s\n",val,string(seed));
  }
}
