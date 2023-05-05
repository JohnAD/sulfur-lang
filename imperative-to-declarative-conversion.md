imperative-to-declarative conversion

So,

```sulfur
#! sulfur type 2022.0.1 en
#@ std::hd_terminal as t

t.print( "The hello program!" )
while ( name != "exit" ) [[
  var name = t.input( "Enter your name or say \"exit\":" )
  if ( name != "exit" ) [[
    t.print( "Hello " & name )
  ]]
]]
var name = t.input( "Enter a final name:" )
t.print( "Goodbye " & name )
```

Basically, the "waiting" bit would create boundaries the states. In this case, that would be any `t.input` call. There are two and I will call them "hello" and "goodbye".

But, the idea would be to create a "super macro" that took this:


```sulfur
#! sulfur type 2022.0.1 en
#@ std::hd_terminal as t

using imperative_to_delarative [[
  macro converter
]]

invoke se = state_engine {{
  body = [[
    t.print( "The hello program!" )
    while ( name != "exit" ) [[
      var name = t.input( "Enter your name or say \"exit\":" )
      if ( name != "exit" ) [[
        t.print( "Hello " & name )
      ]]
    ]]
    var name = t.input( "Enter a final name:" )
    t.print( "Goodbye " & name )
  ]]
}}
```

And somehow creates this with a state-engine:

```sulfur
#! sulfur type 2022.0.1 en
#@ std::hd_terminal_reactive as t

using std::state_engine [[
  macro state_engine
]]

{##

  hd_terminal framework code

##}

bind t {{
  on_load = start_engine
  on_input = value_entered
}}

procedure value_entered {{
  parameters = {{
    text :str
    context :str
  }}
  body = [[
    se.name = text
    se.think()      // indirectly runs the "body" of whatever happens to be the current state
  ]]
}}

procedure start_engine {{
  body = [[
    se.start()
  ]]
}}

{##

  the state engine (SE)

  an SE is very self-referential. Externally, an app can ONLY do the following:
    a. update or read the SE's tracked variables
    b. call `.start()` to reset the SE to it's initial state; this should be done at least once.
    c. call `.think()` to make the SE think about things
    d. call `.get_state()` to VIEW the current state. external routines CANNOT change the state.

##}

invoke se = state_engine {{   # invoking a macro
    vars = {{
      name :str = void
    }}
    states = {{
        init = {{                # the state_engine always starts with the first entry; on_entry and on_body are called once on `.start()`
          on_entry = init_entry
          body = init_body
          on_exit = void
        }}
        hello = {{
          on_entry = hello_entry
          body = hello_body
          on_exit = hello_exit
        }}
        goodbye = {{
          on_entry = goodbye_entry
          body = goodbye_body
          on_exit = goodbye_exit
        }}
        done = {{
          on_entry = void
          body = void
          on_exit = void
        }}
    }}
}}

{##

  se: INIT

##}


procedure init_entry {{
  parameters = {{
    name :str
  }}
  body = [[
    t.print( "The hello program!" )
  ]]
}}

procedure init_body {{
  parameters = {{
    name :str
  }}
  returns = {{
    new_state :se = void
  }}
  body = [[
    new_state = se.hello
  ]]
}}

{##

  se: HELLO

##}


procedure hello_entry {{
  parameters = {{
    name :str
  }}
  #  NOTE: entry() routines cannot return a new state; ever.
  body = [[
    t.request_input( "Enter your name or say \"exit\":", context="hello" )
  ]]
}}

procedure hello_body {{
  parameters = {{
    name :str
  }}
  returns = {{
    new_state :se = void      # if left void, the state does not change
  }}
  body = [[
    if ( name == "exit" ) [[
      new_state = se.goodbye
    ]] else [[
      t.print( "Hello " & name )
      new_state = se.hello     # this will cause the `on_exit` and `on_entry` to run in sequence again
    ]]
  ]]
}}

procedure hello_exit {{
  parameters = {{
    name :str
  }}
  #  NOTE: exit() routines cannot return a new state; ever.
  body = [[
    t.clear_input_requests()
  ]]
}}


{##

  se: GOODBYE

##}

procedure goodbye_entry {{
  parameters = {{
    name :str
  }}
  body = [[
    t.request_input( "Enter a final name:", context="goodbye" )
  ]]
}}

procedure goodbye_body {{
  parameters = {{
    name :str
  }}
  returns = {{
    new_state :se = void
  }}
  body = [[
    t.print( "Goodbye " & name )
    new_state = se.done
  ]]
}}

procedure goodbye_exit {{
  parameters = {{
    name :str
  }}
  body = [[
    t.clear_input_requests()
  ]]
}}
```


This would take some serious convolutions. Lets not do this right now.

It might be done as a CLI utility rather than a macro. Even so, not easy.
