package main
import "fmt"

func substitute_c_g(original string, target rune) string {
  const OFFSET = 32
  var result = "";
  
  for _,r := range original {
    if (r == 'C' || r == 'C' + OFFSET) {
      if r >= 97 {
        result += "g";
      }else {
        result += "G";
      }
    }else if (r == 'G' || r == 'G' + OFFSET){
      if r >= 97 {
        result += "c";
      }else {
        result += "C";
      }
    }else {
      result += string(r);
    }
  }
  return result;
}

func main() {
  var string_to_convert string;
  fmt.Print("Inserire una stringa (senza spazi) da convertire: "); 
  _,err := fmt.Scan(&string_to_convert);
  if(err != nil) {
    fmt.Println("Input non valido");
  }else {
    fmt.Println(substitute_c_g(string_to_convert,'u'));
  }
}
