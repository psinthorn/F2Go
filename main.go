package main

import "github.com/psinthorn/F2Go/app"

// Data type preparing

// var tmpl *template.Template

// func init(){
// 	tmpl = template.Must(template.ParseGlob("./views/welcome/*.html"))
// }

func main() {
	// Set PORT for local env
	// Comment this line when you production
	//os.Setenv("PORT", "8008")

	app.StartApp()
}

// func HandleRoutes() {
// 	// Port := "8008"
// 	port := os.Getenv("PORT")
// 	r := mux.NewRouter()
// 	r.HandleFunc("/", WelcomeIndex)
// 	r.HandleFunc("/user/{fname}/{lname}", UserIndex)
// 	r.HandleFunc("/todos", TodoIndex)
// 	r.HandleFunc("/about", AboutIndex)
// 	r.HandleFunc("/contact", ContactIndex)
// 	r.HandleFunc("/version", VersionIndex)

// 	fmt.Println("Server running on " + port)
// 	err := http.ListenAndServe(":"+port, r)
// 	if err != nil {
// 		log.Fatal("ListenAndServe:"+port, err)
// 	}

// }

// func Hello() string {
// 	return "Hello F2Code"
// }

//Welcome Index Page
// func WelcomeIndex(w http.ResponseWriter, r *http.Request) {

// 	//Used template
// 	tmpl := template.Must(template.ParseFiles("./views/welcome/index.html"))

// 	//Excute template and response
// 	tmpl.Execute(w, nil)

// }

// //User Index page
// func UserIndex(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	fname := vars["fname"]
// 	lname := vars["lname"]

// 	fmt.Fprintf(w, "Hello Myname: %s %s", fname, lname)
// }

// //Todo Index Page
// func TodoIndex(w http.ResponseWriter, r *http.Request) {

// 	/* Step and Requirment Preparing
// 	1. Template
// 	2. Data
// 	3. Render
// 	*/

// 	//Template
// 	tmpl := template.Must(template.ParseFiles("./views/todos/index.html"))
// 	// Data for display
// 	data := TodoPageList{
// 		PageTitle: "Todos Lists",
// 		Todos: []Todo{
// 			{Title: "Learing Go Fundamental", Done: true},
// 			{Title: "Learning Go by TDD", Done: false},
// 			{Title: "Create Web Application by Go with MVC structure", Done: false},
// 			{Title: "Personal Web Project", Done: false},
// 			{Title: "Develope Personal web framework", Done: false},
// 		},
// 	}

// 	tmpl.Execute(w, data)

// }
