
package main

import "github.com/go-martini/martini"

func main() {
  w := martini.Classic()
  w.Get("/", func() string {
    return "hello world"
  })
  w.Run()
}
