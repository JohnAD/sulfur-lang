# The Parser and the Abstract Syntax Tree (AST)

This step in the compilation takes a simple list of tokens for a file and converts it into
a tree of "nodes". See AstNode in the `ast.go` file.

Each AstNode has:

* a **Kind** enum describing the general nature of the AstNode. Specifically, the Kind is
  used to manage the parsing at this first use of the tree.
* a **Src**, which is the corresponding token for the AstNode. Not all nodes have a Src.
* zero or more children under **Children**, representing the ordered branches to other AstNodes

There is also a "shortYaml" flag used for debugging and testing. It does NOT change the tree,
it is simply for making a shorter and more human-readable version of the YAML expression.

## AstNode Kind roles

Not all AstNode's are the same. Specifically they can have the following roles based on their Kind:

### Kind role = Parent

A node that is a "Parent" has NO Src token, but collects many ordered children underneath it.
Each child node is of *X-ITEM* kind; where "X" is a one or two-letter abbreviation of the parent.
For example, `STMT` parent AstNodes have `S-ITEM` AstNode children.

### Kind role = Grouper

A node that is a "Grouper" groups together ZERO OR MORE children. The binder DOES have a Src token
representing the "source" of the binding. If the Src is an "opening operator", such as `{` or `(`
then the finishing of the AstNode occurs when closing operator token IS EXPECTED but DISCARDED
during parsing. Examples:

* `{ x y z }` becomes `ROLNE` with Src=`{` and has three children of type `R-ITEM`.

### Kind role = Binder

A node that is a "Binder" groups together a "fixed" number of children. If a binder has
a Src token, then it represent the "source" of the binding. Otherwise, the source of the binding
is likely the nature of it's type:

* `foo.bar` becomes `OBIND` with Src=`.` and has 2 children `OB-LEFT-ITEM` and `OB-RIGHT-ITEM`.
* A `R-ITEM` node is also a Binder and has 3 children `R-ITEM-NAME`, `R-ITEM-TYPE`, and `R-ITEM-VALUE`.

### Kind role = Hand

A node that is a "Hand" has NO Src token and 0 or 1 children. If it has 0 children, then it is an
"empty hand". If it has one item, then it is a "full hand".

For example, parsing `{ x }` produces a Grouper `ROLNE` with one Binder child `R-ITEM` with three
Hand children:

* `R-ITEM-NAME` which is an empty hand
* `R-ITEM-TYPE` which is an empty hand
* `R-ITEM-VALUE` which is a full hand, the one child is an `IDENT` node with Src=`x`.

### Kind role = Core

A node that is a "Core" has a Src token and 0 OR MORE children. It represents a self-contained node
that can be extended by it's children.

## An example first-pass parsing with notes

Source code:

```sulfur
#! bling

x = 4
print( x, "hello" )
```

Resulting AstNode tree (express in short YAML):

```yaml
kind: ROOT
children:
  - kind: STMT
    children:
      - kind: S-ITEM
        child:
          - kind: OP
            name: "#!"
      - kind: S-ITEM
        child:
          - kind: IDENT
            name: "bling"
  - kind: STMT
    children:
      - kind: S-ITEM
        child:
          - kind: IDENT
            name: "x"
      - kind: S-ITEM
        child:
          - kind: OP
            name: "="
      - kind: S-ITEM
        child:
          - kind: NUM-STR-LIT
            name: "4"
  - kind: STMT
    children:
      - kind: S-ITEM
        child:
          - kind: IDENT
            name: "print"
            children:
              - kind: ROLNE
                name: "("
                children:
                  - kind: R-ITEM
                    children:
                      - kind: R-ITEM-NAME
                      - kind: R-ITEM-TYPE
                      - kind: R-ITEM-VALUE
                        child:
                          - kind: IDENT
                            name: "x"
                  - kind: R-ITEM
                    children:
                      - kind: R-ITEM-NAME
                      - kind: R-ITEM-TYPE
                      - kind: R-ITEM-VALUE
                        child:
                          - kind: STR-LIT
                            name: "hello"
```

Please note that the above YAML is the easier-to-read "ShortYAML" version. The full version that is stored to disk
and used by later compiler steps is far more exhaustive and contains details such as the file/line/column references
to each token.

Notes:

1. Starting at the top, the initial parse of a file begins with a Parent `ROOT` node. That node only accepts a list 
   of statements (`STMT`).
2. There are two `STMT` nodes. A `STMT` has a Parent role. So it contains a sequence of children of kind `S-ITEM`.
3. The first `STMT` has three `S-ITEM` nodes. The second `STMT` has just one.
4. Each eol `INDENT` token seen while on a `S-ITEM` sequence starts a new `STMT`. Later, it is possible that stand-alone
   semicolons (";") will have the same effect as an eol `INDENT`
5. A `S-ITEM` is simply a Hand. All it does is hold the node representing that item.
6. The first `STMT` has three simple `S-ITEM` Hands; each of which hold a Core node. Specifically, an identifier ("x"),
   an operator ("="), and a numeric literal ("4"). In later compilation stages this statement will be converted to
   a `__SET` command. But that is later.
7. In the "ShortYAML" version, a Hand node uses the singular "child" (rather than "children") label if non-empty. If
   empty, nothing is expressed at all.
8. The second `STMT` has one `S-ITEM` which holds the Core node IDENT of "print". In later compilation stages a
   `__CALL` node will be prefixed to this. But that is later. 
9. The first child of the Core "print" node is a `ROLNE` representing the parameters for "print". At this stage the 
   compiler does NOT know that "print" is a function call, it simply knows that the rolne of "(" has been bound to the
   identifier. The lexer already marked the "(" as a "bound" op, not a "standalone" op.
10. The `ROLNE` is parsed in context. The trailing ")" token completes this parsing (but is not stored anywhere). A
    `ROLNE` only contains `R-ITEM` children. In this case, two of them.
11. Each `R-ITEM` is a Binder hand has exactly three Hand children for the name, type, and value.
12. Both `R-ITEM`s have empty hands for name and type. But have full hands for value: an identifier ("x") and a string
    literal ("hello").
13. The "bound" comma (",") operator is simply ignored as it is has no meaning here but is allowed for programmer comfort.
14. Had the closing ")" not been in the source code, an error would have been generated. That is because a file parse
    must end on a "empty" `STMT` just below the `ROOT`. The parser removes this final empty statement to clean things
    up.

