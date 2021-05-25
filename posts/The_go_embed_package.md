---
Index: 01
Dept: Web Dev
Title: The Go Embed Package
Subtopic: golang.binary.embed
Date: 25-05-2021
Time: 23:40:29.440
Topic: Golang embed
Summary: Embedding static content into the binary has never been easier.
SkimTime: 8
Author: Sumit Kumar
---

# Introduction :smile:

Today we will be having a look at one of the coolest features that came with the go1.16 release, that is the `embed` package.
This package is used to embed your static content in the binary. Earlier, we had to use third party packages like the [`packr`](https://github.com/gobuffalo/packr)
package. But with the introduction of the embed package, we do not need any third party libraries.

Let's have a look at how this package works, but before that, make sure you have the go version >=1.16 installed on your machine.   
You can check the version by entering the command `go version` in your terminal.

## Let's start with a really basic example.

First make a directory named `embedpg` and open this directory in your editor of choice. Now, initialize a go mod in this
directory by the `go mod init embedpg`. Now create a main.go file, and a text file that we will embed into the binary. 
Now, add the following code to your main.go file - 

```go
// Declaring this as the main package to run the main function
package main

import (
	// This is important for the embed package to work, make sure it is added to the imports
	_ "embed"
	"fmt"
)

// Here, we are using the `go:embed` directive to embed our file into the binary.
//go:embed content.txt
var content []byte

func main() {
	// Printing the content of the file by converting the array of bytes into string.
	fmt.Print(string(content))
}

```
Make sure the content.txt file is at the same directory level as our main.go file. Now running the file by `go run main.go`
should print the content of our `content.txt` file to the terminal. Magic, right?. Now you don't need to worry about shipping
the content file with the binary to anywhere you would want to use the file, because when you run the go build command and 
build a binary, you will have the content embedded into the binary.
<br/>
<br/>
We can embed directories with the help of the embed package as well. Let's next have a look at that and see how this package
powers this website - 

```go
package main

import (
	"embed"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"log"
	"net/http"
)

// Embed the static/ directory to the app to
//serve favicons, custom css and all other `static` stuff.
//go:embed static
var static embed.FS

func main() {
	app := fiber.New(
		fiber.Config{
			Views: engine,
		})
	// Here, I am using the static variable to expose the
	// whole static directory from my project onto the web server.
	app.Use("/", filesystem.New(
		filesystem.Config{
			Root:   http.FS(static),
			Browse: false,
		}))

	// Start the fiber app woo-hoo.
	log.Panic(app.Listen(":" + PORT))
}
```

Here, I am embedding multiple directories which contain different items for this website to run. Looking carefully at this,
you will find that the types of these variables are not `[]bytes` but `embed.FS`. This type implements the `fs.FS` interface,
and we can open any file of the directory by using the `Open()` method on this struct. Let's see how this is useful in web
development. In this example I am using [fiber](https://github.com/gofiber/fiber) but you can use the `net/http` package,
and it will work just fine. With the app.Use() code block, I am exposing all the static content that is in embedded into
my binary to webserver. You can use this to expose your CSS files, Favicons and Fonts if you are using the golang's templating
engine for server side rendering of your html just like this blog.

## Conclusion 
The `embed` package is a nice addition to the go library, and we gophers are here for it. If you want to learn more about
this package, I would suggest you to check out the official docs [here](https://golang.org/pkg/embed/). If you are interested
in this website's code, it is open source at GitHub [here](https://github.com/imsk17/gob-log). Make sure to :star: the repo,
open an issue if you find a :bug: and open a PR if you think you can fix it/make it better :wink:.