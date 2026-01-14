package main
import "fmt"
import "strings"
import "strconv"
import "math"
import "os"
import "bufio"

type Wifi struct {
  ssid string;
  channel, frequency, signal_dBm int;
}

func testParams(ssid string, channel int, frequency int, signal_dBm int) bool {
  if ssid != "" && signal_dBm < -20 {
    if frequency >= 2412 && frequency <= 2484 {
        if channel >= 1 && channel <= 14 {
          return true;
        }
        return false;
    } else if frequency >= 5035 && frequency <= 5980 {
        if channel >= 7 && channel <= 196 {
           return true;
        }
        return false;
    } else {
      return false;
    }
  }
  return false;
}

func stringToInt(params []string) []int {
  var numbers []int;
  for _,v := range params {
    n,_ := strconv.Atoi(v);
    numbers = append(numbers, n);
  }
  return numbers;
}

func NewWifi(ssid string, channel int, frequency int, signal_dBm int) (Wifi, bool) {
  if testParams(ssid, channel, frequency, signal_dBm) {
    return Wifi{ ssid: ssid, channel: channel, frequency: frequency, signal_dBm: signal_dBm }, true;
  }
  return Wifi{}, false;
}

func NewWifiDaStringa(line string) (Wifi, bool) {
  params := strings.Split(line, ",");
  int_params := stringToInt(params[1:]);
  return NewWifi(params[0], int_params[0], int_params[1], int_params[2]);
}

func ConvertiDBinWatt(signal_dBm int) float64 {
  return math.Pow(10, float64(signal_dBm) / 10) / 1000;
}

func PiuPotente(elenco []Wifi, banda string) int {
  var max int = elenco[0].signal_dBm;
  var check bool = true;
  for _,s := range elenco[1:] {
    check = false;
    switch banda {
      case "2GHz":
        if s.frequency >= 2412 && s.frequency <= 2484 {
          check = true;
        }
      case "5GHz":
        if  s.frequency >= 5035 && s.frequency <= 5980 {
          check = true;
        }
      default:
        check = true;
    }
    if check {
      if s.signal_dBm > max {
        max = s.signal_dBm;
      }
    }
  }
  return max;
}

func (w Wifi) String() string {
  return fmt.Sprintf("{%s,%d,%d,%d,%v}", w.ssid, w.channel, w.frequency, w.signal_dBm, ConvertiDBinWatt(w.signal_dBm));
}


func main() {
  var nome_file, banda string = os.Args[1], "";
  if len(os.Args) >= 3 {
    banda = os.Args[2];
  }
  var reti []Wifi;
  file,_ := os.Open(nome_file);
  scanner := bufio.NewScanner(file);

  for scanner.Scan() {
    rete, err := NewWifiDaStringa(scanner.Text());
    if err {
      reti = append(reti, rete);
    }
  }

  segnale := PiuPotente(reti, banda);
  var potente Wifi;
  for _,v := range reti {
    if v.signal_dBm == segnale {
      potente = v;
      break;
    }
  }
  fmt.Println(potente);
}

