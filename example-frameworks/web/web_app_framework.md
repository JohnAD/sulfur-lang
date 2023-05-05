# Web router with middleware

The following is a mockup of a possible web framework.

This particular framework uses client-side javascript for web page control and layout.
For saving state, there will also be a backend API built by the same code.

When the build is made, you will eventually see something like:

```filesystem
\build\index.html
\build\index.js
\build\api\todo
```

... where 'index.html' is the HTML frontend for the browser, 'index.js' is the compiled JS for the browser, and 'api\todo' is an executable running the backend.

For this example, we are not going to concern ourselves with authentication.

`main.code.sulfer`:

```sulfur
#! sulfur src 2022.1.0 en
#@ std::web_app_framework [[ 
  router 
]]

const website :directory = "website/"

using lib::sqlite3 [[ sqlite, sqltable ]]

var db = sqlite.init()
var todo_table = db.table("todo")

chain todo_app_front = [
]

chain todo_app_api = [
  sqltable( todo_table )
]

router.frontend( todo_app_front, "/", website )
router.backend( todo_app_api, "/api", website )

router.run()
```


`website/index.code.sulfur`

```sulfur
#! sulfter 2022.1.0 en
#@ std::web_app_framework [[ router ]]

using lib::sqlite3 [[ sqltable ]]

using proto3 {{
  "repo::/models/todo.protobuf" as todo_model
}}

bind router.frontend.events {{
  on_load = index
}}

bind router.backend.calls {{
  get_todos
  toggle_item
}}

let index = vinvoke router.frontend( [ "/", "/index.html" ] ) {{
  making [ request, dom ]
  with [
    cookie_messenger( request, result ) making [ send_message, get_messages ]
  ]
  body [[
    var todo_list = router.backend.get_todos()
    dom.head.title = "The TODO App"
    dom.body = body {
      h1 { "The TODO App" }
      div {
        ul {
          vif ( todo_list.valued() and ! todo_list.empty() ) [[
            vfor todo in todo_list() [[
              li {
                id = todo.id
                input {
                  type = "checkbox"
                  vif (todo.done) [[
                    checked
                  ]]
                  rebind on_click = router.frontend.click_item(id = todo.id)
                }
                " "
                todo.text.$
              }
            ]]
          ]] else [[
            p { i { "no entries" } }
          ]]
        }
      }
    }
  ]]
}}

actor click_item {{
  self dom :router_front_end
  parameters {{
    id :string = ""
    name :string = ""
  }}
  body {{
    let now_checked = router.backend.toggle_item(id)
    let item = dom[id]
    if now_checked [[
      item.input<=check()
    ]] else [[
      item.input<=uncheck()
    ]]
  }}
}}

proc toggle_item = vinvoke router.backend( "todo/{id}" ) {{
  with [
    sqltable() making [ todo_table ]
  ]
  parameters {{
    id :string
  }}
  returns {{
    new_state = false
  }}
  body [[
    let current_state = todo_table.find_first( todo_table.id, id )=>get()
    new_state = ! current_state.active
    todo_table.set( todo_table.active = new_state ).where( todo_table.id, id )<=update()
  ]]
}}

proc get_todos = vinvoke router.backend( "todo" ) {{
  with [
    sqltable() making [ todo_table ]
  ]
  returns {{
    first_ten_todos :array<10, todo_model>
  }}
  body [[
    first_ten_todoes = todo_table.find_all().limit( 10 )=>get()
  ]]
}}

----
```