![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# Predictable

## [STATEFUL-VARS]
### Stateful handling of variables with log wrapping

Any variable will be in one of the following states:

* `valued` (both empty and non-empty)
* `null` (unknown)
* `void` (non-existant)
* `errored` (details)

The "default" value for a variable is "valued and empty".

A variable also has a "log" of events that can be attached to it. If the log isn't actually output somewhere, compile-time code removal will remove both the logging structure and supporting code.

Having state and log wrapping helps prevent side effects.

Example:

```sulfur
#! sulfur src 2022.0.1 en
using std::integers [[ type int32 as :int ]]

vars {{
  a_filled :int = 1   # valued
  a_empty :int = 0    # also valued, but also empty
  a_unknown :int = null
  a_non :int = void
  a_error :int = error.ValueOutOfRange
}}

var five :int = 5

assert ( five + a_filled ) == 6<int>
assert ( five + a_empty ) == 5<int>
assert ( five + a_unknown ) == unknown<int>         # 5 plus a mystery number is just a bigger mystery number
assert ( five + a_none ) == 5<int>                  # 5 was never added to
assert ( five + a_error ) == error.CannotOperateOnError  # adding an error to five results in a new error

assert ( five + a_error ).log( 0 ).type = log_type.error  # this is original ValueOutOfRange error
assert ( five + a_error ).log( 1 ).type = log_type.error  # this is the new CannotOperateOnError error
assert ( five + a_error ).log.size == 2
```

## [INLINE-ERR]
### In-line error handling

Sulfur does not allow the simulated "throwing of exceptions" that many other languages do. Handling errors is considered part of the program's business logic.

The benefit, philosophically, is that error handling is dealt with explicitly. Admittedly, this is a philosophical value not a provable one.

```sulfur
#! sulfur src 2022.0.1 en
using std::integers [[
  type int32 as :int
]]

var a :int = 44
var b :int = 0

var c :int = a / b 

assert c.is_errored()
assert c == error.DivideByZero
assert ! c.is_valued()
```

----

[return to README.md](README.md)
