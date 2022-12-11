# Time To Practice: a terminal writer

> NOTE: File is missing therefore I cannot do the exercise.

## Intro

The io.Writer interface is as simple as can get, yet it has great abstraction capabilities. Whatever you want to write
to–standard output, a file, a network connection, a byte buffer, and so on–is abstracted away by a single, uniform
method named Write().

And you can even write your own output facilities that can be used anywhere an io.Writer interface is accepted.

For example, imagine you frequently need to write to the terminal but not across the entire width of the terminal
window. To solve this problem once and for all, you decide to implement a TerminalWriter that writes text to the
terminal at a given width.

## Your task

Start from the code below (also downloadable as “iowriter.go”) and implement the Write method of TerminalWriter.

It must meet the following requirements:

It writes to os.Stdout. (You can use the standard fmt.Print/f/ln family here, or os.Stdout.Write().)
After tw.width bytes, start a new line.
When the function encounters any error during writing, return an appropriate error, and the number of bytes successfully
written.

## Bonus task 1: The poet

Make the text break at word boundaries.

(No solution provided.)

## Bonus task 2: The world citizen

Start the next line after n Unicode characters rather than n bytes.

To simplify this task, only consider Latin scripts (where all glyphs have the same width - at least with a
non-proportional font as used in terminals).

The packages “unicode/utf8” from the standard library (most notably) and “golang.org/x/text/width” may prove useful. (As
well as maybe other packages of golang.org/x/text.)

(No solution provided.)