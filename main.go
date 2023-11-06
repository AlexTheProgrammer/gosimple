package main

import (
	d "github.com/AlexTheProgrammer/sail/dom"
)

func main() {

	ch := make(chan bool)
	d := d.DOM{
		Body: d.Body(
			Section1(),
			Section2(),
			Section3(),
		),
	}

	d.Render()

	<-ch // block indefinitely to keep program alive

}

func Section1() d.Node {
	return Section(
		"golang - the simple guide",
		"bg-slate-100",
		d.Div(d.Props{InnerHTML: "simple guide for getting started with golang"}),
		d.Div(d.Props{Class: d.NewClass("flex flex-row")},
			d.Div(d.Props{ID: "center svg left", Class: d.NewClass("basis-5/12")}),
			d.Div(d.Props{ID: "svg", Class: d.NewClass("basis-2/12")},
				d.SVG(
					d.Props{
						XMLNS:   "http://www.w3.org/2000/svg",
						ViewBox: "0 0 24 24",
					},
					d.Path(d.Props{D: "m12 18.414-4.707-4.707 1.414-1.414L12 15.586l3.293-3.293 1.414 1.414L12 18.414z"}),
					d.Path(d.Props{D: "M11 6h2v11h-2z"}),
				),
			),

			d.Div(d.Props{ID: "center svg right", Class: d.NewClass("basis-5/12")}),
		),
	)
}

func Section2() d.Node {
	return Section(
		"Setup",
		"bg-blue-400",
		d.Div(d.Props{Class: d.NewClass("pb-20")},
			d.A(d.Props{
				Class:     d.NewClass("hover:underline", "bg-blue-800", "p-2", "rounded-lg", "text-slate-200"),
				HRef:      "https://go.dev/doc/install",
				InnerHTML: "Download & install go"},
			),
		),
		d.Div(d.Props{}),
	)
}

var writeAProgram string = `package main

import "fmt"

func main() {
	fmt.Println("my go program")
}`

func Subsection3(text string) d.Node {

	return d.Div(d.Props{Class: d.NewClass("flex flex-row pb-10")},
		d.Div(d.Props{ID: "center svg left", Class: d.NewClass("basis-2/6")}),
		d.Pre(d.Props{
			Class: d.NewClass("bg-slate-800", "p-2", "rounded-lg", "text-slate-200", "basis-2/6"),
		},
			d.Code(d.Props{
				Class:     d.NewClass("text-2xl text-left"),
				InnerHTML: text,
			}),
		),
		d.Div(d.Props{ID: "center svg left", Class: d.NewClass("basis-2/6")}),
	)
}

func Section3() d.Node {
	return SectionWithNode(
		"write a program",
		"bg-sky-700",
		Subsection3("filename: main.go"),
		Subsection3(writeAProgram),
		Subsection3("command: go run main.go"),
	)

	// d.Div(d.Props{ID: "Section 3"},
	// 	d.Div(d.Props{InnerHTML: "write a program"}),
	// 	d.Div(d.Props{},
	// 		d.Div(d.Props{
	// 			InnerHTML: `
	// 					# file main.go
	// 					package main

	// 					import "fmt"

	// 					func main() {
	// 						fmt.Println("my go program")
	// 					}`,
	// 		}),
	// 		d.Div(d.Props{InnerHTML: "then call"}),
	// 		d.Div(d.Props{
	// 			InnerHTML: `go run main.go`,
	// 		}),
	// 	),
	// ),
}

// d.Div(d.Props{InnerHTML: "Setup"}),
// d.Div(d.Props{},
// 	d.A(d.Props{HRef: "https://go.dev/doc/install", InnerHTML: "Download & install go"}),

func Section(heading string, bgColor string, content d.Node, diagram d.Node) d.Node {
	return d.Div(d.Props{ID: heading, Class: d.NewClass("w-full font-serif pt-20", bgColor)},
		d.B(d.Props{},
			d.Div(d.Props{
				InnerHTML: heading,
				Class:     d.NewClass("w-full text-center text-6xl font-thin  mb-10"),
			}),
		),
		d.Div(
			d.Props{Class: d.NewClass("w-full text-center text-3xl")},
			content,
		),
		diagram,
	)
}

func SectionWithNode(heading string, bgColor string, nodes ...d.Node) d.Node {
	firstNode := d.B(d.Props{},
		d.Div(d.Props{
			InnerHTML: heading,
			Class:     d.NewClass("w-full text-center text-6xl font-thin  mb-10"),
		}),
	)
	nodes = append([]d.Node{firstNode}, nodes...)

	return d.Div(d.Props{ID: heading, Class: d.NewClass("w-full font-serif pt-20", bgColor)}, nodes...)

}

func SectionAll() d.Node {
	return d.Div(d.Props{},

		d.Div(d.Props{ID: "Section 4"},
			d.Div(d.Props{InnerHTML: "setup a web server"}),
			d.Div(d.Props{},
				d.Div(d.Props{
					InnerHTML: `
							# file main.go
							package main

							import (
							    "fmt"
							    "net/http"
							)
							
							func main() {
							
							    handler := func(w http.ResponseWriter, r *http.Request) {
							        fmt.Fprintf(w, "hello from server")
							    }
							
							    http.HandleFunc("/", handler)
							
							    if err := http.ListenAndServe(":8080", nil); err != nil {
							        fmt.Println("Error:", err)
							    }
							}`,
				}),
				d.Div(d.Props{
					InnerHTML: `go run main.go and now visit localhost:8080`,
				}),
			),
		),
	)
}
