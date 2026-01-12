package main
import "fmt"
import "os"
import "bufio"
import "strconv"

func printList(l []string) {
  for _,el := range l {
    fmt.Println(el);
  }
}

func main() {
  n_rows,_ := strconv.Atoi(os.Args[1]);
  var file_name string = os.Args[2];
  
  f,_ := os.Open(file_name);
  defer f.Close();

  var rows []string;
  
  scanner := bufio.NewScanner(f);
  for scanner.Scan() {
    rows = append(rows, scanner.Text());
  }
  if n_rows > len(rows) {
    printList(rows);
  } else {
    index := len(rows) - n_rows;
    printList(rows[index:]);
  }
}
