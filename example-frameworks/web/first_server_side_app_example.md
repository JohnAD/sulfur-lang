first server side app example

```sulfur

#! sulfur src 2022.1.0 en
#@ std::server_side_web_framework importing [[
  header, title, body, h1
]]

using lib::SSWF_cookie_messenger importing [[
  cookie_messenger
  cookie
  category
  audience
  level
]]

@bind = [
  self( "Markdown Viewer", root = "/" )
  cookie_messenger()
]

@declare [[
  GET( [ "/", "/index.html" ] ) making [ request ] wanting [ result ]
  cookie( request, result ) making [ inbound_messages ] wanting [ outbound_messages ]
]] do [[
  result.html = {
    header = {
      title = "Markdown Viewer"
    }
    body = {
      h1 = "Markdown Viewer"
    }
  }
  # perhaps ".." could mean "enumeration or field derived by type context"
  outbound_messages.send("here!", category = ..info, audience = ..public, level = ..normal)
]]

```