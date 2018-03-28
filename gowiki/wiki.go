// https://golang.org/doc/articles/wiki/

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte //byte instead of string as this is the type expected by io/ioutil
}

//This method's signature reads: "This is a method named save that takes as its receiver p, a pointer to Page . It takes no parameters, and returns a value of type error."
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	//the "blank identifier" represented by the underscore (_) symbol is used to throw away the error return value (in essence, assigning the value to nothing).
	// body, _ := ioutil.ReadFile(filename)
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*
//This main is used for the first part of the tutorial
func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample page")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
*/

/*
// This viewHandler is used for first part of tutorial without templates
func viewHandler(w http.ResponseWriter, r *http.Request) {
	// The Path is re-sliced with [len("/view/"):] to drop the leading "/view/" component of the request path. This is because the path will invariably begin with "/view/", which is not part of the page's title.
	title := r.URL.Path[len("/view/"):]
	// println(r.URL.Path)
	// println(r.URL.Path[len("/view/"):])
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
	//Again, note the use of _ to ignore the error return value from loadPage. This is done here for simplicity and generally considered bad practice. We will attend to this later.
}
*/

//This viewHandler uses Templates
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	//The http.Redirect function adds an HTTP status code of http.StatusFound (302) and a Location header to the HTTP response
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//to render directly in function see edithandler
	renderTemplate(w, "view", p)
}

/*
// This editHandler is used for first part of tutorial without templates
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	fmt.Fprintf(w, "<h1>Editing %s</h1>"+
		"<form action=\"/save/%s\" method=\"POST\">"+
		"<textarea name=\"body\">%s</textarea><br>"+
		"<input type=\"submit\" value=\"Save\">"+
		"</form>",
		p.Title, p.Title, p.Body)
}
*/

//This editHandler uses Templates
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		// The & operator generates a pointer to its operand.
		// The * operator denotes the pointer's underlying value
		p = &Page{Title: title}
	}
	t, _ := template.ParseFiles("edit.html")
	//The method t.Execute executes the template, writing the generated HTML to the http.ResponseWriter. The .Title and .Body dotted identifiers refer to p.Title and p.Body.
	//Template directives are enclosed in double curly braces.
	t.Execute(w, p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
