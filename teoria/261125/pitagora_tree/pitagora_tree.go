package main

import "github.com/holizz/terrapin"
import "math"
import "image"
import "image/color"
import "image/png"
import "os"
import "strconv"


func pitagora_tree(t *terrapin.Terrapin, lung float64, liv int) {
    if liv == 0 {
      t.Forward(lung);
    } else {
      const ang = math.Pi / 4;
      t.Forward(lung);
      pos, or := t.Pos, t.Orientation;

      t.Right(ang);
      pitagora_tree(t, lung * math.Cos(ang), liv - 1);

      t.Pos, t.Orientation = pos, or;

      t.Left(ang);
      pitagora_tree(t, lung * math.Cos(ang), liv - 1);
    }
}




func main() {
  nomeFile := os.Args[1];
  lung, _ := strconv.ParseFloat(os.Args[2], 64);
  liv, _ := strconv.Atoi(os.Args[3]);

  // Immagine vuota
  img := image.NewRGBA(image.Rect(0, 0, 900, 900));
  for x := 0; x < 900; x++ {
    for y := 0; y < 900; y++ {
      img.Set(x, y, color.RGBA{255, 255, 255, 255});
    }
  }
  
  // Tartaruga
  t := terrapin.NewTerrapin(img, terrapin.Position{450.0, 450.0});
  t.Color = color.RGBA{26, 151, 240, 250};
  pitagora_tree(t, lung, liv);

  // Salvo
  f, _ := os.Create(nomeFile);
  png.Encode(f, img);
  f.Close();
}
