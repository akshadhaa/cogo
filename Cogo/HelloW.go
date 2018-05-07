package main

import(
	"html/template"
	"net/http")

var temp *template.Template

func init(){

temp = template.Must(template.ParseGlob("Templates/*.gohtml"))

}


func index(w http.ResponseWriter, r *http.Request){

	temp.ExecuteTemplate(w, "index.gohtml", nil)

}

func main(){

	http.HandleFunc( "/", index)
	http.HandleFunc( "/process", processor)
	http.ListenAndServe(":8000" , nil)

}

func processor(w http.ResponseWriter, r *http.Request){

	if r.Method != "POST"{
	
		http.Redirect(w , r, "/", http.StatusSeeOther)
		return
	
	}

	fname := r.FormValue("first_name" )
	lname := r.FormValue("last_name" )

	d := struct {

		First string
		Last string
	}{
		First : fname,  
		Last : lname,
	}

	temp.ExecuteTemplate(w, "processor.gohtml", d)

}
