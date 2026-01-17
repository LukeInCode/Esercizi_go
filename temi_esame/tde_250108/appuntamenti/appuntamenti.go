package main
import "fmt"
import "bufio"
import "os"
import "strconv"
import "strings"

type Appuntamento struct {
  giorno, oraInizio, oraFine int;
}

func NuovoAppuntamento(gg, h1, h2 int) (Appuntamento, bool) {
  if gg < 1 || gg > 366 {
    return Appuntamento{}, false;
  }
  if h1 > h2 {
    return Appuntamento{}, false;
  }
  if (h1 < 0 || h1 > 24) || (h2 < 0 || h2 > 24) {
    return Appuntamento{}, false;
  }
  return Appuntamento{ giorno: gg, oraInizio: h1, oraFine: h2 }, true;
}


func SovrapposizioneAppuntamenti(app1, app2 Appuntamento) bool {
  if app1.giorno != app2.giorno {
    return false;
  }
  return app1.oraInizio < app2.oraFine && app2.oraInizio < app1.oraFine;
}

func AggiungiAppuntamento(app Appuntamento, agenda *[]Appuntamento) bool {
  var check bool = true;
  for _,v := range *agenda {
    if SovrapposizioneAppuntamenti(app, v) {
      check = false;
      break;
    }
  }
  if check {
    *agenda = append(*agenda, app);
  }
  return check;
}

func (app Appuntamento) String() string {
  var temp_line = "";
  const HOURS int = 24;
  
  for i := 0; i < HOURS; i++ {
    if i >= app.oraInizio && i < app.oraFine {
      temp_line += "*";
      continue;
    }
    temp_line += "-";
  }
  return fmt.Sprintf("il giorno %.2d dalle %.2d alle %.2d: %s", app.giorno, app.oraInizio, app.oraFine, temp_line);
}

func main() {
  var agenda []Appuntamento;
  scanner := bufio.NewScanner(os.Stdin);

  for scanner.Scan() {
    var params []int;
    for _,s := range strings.Fields(scanner.Text()) {
      n,_ := strconv.Atoi(s);
      params = append(params, n);
    }
    app,err := NuovoAppuntamento(params[0], params[1], params[2]);
    if err {
      AggiungiAppuntamento(app, &agenda);
    }
  }
  for _,v := range agenda {
    fmt.Println(v);
  }
}
