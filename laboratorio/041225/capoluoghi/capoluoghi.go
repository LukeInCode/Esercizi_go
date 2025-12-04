package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "strconv"


type Capoluogo struct {
  Nome, Sigla, Regione string;
  Superficie int;
};


func main() {
  mp := make(map[string]Capoluogo);

  scanner := bufio.NewScanner(os.Stdin);
  if scanner.Scan() {
    scanner.Text();
  }
  for scanner.Scan() {
    fields := strings.Split(scanner.Text(), ",");
    superficie,_ := strconv.Atoi(fields[3]);
    mp[fields[1]] = Capoluogo{fields[0], fields[1], fields[2], superficie};
  }

  if len(os.Args) > 1 {
    for i := 1; i < len(os.Args); i++ {
      fmt.Println(mp[os.Args[i]]);
    }
  }
}
