package main
import "fmt"
import "os"

func isLuhn(s string) bool {
	if len(s) != 16 {
		return false;
	}
	var sum int = 0;
	for i,v := range s {
		if v < '0' || v > '9' {
			sum = -1;
			break;
		}
		var n int = int(v - '0');
		if i % 2 == 0 {
			if n * 2 > 9 {
				sum += n * 2 - 9;
			} else {
				sum += n * 2;
			}
		} else {
			sum += n;
		}
	}
	if sum == -1 {
		return false;
	}
	if sum % 10 == 0 {
		return true;
	}
	return false;
}

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("nessun argomento");
	} else {
		var ok, not_ok int = 0, 0;
		for _,s := range os.Args[1:] {
			valid := isLuhn(s);
			if valid {
				fmt.Println(s, "valido");
				ok++;
			} else {
				fmt.Println(s, "non valido");
				not_ok++;
			}
		}
		fmt.Println("Validi:", ok);
		fmt.Println("Non validi:", not_ok);
	}
}
