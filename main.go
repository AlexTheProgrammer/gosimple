package main

import (
	d "github.com/AlexTheProgrammer/sail/dom"
)

func main() {

	ch := make(chan bool)
	d := d.DOM{
		Body: d.Body(
			Title(),
			DownloadAndInstall(),
			CreateAModule(),
			WriteAProgram(),
			Debugging(),
			WebServer(),
			ReadFile(),
			WriteFile(),
			Types(),
			Credits(),
		),
	}

	d.Render()

	<-ch // block indefinitely to keep program alive

}

var debugProgram string = `package main

import "log"

func main() {

	log.Printf("here is a log, but what line is it on?")
	log.SetFlags(log.Lshortfile) // configure logger
	log.Printf("<- here is the line number")
}
`

func Debugging() d.Node {
	return SectionWithNode(
		"debugging a program",
		"bg-rose-300",
		midText("often you will want logs, but the default logger needs to be configured to return file name/line number"),
		codeSection(debugProgram),
		codeSection("go run main.go"),
	)

}

var typesProgram string = `package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	anInt := 1
	aFloatAsaString := "2.3"
	aIntAsaString := "4"

	f, err := strconv.ParseFloat(aFloatAsaString, 64)
	if err != nil {
		log.Fatalf("could not convert to float, error: %v", err)
	}

	i, err := strconv.Atoi(aIntAsaString)
	if err != nil {
		log.Fatalf("could not convert to int, error: %v", err)
	}

    // when golang knows how to convert you can use the Type(var) syntax
	convertingIsTricky := float64(anInt) + f + float64(i)

	fmt.Printf(
		"anInt +  aFloatAsaString + aIntAsaString = %v but is tricky to code\n", 
		convertingIsTricky,
	)

	fmt.Printf(
		"anInt +  aFloatAsaString + aIntAsaString = %v%v%v is easy to concatenate\n", 
		anInt, aFloatAsaString, aIntAsaString,
	)

}
`

func Types() d.Node {
	return SectionWithNode(
		"handling types",
		"bg-sky-700",
		midText("go is strongly typed. So sometimes you need to wrangle your data into the right type"),
		codeSection(typesProgram),
		codeSection("go run main.go"),
	)

}

func CreateAModule() d.Node {

	return SectionWithNode(
		"create a module",
		"bg-purple-700",
		midText("every go program exists in it's own module to manage dependencies"),
		codeSection("go mod init myModule"),
		midText("you will now have a file called go.mod equilavent to package.json for web devs"),
	)
}

func Title() d.Node {
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

func DownloadAndInstall() d.Node {
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

func codeSection(text string) d.Node {

	return d.Div(d.Props{Class: d.NewClass("flex flex-row pb-10")},
		d.Div(d.Props{ID: "center svg left", Class: d.NewClass("basis-2/6")}),
		d.Pre(d.Props{
			Class: d.NewClass("bg-slate-800", "p-4", "rounded-lg", "text-slate-200", "basis-2/6"),
		},
			d.Code(d.Props{
				Class:     d.NewClass("text-2xl text-left"),
				InnerHTML: text,
			}),
		),
		d.Div(d.Props{ID: "center svg left", Class: d.NewClass("basis-2/6")}),
	)
}

func midText(innerHTML string) d.Node {
	return d.Div(d.Props{InnerHTML: innerHTML, Class: d.NewClass("w-full text-center text-3xl pb-10")})
}

func WriteAProgram() d.Node {
	return SectionWithNode(
		"write a program",
		"bg-sky-700",
		midText("filename - main.go"),
		codeSection(writeAProgram),
		codeSection("go run main.go"),
	)

}
func Credits() d.Node {

	return Section(
		"Credits",
		"bg-slate-200",
		d.Div(d.Props{Class: d.NewClass("pb-20")},
			d.Div(d.Props{InnerHTML: "highly influenced by one of the best sites on the web ", Class: d.NewClass("pb-5")}),
			d.B(d.Props{}, d.A(d.Props{
				Class:     d.NewClass("hover:underline", "p-2", "rounded-lg"),
				HRef:      "https://rogerdudler.github.io/git-guide/",
				InnerHTML: "git - the simple guide"},
			)),
		),
		d.Div(d.Props{}),
	)
}

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

var webserverCode string = `package main

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
        log.Printf("server crashed, error:%v", err)
    }
}`

func WebServer() d.Node {
	return SectionWithNode(
		"create a web server",
		"bg-orange-500",

		midText("filename - main.go"),
		codeSection(webserverCode),
		midText("run your web server and see your message"),
		codeSection("go run main.go"),
		midText("visit localhost:8080 in your browser"),
	)

}

var readfileProgram string = `
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    file, err := os.Open("myfile.txt")
    if err != nil {
		// this will exit your program 
        log.Fatalf("open file error: %v", err) 
    }
	// close when you're done
    defer file.Close() 

    data, err := ioutil.ReadAll(file)
    if err != nil {
        log.Fatalf("read error: %v", err)
    }

    fmt.Println("%s", data)
}
`

func ReadFile() d.Node {
	return SectionWithNode(
		"read a file into a program",
		"bg-violet-600",
		midText("filename - main.go with file myfile.txt stored alongside"),
		codeSection(readfileProgram),
		codeSection("go run main.go"),
	)

}

// # write files

var writefileProgram string = `
package main

import (
    "fmt"
    "os"
)

func main() {
    filePath := "myFile.txt"

    file, err := os.Create(filePath)
    if err != nil {
        log.Fatalf("create file error: %v", err)
    }

	// Ensure the file is closed when we're done
    defer file.Close() 

    data := []byte("go wrote this.\n")

    _, err = file.Write(data)
    if err != nil {
        log.Fatalf("write file error: %v", err)
    }
}
`

func WriteFile() d.Node {
	return SectionWithNode(
		"write to a file",
		"bg-emerald-500",
		midText("filename - main.go"),
		codeSection(writefileProgram),
		codeSection("go run main.go"),
		midText("and you should get myfile.txt"),
	)
}
