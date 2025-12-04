package main
import "fmt"
import "os"

func generateAllAnagrams(chs []rune, mp map[string]bool, index int) {
    if index == len(chs) {
        mp[string(chs)] = true;
        return;
    }
    for i := index; i < len(chs); i++ {
        chs[index], chs[i] = chs[i], chs[index];
        generateAllAnagrams(chs, mp, index + 1);
        chs[index], chs[i] = chs[i], chs[index];
    }
}

func isAnagram(s1, s2 string) bool {
  possbile_anagrams := make(map[string]bool);
  generateAllAnagrams([]rune(s1), possbile_anagrams, 0);
  return possbile_anagrams[s2];
}


func main() {
  if len(os.Args) != 3 {
    fmt.Println("input errato");
  } else {
    var s1, s2 string = os.Args[1], os.Args[2];

    if isAnagram(s1, s2) == true {
      fmt.Printf("%s e %s sono anagrammi\n", s1, s2);
    }else {
      fmt.Printf("%s e %s non sono anagrammi\n", s1, s2);
    }
  }
}
