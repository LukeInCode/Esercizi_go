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
  var B = height * 2 + 1;
  var lenght_ch int;
  var sand bool = true;
  
  for i := 0; i < height; i++ {
    lenght_ch = height - 2*i;
    if lenght_ch < 1 {
      lenght_ch = 0;
    }
    fmt.Print(strings.Repeat(" ",i));
    fmt.Print("\\");
    spaces := (B - i - 2 - lenght_ch) / 2;
    
    if spaces < 0 {
      spaces = 0;
    }
    if i == height - 1 {
      sand = false;
    }
    fmt.Print(strings.Repeat("-",spaces));
    fmt.Print(rigaClessidra(lenght_ch, sand, byte(char)));
    fmt.Print(strings.Repeat("-",spaces));
    fmt.Println("/");
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
