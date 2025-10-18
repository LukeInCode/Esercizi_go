package main
import "fmt"

func main() {
  var x,y float64;

  _,err := fmt.Scan(&x,&y);

  if err == nil {
    var primo_quad, secondo_quad, terzo_quad, quarto_quad bool = x > 0 && y > 0, x < 0 && y > 0, x < 0 && y < 0, x > 0 && y < 0;

    if primo_quad {
      fmt.Printf("Il punto (%.2f,%.2f) si trova nel primo quadrante\n",x,y);
    }else if secondo_quad {
      fmt.Printf("Il punto (%.2f,%.2f) si trova nel secondo quadrante\n",x,y);
    }else if terzo_quad {
      fmt.Printf("Il punto (%.2f,%.2f) si trova nel terzo quadrante\n",x,y);
    }else if quarto_quad {
      fmt.Printf("Il punto (%.2f,%.2f) si trova nel quarto quadrante\n",x,y);
    }else {
      if x == 0 && y < 0 {
        fmt.Printf("Il punto (%.2f,%.2f) si trova nel quarto quadrante\n",x,y);
      }else if x == 0 && y > 0 {
        fmt.Printf("Il punto (%.2f,%.2f) si trova nel primo quadrante\n",x,y);
      }else if y == 0 && x < 0{
        fmt.Printf("Il punto (%.2f,%.2f) si trova nel secondo quadrante\n",x,y);
      }else if y == 0 && x > 0 {
        fmt.Printf("Il punto (%.2f,%.2f) si trova nel primo quadrante\n",x,y);
      }else {
        fmt.Printf("Il punto (%.2f,%.2f) si trova nell'origine\n",x,y);
      }
    }
  }
}
