package main
import "fmt"
import "strings"
import "os"


func main() {
  csv := os.Args[1];
  
  csvParams := strings.Split(csv, ",");
  fmt.Println(csvParams);

  csvParamsComma := strings.SplitAfter(csv, ",");
  fmt.Println(csvParamsComma);


  fieldsParam := os.Args[2];

  fields := strings.Fields(fieldsParam);
  fmt.Println(fields);

    
  fmt.Println([]rune(os.Args[3]));


  fmt.Println([]rune(os.Args[4]));

  timeParam := os.Args[5];
  time := strings.Split(timeParam, ":");
  
  fmt.Println(time);
}
