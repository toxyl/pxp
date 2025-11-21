package language

import (
	"fmt"
	"strings"
	"time"
)

// @Name: printf
// @Desc: Prints formatted strings to the console
// @Param:      str        - - -   The format string
// @Param:      a          - - -   A slice with the arguments
// @Returns:    result     - - -   Number of bytes written
func printf(str string, a []any) (int, error) {
	return fmt.Printf(str, a...)
}

// @Name: printfln
// @Desc: Prints formatted strings to the console
// @Param:      str        - - -   The format string
// @Param:      a          - - -   A slice with the arguments
// @Returns:    result     - - -   Number of bytes written
func printfln(str string, a []any) (int, error) {
	return fmt.Printf(str+"\n", a...)
}

// @Name: sprintf
// @Desc: Creates formatted strings
// @Param:      str        - - -   The format string
// @Param:      a          - - -   A slice with the arguments
// @Returns:    result     - - -   The formatted string
func sprintf(str string, a []any) (string, error) {
	return fmt.Sprintf(str, a...), nil
}

// @Name: uppercase
// @Desc: Uppercases a string
// @Param:      str        - - -   The string to uppercase
// @Returns:    result     - - -   The uppercased string
func uppercase(str string) (string, error) {
	return strings.ToUpper(str), nil
}

// @Name: lowercase
// @Desc: Lowercases a string
// @Param:      str        - - -   The string to lowercase
// @Returns:    result     - - -   The lowercased string
func lowercase(str string) (string, error) {
	return strings.ToUpper(str), nil
}

// @Name: time
// @Desc: Returns the current time according to the given layout
// @Param:      layout     - - "2006-01-02 15:04:05"    The layout
// @Returns:    result     - - -   						The current time as string
func timeStr(layout string) (string, error) {
	return time.Now().Format(layout), nil
}
