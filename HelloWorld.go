package main

import(

  "github.com/hoisie/web"
) 

func hello(val string) string{
   return "hello " + val
}

func main(){
   web.Get("/(.*)", hello)
   web.Run("127.0.0.1:9999")
}
