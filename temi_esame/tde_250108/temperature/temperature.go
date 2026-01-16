package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"
import "slices"

func maxFreq(temps map[int]int) []int {
  var max = -1;
  var keys []int;
  for k,v := range temps {
    if v > max {
      keys = append([]int{}, k);
      max = v;
      continue;
    } 
    if v == max {
      keys = append(keys, k);
    }
  }
  return keys;
}

func findAllOccurrences(file *os.File) map[int]int {
  occ_map := make(map[int]int);
  scanner := bufio.NewScanner(file);
  for scanner.Scan() {
    temps := strings.Fields(scanner.Text());
    for _,t := range temps {
      n,_ := strconv.Atoi(t);
      occ_map[n]++;
    }
  }
  return occ_map;
}

func printMap(temps map[int]int) {
  for k,v := range temps {
    fmt.Print(k, ":", v, " ");
  }
  fmt.Println();
}


func main() {
  if len(os.Args) <= 2 {
    fmt.Println("Mancano parametri");
  } else {
    file,err_file := os.Open(os.Args[1]);
    if err_file != nil {
      fmt.Println("File non esistente");
    } else {
      occ_map := findAllOccurrences(file);
      printMap(occ_map)
      
      most_frequent := maxFreq(occ_map);
      slices.Sort(most_frequent);
      fmt.Print("freq max: ");
      for _,v := range most_frequent {
        fmt.Printf("%d:%d ", v, occ_map[v]);
      }
      fmt.Println("\n---");

      for _,v := range os.Args[2:] {
        n,_ := strconv.Atoi(v);
        if freq,ok := occ_map[n]; !ok {
          fmt.Print(n, ":", 0, " ");
        } else {
          fmt.Print(n, ":", freq, " ");
        }
      }
      fmt.Println();
    }
  }
}
