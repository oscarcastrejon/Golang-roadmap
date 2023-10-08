# Golang Roadmap

# Errors

# Wrapping Errors

In these examples so far, the errors have been created, returned, and handled with a single function call. In other words, the stack of functions involved in “bubbling” up the error is only a single level deep.

Often in real-world programs, there can be many more functions involved - from the function where the error is produced, to where it is eventually handled, and any number of additional functions in-between.

In Go 1.13, several new error APIs were introduced, including errors.Wrap and errors.Unwrap, which are useful in applying additional context to an error as it “bubbles up”, as well as checking for particular error types, regardless of how many times the error has been wrapped.