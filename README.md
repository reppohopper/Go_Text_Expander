# Go Text Expander

Go Text Expander is a text expansion tool written in Golang. It captures keyboard input, matches specific patterns (like `/shortcut*`), and expands them into predefined longer texts. See the demo: 

![Demo GIF hosted on Giphy](https://media.giphy.com/media/v1.Y2lkPTc5MGI3NjExMzU5ZWM0OTQ2YmM5OTMxNTAwNWE3YWJjMWVjMDE4NzA1ZmMyMjAwNSZlcD12MV9pbnRlcm5hbF9naWZzX2dpZklkJmN0PWc/oUglZqjmKdjH2n4CsD/giphy.gif)

## Dependencies

The program has the following dependencies:

- go-hook
- keybd_event
- clipboard

Install these dependencies using `go get`:

```bash
go get github.com/moutend/go-hook/pkg/keyboard
go get github.com/moutend/go-hook/pkg/types
go get github.com/micmonay/keybd_event
go get github.com/atotto/clipboard
```

## Usage

Set your expansions file path as an environment variable named EXPANSIONS_PATH.

Run the program with the command go run main.go.

Press Ctrl+C to stop the program.
