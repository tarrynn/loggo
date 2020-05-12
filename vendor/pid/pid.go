package pid

import (
  "error"
  "os"
  "strconv"
)

type Conf struct {
  path string
}

func CreatePidFile (path string) {
  f, err := os.Create(path)
  error.Check(err)
  f.WriteString(strconv.Itoa(os.Getpid()))
  f.Close()
}

func RemovePidFile (path string) {
  err := os.Remove(path)
  error.Check(err)
}