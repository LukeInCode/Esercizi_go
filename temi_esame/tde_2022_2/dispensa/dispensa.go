package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"


func AggiornaDispensa(dispensa map[string]float64, filename string) bool {
  file, err_file := os.Open(filename);
  if err_file != nil {
    return false;
  }
  var check bool = true;
  scanner := bufio.NewScanner(file);
  scanner.Text();
  for scanner.Scan() {
    params := strings.Split(scanner.Text(), ",");
    q,_ := strconv.ParseFloat(params[2], 64);
    if _, ok := dispensa[params[1]]; !ok {
      dispensa[params[1]] = 0;
    }
    if params[0] == "approv" {
      dispensa[params[1]] += q;
    }else {
      if dispensa[params[1]] - q < 0 {
        check = false;
        break;
      }
      dispensa[params[1]] -= q;
    }
  }
  return check;
}

func Rimanenza(dispensa map[string]float64, alimento string) float64 {
  if _,check := dispensa[alimento]; !check {
    return 0;
  }
  return dispensa[alimento];
}


func main() {
  dispensa := make(map[string]float64);
  check := AggiornaDispensa(dispensa, os.Args[1]);

  if len(os.Args) > 2 {
    for _,v := range os.Args[2:] {
      fmt.Printf("%s %.3f\n", v, Rimanenza(dispensa, v));
    }
  }else {
    if check {
      fmt.Println("file corretto");
    }else {
      fmt.Println("file non corretto");
    }
  }
}

