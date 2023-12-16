# Meta STATEMENT Structures in C

## DOING

    __WHILE_TRUE bool [[ block ]]

    __DO_WHILE_TRUE [[ block ]] bool

    __IF_TRUE bool [[ block ]] ELSE [[ block ]]

    __ASSIGN VAR value

    __FIRST_TRUE [[ boolexpr [[ block ]] boolexpr [[ block ]] ... ]]  // possibly later

## TEMPORARY STRUCTURES

    CALL ident ( args )      // no return value, some pass-by-ref

    DEFINE ident ( params )  // no return value, some pass-by-ref, others copy

## WE ARE NOT DOING:

    FOR ...   // use __WHILE_TRUE or __DO_WHILE_TRUE instead

    SWITCH - CASE   // use chained __IF_TRUE or __FIRST_TRUE statements instead

    CONTINUE // breaks structure rules

    BREAK // breaks structure rules

    GOTO  // breaks structure rules

