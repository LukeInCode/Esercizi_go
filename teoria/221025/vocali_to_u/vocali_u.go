package main
import "fmt"


func substitute_vowels(original string, target rune) string {
  const OFFSET = 32
  var result = "";
  
  for _,r := range original {
    if (r == 'A' || r == 'A' + OFFSET) || (r == 'E' || r == 'E' + OFFSET) || 
       (r == 'I' || r == 'I' + OFFSET) || (r == 'O' || r == 'O' + OFFSET) || 
       (r == 'U' || r == 'U' + OFFSET) {
      
      if r >= 97 {
        result += "u";
      }else {
        result += "U";
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
    fmt.Println(substitute_vowels(string_to_convert,'u'));
  }
}
