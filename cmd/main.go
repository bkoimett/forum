package main

import (
	"fmt"
	"net/http"

	
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w,r)
	} else {
	w.Write([]byte(`
		<h1></h1>
		<a href="/login">Login</a> | <a href="/register">Register</a>	
	`))
	}


}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`
        <form action="/login" method="POST">
            <input type="text" placeholder="Email" name="email">
            <input type="password" placeholder="Password" name="password">
            <button type="submit">Login</button>
        </form>		
	`))
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    w.Write([]byte(`
        <form action="/register" method="POST">
            <input type="text" placeholder="Username" name="username">
            <input type="email" placeholder="Email" name="email">
            <input type="password" placeholder="Password" name="password">
            <button type="submit">Register</button>
        </form>
    `))
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/register", registerHandler)

	fmt.Println("server running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}