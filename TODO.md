
# things I want to cover
# I can add to this later, but I think this will be helpful.

# getting a go program working
## initialise the project..... think `npm init`
```
    go mod init mymodule # you can change this later
```

## write a program that can be compiled
```
# file main.go
package main

import "fmt"

func main() {
	fmt.Println("my go program")
}
```


then call 
```
go run main.go
```


# setup a web server
```
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
}
```

now visit localhost:8080


# read files

say you have a file 

myfile.txt
```
Hi there, 
this is some text in a file
```


```
package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() {

    file, err := os.Open("myfile.txt")
    if err != nil {
        log.Fatalf("could not open file error: %v", err) # this will exit your program
    }
    defer file.Close() // close when you're done

    // Read the file's contents
    data, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    // Print the file's contents as a string
    fmt.Println("%s", data)
}

```


# write files

```
package main

import (
    "fmt"
    "os"
)

func main() {
    // Specify the file path
    filePath := "myFile.txt"

    // Create or open the file for writing
    file, err := os.Create(filePath)
    if err != nil {
        log.Fatalf("could not create file, error: %v", err)
    }
    defer file.Close() // Ensure the file is closed when we're done

    // Data to write to the file
    data := []byte("this was written through golang.\n")

    // Write the data to the file
    _, err = file.Write(data)
    if err != nil {
        log.Fatalf("could not write file, error: %v", err)
    }
}

```

```
go run main.go 
```
should now have a file called `myFile.txt` in the same folder you ran `go run main.go`


# handle types 
```

package main

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

	fmt.Printf("anInt +  aFloatAsaString + aIntAsaString = %v but is tricky to code\n", convertingIsTricky)

	fmt.Printf("anInt +  aFloatAsaString + aIntAsaString = %v%v%v is easy to concatenate\n", anInt, aFloatAsaString, aIntAsaString)

}
```



# where is that log 
```
package main

import "log"

func main() {

	log.Printf("here is a log, but what line is it on?")
	log.SetFlags(log.Lshortfile)
	log.Printf("<- here is the line number")
}
```

