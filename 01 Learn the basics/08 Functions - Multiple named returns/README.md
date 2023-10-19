# Golang Roadmap

# Functions

## What is Function in Golang

A function is a group of statements that exist within a program for the purpose of performing a specific task. At a high level, a function takes an input and returns an output.

Function allows you to extract commonly used block of code into a single component.

The single most popular Go function is main(), which is used in every independent Go program.

## Naming Conventions for Golang Functions

* A name must begin with a letter, and can have any number of additional letters and numbers.
* A function name cannot start with a number.
* A function name cannot contain spaces.
* If the functions with names that start with an uppercase letter will be exported to other packages. If the function name starts with a lowercase letter, it won't be exported to other packages, but you can call this function within the same package.
* If a name consists of multiple words, each word after the first should be capitalized like this: empName, * EmpAddress, etc.
* function names are case-sensitive (car, Car and CAR are three different variables).