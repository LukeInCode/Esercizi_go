package main
import "fmt"

func main() {
 var s, tot string;
 
  for {
    fmt.Scan(&s);
    if s == "stop" {
      break;
    }
    for _,v := range s {
      if v != ':' && v != ',' && v != '.' && v != ';' && v != '?' && v != '!' {
        tot += string(v);
      }
    }
    tot += " ";
  }
  fmt.Println(tot);
}
