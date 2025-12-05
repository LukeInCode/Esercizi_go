package main
import "fmt"
import "os"
import "bufio"
import "strings"
import "strconv"

type Log struct {
  Date, Time, Method, Url      string;
  ResponseCode, Milliseconds   int;
}

func (l Log) TotalRequests(logs []Log) int { return len(logs) };

func (l Log) DifferentRequests(logs []Log) (mp map[string]int) {
  mp = make(map[string]int);
  for _,v := range logs {
    mp[v.Method]++;
  }
  return;
}

func (l Log) MostSearchedUrl(logs []Log) string {
  mp := make(map[string]int);
  var max int = -1;
  var url string;
  
  for _,v := range logs {
    mp[v.Url]++;
    if mp[v.Url] > max {
      max = mp[v.Url];
      url = v.Url;
    }
  }
  return url;
}

func (l Log) FindSlowestResponse(logs []Log) (log Log, ms int) {
  ms = -1;
  
  for _,v := range logs {
    if v.Milliseconds > ms {
      ms = v.Milliseconds;
      log = v;
    }
  }
  return;
}


func (l Log) AverageResponseTime(logs []Log) float64 {
  var sum int = 0;

  for _,v := range logs {
    sum += v.Milliseconds;
  }
  return float64(sum) / float64(len(logs));
}

func GetAllLogs(scanner *bufio.Scanner) (logs []Log) {
  for scanner.Scan() {
    row := scanner.Text();
    if strings.TrimSpace(row) == "" {
      continue;
    }
    fields := strings.Fields(row);
    respCode,_ := strconv.Atoi(fields[len(fields) - 2]);
    ms,_ := strconv.Atoi(strings.ReplaceAll(fields[len(fields) - 1],"ms", ""));
    logs = append(logs, Log{fields[0], fields[1], fields[2], fields[3],respCode, ms});
  }
  return;
}

func main() {
  if len(os.Args) == 1 {
    fmt.Print(
      "The params you passed are incorrect\n",
      "That's how you use the program: ./analisi_log <file_to_scan>\n",
      "The result of the scan will be in the file named res.txt\n\n",
    );
    os.Exit(1);
  }
  
  const TARGET_DEFAUT string = "res.txt";
  fin,errIn := os.Open(os.Args[1]);
  defer fin.Close();
  
  if errIn != nil {
    fmt.Println("Something got wrong");
    os.Exit(1);
  }

  fout, errOut := os.Create(TARGET_DEFAUT);
  defer fout.Close();
  
  if errOut != nil {
    fmt.Println("Something got wrong");
    os.Exit(1);
  }

  var reporter Log;
  scanner := bufio.NewScanner(fin);
  var logs []Log = GetAllLogs(scanner);
  errors := make(map[int]int);

  fmt.Fprintln(fout, "===== REPORT ANALISI FILE DI LOG =====\n\n");
  fmt.Fprintf(fout, "Richieste totali: %d\n", reporter.TotalRequests(logs));

  fmt.Fprintln(fout, "Metodi HTTP utilizzati: ");
  methods := reporter.DifferentRequests(logs);
  for k,v := range methods {
    fmt.Fprintf(fout, "\t%s: %d\n", k, v);
  }

  fmt.Fprintln(fout, "\nErrori:");
  for _,v := range logs {
     if v.ResponseCode > 400 {
      errors[v.ResponseCode]++;
    }
  }
  for k,v := range errors {
    fmt.Fprintf(fout, "\t%d: %d\n", k, v);
  }
  
  fmt.Fprintf(fout, "\nUrl maggiormente ricercata: %s\n", reporter.MostSearchedUrl(logs));

  _,slowest := reporter.FindSlowestResponse(logs);
  fmt.Fprintf(fout, "\nMaggior tempo di risposta: %d\n", slowest);
  fmt.Fprintf(fout, "\nTempo di risposta medio: %.3f\n", reporter.AverageResponseTime(logs));
}
