package main
import "fmt"
import "strconv"
import "strings"

func main() {
  var s1,s2 string = "s","sc";
  var str string;
  fmt.Print("Inserire una stringa: ");
  fmt.Scan(&str);

  if strings.Contains(str,s2) {
    fmt.Printf("%s � contenuta in %s\n",s2,s1);
  } else {
    fmt.Printf("%s non � contenuta in %s\n",s2,s1)
  }
  if strings.ContainsAny(strings.ToLower(str), "aeiou") {
    fmt.Printf("La stringa %s contiene almeno una vocale\n",s2);
    fmt.Printf("La prima vocale trovata si trova in posizione: %d\n",strings.IndexAny(str,"aeiou"));
    fmt.Printf("L'ultima vocale trovata si trova in posizione: %d\n",strings.LastIndexAny(str,"aeiou"));
  }else {
    fmt.Printf("La stringa %s non contiene vocali\n",str);
  }
  fmt.Printf("la stringa %s � contenuta %d volte in %s\n",s1,strings.Count(str,s1),str);
  fmt.Printf("La stringa %s concatenata con se stessa per tre volte �: %s\n",s2,strings.Repeat(s2,3));
  fmt.Printf("La stringa %s concatenata con se stessa per tre volte �: %s\n",s1,strings.Repeat(s1,5));

  var s string = "";
  for _,r := range str {
    ascii := r - '0';
    if ascii >= 0 && ascii <= 9 {
      s += string(r);
    }
  }
  if s != "" {
    fmt.Printf("La stringa %s contiene queste cifre: %s\n",str,s);
    i,_ := strconv.Atoi(s);
    fmt.Printf("Intero: %d\n", i);
    f,_ := strconv.ParseFloat(s,64)
    fmt.Printf("float %f\n", f)
  }
}

