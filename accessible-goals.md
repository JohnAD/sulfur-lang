![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# Accessible

## ML-CODE
### Multi-lingual source code

Programmers are human. Humans speak different languages.

Examples:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hdterminal as t

t.print( "Hello world!\n" )
```

```sulfur
#! sulfur src 2022.0.1 zh
#@ std::hdterminal 叫 屏幕

屏幕.节目( "你好世界\n" )
```

(note: I don't know Mandarin so the above code example might be nonsensical. I apologize for that.)

First three core languages:

* English `en-US`
* Chinese `zh-CN`
* Spanish `es-ES`

Later goal of ten core languages:

* English `en-US`
* Chinese `zh-CN`
* Spanish `es-ES`
* Arabic `ar-EG`
* Hindi `hi-IN`
* Russian `ru-RU`
* Portuguese `pt-BR`
* French `fr-FR`
* Bengali `bn-IN`
* German `de-DE`
* Japanese `ja-JP`

The if a language is used without a country code specification it defaults to the above list. For example, `en` defaults to be `en-US`.

## [ML-STR] 
### Multi-lingual string handling

The users of an app might be human. Humans speak different languages. The support for both compile-time and run-time translation of UTF8 strings will be supported in the very structure of the language.

Examples:

a file called `00_main.sulfur`:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal as t

# compile this program with `--lang=es` to get the spanish version of the app

action say_day {
  doc = {{
    """
      This action prints the day to the terminal.
    """
  }}
  parameters = {
    day<str> = "unknown"
  }
  code = {
    t.print( "Today is {{ day }}.\n"<en> )  # explicit declaration
  }
}

  
t.print( "Hello world!\n" )          # since the declaration line is `en`, this defaults to `en`
const today<str> = "Tuesday"
say_day(today)
```

a called `00_main.sulfur.zh.po`:

```sulfur
msgid "Hello world!\n"
msgstr "你好世界\n"

# {{ day }} is the current day of the week
msgid "Today is {{ day }}.\n"
msgstr "今天是星期${day}\n"

# notice that source code names and docs can be translated also
msgid "__say_day"
msgstr "说那一天"

msgid "__say_day.__parameters.__day"
msgstr "星期几"

msgid "__say_day.__docs.This action prints the day to the terminal."
msgstr "此操作将日期打印到终端。"
```

Here is a alternate version of the `00_main.sulfur` that does translation at run-time:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal as t

var choice :str = t.input("enter language code:")
t.set_language(choice)
t.print( "Hello world!\n".$ )             # the '.$' means "translate this to whatever the language is"
const day<str> = "Tuesday"
t.print( "Today is ${day}.\n".$ )
```

If translations occur at run-time, the all the translated strings are stored in the object code by default.

## `GIT`
### Git-oriented 

aka *line/file/directory oriented*

Specifically, it is optimized for use with GIT for predictable and visible changes for other and future programmers to see.

A sequential series of elements or statements, when separated vertically, does not have "separators" such as commas or semicolons. The helps with tracking line insertions/deletions.

The language isn't bracket-oriented or-indentation oriented: it is BOTH. Failure to properly indent **or** close bracketed operators will result in a compile-time error.

good:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal

if true [[
  var a<byte> = 10
  var b<byte> = 99
]]
```

good:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal

if true [[ var a<byte> = 10l ; var b<byte> = 99 ]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal

if true 
[[
  var a<byte> = 10
  var b<byte> = 99
]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal

if true [[
var a<byte> = 10
var b<byte> = 99
]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal

if true [[ var a<byte> = 10
  var b<byte> = 99
]]
```

compile-time error:

```sulfur
#! sulfur src 2022.0.1 en
#@ std::hd_terminal

if true
  var a<byte> = 10
  var b<byte> = 99
```

----

[return to README.md](README.md)
