package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "slices"

func dictSinonimi(f *os.File) map[string][]string {
  sinonimi := make(map[string][]string);
  scanner := bufio.NewScanner(f);
  
  for scanner.Scan() {
    campi := strings.Split(scanner.Text(), ":");
    sin := strings.Split(strings.ReplaceAll(campi[1], " ", ""),",");
    slices.Sort(sin);
    sinonimi[campi[0]] = sin;
  }

  return sinonimi;
}

func main() {
  if len(os.Args) != 3 {
    fmt.Println("2 file names, please");
    os.Exit(0);
  }
  file_sinonimi, err_sinonimi := os.Open(os.Args[1]);
  defer file_sinonimi.Close();
  
  if err_sinonimi != nil {
    fmt.Println("file 1 not found");
    os.Exit(0);
  }
  file_elab, err_elab := os.Open(os.Args[2]);
  defer file_elab.Close();
  
  if err_elab != nil {
    fmt.Println("file 2 not found");
    os.Exit(0);
  }
  mappa_sinonimi := dictSinonimi(file_sinonimi);
  scanner := bufio.NewScanner(file_elab);
  for scanner.Scan() {
    campi := strings.Fields(scanner.Text());
    for i,el := range campi {
      fmt.Print(el);
      if sin, check := mappa_sinonimi[el]; check == true {
        fmt.Printf("[%s]", strings.Join(sin, " "));
      }
      if i < len(campi) - 1 {
        fmt.Print(" ");
      }
    }
    fmt.Println();
  }
}
