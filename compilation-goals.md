![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# Fast and Small Object Code

## [FLATTEN]
### We remove all the procedures ...

aka *full source deconstruction and extensive multi-pass code removal.*

All functions, procedures, and other abstractions are "flattened" into a serial instruction stream prior to code removal. The flattening
can include some loops.

Many passes of the source tree remove as much code and variables as possible. Even object members are removed if not used. The order
of some statements are also reorganized.

This makes the program faster and use less memory.

Example:

This (psuedo-code)

```psuedocode
proc hello(a, b)
  if (b == true)
    echo "Hello " & a
  else
    echo "Hi " & a

hello("Joe", true)
n = "Sally"
hello(n, true)
hello("Larry", false)
```

becomes

```psuedocode
echo "Hello Joe"
echo "Hello Sally"
echo "Hi Larry"
```

## [CTRUN]
### ... run any functions passed constants at compile-time ...

## [ALGO-PROC]
### ... and then we put functions back in

aka *Algorithmic restatement of procedures.*

After code removal, procedures as added back into the source based on the reduced source tree. The grouping of those procedures is determined by reproducable algorithms and the original organization in the source code is ignored.

This, in theory, makes the final object code smaller.

*TBD: create a small-enough example that still survives the code-removal process*


----

[return to README.md](README.md)