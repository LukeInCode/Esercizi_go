package main
import "fmt"
import "slices"
import "bufio"
import "os"
import "strings"

type StringSet map[string]struct{};
type EmptyStruct struct{};


func NewStringSet() StringSet {
 return make(StringSet);
}

func ToSlice(s StringSet) []string {
  var values []string;
  for k,_ := range s {
    values = append(values, k);
  }
  slices.Sort(values);
  return values;
}

func Union(set1, set2 StringSet) StringSet {
  union := make(StringSet);
  
  for k,v := range set1 {
    union[k] = v;
  }

  for k,v := range set2 {
    union[k] = v;
  }
  return union;
}

func Intersection(set1, set2 StringSet) StringSet {
  intersection := make(StringSet);

  for k,_ := range set1 {
    _, in := set2[k];
    if in {
      intersection[k] = EmptyStruct{};
    }
  }
  return intersection;
}

func Difference(set1, set2 StringSet) StringSet {
  diff := make(StringSet);

  for k,_ := range set1 {
    _, in := set2[k];
    if !in {
      diff[k] = EmptyStruct{};
    }
  }
  return diff;
}

func main() {
  a := NewStringSet();
  b := NewStringSet();
  
  scanner := bufio.NewScanner(os.Stdin);
  for scanner.Scan() {
    row := strings.Fields(scanner.Text());
    for _,v := range row {
      fmt.Println(strings.Split(v, ":"))
      set, value := strings.Split(v, ":")[0], strings.Split(v, ":")[1];
      if set == "a" || set == "A" {
        a[value] = EmptyStruct{};
      } else if set == "b" || set == "B" {
        b[value] = EmptyStruct{};
      }
    }
  }
  fmt.Println("set A:", ToSlice(a));
  fmt.Println("set B:", ToSlice(b));
  fmt.Println("unione:", ToSlice(Union(a,b)));
  fmt.Println("intersezione:", ToSlice(Intersection(a,b)));
  fmt.Println("differenza:", ToSlice(Difference(a,b)));
}
