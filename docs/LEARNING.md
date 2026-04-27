## What I have learnt

** if you do not set Content-Type, the browser will fail to infer that it is a html doc

```go
// set content type
w.Header().Set("Content-Type", "text/html")
```

** response methods
- r.Method
- r.ParseForm() / r.FormValue("email)


** write methods
- 


** http methods
- http.Error(w, "bad request", http.StatusBadRequest)


-- directives / template actions

{{.Author}} - replaces it with the actual value from the data passed in


## packages

package main - Creates an executable program (not a library)

html/template - For rendering HTML templates with dynamic data

net/http - Provides HTTP client/server functionality

time - For handling dates, times, and cookies expiration

## template initialization

    templates - Global variable holding all parsed HTML templates

    init() - Runs automatically before main()

    template.Must() - Panics if error occurs (fail fast approach)

    template.ParseGlob() - Loads all .html files from the templates folder

    This allows using templates.ExecuteTemplate() with any template name

## homepage handler

## login handler

GET request: Show empty login form

POST request: Process submitted credentials

Form handling: Extract email/password from submitted form

Validation: Check fields aren't empty

Cookie creation: Store user session (temporary - should use proper auth)

Redirect: Send user back to homepage after login