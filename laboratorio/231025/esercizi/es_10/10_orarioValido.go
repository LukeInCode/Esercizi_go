package main
import "fmt"

func main() {
  var check = true;

  for check {
    var h,m int;
    
    fmt.Print("Inserire un orario (ore e minuti): ");
    _,err := fmt.Scan(&h,&m);
    
    if err == nil {
      if (h >= 0 && h <= 23) && (m >= 0 && m <= 59) {
        fmt.Println("orario valido");
        check = false;
      }
    } 
  }
}
