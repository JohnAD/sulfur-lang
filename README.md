Sulfur is a compiled programming language.

## Key Positive Characteristics

### Full source deconstruction and extensive multi-pass code removal.

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

### Algorithmic restatement of procedures.

After code removal, procedures as added back into the source based on the reduced source tree. The grouping of those procedures is determined by reproducable algorithms and the original organization in the source code is ignored.

This, in theory, makes the final object code smaller.

*TBD: create a small-enough example that still survives the code-removal process*

### Multi-lingual source code.

Programmers are human. Humans speak different languages.

Examples:

```sulfur
#! sulfur src 2022.0.1 en
using module std::vterm as t
t.print( "Hello world!\n" )
```

```sulfur
#! sulfur src 2022.0.1 zh-cn
使用模块 水平::vterm 作为 屏幕
屏幕.节目( "你好世界\n" )
```

(note: I don't know Mandarin so the above code example might be nonsensical. I apologize for that.)

First three core languages:

* English `en`
* Mandarin Chinese `zh-cn`
* Spanish `es`

Later goal of ten core languages:

* English `en`
* Mandarin Chinese `zh-cn`
* Spanish `es`
* Arabic `ar`
* Hindi `hi`
* Russian `ru`
* Portuguese `pt`
* French `fr`
* Bengali `bn`
* German `de`
* Japanese `ja`

### Multi-lingual UTF8 string handling.

The users of an app are sometimes human. Humans speak different languages. The support for both compile-time and run-time selection of translation will be supported in the very core and structure of the language.

Examples:

a file called `main.src.sulfur`:

```sulfur
#! sulfur src 2022.0.1 en
using module std::vterm as t

# compile this program with `--lang=es` to get the spanish version of the app
  
t.print( "Hello world!\n" )          # since the declaration line is `en`, this defaults to `en`
const day<str> = "Tuesday"
t.print( "Today is ${day}.\n"<en> )  # explicit declaration
```

a file called `main.tr.sulfur`:

```sulfur
#! sulfur tr 2022.0.1 en
  
"Hello world!\n" = [
  "es" = "Hola Mundo!\n"
  "zh-cn" = "你好世界\n"
]
"Today is ${day}.\n" = [
  "es" = "Hoy es ${day}.\n"
  "zh-cn" = "今天是星期${day}\n"
]
"Tuesday" = [
  "es" = "martes"
  "zh-cn" = "二"
]
# You can include variable names as well, though they will not be part of the final object file
# use the fake language of "src" to do this.
"day"<src> = [
  "es" = "día"
  "zh-cn" = "星期几"
]
```

a run-time translation version of the `main.src.sulfur`:

```sulfur
#! sulfur src 2022.0.1 en
using module std::thread_vars as g
using module std::vterm as t
var choice<str> = t.input("enter language code:")
g.lang(choice)
t.print( "Hello world!\n".$ )
const day<str> = "Tuesday"
t.print( "Today is ${day}.\n".$ )
```

### Standard protocol support.

The users of an app are sometimes other machines. The *standard* library will support as many intermachine standards has possible to allow consistent communications with other machines.

* serializations/encodings: JSON, YAML, XML, BSON, JPG, PNG, WAV, PDF, etc.
* transport: UDP, TCP, HTTP/HTTPS, etc.
* templates/specs: OpenAPI, mustache, Jinja, etc.

### Line-oriented.

Specifically, it is optimized for use with GIT for predictable and visible changes for other and future programmers to see.

A sequential series of elements or statements, when separated vertically, does not have "separators" such as commas or semicolons.

The language isn't bracket-oriented or-indentation oriented: it is BOTH. Failure to properly indent **or** close bracketed operators will result in a compile-time error.

good:

```sulfur
#! sulfur src 2022.0.1 en
if true [[
  var a<byte> = 10
  var b<byte> = 99
]]
```

good:

```sulfur
#! sulfur src 2022.0.1 en
if true [[ var a<byte> = 10l ; var b<byte> = 99 ]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
if true 
[[
  var a<byte> = 10
  var b<byte> = 99
]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
if true [[
var a<byte> = 10
var b<byte> = 99
]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
if true [[ var a<byte> = 10
  var b<byte> = 99
]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
if true
  var a<byte> = 10
  var b<byte> = 99
```

### Predictable object code contruction and strict library handling.

Specifically, compiling on another machine with the same source code produces EXACTLY the same native object code. To do this,
library handling and version dependencies are very strict and reproducable.

Any library written must handle version different with past implementations. The only exception is `major` version differences.

For example if:

* the main program uses the `foo` and `bar`
* the `foo` library uses `Fish` version `1.0.1`
* the `bar` library uses `Fish` version `1.0.2`

Then any data passed by the main between `foo` and `bar` will be handled in a controlled manner because all libraries *MUST* support the conversion. For example, `fish` version `1.0.2` would be something like:

```sulfur
#! sulfur type 2022.0.1 en
#% type_library Fish 1.0.2

common_name<str> = ""
species_name<str> = ""
fin_color<str> = ""          # fin_color added in version 1.0.2

convert `1.0.1` {{
  parameters = {
    older<Fish>
  }
  up = [[
    common_name = older.common_name
    species_name = older.species_name
    fin_color = null<str>
  ]]
  down = [[
    older = error("cannot convert from Fish 1.0.2 to Fish 1.0.1")
    # alternatively, we could have simply "dropped" the fin_color; or we could have thrown a compile-time error.
  ]]
}}
```

### Stateful handling of variables with log wrapping.

A variable will one of the following states:

* `valued` (both empty and non-empty)
* `null` (unknown)
* `void` (non-existant)
* `errored` (details)

The "default" value for a variable is "valued and empty".

A variable also has a "log" of events that can be attached to it. If the log isn't actually output somewhere, compile-time code removal will remove both the logging strucure and function.

Having state and log wrapping helps the language prevent side effects.

Example of state:

```sulfur
#! sulfur src 2022.0.1 en
using std::integers {{ type int32 as int }}

vars {{
  a_filled<int> = 1   # valued
  a_empty<int> = 0    # also valued, but also empty
  a_unknown<int> = null
  a_non = void
  a_error = error.ValueOutOfRange
}}

var five = 5<int>

assert ( five + a_filled ) == 6<int>
assert ( five + a_empty ) == 5<int>
assert ( five + a_unknown ) == unknown<int>         # 5 plus a mystery number is just a bigger mystery number
assert ( five + a_none ) == 5<int>                  # 5 was never added to
assert ( five + a_error ) == error.CannotOperateOnError  # adding an error to five results in a new error

assert ( five + a_error ).log( 0 ).type = log_type.error  # this is original ValueOutOfRange error
assert ( five + a_error ).log( 1 ).type = log_type.error  # this is the new CannotOperateOnError error
assert ( five + a_error ).log.size == 2
```

### In-line error handling

Sulfur does not allow the simulated "throwing of exceptions" that many other languages do. Handling errors is considered part of the program's business logic.

The benefit, philosophically, is that error handling is dealt with explicitly. It is, admittedly, a philosophical value not a provable one.

```sulfur
#! sulfur src 2022.0.1 en
using std::integers {{
  type int32 as int
}}

var a<int> = 44
var b<int> = 0

var c<int> = a / b 

assert c.is_errored()
assert c == error.DivideByZero
assert ! c.is_valued()
```

### Frameworks for common uses in standard library

Focusing community work on a common framework can have many benefits. As such, an attempt will be made to create "default" frameworks for potentially common uses of the language. Examples:

* Web Server
* Javascript Client
* 2D Game
* State Engine

To truly encourage general-purpose use and involvement, all frameworks should have a reasonable manner of adding "middleware" to expand it.

## Notable Downsides

### Somewhat larger vocabulary and learning time.

There are more reserved keywords and structures than most language. Sulfur is attempting to gleam the "intent" of the code being written. The different versions of conditionals and loops helps with this.

### Recursion is forbidden.

The nature of the compilation process does not support recursion. Thus, some algorithms will be less intuitive to the programmer.

From the beginning, the compiler will detect recursion and will throw an error when detected.

### Early exit functions are forbidden.

Loops, routines, and other structures in the language have a specific beginning, middle, and end. As such, common statements such as `return`, `continue`, and `break` are now allowed. This will force the programmer to refactor more often to keep code readable.

## Target Audience

The target audience for this language is skilled computer programmers wishing to write general-purpose applications such as mobile app or websites.

It is not designed for embedded systems programming, though that is certainly possible given it's heap management.

It is not designed for new programmers. Sadly, the learning curve of this language is a bit too high for that.

It is not designed to be a scripting language. In fact, it's directory/file/line orientation makes it a really poor scripting language. You could use it, however, to make utilities that are called by a scripting language.

