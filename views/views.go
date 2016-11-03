package views

import "html/template"

func NewView(layout string, files ...string) *View { // this is called a variadric parameter, it means that it takes in any numbber of args as long as it is a string the *View to the next of the params that are passed is the
	files = append(files,
		"views/layout/footer.gohtml",
		"views/layout/bootstrap.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

type View struct { // this is a struct that contains a pointer to a template that contains our compiled template
	Template *template.Template
	Layout   string
}
