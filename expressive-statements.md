# Expressive Statements

## concept

sulfur's "core" has almost no statements and they are mostly forbidden dunder types anyway.

But the "internal" library has a *huge* library of statements that allow for clear expression.

The variations are there to allow the programmer to describe what they WANT.

So, for example, if creating a list of items, the following is mildly ambigous:

  `[ 1 ... 6 ]`

whereas:

  `[ 1 ...including 6 ]`

I now am certain the it has the number 6 in the list. But,

  `[ 1 ...before 6 ]`

tells me that it includes 5 but NOT 6.

This is even more useful for nebulous bounds. For example, let's say
the list is always for integers but the upper bound is not an integer, such as:

  `[ 1 ... 8.1 ]`

Does it include 9 or does it stop at 8?

  `[ 1 ...before 8.1 ]`

I now know that it stops at 8. Whereas:

  `[ 1 ...including 8.1]`

sets an error: `Error: can't include 8.1 into list`. In fact, in this particular case, because
literals are being used, it would be caught at compile-time.

## examples

### equivalent of select/switch statements

```sulfur
# only the FIRST matching block will be executed

first_match {{
  (x > 4) :: [[
    # stuff
  ]]
  (y < 2) :: [[
    # other stuff
  ]]
  else :: [[
    # otherwise stuff
  ]]
}}

# the else can be skipped if "doing nothing" is okay when nothing matches
# The `::` means "associates with". Using `=` would be kinda confusing.
```

```sulfur
switch(blah) {{
  [ 3, 9 ] :: [[
    # stuff
  ]]
  [ 4 ] :: [[
    # other stuff
  ]]
  else :: [[
    # otherwise stuff
  ]]
}}

# the else can be skipped if "doing nothing" is okay when nothing matches
```

```sulfur
case( arrow_key ) {{
  [ up ] :: [[
    # stuff
  ]]
  [ down, left] :: [[
    # other stuff
  ]]
  [ right ] :: [[
    # right stuff
  ]]
}}

# 'else' is NOT supported and blocks must account for ALL possible cases; no overlap. Mostly used for enums.
# essentially, the type needs a `case_count` of less than 255.
```

### equivalent of loop statements


```sulfur
for_each( color in colors ) [[
  if (color == red) [[
    color.exit_each()  # `color` is a normal const, but an 'in-context' method has been added called `exit_each` inside the `each`
  ]]
  t.println( "item {i}: color is {c}".fmt( i = color.index, c = color.$ ) )  # index is also a 'in-context' value.
]]

# note: sulfur does NOT allow middle-of-the-block exit. In the above example, "red" IS still printed. It only actually exits with
# `exit_each` when the block is finished and the loop is again considered.
# note: 'colors' becomes read-only in the context of the loop. You cannot change anything within the colors list.
```

There is also a key/value version: `each_item( k, i, v in some_object ) [[ ...` where k is the string key, 
i is the count of key, v is the value. For a dictionary, `i` is always 0. 

Here also, `some_object` temporarily becomes immutable.

```sulfur
for_index( i in colors ) [[
  if (colors[ i ] == red) [[
    i.exit_each()
  ]]
]]
t.println( "item {i}: color is {c}".fmt( i = i, c = colors[ i ].$ ) )

# note: while similar to `each`, the `colors` does NOT become immutable. However, deletion of an entry is still forbidden.

for_edit(val in colors ) [[
  if (val == red) [[
    val.exit_each()
  ]]
]]
t.println( "color {c} found".fmt( c = $val ) )
```

key/value version: `reference_items(k, i, v in some_object)`. `some_object` is left writable. Keep in mind that the "key list"
is pulled prior to the start of the loop.

```sulfur
for_alter( color in colors ) [[
  # until the first red is found, change the colors to blue; delete the first red
  if (color == red) [[
    color.delete()
    color.exit_change_list()
  ]] else [[
    color.set( blue )  # `color` is still immutable, but `set` changes the entry in `colors`
  ]]
]]
```
Ironically, `colors` is not even _visible_ inside the loop. Even reading it is not allowed.

Other methods in `alter`: `push_before`, `push_after`. These add items just before
or just after the current location. (Even with a `push_after` you will NOT see
the new item in the loop until the loop is finished.)

alternate:

```sulfur
alter( color in colors ) [[
  # until the first red is found, change the colors to blue; delete the first red
  if (color == red) [[
    color = nothing  # causes deletion
    color<-finished() # simply appends the remainder of list, loop complete at END
  ]] else [[
    color = blue  # `color` is still immutable, but `set` changes the entry in `colors`
  ]]
]]
```


2023-12-03 idea

match up functional languages for for-loop name:
  map
  reduce
  filter

