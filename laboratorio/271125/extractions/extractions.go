package main

func estraiPari(in []int) (out []int) {
  for _, v := range in {
    if v % 2 == 0 {
      out = append(out, v);
    }
  }
  return;
}


func rimuoviMultipli(m int, in []int) (out []int) {
  out = make([]int, len(in));
  copy(out, in);
  for i, v := range in {
    if v % m == 0 {
      if i < len(out) - 1 {
        out = append(out[:i], out[i + 1:]...);
      }else {
        out = out[:i - 1];
      }
    }
  }
  return;
}


func main() {
}
