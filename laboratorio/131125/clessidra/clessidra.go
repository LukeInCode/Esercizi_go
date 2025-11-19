package main
import "fmt"
import "strings"
import "os"
import "os/exec"
import "time"


func rigaClessidra(length int, sand bool, sandChar byte) (s string) {
  if sand {
    s = strings.Repeat(string(sandChar), length);
  }else {
    s = strings.Repeat(" ",length);
  }
  return s;
}


func cancella() {
   cmd := exec.Command("clear")
   cmd.Stdout=os.Stdout
   cmd.Run()
}


func attendi(n_seconds float64){
   time.Sleep(time.Duration(n_seconds) * time.Second)
}


func clessidra(height int, time int, char byte) {
  var BASE = height * 2 + 1;
  var sand bool = true;
  var spazi_interni int;

  var j int = BASE - 2;
  for i := height; i > 0; i-- { // Ciclo per la parte superiore della clessidra
    fmt.Print(strings.Repeat(" ", height - i));
    fmt.Print("\\");
    if i == 1 {
      spazi_interni = 1;
    }else {
      spazi_interni = j;
      j -= 2;
    }
    fmt.Print(rigaClessidra(spazi_interni,sand,char));
    fmt.Println("/");
  }

  var c int = 3;
  for i := height - 1; i >= 0; i-- { // Ciclo per la parte inferiore della clessidra
    fmt.Print(strings.Repeat(" ", i));
    fmt.Print("/");
    if i == height - 1 {
      spazi_interni = 1;
    } else {
      spazi_interni = c;
      c += 2;
    }
    fmt.Print(rigaClessidra(spazi_interni,false,char));
    fmt.Println("\\");
  }
}


func main() {
  var secondi, counter int;
  fmt.Scan(&secondi);
  
  //cancella();
  clessidra(secondi, counter, '*');
  counter++;
  //attendi(1);
}
