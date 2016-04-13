package main

import (
  "log"
  "net/http"
  "github.com/codegangsta/negroni"
)

func main()  {
  n := negroni.New(
    negroni.NewRecovery(),
    negroni.HandlerFunc(MyMiddleware),
    negroni.NewLogger(),
    negroni.NewStatic(http.Dir(".")),
  )
  n.Run(":8080")
}

func MyMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc)  {
  log.Println("logging")

  if r.URL.Query().Get("key") == "123" {
    next(w, r)
    return
  } else {
    http.Error(w, "not authorized", 401)
  }

  log.Println("logging again")
}
