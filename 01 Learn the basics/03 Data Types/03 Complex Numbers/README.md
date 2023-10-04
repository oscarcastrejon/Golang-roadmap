# Golang Roadmap

# Complex Numbers

Complex Numbers are of two types:

| Type | Size |
| :---:   | :---: |
| complex64 |	Both real and imaginary part are float32
| complex128 |	Both real and imaginary part are float64

The default complex type is complex128

Initialization

Complex Numbers can be initialized in two ways

Using complex function. It has below signature. Do make sure that both a and b should be of same type , meaning either they both should be float32 or both should be float64
```
complext(a, b)
```
Using the shorthand syntax. This is used when creating a complex number with direct numbers. The complex type created using below method will be of type complex128 if type is not specified
```
a := 5 + 6i
```
