# Part 1

In this section, you have learned a lot about advanced data types, structs, methods, and interfaces, and now it is time
to put this into practice. So here is our exercise for section 3.

## Your task

Here is what we are going to do: We build a small Web service that lets us fetch quotes from a datastore, as well as
create new quotes or update existing ones.

In the first video of this exercise, we start by building a simple Web server. It is really simple, so this video will
be more like a short walkthrough than a full exercise. But watch closely, as you will now learn how to write a Web
server with just two functions that consist of eleven lines of code.

So here we go.

First, we write a function that receives two parameters, an http.ResponseWriter and an http.Request.

An http.ResponseWriter is an interface that implements the io.Writer interface, so we can call the Write() method on it,
as well as use all functions that expect an io.Writer as a parameter.

So our function can, for example, use the convenience function

`io.WriteString(Writer, string)`
to write a string into the response.

```go
package main

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
}
```

An http.Request is a struct that contains all aspects of an HTTP request, such as the URL, the request header, the body,
and more.

Let’s fetch the URL path from the request and write it into the response.

```go
package main

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!\n")
	io.WriteString(w, "URL: "+r.URL.Path)
}
```

Now we can use this function to handle an HTTP request that matches a given URL path. For this, we can use the method
http.HandleFunc() that takes a path and a function and turns that function into a handler for the given path.

```go
package main

func main() {
	http.HandleFunc("/", hello)
}
```

Then we only need to start the Web server, add some error handling, and we are done.

```go
package main

func main() {
	http.HandleFunc("/", hello)
	err := http.ListenAndServe("localhost:8000", nil)
	if err != nil {
		log.Fatalln("ListenAndServe:", err)
	}
}
```

`http.ListenAndServe()` starts an endless loop that waits for incoming requests and calls our hello function for each
one.

The first parameter specifies the interface and the port to listen on. We want to listen only to connections from the
local machine, so we use “localhost” as address, which binds to the loopback interface. To bind to all interfaces (and
listen to connection requests from anywhere), leave the address part empty and just specify a colon and the port number.

Let’s try this with a test client. We could use a browser here, but as we want to build a Web service, let’s use a
client that lets us inspect the traffic details. I will use a command-line client called “wuzz”, which, of course!, is
written in Go.

Installation is quick, via go get.

`go get github.com/asciimoo/wuzz`
If you have added GOPATH/bin to your PATH, you can now use wuzz right away.

After starting the server…

`go run main.go`
…we can type in the URL (using plain HTTP) and then we should see the answer from the server.

And that’s it! Now we have a very basic but working HTTP server.

As the next step, we will add routing capabilities, and two more handlers.

Download the code right from the lecture, and see the links section for a small selection of HTTP/REST clients.

## Links

wuzz - a command-line client for HTTP inspection, with an interactive UI.

bat - similar to wuzz but even more command-line-ish; that is, without an interactive UI.

Advanced REST client - a Chrome app with a convenient GUI and some useful features, like saving requests.

# Part 2

Now that we have a basic Web server up and running, our next step is to add some routing.

Routing means that we want to map URL paths to different handlers, and HTTP verbs to different operations. After all,
our server shall work as a REST service. But what is REST exactly? Here is a brief introduction, taken from the Applied
Go playist on YouTube.

## REST Basics

### Basic Operations

The REST protocol is based on four basic operations: Create, Read, Update, and Delete. These operations are often
described by the acronym CRUD.

### REST and HTTP

REST is often but not always based on the Hypertext Transfer Protocol. Methods are built from HTTP verbs like GET, POST,
PUT, DELETE, et cetera.

URL’s just point to resources that the methods can operate on. A URL never includes a method name.

This is how the CRUD operations can be mapped onto HTTP verbs:

HTTP verb CRUD operation
POST CREATE
GET READ
PUT UPDATE
DELETE DELETE
This assignment is not absolute. For example, both PUT and POST could be used to create or update data.

In RESTful Web API’s, a URL can describe a collection of elements, or an individual element. An HTTP verb may therefore
represent two different operations, depending on whether the resource is a collection or a single element.

### How To Pass Data

For sending data to the server, there are two options.

First, you can send small amounts of data within the URL itself.
Second, data can reside in the body of the HTTP request.
The server always returns data via the body of the HTTP response.

## Adding REST - Step 1: Routing capabilities

With the basics in mind, we will now add REST functionality to our server.

In a first step, we will add routing. This works pretty much the same as with the hello handler, except that we now have
to turn our resources into paths, and the operations into HTTP verbs.

Here is how to add routing for handling operations on a single quote.

First, we want to access our REST API at a distinct path, say, /api. For easier updating, our API needs a version
number, so we add /v1 to the path:

`prefix := "/api/v1/"`
Now we want to handle a single quote, so we register another handler function in main():

`http.HandleFunc(prefix+"quote/", handleQuote)`
Since we want to handle single quotes, we call the path simply “quote”. Note the trailing slash. A path that ends with a
slash names a rooted subtree, which means that everything below that path is routed to the function named handleQuote().
Without the trailing slash, only the exact path named /api/v1/quote would match.

After taking care of the resource, we now turn to implementing the operations for the resource named “quote”.

For this we create a handler function named handleQuote(). This function receives the request object, which contains the
HTTP method as a property named Method. We use this to invoke the respective method via a switch statement.

```go
package main

func handleQuote(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		io.WriteString(w, "Create: "+r.URL.Path+"n")
	case "GET":
		io.WriteString(w, "Read: "+r.URL.Path+"n")
	case "PUT":
		io.WriteString(w, "Update: "+r.URL.Path+"n")
	case "DELETE":
		io.WriteString(w, "Delete: "+r.URL.Path+"n")
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
```

For now, the handler shall do nothing but writing out the detected CRUD method and the resource path to the HTTP
response.

For any other HTTP verb than POST, GET, PUT, or DELETE, we want to answer the request by setting the HTTP status to “Bad
Request”. For this we can use the method WriteHeader and the constant value StatusBadRequest that is defined in the http
package, along with other HTTP status values.

When we run the server, we can now send HTTP requests to the quote/ path, and we should see the path and the requested
CRUD operation: “Read” for “HTTP GET”, “Create” for “HTTP POST”, and so on.

We can also verify that anything beneath quote/ is sent to our new handler function. (If we omit the trailing slash, the
server adds it automatically, since there cannot be a subtree and a single resource of the same name.)

## Your task

Here is your task for step 2 of this exercise. Besides working with single quotes, we also want to retrieve a list of
all existing quotes.

Write a function handleQuotesList() that ensures that the received method is an HTTP GET, and then writes the URL Path
to the request object.
If the HTTP verb is not GET, set the HTTP status to “Bad Request”.
Register this function as a handler for the API path that ends in quotes/.
You can start from the source file that comes with this lecture.

## Links

RESTful Web API Basics in Go - the blog article at appliedgo.net

RESTful Web API Basics in Go - the video from the blog article on YouTube