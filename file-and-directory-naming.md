# File and Directory Naming

# filenames

naming limits:
  form: base_name.extension
  32 UTF8 characters in base_name
  No punctuation or spacing other than underscore.
  No double-underscore pairs.
  No leading or trailing underscores.
  Leading numbers permitted, but not recommended.

filenames determine class name, so they are very imported

extensions limited to types recognized by the compiler. The exception to this rule:

1. data files pulled in at compile time can have any extension
2. source & content translations have an additional `.<lang>.po` extensions.

Note: `.pot` files are generated in the `__pot` directory (and subdirectories).

One or more source files that are compiler targets are prefixed with `00_`. Examples:

```bash
00_main.sulfur                   # `sulfer main` would start compilation here
00_server.sulfur                 # `sulfur server` would start compilation here
00_server.sulfur.es.po           # `sulfur server` would use this file for Spanish translation; should it be needed
00_some_utility.sulfur
```

Alternatively, `.<lang>.po` files can go into a subtending `__tr` directory. A 1:1 mapping is made. So, `./xyz.sulfur` would map to `./__tr/xyz.sulfur.es.po` but only if a local `./xyz.sulfur.es.po` is missing.

# directories

naming limits:
  limit of 32 UTF8 characters
  No punctuation or spacing other than underscore.
  No double-underscore pairs.
  No leading or trailing underscores.
  Leading numbers permitted, but not recommended.


The compiler itself auto-creates some directories that are prefixed with dunder.
You will probably want to have your VCS ignore those directories. Examples

\__libs       : cached copies of the libraries

\__work       : a temporary "scratch" directory (and subdirectories) used by the compiler.

\__frameworks : cached copies of the frameworks

\__build      : destination of the compiler output in the context of the target

Not necessarily meant to be ignored by the VCS:

\__pot        : this is where new or updates to the `.pot` files are made.

