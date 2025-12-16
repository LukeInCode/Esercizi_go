package main
import "fmt"
import "os"
import "strconv"

type Punto struct {
  x,y int;
}

func (p Punto) String() string {
  return fmt.Sprintf("(%d,%d)", p.x, p.y);
}

func riflessioneY(p Punto) Punto {
  return Punto{ p.x, p.y * -1 };
}

func main() {
  x,_ := strconv.Atoi(os.Args[1]);
  y,_ := strconv.Atoi(os.Args[2]);

  p1 := Punto{x,y};
  fmt.Println(p1);

  p2 := riflessioneY(p1);
  fmt.Println(p2);
}
