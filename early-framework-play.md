
Attachment save-as.png added.Conversation opened. 1 read message.

Skip to content
Using Gmail with screen readers
4 of 818
sulfur stuff 3
Inbox

John Dupuy <jdupuy98@gmail.com>
Tue, Oct 25, 2:29 PM (3 days ago)
to me

invoke foo_bar(leo, "xo") making [
  x
  y as zippy  # the 'as' system can help prevent name collisions
  blah_function
  # you do NOT have to include ALL of the FooBar's makings
] body [[
  # do something with x, zippy, and blahFunction etc.
]]  # outside of the body, `x`, `zippy`, and `blahFunction` do not exist

----

macro foo_bar {{
  parameters = {{
    bob :str
    zog :str *static           # either a const or a literal
  }}
  making = {{
    x :int
    y :int
    z :int
    blah_function :proc        # this is a psuedo-type only allowed in macros
    bingo :proc
  }}
  parse = void  # if non-void, then the passed-in body is parsed and rewritten; only static parameters are visible
  before = [[
    t.print("before the body starts")
    x = 14
    y = 7
    z = bob.len()
    invoke foo_bar(z) body [[ t.print("something else") ]]
  ]]
  after = [[
    t.print("after the body ends")
  ]]
  proc = blah_function {{
    # stuff
  }}
  proc = bingo {{
    # stuff
  }}
}}

----

# also support vinvoke, which has a `result`

var b = vinvoke foo_bar(leo, "xo") [   # ...

----

Recursion in macros is also forbidden; both in `main` and `body`.

So, foo_bar -> nappy -> foo_bar  will trigger a compile-time error.

----

# Web router example:

var r = router()

invoke route(r, "/index.html") making [
  request
  result
] body [[
  if request.style == GET [[
    result.content = "<h1>Hello</h1>"
  ]] else [[
    result.content = "whoops"
    result.code = 500
  ]]
]]

r.start()

----

# Web router example (even simpler?):

var r = router()

invoke route(r, GET, "/index.html") making [
  request
  result
] body [[
  result.content = "<h1>Hello</h1>"
]]

r.start()

----

# Web router example (way better from systemic framework):

```sulfur
#! sulfur type 2022.0.1 en
#% framework ssr_web 1.0.2

# NOTE: ssr_web is a declarative framework; not an imperative one.

from framework {{
  GET
  route
  myconfig = config  # this is how you do aliasing
}}

myconfig.port = 8088

route(GET, "/hello/{name}") making [[
  let name = request.param.name
  result.content = "<h1>Hello ${name}</h1>".fmt( name = name )
]]
```

----

# Web router with middleware

var router = router()

var db = db_init()

var login_stuff = login_handler(db, "blah", 14)

chain metaroute = [
  route
  login_stuff
  cookie_messenger
]

metaroute [
  route(r, "/index.html") making [ request, result ]
  login_stuff(request) making [ logged_in, user ]             # login_stuff can have `request` passed in (and modified), but the opposite is NOT true. route cannot have 'user' passed in
  cookie_handler( request, result ) making [ send_message, get_messages ]
] body [[
  if request.style == GET [[
    if logged_in [[
      result.content = "<h1>Hello {name}</h1>".fmt(name = user.name)
    ]] else [[
      result.content = "<h1>Hello</h1>"
    ]]
  ]] else [[
    result.redirect = "/index.html"
    result.code = redirect
    send_message( "you can only GET the home page, you silly human" )
  ]]
]]


# order matters. In the above example, the macro result is:

# at compile-time:
  # `route` body modification
  # `login_stuff` body modification
  # `cookie_messenger` body modification
# then at run-time:
  # `route` before code
  # `login_stuff` before code
  # `cookie_messenger` before code
  # modified `body`
  # `cookie_messenger` after code
  # `login_stuff` after code
  # `route` after code


----

# Web router with middleware and framework

```sulfur
#! sulfur type 2022.0.1 en
#% framework ssr_web 1.0.2

from framework {{ GET, route, chain, redirect }}

from library dblib {{ db }}
from library weblogin {{ login_handler }}
from library cookie_messenger {{ cookie_handler }}

db.init()

var login_stuff = login_handler(db, "blah", 14)

chain [
  ( request, result ) = route(GET, "/index.html")
  ( logged_in, user ) = login_stuff( request )
  ( msg ) = cookie_handler( request, result ) as ch
] do [[
  if logged_in [[
    if (user.name == "joe") [[
      result.redirect = "/logout"
      result.code = redirect
      msg = "Joe can't be here."
      ch.blah("x")
    ]] else [[
      result.content = "<h1>Hello {name}</h1>".fmt(name = user.name)
    ]]
  ]] else [[
    result.content = "<h1>Hello</h1>"
  ]]
]]
```

example of cookie handler:

```sulfur

chain_wrapper chookie_handler {{
  parameters = {
    request
    result
  }
  returns = {
    // variable seen by inner functions, if any, goe here
    msg
  }
  local = {
    // persistent variables between entry and exit go here
  }
  entry = [[
    // stuff goes here, sees request, result, and msg
  ]]
  exit = [[
    // stuff goes here, sees request, result, and msg
  ]]
  methods = {{
    // just like a class, a chain_wrapper can have methods
    blah = {{
      parameters = {
        something
      }
      returns = { }
      body = [[
        // .... stuff goes here
        // sees request, result, msg, and something
      ]]
    }}
  }}
}}
```
