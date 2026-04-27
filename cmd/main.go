package main

import (
    "html/template"
    "net/http"
    "time"
)

type Post struct {
    ID        int
    Title     string
    Content   string
    Author    string
    CreatedAt time.Time
}

type PageData struct {
    Title       string
    IsLoggedIn  bool
    Username    string
    Posts       []Post
    Error       string
    Success     string
}

var templates *template.Template

func init() {
    // Parse all templates at startup
    templates = template.Must(template.ParseGlob("templates/*.html"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    
    // Sample posts (you'll replace with database later)
    posts := []Post{
        {ID: 1, Title: "Welcome to the forum!", Content: "Feel free to introduce yourself.", Author: "Admin", CreatedAt: time.Now()},
        {ID: 2, Title: "Go programming tips", Content: "Remember to handle your errors!", Author: "Gopher", CreatedAt: time.Now()},
    }
    
	// bundle data to send to template
    data := PageData{
        Title:      "Forum Home",
        IsLoggedIn: false,
        Posts:      posts,
    }

	// render templates with post data    
    templates.ExecuteTemplate(w, "index.html", data)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        templates.ExecuteTemplate(w, "login.html", PageData{Title: "Login"})
        return
    }
    
    // Handle POST
    r.ParseForm()
    email := r.FormValue("email")
    password := r.FormValue("password")
    
    // Basic validation
    if email == "" || password == "" {
        data := PageData{Title: "Login", Error: "All fields are required"}
        templates.ExecuteTemplate(w, "login.html", data)
        return
    }
    
    // Here you would check against database
    // For now, accept any non-empty credentials
    http.SetCookie(w, &http.Cookie{
        Name:    "demo_user",
        Value:   email,
        Expires: time.Now().Add(24 * time.Hour),
        Path:    "/",
    })
    
    http.Redirect(w, r, "/", http.StatusSeeOther) // 303 redirect to ome
}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
    http.SetCookie(w, &http.Cookie{
        Name:    "demo_user",
        Value:   "",
        Expires: time.Now().Add(-1 * time.Hour),
        Path:    "/",
    })
    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func createPostHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "GET" {
        templates.ExecuteTemplate(w, "create-post.html", PageData{Title: "Create Post"})
        return
    }
    
    // Handle POST
    r.ParseForm()
    title := r.FormValue("title")
    content := r.FormValue("content")
    
    if title == "" || content == "" {
        data := PageData{Title: "Create Post", Error: "Title and content are required"}
        templates.ExecuteTemplate(w, "create-post.html", data)
        return
    }
    
    // Here you would save to database
    data := PageData{Title: "Create Post", Success: "Post created successfully!"}
    templates.ExecuteTemplate(w, "create-post.html", data)
}

func main() {
    // Serve static files
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
    
    // Register routes
    http.HandleFunc("/", homeHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/logout", logoutHandler)
    http.HandleFunc("/create-post", createPostHandler)
    
    println("Server starting on http://localhost:8080")
    println("Available routes:")
    println("  /            - Home page")
    println("  /login       - Login page")
    println("  /logout      - Logout")
    println("  /create-post - Create a new post")
    
    http.ListenAndServe(":8080", nil)
}