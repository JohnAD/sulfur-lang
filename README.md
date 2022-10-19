![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")


Sulfur is a programming language.

# Good Goals

The following are the major goals for the language. These goals are what set this language apart from existing ones.

### Fast and Small Object Code

* `FLATTEN` We remove all the procedures ...

  aka *full source deconstruction and extensive multi-pass code removal.*

  [more](compilation-goals.md#flatten)

* `CTRUN` ... remove functions with resolved inputs ...

  aka *upgrade `var` to `let` to `const`*

  aka *run crude interpreter on non-variable parameters*

  [more](compilation-goals.md#ctrun)

* `ALGO-PROC` ... and then we put functions back in

  aka *Algorithmic restatement of procedures.*

  [more](compilation-goals.md#algo-proc)

### Accessible

* `ML-CODE` Multi-lingual source code

  [more](accessible-goals.md#ml-code)

* `ML-STR` Multi-lingual string handling

  This includes both compile-time translation and run-time translation.

  [more](accessible-goals.md#ml-str)

* `GIT` Git-oriented 

  aka line/file/directory oriented

  [more](accessible-goals.md#git)

### Scalable

* `PROTOCOL` Common protocol support

  [more](scalable-goals.md#protocol)

* `FRAMEWORKS` Frameworks for common uses inside standard library

  [more](scalable-goals.md#frameworks)

* `TYPE-VERSIONING` Predictable and strict type versioning

  [more](scalable-goals.md#type-versioning)

Frankly, of the three, the type versioning is the most important in a 100,000+ line ecosystem.

### Predictable

* `STATEFUL-VARS` Stateful handling of variables with log wrapping

  [more](predictable-goals.md#type-versioning)

* `INLINE-ERR` In-line error handling

  [more](predictable-goals.md#type-versioning)

## Minor Goals

[This document](minor-goals.md) also describes some additional minor goals. Any of these might change or be ignored for various practical reasons.

# Notable Downsides

Life is about trade-offs. There are definitely some downsides to this language.

* Somewhat larger vocabulary and learning time.

  There are more reserved keywords and structures than most languages. Sulfur is attempting to gleam the "intent" of the code being written. The different versions of conditionals and loops help with this.

  Simply picking up this language and getting to work will be pretty easy if you already know a few other languages. But getting it to work _well_ might take a while.

* Recursion is forbidden.

  The nature of the compilation process does not support recursion. Thus, some algorithms will be less intuitive to the programmer.

  From the beginning, the compiler will detect recursion and will throw an error when detected.

* Early exit functions are forbidden.

  Loops, routines, and other structures in the language have a specific beginning, middle, and end. As such, common statements such as `return`, `continue`, and `break` are now allowed. This will force the programmer to refactor more often to keep code readable.

# Target Audience

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

In general, the versioning follows the semver standard. More specifically though, the version is in the form of `<year>.<quarter>.<minorversion>`.
