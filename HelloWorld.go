package main

import(

  "./web"
  "os"
  "strconv"
)

var(
   addr = 9999
   listenaddr string
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
} 

func hello(val string) string{
   return "hello " + val
}

func main(){
   web.Get("/(.*)", hello)
   web.Run("127.0.0.1:" + listenaddr)
}
