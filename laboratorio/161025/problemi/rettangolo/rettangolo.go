package main
import "fmt"
import "math"

func main() {
  var p1_x, p2_x, p3_x int;
  var p1_y, p2_y, p3_y int;

  fmt.Print("P1: "); _,err_p1 := fmt.Scan(&p1_x, &p1_y);

  fmt.Print("P2: "); _,err_p2 := fmt.Scan(&p2_x, &p2_y);
  
  fmt.Print("P3: "); _,err_p3 := fmt.Scan(&p3_x, &p3_y);

  if err_p1 == nil && err_p2 == nil && err_p3 == nil {
    var vertice_basso_x, vertice_basso_y int = int(math.Min(float64(p1_x),float64(p2_x))), int(math.Min(float64(p1_y),float64(p2_y)));
    var vertice_alto_x, vertice_alto_y int = int(math.Max(float64(p1_x),float64(p2_x))), int(math.Max(float64(p1_y),float64(p2_y)));
    
    if (p3_x < vertice_basso_x || p3_x > vertice_alto_x) || (p3_y < vertice_basso_y || p3_y > vertice_alto_y) {
      fmt.Println("esterno");
    }else if (p3_x > vertice_basso_x && p3_x < vertice_alto_x) && (p3_y > vertice_basso_y && p3_y < vertice_alto_y) {
      fmt.Println("interno");
    }else {
      fmt.Println("perimetrale");
    }
  }
}
