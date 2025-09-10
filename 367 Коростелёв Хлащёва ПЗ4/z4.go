
package main

import (
 "fmt"
 "net/http"
 "sync"
 "time"
)

func main() {
 uri := []string{
  "https://www.avito.ru/kurgan/kvartiry/sdam",
  "https://web.telegram.org",
  "https://github.com",
  "https://www.codingexplorations.com",
  "https://yandex.ru/games/app/237704?",
 }

 jobs := make(chan string, len(uri))
 results := make(chan string, len(uri))

 var fignya sync.WaitGroup

 for i := 1; i <= 3; i++ {
  fignya.Add(1)
  go func(workerID int) {
   defer fignya.Done()

   for url := range jobs {
    fmt.Printf("%d проверяет: %s\n", workerID, url)
    

    client := http.Client{
     Timeout: 10 * time.Second,
    }

    resp, err := client.Get(url)
    var status string
    if err != nil {
     status = fmt.Sprintf("ОШИБКА: %v", err)
    } else {
     defer resp.Body.Close()
     status = fmt.Sprintf("Статус: %d %s", resp.StatusCode, http.StatusText(resp.StatusCode))
    }
    

    results <- fmt.Sprintf(" %d: %s - %s", workerID, url, status)
   }
  }(i) 
 }

 for _, url := range uri {
  jobs <- url
 }
 close(jobs)


 fignya.Wait()
  close(results)

 fmt.Println("\n URL:")
 for result := range results {
  fmt.Println(result)
 }
}