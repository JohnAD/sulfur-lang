![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# minor goals

The following are _minor_ goals for the language. They can take a back seat to more practical concerns, but it would be very nice if we could meet these goals also.

## [ONEWAY]
### There should only be one way to do things

Having five different ways to do the same things makes for less predictable code.

However, since programming is partly an art, some flexibility needs to be allowed for. Such exceptions should be a very purposeful thing and justified.

## [NAMING]
### Naming conventions are based on the _human_ language

Naming standards for identifiers such as types and variables should be based on the human language the source code is following.

Underscores are a substitute for spacing if spacing is needed by the language.

Violating these principles will generate a "warning" rather than an error, however.

#### English

    In English, proper names start with an uppercase letter followed by lower case. Common nouns are all lowercase. Acronyms are all uppercase.

    Since a type is a "commonly shared name", it is a noun. Most variables are lower case unless they happen to contain a proper name or acronym.

    A singleton directly shared between threads would be a proper name since it is systemically unique (and thus start with an upper case letter). In other words, `actor` and `thread` libraries start with an upper case letter. Otherwise libraries are generally all lower case.

    Example variable names:

    ```sulfur
    var pet_fish = fish( name = "Zippy" )                   # pet_fish is a noun for any pet_fish
    var fish_belonging_to_Bob = fish( name = "Foo" )        # Bob is a person's name, so that word starts with an uppercase letter.
    var doc = JSON()                                        # JSON is an acronym for JavaScript Object Notation
    ```

    The `sulfur` extension (and command-line name) are a known exceptions to this rule as keeping extensions and utilities lowercase is a common convention.

## [COMMENTS]
### Comments and spacing are code and part of the language.

Functions, modules, etc. should have an explicit structure for coding a description (required) and providing details about parameters (optional), etc.

Support for "vertical" commenting with blank lines.

  * Optional one or zero blank lines inside code blocks. Such separation can guide/group the reader of the source.
  * Exactly two lines between functions, procedures, etc. This vertacle separation makes it easier for humans to parse the file at a glance.
  * Comment blocks. This allows for explanitory text or sections that are not part of the general description.

In general, comments are a permitted exception to the ONEWAY minor goal.

Comments should be written in GFM Markdown since that is the emerging standard for the industry.

## [TESTING] 
### Strong testing support.

Specifically, unit-level testing is explicitly part of the language and compiler. Other forms of testing are encouraged but not embraced directly.

```
main.test.sulfur
```

## [HUMAN-LANG-EXPORT]
### Tools for exporting to other human languages

`lang_convert main.code.sulfur es -o=primero.code.sulfur`

Basically, the tool will invert the role of the source file (`*.code.sulfur`) and the translation document (`*.tr.sulfur`).

## [PROG-LANG-EXPORT]
### Tools for exporting to other coding languages

`code_convert main.code.sulfur js -o=main.js`

This is NOT the same thing as using a language as a compiler target ( see the [next section](#obj-targets) ). The exported code should not do code removal and should reflect the structure and intend of the sulfur code.

The primary goal is human readability; so sometimes it might have to resort leaving legacy code in comment blocks.

## [OBJ-TARGETS]
### Compiler can target other languages not just object code.

`sulfur main -env=linux -targ=js`

This is NOT the same thing as exporting to a language (see the [previous section](#prog-lang-export) ). The compiled code is in the other language but there is *no expectation of human readability*. The audience for the generated code is another compiler.

At first, JavaScript and C are good targets as they are very cross-platform in affect. Nearly every CPU ever made has support from a C compiler. And JavaScript, of course, is the lingua franca of web browsers.

----

[return to README.md](README.md)