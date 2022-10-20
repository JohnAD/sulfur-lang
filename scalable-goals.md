![elemental-sulfur](https://upload.wikimedia.org/wikipedia/commons/thumb/8/88/Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg/220px-Sulfur_-_El_Desierto_mine%2C_San_Pablo_de_Napa%2C_Daniel_Campos_Province%2C_Potos%C3%AD%2C_Bolivia.jpg "Elemental Sulfer as seen on Wikipedia. Credit: Iifar")

# Scalable

## [PROTOCOL]
### Common protocol support

The users of an app are sometimes other machines. The *standard* library will support as many intermachine standards has possible to allow consistent communications with other machines.

* serializations/encodings: JSON, YAML, XML, BSON, JPG, PNG, WAV, PDF, etc.
* transport: UDP, TCP, HTTP/HTTPS, etc.
* templates/specs: OpenAPI, mustache, Jinja, etc.

This major goal is also supported by the [[PROTOBUF]](minor-goals.md#protobuf) direct compilation, a minor goal.

## [FRAMEWORKS]
### Frameworks for common uses inside standard library

Focusing community work on a common framework can have many benefits. As such, an attempt will be made to create "default" frameworks for potentially common uses of the language. Examples:

* Console/Shell
* Web Server
* Javascript Client
* 2D Game
* State Engine

To truly encourage general-purpose use and involvement, all frameworks should have a reasonable manner of adding "middleware" to expand it.

Frameworks are also specific to their operating environment. Some of the earliest frameworks are written for a limited half-duplex terminal that will often be used in examples:

* std::**hd_terminal_Linux_bash_x64** : a half-duplex UTF8 terminal that runs in linux from a bash shell
* std::**hd_terminal_Linux_x11_x64** : a half-duplex UTF8 terminal that runs in linux as a GUI in a X11 shell
* std::**hd_terminal_web_wasm** : a half-duplex UTF8 terminal that runs in a web browser as a single page WASM app and HTML doc
* std::**hd_terminal_Win10_x64** : a half-duplex UTF8 terminal that runs in MS Windows 10 64-bit GUI (or compatible)

In this particular case, the hdterminal has the same methods and functions. The 'hdterminal' variations are cross-platform. But not all frameworks are expected to be cross-platform.

For the half-duplex terminal, the cross-platform reference is simply:

* std::**hd_terminal**

The language willl support passing the final cross-platform framework details via command-line to the compiler.

All `src` documents must have a framework chosen. Libraries need it if they are limited to specific frameworks.

## [BOTH-I-R]
### Supports both imperative and reactive models

Some operating environments are [Reactive](https://en.wikipedia.org/wiki/Reactive_programming) (such as the DOM in web programming written in JS). Some are [Imperative](https://en.wikipedia.org/wiki/Imperative_programming) (such as a Windows App written in C#).

In Sulfur the distinction is determined by the framework chosen. Internally, the compiler will adjust it's behavior based on the two models.

## [TYPE-VERSIONING]
### Predictable and strict type versioning

Compiling on another machine with the same source code should produce EXACTLY the same native object code. Yet, any library or type managment that allows for wildcards can't really insure this. This language solves this problem by concentrating on the `type` libraries and enforcing a predictable way to convert to/and from shared types between the app's code and it's various dependent libraries.

Any `type` library written must handle version differences with past implementations. The only exception is a semver `major` difference. Those are assumed to be ALWAYS incompatible.

For example if:

* the main program uses the `foo` and `bar`
* the `foo` library uses `fish` type version `1.0.1`
* the `bar` library uses `fish` type version `1.0.3`

Then any data passed by the main between `foo` and `bar` will be handled in a controlled manner because all libraries *MUST* support the conversion. For example, `fish` version `1.0.3` would be something like:

```sulfur
#! sulfur type 2022.0.1 en
#% type_library fish 1.0.3

var common_name :str = ""
var species_name :str = ""
var fin_color :str = ""          # fin_color added in version 1.0.2

conversions {{
  ["1.0.1", "1.0.2"] = {{
    up = [[
      # unless you declare differently,
      #   `old` is the older object ("1.0.1" in this case)
      #   `this` is the newer object ("1.0.2" in this case)
      this.common_name = old.common_name
      this.species_name = old.species_name
      this.fin_color = null<str>
    ]]
    down = [[
      old = error("cannot convert from Fish 1.0.2 to Fish 1.0.1")
      # alternatively, we could have silently "dropped" the fin_color (no error):
      #    old.common_name = this.common_name
      #    old.species_name = this.species_name
      # or we could have thrown a compile-time error:
      #    compiler_stop("cannot convert from Fish 1.0.2 to Fish 1.0.1")
    ]]
  }}
  ["1.0.2", "1.0.3"] = void     # setting to void means that nothing needs to be done
}}
```

In the above example, the compiler would run each conversion from 1.0.1 -> 1.0.2 and 1.0.2 -> 1.0.3 sequentially when upconverting from 1.0.1 to 1.0.3. Failing to find a conversion path would result in a compile-time error. If multiple paths are found, the shortest one is chosen. Be careful when allowing multiple paths.

It is possible that multiple paths may be forbidden; it is a point to discuss.

----

[return to README.md](README.md)
