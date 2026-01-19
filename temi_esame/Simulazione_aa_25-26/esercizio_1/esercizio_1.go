package main
import "fmt"
import "unicode"
import "strconv"
import "math"
import "os"

func estraiCoefficienti(eq string) (int, int, int) {
  var coeff []string;
  var digit string = "";
  for i,r := range eq {
    if  i > 0 && eq[i - 1] == '^' {
      i++;
      continue;
    }
    s := string(r);
    if s == "-" || unicode.IsDigit(r) {
      digit += s;
    } else {
      if digit != "" {
        coeff = append(coeff, digit);
        digit = "";
      }
    }
  }
  a,_ := strconv.Atoi(coeff[0]);
  b,_ := strconv.Atoi(coeff[1]);
  c,_ := strconv.Atoi(coeff[2]);
  
  return a,b,c;
}

func soluzioni(a int, b int, c int) ([]float64, int) {
  var a_float, b_float, c_float = float64(a), float64(b), float64(c);
  var delta float64 = math.Pow(b_float, 2) - 4 * a_float * c_float;
  if delta < 0 {
    return []float64{}, 0;
  } else if delta == 0 {
    s := -1 * b_float / 2 * a_float;
    return []float64{s}, 1;
  } else {
   s1 := (-1 * b_float + math.Sqrt(delta)) / (2 * a_float);
   s2 := (-1 * b_float - math.Sqrt(delta)) / (2 * a_float);
   return []float64{s1, s2}, 2;
  }
}

func main() {
  eq := os.Args[1];
  soglia,_ := strconv.ParseFloat(os.Args[2], 64);
  epsilon,_ := strconv.ParseFloat(os.Args[3], 64);
  
  a,b,c := estraiCoefficienti(eq);
  soluzioni,n_soluzioni := soluzioni(a,b,c);

  if n_soluzioni == 2 {
    fmt.Printf("Esistono %d soluzioni reali: %.3f e %.3f\n", n_soluzioni, soluzioni[0], soluzioni[1]);
  } else if n_soluzioni == 1 {
    fmt.Printf("Esiste un’unica soluzione reale: %.3f\n", soluzioni[0]);
  } else {
    fmt.Println("Non ci sono soluzioni reali");
  }

  var valid []float64;
  if len(soluzioni) >= 1 && soluzioni[0] > soglia + epsilon {
    valid = append(valid, soluzioni[0]);
  }
  if len(soluzioni) >= 2 && soluzioni[1] > soglia + epsilon {
    valid = append(valid, soluzioni[1]);
  }
  for _,n := range valid {
    fmt.Printf("La soluzione %.3f è maggiore della soglia\n", n);
  }
}
