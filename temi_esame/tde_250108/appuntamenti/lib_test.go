/*
	"libreria" di test per gli esami, attenzione a modificare questo file!
	(è hardlinked ovunque)

	TODO fattorizzare ulteriormente, ci sono ancora tante duplicazionie

	rev: ottobre 2025
*/

package main

import (
	"fmt"
	"io"
	"math"
	"net"
	"os"
	"os/exec"
	"path"
	"runtime"
	"strings"
	"testing"
	"time"
)

//var HEADER string = "\n" + strings.Repeat(" ", (diffwidth-64)/2) + "___   ---   ===   ^^^   ***   TEST   ***   ^^^   ===   ---   ___"
//var HEADER string = "\n\n\n" + "___   ---   ===   ^^^   ***   TEST   ***   ^^^   ===   ---   ___"

func Test_Compila(t *testing.T) {
	//fmt.Println(HEADER)
	//fmt.Println()
	//fmt.Print("Verifico compilazione... ")

	// assumo che il go si chiami come la dir e il test sia <nomesorgente>_test.go

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("*** errore nella lettura della directory corrente ***")
		t.Fail()
		return
	}
	/*
		if strings.Contains(wd, "trent") {
			return
		}
	*/
	//fmt.Println(wd)
	nomeexe := path.Base(wd) // strippato diventa nome eseguibile
	nomego := nomeexe + ".go"
	//nometest := nomeexe + "_test.go"

	fexe, err := os.Stat(nomeexe)
	if fexe == nil {
		fmt.Println("*** c'è qualche problema sul nome della directory o del file .go (non corrispondenza con le specifiche?) ***")
		fmt.Println(err)
		t.Fail()
		return
	}

	tExe := fexe.ModTime()
	//fmt.Println(nomeexe, tExe)

	fgo, _ := os.Stat(nomego)
	tGo := fgo.ModTime()
	//fmt.Println(nomego, tGo)

	//ftest, _ := os.Stat(nometest)
	//tTest := ftest.ModTime()
	//fmt.Println(nometest, tTest)

	if tGo.After(tExe) {
		fmt.Println("**************************************************************************")
		fmt.Println("*** ATTENZIONE! il sorgente non è stato compilato dopo le modifiche!!! ***")
		fmt.Println("**************************************************************************")
		t.Fail()
	}
	/* else {
		fmt.Println("L'eseguibile è AGGIORNATO")
	}*/
}

/* a tendere supportare anche mac (darwin)
 */
func Test_Linux(t *testing.T) {
	//t.Error("PROVA LOG")
	//fmt.Println(HEADER)
	//fmt.Println()
	//fmt.Print("Controllo sistema operativo...", runtime.GOOS)
	fmt.Print("Data e ora attuale: ", time.Now(), " - IPs: ")

	//fmt.Println("Indirizzi IP della macchina:")
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
				if ipNet.IP.To4() != nil {
					fmt.Print(ipNet.IP.String(), " ")
				}
			}
		}
	}
	fmt.Println()

	if runtime.GOOS != "linux" {
		fmt.Println()
		fmt.Println("*************************************************")
		fmt.Println("* ATTENZIONE! sistema operativo NON supportato! *")
		fmt.Println("*************************************************")
	}
	//fmt.Println("--------------------------------------")
}

/*
NUOVA base?
lancia stud e oracolo (il nome è hardcoded 'oracolo') e confronta

- BISOGNA lasciare nelle dir del tema il file `oracolo` eseguibile compilato dal nostro sorgente e impacchettarglielo nel tar
- si possono scrivere SIA i test classici che alcuni con questa `confronta`
- per ora vedi cancellaParole per un esempio d'uso
*/
func ConfrontaConOracolo(
	t *testing.T,
	progname string,
	filestdinput string, // se nome vuoto viene creato un contenuto a "NIENTE"
	args ...string) {

	//fmt.Println(HEADER)
	//fmt.Println()
	fmt.Println(">>> Questo test confronta l'output studente con l'output atteso")
	//fmt.Println()
	fmt.Println(">>> L'eseguibile da testare (", progname, ") deve essere stato compilato!")
	fmt.Println()

	const oracleExe = "./oracolo"

	mustExist := func(path string) {
		if _, err := os.Stat(path); err != nil {
			t.Fatalf("manca: %s", path)
		}
	}
	mustExist(oracleExe)
	mustExist(progname)

	//////////////////////////////////////////////////////
	fmt.Printf("/// Argomenti a linea di comando:\t%s\n", args)
	fmt.Printf("/// File per StdInput (se vuoto non era previsto stdin):\t%s\n", filestdinput)
	fmt.Println()
	fmt.Println("### eseguo diff...")
	fmt.Println()

	l1 := "studente"
	l2 := "oracolo"
	fmt.Println("[", l1, "]", strings.Repeat(" ", diffwidth-len(l1)-len(l2)-10), "[", l2, "]")

	fmt.Println(strings.Repeat("-", diffwidth))
	diffResult, diffCode := lib_RunAndDiff(filestdinput,
		progname,
		oracleExe,
		args...)

	fmt.Println(diffResult)
	fmt.Println(strings.Repeat("-", diffwidth))

	if diffCode != nil { //&& diffCode.ExitCode() > 0 {
		fmt.Println("+++ ERRORE >>> FAIL! differisce da output atteso, diff return:", diffCode)
		fmt.Printf("%T\n", diffCode)
		t.Fail()
	}

	//oracolo.Process.Kill()
	//studente.Process.Kill()
}

func lib_WrapStdin(filestdinput string) (stdin io.Reader, err error) {
	if len(filestdinput) == 0 {
		stdin = strings.NewReader("NIENTE") // dummy
		//newname = "<non era previsto input da stdin>"
	} else {
		stdin, err = os.Open(filestdinput)
		/*
			if err != nil {
				fmt.Println(">>> Non posso aprire file stdin:", filestdinput)
			}
		*/
	}
	return
}

func lib_intorno(a, b float64) bool {
	return math.Abs(a-b) < 10e-6
}

func lib_min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func lib_max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
// RunWithStdoutCapture esegue un programma con i suoi argomenti
// e restituisce stdout+stderr in una stringa.
func RunWithStdoutCapture(prog string, args ...string) (string, error) {
	// costruiamo: stdbuf -oL -eL prog arg1 arg2 ...
	cmdArgs := append([]string{"-oL", "-eL", prog}, args...)
	cmd := exec.Command("stdbuf", cmdArgs...)

	out, err := cmd.CombinedOutput()
	return string(out), err
}
*/

// runAndDiff esegue due comandi e ritorna il risultato di diff tra i loro stdout
func lib_RunAndDiff(filestdinput string, stud string, orac string, args ...string) (string, error) {
	// prep i due stdin (anche se non saranno usati veramente)
	stdin1, _ := lib_WrapStdin(filestdinput)
	stdin2, _ := lib_WrapStdin(filestdinput)

	/*
		if err1 != nil || err2 != nil {
			fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++ errore creazione stdin (avvisare docente)")
			return "ERRORE WRAP STDIN", nil
		}
	*/

	// esegui il primo comando
	c1 := exec.Command(stud, args...)
	c1.Stdin = stdin1
	out1, err := c1.CombinedOutput()
	defer c1.Wait()
	/*
		if err != nil {
			fmt.Printf("[STUDENTE exit code: %s (non è un test fallito se si termina il programma con un esplicito os.Exit)]\n", err)
		} // TODO in futuro fattorizzare meglio codice in modo che questi messaggi li stampi fuori dall'output del diff
	*/

	// esegui il secondo comando
	c2 := exec.Command(orac, args...)
	c2.Stdin = stdin2
	out2, err := c2.CombinedOutput()
	defer c2.Wait()
	/*
		if err != nil {
			fmt.Printf("[ORACOLO exit code: %s (non è un test fallito se si termina il programma con un esplicito os.Exit)]\n", err)
		}
	*/

	// crea due file temporanei
	f1, err := os.CreateTemp("", "cmd1-*.txt")
	if err != nil {
		return "****************** ERRORE CREAZIONE TMP", err
	}
	defer os.Remove(f1.Name())
	defer f1.Close()

	f2, err := os.CreateTemp("", "cmd2-*.txt")
	if err != nil {
		return "****************** ERRORE CREAZIONE TMP", err
	}
	defer os.Remove(f2.Name())
	defer f2.Close()

	// scrivi gli output nei file
	if _, err := f1.Write(out1); err != nil {
		return "****************** ERRORE CREAZIONE OUT", err
	}
	if _, err := f2.Write(out2); err != nil {
		return "****************** ERRORE CREAZIONE OUT", err
	}

	// lancia diff
	//ORIG diffCmd := exec.Command("diff", "-u", f1.Name(), f2.Name())

	// OK
	diffCmd := exec.Command("diff", "-b", "-y", "-W", fmt.Sprint(diffwidth), f1.Name(), f2.Name())

	// PROVA
	//diffCmd := exec.Command("cat", f1.Name(), f2.Name())

	output, err := diffCmd.CombinedOutput()

	//fmt.Println("+++++++++++++++++++++++++++++++++DEBUG:", string(output), err)

	return string(output), err
}
