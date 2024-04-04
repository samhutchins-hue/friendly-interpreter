package main

import (
  "fmt"
  "os"
  "os/user"
  "friendly-interpreter/repl"
)

func main()  {
  user, err := user.Current()
  if err != nil {
    panic(err)
  }
  fmt.Printf("Hell %s! This is the Friendly programming language!\n",
    user.Username)
  fmt.Printf("Feel free to type in commands\n")
  repl.Start(os.Stdin, os.Stdout)
}
