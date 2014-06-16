package main

import(

  "./web"
) 

func hello(val string) string{
   return "hello " + val
}

func main(){
   web.Get("/(.*)", hello)
   web.Run("localhost:61074")
}
