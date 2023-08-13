package greetings

import (
	"errors"
	"fmt"
    "math/rand"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
    // If no name given, return error with message
    if name == "" {
        return "", errors.New("Name is empty")
    }
    // Return a greeting that embeds the name in a message.
    message := fmt.Sprintf(randomGreeting(), name)
    // Test faile case with no name in f string to replace parameter
    // message := fmt.Sprint(randomGreeting())
    return message, nil //nil means no error
}

// New function that calls on Hello for more people

func Hellos(names []string) (map[string]string, error) {
    // A map to associate names with messages
    messages := make(map[string]string)

    //loop through the received slice of names, calling the hello 
    //function for a name to get the message
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            //if there is an error, return the error, nil to pad return type
            return nil, err
        }

        //in loop, in the map, associate the message to a name with the name
        messages[name] = message
    }
    return messages, nil
}

func randomGreeting() string{

    //slice with message formats
    formats := []string{"Hi, %v. Welcome!",
                        "Sup %v, you my basbousa",
                        "Salam %v",
                        "Ily %v"}
    return formats[rand.Intn(len(formats))] //Index with a random integer between start number and length of formats 
}