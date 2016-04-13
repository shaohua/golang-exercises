package main

import "fmt"

type Door struct {
  opened bool
}

func (d *Door) Open()  {
  d.opened = true
}

func (d *Door) Close()  {
  d.opened = false
}

type OpenCloser interface {
  Open()
  Close()
}

func main()  {
  door := Door{opened: true}
  fmt.Println("door", door)
  door.Close()
  fmt.Println("door", door)
}
