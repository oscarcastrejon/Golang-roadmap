// https://earthly.dev/blog/golang-errors/

// Errors

// Wrapping Errors

// Errors Are Better Wrapped

// The snippet below is refactored so that is uses fmt.Errorf with a %w
// verb to “wrap” errors as they “bubble up” through the other function
// calls. This adds the context needed so that it’s possible to deduce
// which of those database operations failed in the previous example.

// If we re-run the program and encounter the same error, the log should
// print the following:

// $ failed finding or updating user: FindAndSetUserAge: SetUserAge:
// failed executing db update: malformed request

// Now our message contains enough information that we can see the
// problem originated in the db.SetUserAge function. Phew! That definitely
// saved us some time debugging!

// If used correctly, error wrapping can provide additional context
// about the lineage of an error, in ways similar to a traditional
// stack-trace.

// Wrapping also preserves the original error, which means errors.Is
// and errors.As continue to work, regardless of how many times an
// error has been wrapped. We can also call errors.Unwrap to return
// the previous error in the chain.

// =====================================================================
// When To Wrap

// Generally, it’s a good idea to wrap an error with at least the
// function’s name, every time you “bubble it up” - i.e. every time
// you receive the error from a function and want to continue returning
// it back up the function chain.

// There are some exceptions to the rule, however, where wrapping an
// error may not be appropriate.

// Since wrapping the error always preserves the original error
// messages, sometimes exposing those underlying issues might be a
// security, privacy, or even UX concern. In such situations, it could
// be worth handling the error and returning a new one, rather than
// wrapping it. This could be the case if you’re writing an open-source
// library or a REST API where you don’t want the underlying error
// message to be returned to the 3rd-party user.

// =====================================================================
// To run it:
// go run main.go

package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID       string
	Username string
	Age      int
}

func FindUser(username string) (*User, error) {
	var err error = errors.New("not connected")

	return nil, fmt.Errorf("FindUser: failed executing db query: %w", err)
}

func SetUserAge(u *User, age int) error {
	var err error = errors.New("not connected")

	return fmt.Errorf("SetUserAge: failed executing db update: %w", err)
}

func FindAndSetUserAge(username string, age int) error {
	var user *User
	var err error

	user, err = FindUser(username)
	if err != nil {
		return fmt.Errorf("FindAndSetUserAge: %w", err)
	}

	if err = SetUserAge(user, age); err != nil {
		return fmt.Errorf("FindAndSetUserAge: %w", err)
	}

	return nil
}

func main() {
	if err := FindAndSetUserAge("bob@example.com", 21); err != nil {
		fmt.Printf("failed finding or updating user: %s", err)

		return
	}

	fmt.Println("successfully updated user's age")
}
