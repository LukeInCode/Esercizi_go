package main
import "fmt"

func main() {
  var mese,giorno,anno int;

  fmt.Print("Inserire rispettivamente giorno, mese e anno: ");
  _,err := fmt.Scan(&giorno,&mese,&anno);

  if err != nil {
    fmt.Println("Valori inseriti non validi!");
  }else {
   const MAX_MESE = 12;
   var controllo_data bool = false;
    if mese > 0 && mese <= MAX_MESE && giorno > 0 && anno >= 0 {
      if mese == 2 {
        if giorno <= 28 {
          controllo_data = true;
        } else if giorno == 29 {
            if anno % 100 == 0 {
              if anno % 400 == 0 {
                controllo_data = true;
              } else {
                controllo_data = false;
              } 
            }else {
              if anno % 4 == 0 {
                controllo_data = true;
              } else {
                controllo_data = false;
              }
            }
        } else {
          controllo_data = false;
        }
      
      } else if mese == 4 || mese == 6 || mese == 11 || mese == 9  {
        if giorno <= 30 {
          controllo_data = true;
        }else {
          controllo_data = false;
        }
      }else {
        if giorno <= 31 {
          controllo_data = true;
        }else {
          controllo_data = false;
        }
      }
    }else {
      controllo_data = false;
    }
    if controllo_data {
      fmt.Println("Data valida");
    }else {
      fmt.Println("Data non valida");
    }
  }
}
