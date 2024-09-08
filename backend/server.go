package main

// import (
// 	"fmt"
// 	"net/http"
// )

// // reating a simple net/http server

// func setupRoutes(){
// 	/*
// 	what handleFunc does is it takes a path and a function as arguments and maps the path to the function called handler in this case
// 	// the path is "/" and the function is handler and the handler is a function that takes a response writer and a request as arguments
// 	and returns nothing all it does is write "simple server" to the response writer

// 	the response writer is the object that the server uses to write the response to the client

// 	*/
// 	 var handler = func(w http.ResponseWriter, r *http.Request){
// 		fmt.Fprintf(w, "Ahmed;")
// 	 }
// 	http.HandleFunc("/", handler)
// }
