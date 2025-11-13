package main
import "fmt"

type Punto struct {
  x, y float64
}

func funzUno(p1, p2 Punto) Punto {
  var p Punto
  p.x = (p1.x + p2.x) / 2
  p.y = (p1.y + p2.y) / 2
  return p
}

func newPunto(x, y float64) Punto {
  return Punto{x,y};
}

func specchiaPunto(p *Punto) {
  (*p).x = (*p).x * -1; 
}

func funzDue(p *Punto, horiz, vert float64) {
  (*p).x += horiz
  (*p).y += vert
}

func main() {
  var p1, p2 Punto

  fmt.Print("x e y di P1: ")
  fmt.Scan(&p1.x, &p1.y)

  fmt.Print("x e y di P2: ")
  fmt.Scan(&p2.x, &p2.y)

  fmt.Println(funzUno(p1, p2))

  funzDue(&p1, 2, 3)
  fmt.Println(p1)

  var px Punto = newPunto(10,20);
  fmt.Println("newPunto x:", px.x, "y:", px.y);
  specchiaPunto(&px);
  fmt.Println("Punto specchiato rispetto asse y, x:", px.x, "y:", px.y);
}


