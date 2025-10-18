package main
import "fmt"

func main() {
  var giorno, mese, anno int;
  var giorno_target, mese_target, anno_target = 1,1,1970;
  
  
  fmt.Print("Inserire rispettivamente gg mm aaaa: "); fmt.Scan(&giorno,&mese,&anno);

  var giorni_totali int = giorno - giorno_target;
  for i := 0; i < (anno - anno_target); i++ {
    for j := 1; j <= 12; j++ {
      if j == 11 || j == 4 || j == 6 || j == 9 {
        giorni_totali += 30;
      }else if j == 2 {
        anno_corrente := anno_target + i;
        if (anno_corrente % 100 == 0 && anno_corrente % 400 == 0) || anno_corrente % 4 == 0 {
          giorni_totali += 29;
        }else {
          giorni_totali += 28;
        }
      }else {
        giorni_totali += 31;
      }
    }
  }
  for i := 1; i <= (mese - mese_target); i++ {
    if i == 11 || i == 4 || i == 6 || i == 9 {
      giorni_totali += 30;
    }else if i == 2 {
      if (anno % 100 == 0 && anno % 400 == 0) || anno % 4 == 0 {
        giorni_totali += 29;
      }else {
        giorni_totali += 28;
      }
    }else {
      giorni_totali += 31;
    }
  }
  fmt.Printf("La distanza tra %d/%d/%d e %d/%d/%d è di %d giorni\n", giorno_target,mese_target,anno_target,giorno,mese,anno,giorni_totali);
}
