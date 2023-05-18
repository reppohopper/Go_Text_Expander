# Go Text Expander

Go Text Expander is a text expansion tool written in Golang. It captures keyboard input, matches specific patterns (like `/shortcut*`), and expands them into predefined longer texts.

## Demo
![Demo GIF hosted on Giphy](https://giphy.com/embed/oUglZqjmKdjH2n4CsD)

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
