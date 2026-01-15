package main
import "fmt"
import "bufio"
import "os"
import "strings"
import "slices"

func findOccurences(text string, word string) int {
  var occurences int = 0;
  for _,v := range strings.Fields(text) {
    if v == word {
      occurences++;
    }
  }
  return occurences;
}

func loadFile(file *os.File) (rows []string) {
  scanner := bufio.NewScanner(file);
  for scanner.Scan() {
    rows = append(rows, scanner.Text());
  }
  return rows;
}

func populateList(row int, n int) (l []int) {
  for j := 0; j < n; j++ {
    l = append(l, row); 
  }
  return l;
}

func getAllWords(rows []string) []string {
  words := make(map[string]bool);
  keys := make([]string, 0);
  for _,v := range rows {
    for _,w := range strings.Fields(v) {
      words[w] = true;
    }
  }
  for k,_ := range words {
    keys = append(keys, k);
  }
  slices.Sort(keys);
  return keys;
}

func populateOcc(rows []string, word string)  []int {
  occ_rows := make([]int, 0);
  for i := 0; i < len(rows); i++ {
      occs := findOccurences(rows[i], word);
      occ_rows = append(occ_rows, populateList(i + 1, occs)...);
  }
  return occ_rows;
}


func main() {
  if len(os.Args) <= 1 {
    fmt.Println("Fornire un nome di file");
  } else {
    file,err := os.Open(os.Args[1]);
    if err != nil {
      fmt.Println("File non accessibile");
    } else {
      scanner := bufio.NewScanner(os.Stdin);
      file_rows := loadFile(file);
      fmt.Println("RICERCHE:");
      var occ_rows []int;
      for scanner.Scan() {
        occ_rows = make([]int, 0);
        input := scanner.Text();
        if !strings.HasPrefix(input, "? ") {
          fmt.Println("richiesta non valida", input);
          continue;
        } 
        prompt := strings.Split(input, "? ")[1];
        for i := 0; i < len(file_rows); i++ {
          occs := findOccurences(file_rows[i], prompt);
          occ_rows = append(occ_rows, populateList(i + 1, occs)...);
        }
        fmt.Println("parola:", prompt);
        fmt.Println("righe", occ_rows);
      }
      fmt.Println("TESTO:");
      words := getAllWords(file_rows);
      for _,v := range words {
        fmt.Println(v, ":", populateOcc(file_rows, v));
      }
    }
  }
}
