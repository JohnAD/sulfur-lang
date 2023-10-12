## TLDR

  Functions only READ the target object. No side affects. Someting is ALWAYS returned.
    uses `.` for UCFS
  Method change the target object. No side effects. Nothing is returned.
    uses `~` for UCFS
  Actors have side effects. They cannot be used with UCFS.
    modified calling parameters are prefixed with `*`
    Actors are prefixed with ! when called.
  Modules are singletons
    uses \


## DETAILS

### Function
    a procedure with no side effects.

    It is usable with UCFS
      separation is done with a dot. "."

   functions can call Methods on member elements, but not on the target object.

   functions cannot call Actors. Ever.

   functions are "desist" by default. That is, it can be code-removed if the compiler does not think it is needed.

### Method

    a procedure with no side effects.

    The procedure modifies the target object.

    It is usable with UCFS
       separation is done with a `~`

    If specifically marked as "%%ALLOW_BANG"?, they can also be called with a ! prefix.

    methods are "desist" by default. That is, it can be code-removed if the compiler does not think it is needed.

    methods cannot call Actors. They can call Functions.

### Actor
    a procedure with side effects; includes anything with external I/O

    It is NOT usuable with UCFS.

    calling parameters that are modifiable MUST be prefixed with a `*`

    Actors are preceded with a ! when called, thus calling out the danger.

    actors are "insist" by default. That is, the structure of the code is not removable unless it is fully emptied out first.
      unless the actor is functional in nature, this isn't likely to happen.

### module
    a module library is a singleton object with optional variables

    if used in a UCFS manner, the separation is backslash '\'

    modules are "insist" by default. That is, the structure of the code is not removable unless it is fully emptied out first.
      However this is a rare of an event if it ever happens.

The insist / desist nature can be overridden with prefixes on the call.

The call to the procedures can then override everything with a prefix on the statement.

```sulfur
#!
#@ lib::blah as b [[ func bar, type zing, func bam, actor larry, method moo ]]

var x = b\foo()
var y = bar()
insist var z = bar()  # both 'z' and the code of `bar` will not be code removed no matter what.

var obj = zing()

var j = obj.b\bing()
var k = obj.bang()   # bang is a function for obj

!larry(*obj)          # larry is an actor for zing
desist !larry(*obj)   # this second call _might_ be removed if obj is not used later

b~moo()              # moo is a method for b (b is likely now changed)
```

## LOGGING

The Logger can, at compile-time, change it's behavior.

In DEFER mode, a global late-binding "logger class" is passed around. `logger\!log("blah")` causes the string to be added to the narrative collection. But, no actual I/O occurs. At program exit, the logger then "acts" on the logs and send the logs to their destination.

The benefit of this is that the `!log` is NOT a procedures and can then, thus, have it's code removed. The downside: a "hard crash" will cause you to lose the logs.

In IMMEDIATE mode, the `!log("blah")` is a Procedure which triggers I/O all throughout the program. As such, much less code is removed.

## GENERAL NOTES

calling procedures is perfectly normal in a typical app. After all, doing I/O of _some_ kind is the purpose of an app.

libraries, however, should be much more careful about them. It is better to NOT 