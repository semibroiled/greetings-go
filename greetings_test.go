package greetings // get greetings to test its funcs

import (
    "testing"
    "regexp"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) { // pointer to testing.T type
    name := "Gladys" // get a name
    want := regexp.MustCompile(`\b`+name+`\b`) // want has name followed and preceded by characters
    msg, err := Hello("Gladys") //call function
    if !want.MatchString(msg) || err != nil { // if not represented in pattern in want or err is nil, we call error
        t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, msg, err, want) // use testing.T parameter method to test and log


    }
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
    msg, err := Hello("")
    if msg != "" || err == nil {
        t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
    }
} // I think this will return an error right msg has empty string instead of name not not nothing
// no quatsch, this checks if msg is non empty. so something no wait
// I was right, beacause we reutrn empty string if no name üassed and give an error

