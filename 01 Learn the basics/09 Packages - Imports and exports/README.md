# Golang Roadmap

# Packages

Packages are the most powerful part of the Go language. The purpose of a package is to design and maintain a large number of programs by grouping related features together into single units so that they can be easy to maintain and understand and independent of the other package programs. This modularity allows them to share and reuse. In Go language, every package is defined with a different name and that name is close to their functionality like “strings” package and it contains methods and functions that only related to strings. 

## Golang Standard Library

* https://pkg.go.dev/std

## Golang search packages or symbols

* https://pkg.go.dev/

## Important Points

1. Import paths: In Go language, every package is defined by a unique string and this string is known as import path. With the help of an import path, you can import packages in your program. For example:

```go
import "fmt"
```

This statement states that you are importing an fmt package in your program. The import path of packages is globally unique. To avoid conflict between the path of the packages other than the standard library, the package path should start with the internet domain name of the organization that owns or host the package. For example:

```go
import "geeksforgeeks.com/example/strings"
```

2. Package Declaration: In Go language, package declaration is always present at the beginning of the source file and the purpose of this declaration is to determine the default identifier for that package when it is imported by another package. For example:

```go
package main
```

3. Import declaration: The import declaration immediately comes after the package declaration. The Go source file contains zero or more import declaration and each import declaration specifies the path of one or more packages in the parentheses. For example:

```go
// Importing single package
import "fmt"

// Importing multiple packages
import(
"fmt"
"strings"
"bytes"
) 
```

When you import a package in your program you’re allowed to access the members of that package. For example, we have a package named as a “sort”, so when you import this package in your program you are allowed to access sort.Float64s(), sort.SearchStrings(), etc functions of that package.

4. Blank import: In Go programming, sometimes we import some packages in our program, but we do not use them in our program. When you run such types of programs that contain unused packages, then the compiler will give an error. So, to avoid this error, we use a blank identifier before the name of the package. For example:

```go
import _ "strings"
```

It is known as blank import. It is used in many or some occasions when the main program can enable the optional features provided by the blank importing additional packages at the compile-time.

5. Nested Packages: In Go language, you are allowed to create a package inside another package simply by creating a subdirectory. And the nested package can import just like the root package. For example:

```go
import "math/cmplx"
```

Here, the math package is the main package and cmplx package is the nested package.

6. Sometimes some packages may have the same names, but the path of such type of packages is always different. For example, both math and crypto packages contain a rand named package, but the path of this package is different, i.e, math/rand and crypto/rand.

7. In Go programming, why always the main package is present on the top of the program? Because the main package tells the go build that it must activate the linker to make an executable file.