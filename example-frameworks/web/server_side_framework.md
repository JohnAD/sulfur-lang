# Web router with middleware

The following is a mockup of a possible web framework. This particular framework uses
server-side rendering. That is, the user's browser is provided with pre-rendered HTML
by the server. If any JS is needed, it is used in tiny pieces in an "island" fashion.

The system will also be using a react-style library on the main page; and a mustach
template library on the login page.


`main.code.sulfer`:

```sulfur
#! sulfur src 2022.1.0 en
#@ std::server_side_web_framework [[ 
  router 
]]

const website :directory = "website/"

using lib::login_stuff [[ macro login_stuff ]]
using lib::cookie_messenger [[ cookie_messenger ]]
using lib::sqlite3 [[ sqlite, sqltable ]]

var db = sqlite.init()
var user_table = db.table("user")
var todo_table = db.table("todo")

let user_table :macro = vinvoke sqltable( db, "todo" )

@bind = [
  router( "ToDo App", todo_app, "/", website )
  login_stuff( user_table )
  cookie_messenger()
  sqltable( todo_table )
]

# the "run" is implied/handled by the framework

```


`website/index.code.sulfur`

```sulfur
#! sulfter 2022.1.0 en
#@ std::server_side_web_framework [[ router, route, html, body, h1, :id ]]

using lib::login_stuff [[ ]]
using lib::cookie_messenger [[ ]]
using lib::sqlite3 [[ sqltable ]]

using proto3 {{
  "repo::/models/user.protobuf" as user_model
  "repo::/models/todo.protobuf" as todo_model
}}

## one way:

invoke router.get( [ "/", "/index.html" ] ) {{
  making [ request, result ]
  with [
    login_stuff( request, required = false ) making [ logged_in, user ]
    cookie_messenger( request, result ) making [ send_message, get_messages ]
    sqltable() making [ todo_table ]
  ]
  body [[
    var inout :id = "login_or_logout"
    var todo_list = todo_table.find_matches( todo_table.user_id, user.id ).limit(10)=>get()
    result.html = {
      title { "The TODO App" }
      body {
        h1 { "The TODO App" }
        div {
          ul {
            vif todo_list.valued() [[
              vfor todo in  [[
                li {
                  vif (todo.done) [[
                    s { todo.text.$ }
                  ]] else [[
                    b { todo.text.$ }
                  ]] 
                }
              ]]
            ]] else [[
              p { i { "no entries" } }
            ]]
          }
        }
        p
        p { id = inout }
      }
    }
    if logged_in [[
      result[inout] = {
        center {
          "Click "
          href { a = "/logout", "here" } 
          " to logout"
        }
      }
    ]] else [[
      result[inout] = {
        center {
          "Click "
          href { a = "/login_form", "here" }
          " to login"
        }
      }
    ]]
  ]]
}}

# another way (code fragment)
# the `at` symbol means "framework invokation"

@declare [[
  get( [ "/", "/index.html" ] ) making [ request, result ]
  login_stuff( request, required = false ) making [ logged_in, user ]
  cookie_messenger( request, result ) making [ send_message, get_messages ]
  sqltable() making [ todo_table ]
]] do [[
]]


```

`website/login_form.code.sulfur`

```sulfur
#! sulfter 2022.1.0 en
#@ std::server_side_web_framework [[ router, http_response_codes ]]

using lib::login_stuff [[ ]]
using lib::cookie_messenger [[ ]]
using lib::sqlite3 [[ sqltable ]]
using std::mustach [[ template ]]

using proto3 {{
  "repo::/models/user.protobuf" as user_model
  "repo::/models/todo.protobuf" as todo_model
}}

invoke router.get( [ "/login_form" ] ) {{
  making [ request, result ]
  with [
    login_stuff( request, required = false ) making [ ]
    cookie_messenger( request, result ) making [ send_message, get_messages ]
    sqltable() making [ todo_table ]
  ]
  body [[
    result.html = template(
      ["""
        <html>
          <head>
            <title>The TODO App : Login<title>
          </head>
          <body>
            <h1>Log In Form</h1>
            <p>
            <form action="/login_form" method="POST">
              Enter username: <input type="text" name="username" value="{ username }" />
              <p />
              Enter password: <input type="password" name="password" />
              <p />
              <input type="submit">Login</input>
            </form>
          </body>
        </html>
      """]
      { username = request.url_vars["username"] }
    )
    result.code = http_response_codes.ok
    # if there are ANY 'cookie_messenger' cookies, the macro will rewrite the
    # html to include those messages and remove them from cookies
  ]]
}}

invoke router.post( [ "/login_form" ] ) {{
  making [ request, result ]
  with [
    login_stuff( request, required = false ) making [ auth, attempt_login ]
    cookie_messenger( request, result ) making [ send_message ]
    sqltable() making [ todo_table ]
  ]
  body [[
    let username = request.form_vars["username", 0]  # actually, the 0 is the default for a rolne and not needed
    var success = auth=>attempt_login( username, request.form_vars["password"] )
    # the 'cookie_messenger' will store the message as a cookie
    if success {{
      result.redirect = "/"
      send_message( category_success, "You are logged in!" )
    }} else {{
      result.redirect = "/login_form?username={username}".replace( { username = username } )
      send_message(
        category_danger,
        "Bad username or password for {username}.".replace( { username = username } )
      )
    }}
    result.code = http_response_codes.redirect
  ]]
}}

----
```