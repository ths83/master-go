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

# Part 3

In the third part of this exercise, we will implement a very simple data store and rewrite our handler functions to
store and retrieve quotes to and from that store.

That is, I write the simple data store, and you write the rest! :-P

So let’s go!

A quote usually includes a text, an author’s name, and optionally the source of the quote, like a book or an article in
a magazine. We can pack all this into a struct.

```go
package main

type Quote struct {
	Author string
	Text   string
	Source string
}
```

And for converting the data to and from JSON, we need to add field tags. Because I am lazy, I do this via the editor’s
context menu. Now we just need to remove the omitempty tags for Author and Text, because those two fields are not
optional.

```go
package main

type Quote struct {
	Author string `json:"author"`
	Text   string `json:"text"`
	Source string `json:"source,omitempty"`
}
```

Now we can define a data store. But here is a caveat. We don’t want to create the data store as a global object, but on
the other hand, we cannot pass it as an additional parameter to our handler functions, because the parameter list of a
handler function cannot be changed.

But there is a solution for that. Remember that methods are values and can take their receiver with them. We therefore
can change all our handler functions into methods without touching the parameter list. And as the receiver of these
methods, we create a struct named App that holds our data store.

As already mentioned, the data store itself is very simple: It is just a map from string to *Quote. The key of the map
is the author’s name, as we want to query for authors. This duplicates the author name, because it is already contained
in the Quote struct, but we take this into account for easier JSON marshalling and unmarshalling.

```go
package main

type App struct {
	storage map[string]*Quote
}
```

After turning the handler functions into methods…

`func (app *App) handleQuote(w http.ResponseWriter, r *http.Request) { ... }`

`func (app *App) handleQuotesList(w http.ResponseWriter, r *http.Request) { ... }`
…we also need to instantiate the App object in the main function…

```go
package main

func main() {
	app := &App{
		storage: map[string]*Quote{},
	}
	...
```

…and then modify the HandleFunc invocations slightly…

`http.HandleFunc(prefix+"quote/", app.handleQuote)`
`http.HandleFunc(prefix+"quotes/", app.handleQuotesList)`

…and now the handlers have access to the App object and thus also to the data store.

## Your task

That was my part, and now here is your part.

Implement handlers for:

- Creating a quote,
- Reading a quote, and
- Listing all quotes.
  For creating a quote, the client sends a POST request to the quote/ path, with a single quote encoded as a JSON object
  in the request body:

```go
package main
{
"author": "Homer J. Simpson",
"text": "D'oh!",
"source": "The Simpsons"
}
```

From the lecture about struct tag fields, you already know how to marshal and unmarshal JSON object from and to structs.
I added a few hints to the TODO comments, and remember that you always have the package documentation at golang.org
available.

Happy coding!

# Part 4

> NOTE: Cannot find files to complete the exercise.

In the last part of this exercise, we will replace our overly simple storage map with a full-fledged, concurrency-safe
key value store.

## Bolt DB

I chose BoldDB because of its simplicity and maturity. Bolt DB is used by many large projects, and is known for being
fast, robust, and stable.

So before diving into the assignment, let’s have a brief look at Bolt DB.

### Storing and retrieving data

Bolt DB is a key-value store. Key-value stores are like the big brother of maps, with persistent storage and some other
goodies included. Basically, data is stored as key-value pairs, just like with maps, and in Bolt DB the key and the
value both are byte slices.

So in order to store a Quote struct in Bolt DB, we need to serialize it. For this purpose, we use a binary format called
gob which is supported in the standard library by a package of the same name.

gob uses encoders and decoders that pretty much work like JSON’s marshal and unmarshal operations but with binary data.

```go
enc := NewEncoder(aWriter)
enc.Encode(someData)
```

gob has a very small API, and you can look into the examples within the package documentation to get familiar with how
encoding and decoding works.

### Buckets

Data in Bolt DB is organized in so-called buckets. In each bucket, the keys must be unique. Buckets can exist
side-by-side, and they can be nested. Think of a slice of maps, and a map of maps as an analogy.

### Transactions

Bolt DB supports transactions. Transactions ensure data consistency. You can run multiple database operations within a
transaction, and if one of the operation fails, all previous operations within that transaction are rolled back, in
order to restore the previous consistent state.

Bolt DB provides three types of transactions, of which we need two: The Update transaction, and the View transaction.

Both are designed as methods that take a function as a parameter; this is usually an anonymous function, or closure,
declared right within the call. Within this closure, we start by creating or opening a bucket. Then we can use Put and
Get to insert, update, and retrieve values to and from the bucket.

To iterate over all keys in a bucket, use the ForEach method.

Now with this brief intro to Bolt DB and gob in mind, lets turn to our task.

## Your task

For this final part I want to raise the bar a bit and give you only minimal instructions. Instead, I encourage you to
dig into the documentation of the Bolt DB and gob packages and figure out on your own how things fit together. The task
itself is not too difficult, but you will have to get familiar with two new API’s, so take a bit more time for this part
than for the previous ones.

Your task is to replace the map storage by Bolt DB. For this, I have moved the Quote struct to a separate package called
quotes. In the file quote.go you will find the Quote struct along with two new methods that you shall implement:

Serialize, and
Deserialize.
The former turns the quote it belongs to into a gob, and the latter turns a gob back into a quote.

Then turn over to the file db.go. Here you find a struct called DB and a few functions and methods to implement:

`Open()` shall take a path and return a new or existing database.
`DB.Close()` shall close the open database.
`DB.Create()` takes a quote and shall insert this quote into the database. It is an error to insert a quote that already
exists.
`DB.Get()` takes an author name and shall retrieve the corresponding quote.
And finally `DB.List()` shall list all quotes that are stored in the database.
You may optionally implement “update” and “delete” operations, too, but this is not part of this assignment.

Finally, take main.go and adjust the import path to point to your local quotes package. main.go is already modified to
use your new DB instead of the map, so you don’t have to do this text replacement tasks yourself.

Test the database access now using a REST client. You can create new data via the client, or you can use the file
“quotesdb”, which is a prefilled Bolt DB database so that you don’t need to start from scratch.

This time I provide no walkthrough videos as the concepts used here should already be understood quite well from
previous tasks, and the rest should be fairly close to the documentation.

## Bonus task

As a bonus task, extend the quotes package to allow storing multiple quotes per author. Here is one tip: Nested buckets
might be useful here.

Happy coding!
