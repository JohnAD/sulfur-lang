![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")


Sulfur is a programming language.

## Good Goals

The following are the major goals for the language. These goals are what set this language apart from existing ones.

### Fast and Small Object Code

* [FLATTEN] We remove all the procedures ...

  aka *full source deconstruction and extensive multi-pass code removal.*

* [CTRUN] ... remove functions with resolved inputs ...

  aka *upgrade `var` to `let` to `const`*

  aka *run crude interpreter on non-variable parameters*

* [ALGO-PROC] ... and then we put functions back in

  aka *Algorithmic restatement of procedures.*

### Accessible

* [ML-CODE] Multi-lingual source code

* [ML-STR] Multi-lingual string handling

  This includes both compile-time translation and run-time translation.

* [GIT] Git-oriented 

  aka line/file/directory oriented

### Scalable

* Common protocol support

* Frameworks for common uses inside standard library

* Predictable and strict type versioning

### Predictable

### (dependable) Stateful handling of variables with log wrapping

A variable will be in one of the following states:

* `valued` (both empty and non-empty)
* `null` (unknown)
* `void` (non-existant)
* `errored` (details)

The "default" value for a variable is "valued and empty".

A variable also has a "log" of events that can be attached to it. If the log isn't actually output somewhere, compile-time code removal will remove both the logging strucure and function.

Having state and log wrapping helps with preventing side effects.

Example of state:

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

### (dependable) In-line error handling

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

## Notable Downsides

Life is about trade-offs. There are definitely some downsides to this language.

### Somewhat larger vocabulary and learning time.

There are more reserved keywords and structures than most languages. Sulfur is attempting to gleam the "intent" of the code being written. The different versions of conditionals and loops help with this.

Simply picking up this language and getting to work will be pretty easy if you already know a few other languages. But getting it to work _well_ might take a while.

### Recursion is forbidden.

The nature of the compilation process does not support recursion. Thus, some algorithms will be less intuitive to the programmer.

From the beginning, the compiler will detect recursion and will throw an error when detected.

### Early exit functions are forbidden.

Loops, routines, and other structures in the language have a specific beginning, middle, and end. As such, common statements such as `return`, `continue`, and `break` are now allowed. This will force the programmer to refactor more often to keep code readable.

## Target Audience

The target audience for this language is:

> skilled computer programmers wishing to write large general-purpose applications

It is NOT designed for:

* embedded systems programming

  While it is certainly possible given it's heap management, that is not the focus.

* new programmers

  Sadly, the learning curve of this language is a bit too high for that. I recommend starting with something like Python.

* scripting

 It's directory/file/line orientation makes it a very poor scripting language. You could use it, however, to make utilities that are called by a scripting language such as `bash` or `perl`.

## Development Stages

There are four stages planned for creating the language:

1. 2023.0.n: Bootstrap: Compiler is written in Go and uses LLVM. Only the core of the language and a small number of standard libraries will work. I'm not expecting (or hoping) for help at this stage. Mostly works with small test apps.

2. 2025.0.n: Momentum builds: Rewrite the language compiler in Sulfur itself. Write the first framework (likely "web service"). The first database client library written. Start of the optional 3rd party library system. Support formatting for editors and repository hosting systems.

3. 2027.0.n: Base Building: Write more frameworks and test them together. Start pulling in standards for library writing. Build a C / Rust interface (*every language* should interface with C.) Cross-compile to WASM in addition to LLVM (unless LLVM starts doing that on it's own.)

4. 2029.0.n: Larger Adoption: Fine-tune the System for general production systems and mainstream use. Support cross-compilation to Swift and Kotlin. Add linter and code-suggestion support for editors. Finally concentrate on improving compile-time.

In general, the versioningh follows the semver standard. More specifically though, the version is in the form of `<year>.<quarter>.<minorversion>`.
