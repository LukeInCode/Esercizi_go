package main
import "fmt"
import "strings"
import "bufio"
import "os"
import "strconv"


func main() {
  var price, quantity, discount, total float64;
  scanner := bufio.NewScanner(os.Stdin);
  
  for scanner.Scan() {
    row := scanner.Text();
    order := strings.Split(row, "#"); 
    
    
    for i,v := range order {
      value, _ := strconv.ParseFloat(v, 64);
      *[]*float64{&price, &quantity, &discount}[i] = value;
    }
    total += (price - (price * discount)) * quantity;
  }
  fmt.Println(total);
}
