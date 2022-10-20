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

> In English, proper names start with an uppercase letter followed by lower case. Common nouns are all lowercase. Acronyms are all uppercase.
> 
> Since a type is a "commonly shared name", it is a noun. Most variables are lower case unless they happen to contain a proper name or acronym.
> 
> A singleton directly shared between threads would be a proper name since it is systemically unique (and thus start with an upper case letter). In other words, `actor` and `thread` libraries start with an upper case letter. Otherwise libraries are generally all lower case.
> 
> Example variable names:
> 
```sulfur
var pet_fish = fish( name = "Zippy" )                   # pet_fish is a noun for any pet_fish
var fish_belonging_to_Bob = fish( name = "Foo" )        # Bob is a person's name, so that word starts with an uppercase letter.
var doc = JSON()                                        # JSON is an acronym for JavaScript Object Notation
```
> 
> The `sulfur` extension (and command-line name) are known exceptions to this rule as keeping extensions and utilities lowercase is a common convention.

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

This is NOT the same thing as using a language as a compiler target ( see the [next section](#obj-targets) ). The exported code should not do code removal and should reflect the structure and intent of the Sulfur code.

The primary goal is human readability; so sometimes it might have to resort leaving legacy code in comment blocks.

## [OBJ-TARGETS]
### Compiler can target other languages not just object code.

`sulfur main -env=linux -targ=js`

This is NOT the same thing as exporting to a language (see the [previous section](#prog-lang-export) ). The compiled code is in the other language but there is *no expectation of human readability*. The audience for the generated code is another compiler.

At first, JavaScript and C are good targets as they are very cross-platform in affect. Nearly every CPU ever made has support from a C compiler. And JavaScript, of course, is the lingua franca of web browsers.

## [SELF-FILE]
### Self-contained transparent file contents

If someone were to copy the content of a file and send it to you, without any context, could you interpret it? The answer _should_ be yes. But this is almost never true in the current suite of languages.

Specifically:

* You should be able to see what _kind_ of file it is. Every source file must start with a line like `#! sulfur src 2022.0.1 en`. In one short bit of text, I know that this is the source code for a Sulfur program version 2022.0.1 written in English. No need to guess.

* Do not allow wild-cards in library versions. Either an explicitly stated version is used or it defaults to match the version of sulfur in the top line. There is no need to lookup some kind of project table or mapping file to determine which library version is going to used. As it happens, this also matches major goal [[TYPE-VERSIONING]](scalable-goals.md#type-versioning).

* Do not allow "wild-card" imports of identifiers. If you see an identifier called `foo` you should know exactly where that identifier was defined with nothing but the contents of the file itself. If `foo` came from a library, it is either explictly named in the `using` statement, or it is referenced by something that is explicitly named.

  For example, if you have a two libraries called 'colors' and 'moods', the following would be a bad thing:

  ```sulfur
  #! sulfur src 2022.0.1 en
  using colors [[ type color ]]
  using moods [[ type mood ]]

  var a = red
  var h = blue                    # a "blue" mood, or a "blue" color?
  var y = zork                    # I have no idea which library had this.
  var z = happy
  ```

  There are two solutions. First, you could simply pull in all the different words and deal with aliases.

  ```sulfur
  #! sulfur src 2022.0.1 en
  using colors [[
    type color 
    color [ red, blue, zork ]    # I don't know what "zork" is, but I do know it is from the color library.
  ]]
  using moods [[
    type mood 
    mood happy
    mood blue as feeling_blue    # the "as" alias prevents a name-conflict error
    mood coasting
  ]]

  var a = red
  var h = feeling_blue
  var y = zork
  var z = happy
  ```

  Or, you could make indirect references:

  ```sulfur
  #! sulfur src 2022.0.1 en
  using colors [[ type color ]]
  using moods [[ type mood ]]

  var a = color.red
  var h = mood.blue
  var y = color.zork
  var z = mood.happy
  ```

  Or, mix and match the methods. The key is that the language will not let you do the equivalent of an "import all" from a library.

* A point of discussion: should you also know that you are looking at the whole file? The start of the file is self-evident with the top line declaration. But the bottom is not. Should we allow something like `----` or `#@` to finish the file? Should it be _required_ as the last line?

  An example if done:

  ```sulfur
  #! sulfur src 2022.0.1 en

  using colors [[ type color ]]
  using moods [[ type mood ]]

  var a = color.red
  var h = mood.blue
  var z = mood.happy

  ----
  ```

## [PROTOBUF]
### Direct compilation of ProtoBuf files

ProtoBuf is a very conventient way to specify a serialization schema; a very common thing to share in an API specification. The protobuf schema is support by many language and I'm confident this will also. But, it would be nice if it were taken to the next level: allowing the schema to be read directly into code as part of the language.

Essentially, a protobuf file could be used as part of a `type` library.

Example:

First, the type library wrapping:

`person.type.sulfur`

```sulfur
#! sulfur type 2022.0.1 en
#% type_library person 1.0.1

protobuf "../models/person.protobuf"

conversions {{
}}

method `$` = {{
  description "convert to a readable string"
  returns = {{
    final_string :str
  }}
  body {{
    final_string = "Person( "
    var center_missing = true
    if self.name.has_value() [[
      final_string &= self.name.repr & " "
      center_missing = false
    ]]
    if self.id.has_value() [[
      final_string &= "id=" & self.id & " "
      center_missing = false
    ]]
    if self.email.has_value() [[
      final_string &= "<" & self.email & "> "
      center_missing = false
    ]]
    if center_missing [[
      final_string &= "*empty* "
    ]]
    final_string &= ")"
  }}
}}
```

And it directly references the file:

`person.protobuf`

```protobuf
message Person {
  optional string name = 1;
  optional int32 id = 2;
  optional string email = 3;
}
```

This minor goal also support the [[PROTOCOL]](scalable-goals.md#protocol) major goal.

----

[return to README.md](README.md)