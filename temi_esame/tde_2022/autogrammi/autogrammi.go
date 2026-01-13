package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"

func excludeNumbers(s string) (result string) {
  words := strings.Fields(s);
  for _,v := range words {
    if !strings.ContainsAny(v, "0123456789") {
      result += v + " ";
    }
  }
  return;
}

func cleanWord(s string) string {
  r := strings.NewReplacer(
        ",", "",
        ":", "",
        ".", "",
        " ", "",
  );
  return r.Replace(s);
}

func findShortestWord(s string) (shortest string, l int) {
  var words []string = strings.Fields(s);
  shortest = cleanWord(words[0]);
  l = len(shortest);
    
  for i := 1; i < len(words); i++ {
    actual_word := cleanWord(words[i]);
    actual_len := len(actual_word);
    if actual_len < l {
      shortest = actual_word;
      l = actual_len;
    }
  }
  return;
}

func findLongestWord(s string) (longest string, l int) {
  var words []string = strings.Fields(s);
  longest = cleanWord(words[0]);
  l = len(longest);
      
  for i := 1; i < len(words); i++ {
    actual_word := cleanWord(words[i]);
    actual_len := len(actual_word);
    if actual_len > l {
      longest = actual_word;
      l = actual_len;
    }
  }
  return;
}

func CalcQuanteMinMax(frase string) (quante, min, max int) {
  if frase == "" {
    quante, min, max = 0, 0, 0;
    return;
  }
  
  noNumbers := excludeNumbers(frase);
  quante = len(strings.Fields(noNumbers));
  _,min = findShortestWord(noNumbers);
  _,max = findLongestWord(noNumbers);
  return;
}

func TrovaNumDopoKeyword(frase, keyWord string) (num int) {
  var index int = -1;
  fields := strings.Fields(frase);
  for i,v := range fields {
    if v == keyWord {
      index = i;
      break;
    }
  }
  if index == -1 {
    num = 0;
    return;
  }
  toConvert := cleanWord(fields[index + 1])
  num,_ = strconv.Atoi(toConvert);
  return;
}


func Autogramma(frase string) bool {
    key_words, key_min, key_max := "contiene:", "minima:", "massima:";
    verified_words, verified_shortest, verified_longest := CalcQuanteMinMax(frase);
    words, shortest, longest  := TrovaNumDopoKeyword(frase, key_words), TrovaNumDopoKeyword(frase, key_min), TrovaNumDopoKeyword(frase, key_max);

    if words != verified_words || shortest != verified_shortest || longest != verified_longest {
      return false;
    }
    return true;
}

func StampaAnalisiAutogramma(frase string) {
  key_words, key_min, key_max := "contiene:", "minima:", "massima:";
  verified_words, verified_shortest, verified_longest := CalcQuanteMinMax(frase);
  words, shortest, longest  := TrovaNumDopoKeyword(frase, key_words), TrovaNumDopoKeyword(frase, key_min), TrovaNumDopoKeyword(frase, key_max);

  fmt.Println("=== ", frase);
  fmt.Printf("minimo dichiarato %d contro minimo verificato %d\n", shortest, verified_shortest);
  fmt.Printf("massimo dichiarato %d contro massimo verificato %d\n", longest, verified_longest);
  fmt.Printf("numero parole dichiarato %d contro numero parole verificato %d\n", words, verified_words);

  if Autogramma(frase) {
    fmt.Println("onesto");
  }else {
    fmt.Println("bugiardo");
  }
}

func main() {
  if len(os.Args) <= 1 {
    fmt.Println("file name?");
  }else {
    file,err_file := os.Open(os.Args[1]);
    if err_file != nil {
      fmt.Println("file not found");
    }else {
      scanner := bufio.NewScanner(file);
      for scanner.Scan() {
        row := scanner.Text();
        if row != "" {
          StampaAnalisiAutogramma(row);
        }
      }
    }
  }
}
