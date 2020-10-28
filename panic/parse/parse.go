// parse.go
package parse

import (
	"fmt"
	"strconv"
	"strings"
)

// A ParseError indicates an error in converting a word into an integer.
type ParseError struct {
	Index int    // The index into the space-separated list of words.
	Word  string // The word that generated the parse error.
	Err   error  // The raw error that precipitated this error, if any.
}

// String returns a human-readable error message.
func (e *ParseError) String() string {
	return fmt.Sprintf("pkg parse: error parsing %q as int", e.Word)
	//return e.Err.Error()
}

// Parse parses the space-separated words in in put as integers.
func Parse(input string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			fmt.Println("get recover is ", r)
			// this statement print: get recover is  pkg parse: error parsing "12.5" as int
			// why here auto to call ParseError.String()? ParseError implements Stringer
			// case 'v', 's', 'x', 'X', 'q':
			//			// Is it an error or Stringer?
			//			// The duplication in the bodies is necessary:
			//			// setting handled and deferring catchPanic
			//			// must happen before calling the method.
			//			switch v := p.arg.(type) {
			//			case error:
			//				handled = true
			//				defer p.catchPanic(p.arg, verb)
			//				p.fmtString(v.Error(), verb)
			//				return
			//
			//			case Stringer:
			//				handled = true
			//				defer p.catchPanic(p.arg, verb)
			//				p.fmtString(v.String(), verb)
			//				return
			//			}
			//		}
			err, ok = r.(error)
			fmt.Println("is convert recover to error ok? ", ok)
			fmt.Println("convert recover to error is ", err)
			if !ok {
				err = fmt.Errorf("pkg: %v", r)
			}
		}
	}()
	// Fields splits the string s around each instance of one or more consecutive white space
	// characters, as defined by unicode.IsSpace, returning a slice of substrings of s or an
	// empty slice if s contains only white space.
	fields := strings.Fields(input)
	numbers = fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseError{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
