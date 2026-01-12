package main
import "fmt"
import "strings"
import "math"
import "strconv"
import "bufio"
import "os"

type Cerchio struct {
  nome string;
  x,y,raggio float64;
}

func newCircle(descr string) Cerchio {
  flds := strings.Fields(descr);
  parsed_x,_ := strconv.ParseFloat(flds[1], 64);
  parsed_y,_ := strconv.ParseFloat(flds[2], 64);
  radius,_ := strconv.ParseFloat(flds[3], 64);
  return Cerchio{flds[0], parsed_x, parsed_y, radius};
}

func distanzaPunti(x1,y1,x2,y2 float64) float64 {
  return math.Sqrt(math.Pow((x2 - x1), 2) + math.Pow((y2 - y1), 2));
}


func tocca(cerc1, cerc2 Cerchio) bool {
  const TOLERANCE float64 = 1e-6;
  centers_distance := distanzaPunti(cerc1.x, cerc1.y, cerc2.x, cerc2.y);
  radius_sum := cerc1.raggio + cerc2.raggio;
  radius_diff := math.Abs(cerc1.raggio - cerc2.raggio);
  
  if centers_distance - radius_sum > -1 * TOLERANCE && centers_distance - radius_sum < TOLERANCE {
    return true;
  }
  if centers_distance - radius_diff > -1 * TOLERANCE && centers_distance - radius_diff < TOLERANCE {
    return true;
  }
  return false;
}

func maggiore(cerc1, cerc2 Cerchio) bool {
  return cerc1.raggio > cerc2.raggio;
}

func main() {
  scanner := bufio.NewScanner(os.Stdin);
  var circle Cerchio;
  var previous_circle Cerchio = Cerchio{"", 0, 0, 0};
  
  for scanner.Scan() {
    circle = newCircle(scanner.Text());
    fmt.Print(circle, " ");

    if tocca(circle, previous_circle) {
      fmt.Print("tangente, ");
    }else {
      fmt.Print("non tangente, ");
    }

    if maggiore(circle, previous_circle) {
      fmt.Println("maggiore");
    }else {
      fmt.Println("minore o uguale");
    }

    previous_circle = circle;
  }
}
