package main

import (
	"fmt"
	"log"
	"net/http"
	
)

func main(){
	
     
	fileServer := http.FileServer(http.Dir("./static"))  //used to checkout the static file by default it only
	                                                     // check the index.html

	http.Handle("/",fileServer)    //home page 
    http.HandleFunc("/hello",helloHandler)  //check hello function
	http.HandleFunc("/form",formHandler)    // check form route present in static
    
	fmt.Printf("Go server running on port : 3000")
	if err := http.ListenAndServe(":3000",nil); err != nil{
		log.Fatal("err")
	}
}


func helloHandler(w http.ResponseWriter, r *http.Request){
      // r is pointing towards the request

	  if r.URL.Path != "/hello"{
		http.Error(w,"404 Not Found",http.StatusNotFound)
		return
	  }

	  if r.Method != "GET"{
		http.Error(w,"correct method required",http.StatusBadGateway)
	  }

	  fmt.Fprintf(w,"Hello func worikng fine")
}

func formHandler(w http.ResponseWriter, r *http.Request){
     
	    if err := r.ParseForm(); err != nil{
            fmt.Fprintf(w,"%v",err)
			return
		}
		fmt.Fprintf(w,"Post Request Successful\n")
		name := r.FormValue("name")
		email := r.FormValue("email")

		fmt.Fprintf(w,"Name : %s\n",name)
		fmt.Fprintf(w,"Email : %s\n",email)
}