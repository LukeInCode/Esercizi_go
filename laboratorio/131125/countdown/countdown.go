package main
import "fmt"
import "time"

type Clock struct {
  hour,min,sec int
};


func countdown(ct *Clock) {
    if (*ct).sec == 0 {
      downMin(ct);
    }else {
      (*ct).sec--;
    }
}

func downMin(ct *Clock) {
  if (*ct).min == 0 {
    downHour(ct);
  }else {
    (*ct).min--;
  }
  (*ct).sec = 59;
}

func downHour(ct *Clock) {
  (*ct).hour--;
  (*ct).min = 59;
}


func main() {
  var h,m,s int;
  fmt.Print("time (hh mm ss): ");
  fmt.Scan(&h, &m, &s);

  var ct Clock = Clock{h, m, s};
  for ct.hour > 0 || ct.min > 0 || ct.sec > 0 {
    countdown(&ct);
    fmt.Println(ct);
    time.Sleep(time.Duration(1) * time.Second);
  }
}

