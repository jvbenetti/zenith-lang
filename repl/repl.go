package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/jvbenetti/zenith-lang.git/lexer"
	"github.com/jvbenetti/zenith-lang.git/token"
)

// PROMPT defining how to be the cursor
const PROMPT = ">> "

// Start init infinity loop in the terminal in Zenith
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()

		// If user click ctrl+C
		if !scanned {
			return
		}

		// Read line that user had wrote
		line := scanner.Text()

		// Pass the text to Lexer
		l := lexer.New(line)

		// Loop that asks the Lexer to process and print Token by Token
		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			// Print on screen the Type and the literal Value of the found token
			fmt.Fprintf(out, "{Type: %s, Literal: '%s'}\n", tok.Type, tok.Literal)
		}
	}
}
