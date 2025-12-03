package main
import "fmt"
import "slices"
import "bufio"
import "os"

const LEN_ALFABETO int = 26;
var alfabeto []string = []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"};

func aggiornaConteggi(s string, contaMinu *[LEN_ALFABETO]int) {
  for _,v := range s {
    index := slices.Index(alfabeto, string(v));
    if index != -1 {
      (*contaMinu)[index]++;
    }
  }
}


func main() {
  var contaMinu[LEN_ALFABETO]int;
  scanner := bufio.NewScanner(os.Stdin);

  for scanner.Scan() {
    aggiornaConteggi(scanner.Text(), &contaMinu);
  }
  for i,_ := range contaMinu {
    fmt.Println(alfabeto[i], contaMinu[i]);
  }
}
