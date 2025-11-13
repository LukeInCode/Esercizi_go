/*
  1) La variabile mesi è di tipo array
  2) La variabile mesi serve ad avere i nomi di tutti i mesi ordinati
  3) La funzione num2mese serve ad ottenere il nome del mese a partire dal suo numero
*/
package main
import (
   "fmt"
   "strconv"
   "strings"
)

func main() {
   var dataGMA string
   fmt.Print("data nel formato giorno/mese/anno: ")
   fmt.Scan(&dataGMA)
   g, m, a := stringGMA2GMA(dataGMA)
   fmt.Println("giorno", g)
   fmt.Println("mese", num2mese(m))
   fmt.Println("anno", a)

}

func stringGMA2GMA(data string) (int, int, int) {
   var split []string = strings.Split(data,"/");
   var converted [3]int;
   for i := 0; i < len(converted); i++ {
      res,_ := strconv.Atoi(split[i]);
      converted[i] = res;
   }
   return converted[0],converted[1],converted[2]
}

func num2mese(m int) string {
   var mesi = [13]string{"", "gen", "feb", "mar", "apr", "mag", "giu", "lug", "ago", "set", "ott", "nov", "dic"}
   return mesi[m]
}

