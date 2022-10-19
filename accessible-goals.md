![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# Accessible

## ML-CODE
### Multi-lingual source code

Programmers are human. Humans speak different languages.

Examples:

```sulfur
#! sulfur src 2022.0.1 en
using std::hdti [[ actor terminal as t ]]
t.print( "Hello world!\n" )
```

```sulfur
#! sulfur src 2022.0.1 zh-CN
使用模块 水平::hdti [[ 演员 屏幕 叫 屏幕 ]]
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

## [ML-STR] 
### Multi-lingual string handling

The users of an app might be human. Humans speak different languages. The support for both compile-time and run-time translation of UTF8 strings will be supported in the very structure of the language.

Examples:

a file called `main.src.sulfur`:

```sulfur
#! sulfur src 2022.0.1 en
using std::hdti [[ actor Terminal as t ]]

# compile this program with `--lang=es` to get the spanish version of the app
  
t.print( "Hello world!\n" )          # since the declaration line is `en`, this defaults to `en`
const day<str> = "Tuesday"
t.print( "Today is {{ day }}.\n"<en> )  # explicit declaration
```

a directory-wide file called `directory.tr.sulfur`:

```sulfur
#! sulfur tr 2022.0.1 en
  
"Hello world!\n" = [
  "es" = "Hola Mundo!\n"
  "zh-cn" = "你好世界\n"
]

"Today is {{ day }}.\n" = [
  "es" = "Hoy es {{ day }}.\n"
  "zh-cn" = "今天是星期${day}\n"
  comment = "{{ day }} is the current day of the week"
]

"My name is [[ Bob ]].\n" = [
  comment = "[[ Bob ]] is a person's name"
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

Here is a alternate version of the `main.src.sulfur` that does translation at run-time:

```sulfur
#! sulfur src 2022.0.1 en
using std::hdti [[ actor Terminal as t ]]
var choice<str> = t.input("enter language code:")
tglob.lang(choice)
t.print( "Hello world!\n".$ )
const day<str> = "Tuesday"
t.print( "Today is ${day}.\n".$ )
```

If translations occur at run-time, the translated strings are stored in the object code by default.

## `GIT`
### Git-oriented 

aka *line/file/directory oriented*

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

----

[return to README.md](README.md)
