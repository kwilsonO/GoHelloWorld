package main

import(

  "./web"
  //"database/sql"
  //_ "./go-sqlite3"
  "os"
  "strconv"
  "fmt"
  //"log"
)

var(
   addr = 9999
   listenAddr string
   //db driver.Conn
)

func init() {
   if p := os.Getenv("HTTP_PORT"); p != "" {
      var err error
      addr, err = strconv.Atoi(p)
      if err != nil {
         log.Fatal(err)
      }
    }
  listenAddr = fmt.Sprintf(":%d", addr)
  //db, err := sql.Open("sqlite3", "./local.db")
  //if err != nil {
   //  log.Fatal(err)
 // }
  //defer db.Close()


  
} 

func hello(val string) string{
   return "hello " + val
}

func main(){
   web.Get("/(.*)", hello)
   web.Run("127.0.0.1" + listenAddr)
}
