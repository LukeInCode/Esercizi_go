package main
import "fmt"

func main() {
  var giorno, mese, anno int;
  var giorno_target, mese_target, anno_target = 1,1,1970;
  const GIORNI_ANNO = 365;
  const GIORNI_MESE = 31;
  
  fmt.Print("Inserire rispettivamente gg mm aaaa: "); fmt.Scan(&giorno,&mese,&anno);

  var giorni_totali int = (anno - anno_target) * GIORNI_ANNO;
  giorni_totali += (giorno - giorno_target) + ((mese - mese_target) * GIORNI_MESE);

  fmt.Printf("La distanza tra %d/%d/%d e %d/%d/%d è di %d giorni\n", giorno_target,mese_target,anno_target,giorno,mese,anno,giorni_totali);
  
}
