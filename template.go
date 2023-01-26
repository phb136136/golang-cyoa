package main

import (
	"html/template"
	"os"
)

// <div>{{range $index, $element := .Options}} {{element.Text}} {{end}}</div>
func generateTemplate(name string, storyarc StoryArc) {
	htmlTemplate := `
	<a title='{{.Title}}'>
	<div>
	{{ range .Story }}
		{{.}}
		{{end}}
	</div>
	<div>{{range $index, $element := .Options}} 
	<a href="{{$element.Arc}}">{{$element.Text}} </a> 
	{{end}}</div>
	
	`
	t, err := template.New(name).Parse(htmlTemplate)
	err = t.Execute(os.Stdout, storyarc)
	if err != nil {
		panic(err)
	}
}
