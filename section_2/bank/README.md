# Transcript

We are at the end of section 2, and this exercise summarizes all the stuff we have learned in this section: variables,
control structures, command-line handling, strings and numbers, pointers, functions, error handling, and working with
packages and libraries.

## The task

Here is your task.

Write a command line tool that manages bank accounts. The tool shall implement a few commands, according to this usage()
function:

```go
package main

func usage() {
	fmt.Println(`Usage:

bank create <name>                     Create an account.
bank list                              List all accounts.
bank update <name> <amount>            Deposit or withdraw money.
bank transfer <name> <name> <amount>   Transfer money between two accounts.
bank history <name>                    Show an account's transaction history.
`)
	os.Exit(1)
}
```

Here, bank is the name of the compiled binary. We can safely substitute the familiar go run main.go for that, like,

```shell 
go run main.go create Rob
```

Or if you want, you can run “go build” in the source code folder, and this will then create a binary right in that
folder, with the same name as the folder.

But that’s just as an aside - let’s turn back to the task and break it into smaller parts.

First, we need to read the command line parameters and decide how to act upon them. Remember the os.Args list that
contains all command line parameters, and the switch statement that lets us branch into separate actions.

### Implement the commands

Then, we have to implement each of the commands. To facilitate that, I have created a library package with a basic
banking API. Add the line

```go 
import "github.com/appliedgocourses/bank"
```

to your source file in order to use the package in your code.

And remember, if your editor helps you getting rid of unused imports upon saving, ensure to add code that uses the
package before hitting save.

The bank package has a Load() function (among others), so let’s add a call to this inside main():

```
err := bank.Load()
if err != nil {
    log.Println("Cannot load bank data:", err)
}
```

If you have no go.mod file yet, it would be a good time to run go mod init bank right now. Else try go mod tidy to clean
up and update the go.mod file. You should then see the bank package listed there as a new dependency.

By the way, there is no name clash between our bank module and the bank package we imported. The module determines the
default name of the binary, so I chose bank to have a short and concise name for the binary.

## Inspecting a package API

In order to use this package, you will want to know its API - that is, which functions and types it provides. Here are
two ways of inspecting the API of a package.

### The CLI approach: go doc

On the command line, you can use the go doc tool. Type

```shell
go doc bank
```

and go doc will list the API of the first package named “bank” that it can find in the standard library or in GOPATH (
where all the cached third-party packages live). If you had another package named bank on your system, you can choose
the correct one by supplying more path information, like

```shell
go doc appliedgocourses/bank
```

or the whole import path github.com/appliedgocourses/bank.

```shell
go doc then returns this API information:
```

```go
package bank // import "github.com/appliedgocourses/bank"

Package bank was made for the Master Go course at appliedgo.com.

func Balance(a *Account) int
func Deposit(a *Account, m int) (int, error)
func History(a *Account) func() (int, int, bool)
func ListAccounts() string
func Load() error
func Name(a *Account) string
func Save() (err error)
func Transfer(a, b *Account, m int) (int, int, error)
func Withdraw(a *Account, m int) (int, error)

type Account struct{ ... }

func GetAccount(name string) (*Account, error)
func NewAccount(s string) *Account
```

You can see the name of the package along with the proper import statement, as well as a short package description and a
list of exported functions and types.

To find out more about a specific function; say, about the function Balance(), type

```shell
go doc bank.Balance
```

to get a short description of this function:

```go
package bank // import "github.com/appliedgocourses/bank"

func Balance(a *Account) int
Balance returns the current balance of account a.
Try this for yourself, for example,

go doc bank.Account
go doc bank.History
```

## The online approach: godoc.org

A more convenient way of inspecting a package’s API is godoc.org This online service displays basically the same output
as go doc but nicely rendered on a single page.

So we just need to paste the import path of our bank package into the search field to get the complete API
documentation.

(Link: https://godoc.org/github.com/AppliedGoCourses/bank)

## A quick API walkthrough

To get you up to speed, let’s have a quick look at the available functions.

### Data Persistence

As we write a command line tool, we need a way to preserve data from one call to the next. The two functions Load() and
Save() take care of this. Call Load() at the start of main() and defer Save() to the end, and you are done.

### Account functions

Next, there are also some functions around the account itself.

NewAccount() creates an account by name. GetAccount() returns an existing account by name. You will make use of
GetAccount() quite a lot, as most functions expect an Account object rather than just the name.

ListAccounts() returns a formatted string that lists all accounts and their current balance.

Balance() returns the current balance for a particular account.

The History() function is a special case. It returns a function - a closure that yields one entry of an account’s
transaction history at a time. The last parameter, a boolean, indicates if there are more entries to come.

Call this closure in a loop to build up a formatted transaction history output.

### Transactions

The remaining functions provide transaction functionality.

Deposit() and Withdraw() add or remove an amount of money to or from an account, respectively. No specific currency is
assumed, let’s call it just “credits”.

Transfer() moves an amount of money from one account to another.

## Let’s start!

So let’s start building our banking application. I’ll leave it to you how you approach the task. You may start from
scratch, or use the provided main.go as a starting point.

And remember: It is ok to struggle with this task–this is how the brain learns–, but don’t struggle too much.

The next lecture provides a step-by step solution, along with source code, so if you get stuck at any stage, feel free
to peek into the corresponding solution.

If you want, play with the solution code. Take it apart, and examine the details. Use it for filling some of the blanks
of your own solution. Test the result and try to understand how it works. Analyzing, dissecting, and reassembling is
another way of learning.

Happy coding!

