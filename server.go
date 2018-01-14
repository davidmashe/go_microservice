package main

import "github.com/go-martini/martini"
import "golang.org/x/crypto/bcrypt"
import (
  "fmt"
  "net/http"
)

func main() {
  app := martini.Classic()
  app.RunOnAddr(":3025")

  app.Post("/hash", func(w http.ResponseWriter, r *http.Request) string {

    password := r.FormValue("password")

    // fmt.Println("password:", password)

    pw_bytes := [] byte(password)

    hashed, err := bcrypt.GenerateFromPassword(pw_bytes, bcrypt.DefaultCost)

    if err != nil {
      return "{\"error\": \"" + "something went wrong in hashing" + "\"}"
    }

    // fmt.Println(string(hashed[:]))
    
    return "{\"hashed_password\": " + "\"" + string(hashed[:]) + "\"}"
  })

  app.Post("/compare", func(w http.ResponseWriter, r *http.Request) string {

    password := r.FormValue("password")
    hashed_pw := r.FormValue("hashed")

    // fmt.Println("password:", password, "hashed_pw:", hashed_pw)

    pw_bytes := [] byte(password)
    hashed_bytes := [] byte(hashed_pw)

    err := bcrypt.CompareHashAndPassword(hashed_bytes, pw_bytes)

    if err != nil {
      return "{\"match\": \"false\"}"
    }
    
    return "{\"match\": \"true\"}"
  })  

  app.NotFound(func() string{
  	return "not found"
  })

  app.Run()
}