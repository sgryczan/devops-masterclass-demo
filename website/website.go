package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"gitlab.com/sgryczan/website/models/post"
)

func displayPostHandler(w http.ResponseWriter, r *http.Request) {
	post := post.NewPost(
		&post.NewPostInput{
			Username:     "",
			Mood:         post.Moods["small"],
			Caption:      "",
			MessageBody:  "",
			Url:          "",
			ImageURI:     "/images/santa.gif",
			ThumbnailURI: "",
			Keywords:     []string{"rocket"},
		},
	)

	fmt.Printf("myPost: %+v\n", post)

	renderTemplate(w, "./templates/post.html", post)
}

// Template rendering function
func renderTemplate(w http.ResponseWriter, templateFile string, templateData interface{}) {

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		log.Fatal("Error encountered while parsing the template: ", err)
	}
	t.Execute(w, templateData)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", displayPostHandler)
	http.Handle("/images/", http.FileServer(http.Dir("./static")))
	http.Handle("/css/", http.FileServer(http.Dir("./static")))
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
