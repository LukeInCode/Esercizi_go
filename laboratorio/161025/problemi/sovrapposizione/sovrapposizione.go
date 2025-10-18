package main
import "fmt"

func main() {
  var giorno_app1, oraI_app1, oraF_app1 int;
  var giorno_app2, oraI_app2, oraF_app2 int;

  fmt.Print("appuntamento 1 (gg, start, end): ");
  _,err_1 := fmt.Scan(&giorno_app1,&oraI_app1,&oraF_app1);
  
  fmt.Print("appuntamento 2 (gg, start, end): ");
   _,err_2 := fmt.Scan(&giorno_app2,&oraI_app2,&oraF_app2);

   if err_1 == nil && err_2 == nil {
    if giorno_app1 != giorno_app2 {
      fmt.Println("non si sovrappongono");
    }else {
      rangeI := oraI_app1 - oraF_app1;
      rangeII := oraI_app2 - oraI_app2;
      
      if rangeI == rangeII {
          if (oraI_app1 > oraI_app2 || oraI_app2 > oraI_app1) || (oraF_app1 > oraF_app2 || oraF_app2 > oraF_app1) {
            fmt.Println("si sovrappongono");
          } else {
            fmt.Println("non si sovrappongono");
          }
      } else if rangeI > rangeII {
        if (oraI_app2 >= oraI_app1 && oraI_app2 <= oraF_app1) || (oraF_app2 >= oraI_app1) {
          fmt.Println("si sovrappongono");
        }else {
          fmt.Println("non si sovrappongono");
        }  
      } else {
        if (oraI_app1 >= oraI_app2 && oraI_app1 <= oraF_app2) || (oraF_app1 >= oraI_app2) {
            fmt.Println("si sovrappongono");
          } else {
            fmt.Println("non si sovrappongono");
          }  
        }
      }
   }
}
