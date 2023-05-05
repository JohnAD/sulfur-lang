TLDR

  Actors have side effects on objects in general
    uses <-
  Functions don't have side affects
    if not modifying UCFS object uses .
    if modifying UCFS object uses ::
  Modules are singletons
    uses \


----------------
rather than all the procedure names, make most of it "contextual" in nature a limit it to two types:

  function
    a procedure with no side effects.

    if prefaced with ( self ) parameter, then it is usable with UCFS
      separation is done with a dot. "."

   functions cannot call actors. Ever.

   functions are "desist" by default. That is, then it can be code-removed if the compiler does not think it is needed.

  actor
    a procedure with side effects; includes anything with external I/O

    if prefaced with ( self ) parameter, then it is usable with UCFS

    actors are "insist" by default. That is, the structure of the code is not removable unless it is fully emptied out first.
      unless the actor is functional in nature, this isn't likely to happen.

    the logger module uses desist on most of its methods.

  module
    a module library is a singleton object with optional variables

    if used in a UCFS manner, the separation is backslash '\'

    modules are "insist" by default. That is, the structure of the code is not removable unless it is fully emptied out first.
      However this isn't that rare of an event.

The insist / desist nature can be overridden with prefixes on the definition.

The call to the procedures can then override everything with a prefix on the statement.

```sulfur
#!
#@ lib::blah as b [[ func bar, type zing, func bam, actor larry, actor moo ]]

var x = b\foo()
var y = bar()
insist var z = bar()  # both 'z' and the code of `bar` will not be code removed no matter what.

var obj = zing()

var j = obj.b\bing()
var k = obj.bang()

obj <- larry()   # larry is an actor for zing
desist obj <- larry()  # this second call _might_ be removed if obj is not used later

b <- moo()    # moo is an actor for b itself.
```

logger snippets

```sulfur
log <- danger("something happened")   # might get removed if the generic "log" is never output because `danger` actor is overridden with 'desist` in the library.
insist log <- danger("tada")

var x :byte = 22

x.log.danger("one")  # will get removed if the "log" of x isn't used anywhere
insist x.log.danger("two")  # no removal 

log <- pull(x)  # the log of x is "added" to the the generic log; if there are any entries
x.log.clear()
```