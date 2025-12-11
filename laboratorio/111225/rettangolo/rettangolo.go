package main
import "fmt"
import "os"
import "strconv"

type Rettangolo struct {
  base, altezza int;
}

func (rect Rettangolo) String() string {
  if rect.base == 0 || rect.altezza == 0 {
    return "rettangolo degenere";
  }
  var rectangle string = "";
  for y := 0; y < rect.altezza; y++ {
    for x := 0; x < rect.base; x++ {
      rectangle += ".";
    }
    rectangle += "\n";
  }
  return rectangle;
}


func main() {
  base,_ := strconv.Atoi(os.Args[1]);
  altezza,_ := strconv.Atoi(os.Args[2]); 
  var rect Rettangolo = Rettangolo{base: base, altezza: altezza};
  fmt.Println(rect);
}
