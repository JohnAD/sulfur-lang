# 20230219_flushing_out_target_vs_passive_parameters.md

The idea is a "functional"-style promise:

  - passed in parameters are never modified

Combined with an easily visualized way to change something.

## passive parameters and `functions`

So,

  g = xyz( a, b, c )

The parameters a, b, and c are left alone.

Those three parameter would be called "passive parameters".

If a function modifies a, b, or c, then the compiler makes a "deep copy" of
the parameter for strictly-local usage.

'functions' always return a value.

## target parameters and `procedures`

But,

  d_jkm( a, b, c )

The parameters a, b, c are left alone. But d is expected to be modified.

`d` doesn't have to be modified, but that is the expectation. These parameters
are called "target parameters".

Verbally, it would be read as

  "d changed by jkm with parameters a b and c"

To modify multiple parameters, which should be fairly rare case, it would be
in the form of:

  [d, e, f]~jkm( a, b, c )

Internally in the function, the compiler does NOT make a copy of the target
parameters. It is the equivalent of "pass by reference".

What an 'procedures' based on context:

* when alone on a line, it returns nothing.
* when part of a chain, it passes a reference to itself down the chain
* when the "end" is returned to an assignment, a COPY of itself is made.

### alone on line

So,

  d~jkm( a, b, c )

does not return anything.

### in a chain

So,

  d~jkm( a, b, c )~zing()~blah( r )

has `jkm` modify `d` and pass a REFERENCE to `d` to `zing` which modifies
`d` and passes a reference to `d` to blah which modifies `d` also. But `blah` does
not return anything.

And,

  name = d~jkm( a, b, c )~zing().get_name( a )

has `jkm` modify `d` and pass a REFERENCE of `d` to `zing` which modifies
`d` and passes the reference to method `get_name` which DOES NOT modify `d` but does
return something. See the section below on 'methods'.

## sent to an assignment

So,

  new_d = d~jkm( a, b, c )

the `new_d` variable contains a COPY of `d` (with it's changes). Essentially, `new_d`
and `d` have the same contents; but they are independent copies.

and,

  another_d = d~jkm( a, b, c )~zing()

the `another_d` variable contains a COPY of `d` with the changes from both `jkm` and `zing`.
Essentially, `another_d` and `d` have the same contents; but they are independent copies.

## target AND passive parameters for 'methods'

A 'method' is a function with a target parameter. It is used for object-oriented calling.

Example:

  name = d.get_name( a )

In this example:

* the `a` parameter is a passive parameter and will not be changed.
* the `d` parameter is a target parameter and will not be changed.

The target parameter is READ ONLY. If a change is seen in the method, the compiler
will make a copy of the target parameter for local usage.

All methods return something.

### side note

BTW, the `<-` and `->` punctuation for methods/procedures made by earlier notes
is not being used any more. Why? It breaks two rules:

1. It does not work well with RTL human languages since it implies direction
2. It implies uses "bracing" `<` `>` angle brackets in a non-bracing way.

## examples

```sulfer

struct Person {{
    public = {{            // accessable to everything
        first_name :str
    }}
    private = {{           // accessable to only methods and procedures via target
        last_name :str
        phone_number :str
        age :uint8
    }}
}}

method full_name {{
    parameters = {{}}
    target = person :Person
    returns = final :str
    body = [[
        if person.last_name.len() > 0 [[
            final = person.first_name & " " & person.last_name
        ]] else [[
            final = person.first_name
        ]]
    ]]
}}

function awkward_greeting {{
  parameters = {
    person :Person
    greeting :str = "Hello"
  }
  returns = full_greeting :str
  body = [[
    full_greeting =  greeting & ", " & person.full_name()
    full_greeting &= " I have noticed you are " & person.age.$ & " years old."
  ]]
}}

procedure set_age {{
  parameters = {{
    new_age :uint8
  }}
  target = person :Person  // there is no "returns", the target is returned depending on context
  body = [[
    person.age = new_age
  ]]
}}
```

## "The Questions"

* Q: I want to write something that both changes a parameters and returns a value.
* A: No.

Essentially, you are wanting to do "two things" in single block of code. I recommend writing two things to do two things.

* Q: When writing a procedure/function/method I keep getting a "You cannot call an actor from a procedure/function/method"?
* A: You will want to write an 'actor' instead.

An 'actor' is VERY similar to a procedure. The big difference is that actors are allowed to call other actors. The actor references
are passed in as target parameters. Actors are exempt from code removal because
they change outside resources and generate side-effects. So, use actors sparingly as possible.

