package main

/*
q1. Create a function that converts string to an integer
    if any alphabets are passed, wrap strconv error and ErrStringValue error (create ErrStringValue error)
    ErrStringValue contains a message that 'value is of string type'

	return the wrapped errors
    if error is happening because of something else other than alphabets are passed then return the original error

    use the regex to check if value is of string type or not
    Hint: regexp.MatchString(`^[a-zA-Z]`, s)
    fmt.Errorf("%w %w") // to wrap error

    In main function check if ErrStringValue error was wrapped in the chain or not
    If yes, log a message 'value must be of int type not string' and log original error message alongside as well
*/
