package writer

import (
    "os"
    "sync"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func WriteToFile(path string, channel <-chan string, wg *sync.WaitGroup) {
  defer wg.Done()
  f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
  check(err)
  defer f.Close()

  for a := 0; a < 6; a++{
		m, more := <-channel
		f.WriteString(m+"\n")
		if more == false {
      f.Sync()
			return
		}
	}
}