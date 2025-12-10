package main
import "fmt"
import "strconv"
import "os"

func average(values []int) float64 {
  var sum int = 0;
  for _,v := range values {
    sum += v;
  }
  return float64(sum) / float64(len(values));
}

func median(values []int) float64 {
  l := len(values);
  if l % 2 != 0 {
    return float64(values[l / 2 ]);
  } else {
    right, left := int(float64(l) / 2 + 0.5), int(float64(l) / 2 - 0.5);
    return float64(values[left] + values[right]) / 2;
  }
}

func values_under_average(values []int, average float64) (res []int){
  for _,v := range values {
    if float64(v) < average {
      res = append(res, v);
    }
  }
  return;
}

func minimum_values(values []int) []int {
  var res []int = []int {values[0], values[1], values[2]};
  for i := 3; i < len(values); i++ {
    for j,min := range res {
      if values[i] <= min {
        res[j] = values[i];
        break;
      }
    }
  }
  return res;
}

func maximum_values(values []int) []int {
  var res []int = []int {values[0], values[1], values[2]};
  for i := 3; i < len(values); i++ {
    for j,min := range res {
      if values[i] >= min {
        res[j] = values[i];
        break;
      }
    }
  }
  return res;
}

func String(values []int) (s string) {
    for _,v := range values {
      sValue := strconv.Itoa(v);
      s += sValue + " ";
    }
    return;
}

func main() {
  var n int;
  var values []int;
  for {
    fmt.Print("temperatura: ");
    fmt.Scanf("%d", &n);
    if n == 999 {
      break;
    }
    values = append(values, n);
  }

  if len(values) == 0 {
    fmt.Println("nessuna temperatura");
    os.Exit(0);
  }
  
  avg := average(values);
  
  fmt.Println("media", avg);
  fmt.Println("mediana", median(values));
  fmt.Println("temperature sotto la media", len(values_under_average(values, avg)));
  if len(values) >= 3 {
    fmt.Println("le tre temperature minori", String(minimum_values(values)));
    fmt.Println("le tre temperature maggiori", String(maximum_values(values)));
  }
}
