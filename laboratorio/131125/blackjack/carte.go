package main
import "fmt"
import "math/rand"
import "strconv"

const LEN_SEMI = 4;
const VALORI = 13;
const DECK = 52;

var semi = [LEN_SEMI]string{ "\u2665", "\u2666", "\u2663", "\u2660" }
var valori = [VALORI]string{ "A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K" }

type Carta struct {
  valore, seme string;
};

func getValoreBJ(c *Carta) int {
  switch (*c).valore {
    case "J","K","Q":
      return 10;
    case "A":
      return 11;
    default:
      v,_ := strconv.Atoi( (*c).valore );
      return v;
  }
}

func carta(n int) (Carta, bool) {
  var index int;
  var c Carta;
  var trovato bool = false;
  
  if n < 0 || n >= 52 {
    return c,false;
  }
  for i := 0; i < LEN_SEMI && !trovato; i++ {
    for j := 0; j < VALORI; j++ {
      if index == n {
        c.seme = semi[i];
        c.valore = valori[index % VALORI];
        trovato = true;
        break;
      }
      index++;
    }
  }
  return c, true;
}

func estraiCarta() Carta {
  c,_ := carta(rand.Intn(DECK));
  return c;
}

func dai4Carte() *[4]Carta {
  var c [4]Carta;
  for i := 0; i < 4; i++ {
    c[i] = estraiCarta();
  }
  return &c;
}

func mazzoPoker() []Carta {
  var deck []Carta;
  deck = make([]Carta, DECK)
  var index int;
  
  for i := 0; i < LEN_SEMI; i++ {
    for j := 0; j < VALORI; j++ {
      deck[index] = Carta{semi[i], valori[index % VALORI]};
      index++;
    }
  }
  return deck;
}

func mischia(deck *[]Carta) {
  for i,_ := range (*deck) {
    target := rand.Intn(DECK - i);
    (*deck)[i],(*deck)[target] = (*deck)[target],(*deck)[i]
  }
}


func main() {
  var v string;
  fmt.Print("Inserire il valore: ");
  fmt.Scan(&v);

  var c Carta = Carta{v, ""};
  fmt.Println("valore bj",getValoreBJ(&c));

  var deck []Carta = mazzoPoker();
  fmt.Println("mazzo completo:",deck);

  mischia(&deck);
  fmt.Println("Mazzo mischiato:\n",deck);
  
  fmt.Println(estraiCarta());
  fmt.Println(*dai4Carte());
}
