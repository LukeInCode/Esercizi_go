package main
import "fmt"
import "math"

type Point struct {
  x, y float64;
}

type Rectangle struct {
  pLL, pUR Point;
}


func NewPoint(x, y float64) Point {
  return Point{x, y};
}

func NewRectangle(p1, p2 Point) Rectangle {
  var pLL Point = Point{math.Min(p1.x, p2.x), math.Min(p1.y, p2.y)}; 
  var pUR Point = Point{math.Max(p1.x, p2.x), math.Max(p1.y, p2.y)};

  return Rectangle{ pLL: pLL, pUR: pUR };
}

func Rotate(r *Rectangle, dir byte) {
  var p4 Point = Point{r.pLL.x, r.pUR.y};
  width := r.pUR.x - r.pLL.x;
  height := r.pUR.y - r.pLL.y;

  switch dir {
    case 'L':
      r.pUR.x, r.pUR.y = r.pLL.x + (p4.x - r.pLL.x), r.pUR.x - p4.x;
      r.pLL.x = r.pLL.x - (p4.y - r.pLL.y);

    case 'R':
      r.pLL.y = r.pLL.y - width;
      r.pUR.x, r.pUR.y = r.pLL.x + height,  r.pUR.y - height;
  }
}

func (r Rectangle) String() string {
  return fmt.Sprintf("((%f,%f) (%f,%f))", r.pLL.x, r.pLL.y, r.pUR.x, r.pUR.y);
}


func main() {
  p1 := NewPoint(10.5, 6.9)
	p2 := NewPoint(4, 0)
	r := NewRectangle(p1, p2)
	fmt.Println(r)
	Rotate(&r, 'R')
	fmt.Println(r)
}
