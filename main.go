package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
    "encoding/json"
    "strconv"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func main() {
    if os.Args[1] == "-h" || os.Args[1] == "--help" {
      fmt.Println("Example use:\ngo run main.go <exported .csv from takeout.google.com>")
      os.Exit(0)
   }
    fmt.Println("[+] converting .csv file " + os.Args[1])
    records := readCsvFile(os.Args[1])
    subs := sublist {
      App_version: "0.21.9",
      App_version_int: 975,
      Subscriptions: make([]sub, len(records)-1),
   }
    for i:=1; i < len(records); i++ {
      subs.Subscriptions[i-1].Service_id = 0
      subs.Subscriptions[i-1].Url = records[i][1]
      subs.Subscriptions[i-1].Name = records[i][2]
   }
   fmt.Println("[+] " + strconv.Itoa(len(records)-1) + " subscriptions parsed")

   jsonData, err := json.Marshal(subs)
   //fmt.Println(string(jsonData))
   if err != nil {
    log.Println(err)
   }

   f, err := os.Create("subscriptons.json")
   if err != nil {
    log.Println(err)
   }

   n3, err := f.WriteString(string(jsonData))
   if err != nil {
    log.Println(err)
   }

    f.Sync()
    fmt.Printf("[+] wrote %d bytes to subscriptions.json\n", n3)
    fmt.Printf("[+] Done, import subscriptions.json to NewPipe as \"Previous export\"\n")

}

type sublist struct {
   App_version string `json:"app_version"`
   App_version_int   int64 `json:"app_version_init"`
   Subscriptions []sub `json:"subscriptions"`
}

type sub struct {
   Service_id  int `json:"service_id"`
   Url   string `json:"url"`
   Name    string `json:"name"`
}
