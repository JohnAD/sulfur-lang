![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# Predictable

## [STATEFUL-VARS]
### Stateful handling of variables

Any variable will be in one of the following states:

* `valued` (both empty and non-empty)
* `null` (unknown)
* `void` (non-existant)
* `errored` (details)

The "default" value for a variable is "valued and empty".

Having state helps prevent side effects and makes the variable's nature "more obvious".

Additionally, a variable has a 5-character "reason" attribute. It is short but open-ended. Use it for whatever you need. However, I recommend not passing operational values back in it. That is a bit too meta. It's purpose is describe "why" it has the current value.

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
```

If writing a function and you wish to indicate to the caller that "something didn't work" in a returned value, should you return an error or should you do something else? Ask yourself this: is it not working completely unexpected or a common thing.

If completely unexpected, then return as an error.

If the reason is "this does not apply", then return a `void`.

If the reason is anything else, then return a `null`. If additional data is needed, set a value in the `reason`.

If more common than that, I'd recommend returning it as null. If the caller wants more detail, then consider returning a tuple with a string containing the reason. Only return `void` if the reason is truly "this does not apply".

It will strike some folks that this is a lot of overhead for every single variable. But keep in mind that code removal will remove anything that is not used. If not using state; it will dissapear. If not using the `reason` attribute, it will dissapear also: both the code and memory used.

## [FUNC-LOG]
### A windowed functional in-thread logger

In most languages and/or libraries a logger creates side effects and yet it is used throughout the program. Allowing a logger to be like that has three severe consequences in this language:

1. the logger itself can be a source of side effects modifying the behavior of the code in a given environment, and 
2. the allocation of memory is not predictable at compile-time, and
3. any code that uses the logger cannot, by definition, allow for full code removal.

So, this language has a tightly bound "windowed" logger that bypass these effects. To wit:

1. the quantity of logs and the max message size is set at compile-time. This allows for pre-allocation. When the limit is exceeded, the log window rolls, discarding the oldest logs first.
2. any "character-mix" of log can be pre-assigned to a export filter (at compile-time). That way, only those logs match the filter will prevent code-removal.

A secondary benefit is that third-party libraries can confidently generate logs without concern for impact on the calling code.

A third-pary libary should never write logs to an external location. That decision should be left to the user of the library. The sole exception is when the purpose of the library is to act as a log writer.

A log will store 5 things:

**message**

> The message of the log itself. It is basically a string of limited pre-allocated length. Length determine at compile-time.

**audience**

> an enum determining the expected scope of recipient. The current values are `debugger` (10), `admin` (20), `user` (30), and `public` (40). In this context, `user` means the end-user of the program if there is one.

**nature**

> The good/bad/neutral nature of the log. The current values are `success` (10), `info` (20), `warning` (30), and `danger` (40).

**priority**

> the importance of the log. The current values are `trivia` (10), `normal` (20), `exception` (30), and `emergency` (40).

**context**

> a string limited to 128 characters representing the general "context" of the message. When in a library, the context should be the name of the library.

The default for a type-driven exception, such as divide-by-zero, is:

```text
audience = admin
nature = danger
priority = exception
context = <library-name> or filename
```

All of these can be overridden in code. And, just like variable wrapping, if the logger is never output or used for a decision that is output, then the entire logging system will be eliminated during compilation.

An example:

```sulfur
#! sulfur src 2022.0.1 en
using std::hdti [[ actor Terminal as t ]]

bind logger {{
  {{
    audience = [[public, user]]  # we will ONLY send messages to public or user to the ...
  }} = t.log_binder              # ... terminal's `log_binder` function.
}}

var a :byte = 0
var b :byte = 0

a_str = t.input("enter a: ")
b_str = t.input("enter b: ")

a = byte(t.input("enter"))

c = a / b

log("the division has happened")  # code removal will make this disappear since nothing uses it

if c.is_errored() [[
  log("b can't be zero", audience = user, nature = danger)
]] else [[
  log("a / b = " & str(c), audience = user, nature = success)
]]
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
