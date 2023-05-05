# punctuation.md



## latin punctuation Meanings

33  21  !

    !           boolean not

34  22  "

    "a"         escapement for strings; a static array of utf8 runes

35  23  #

    # a         A comment

36  24  $

    $a          Seriallization of variable "a" to a string (an array of utf8 runes)

37  25  %

    a % b       modulu

38  26  &

    a & b       concatenation of a with b

39  27  '

    'a'         escapment for an individual utf8 runes

40  28  (

    a( x )      parameters for function/actor a. Notice the lack of space between 'a' and '('

    ( x + y )   a summary or grouping of items. Notice the space between '(' and 'x'

41  29  )

    a( x )      parameters for function/actor a. Notice the lack of space between 'a' and '('

    ( x + y )   a summary or grouping of items. Notice the space between ')' and 'y'

42  2A  *

    a * b       multiplication. Must conform to the Distributive, Commutative, Associative, and Common-Factor Removal properties

43  2B  +

    a + b       addition; must conform to Commutative, Associative, Distributive, and Additive Identity properties

44  2C  ,

    a, b        a sequence separator

45  2D  -

46  2E  .

    0.0         after a digit, it is a "number point"

    a.b         a reference to variable b in struct a

    a.b()       a UCFS reference to _function_ b in struct a

    ..a         the 'a' is an identifier derived by type context; so `var x :Level = Level.normal` could simply be `var x :Level = ..normal`

47  2F  /

    a / b       division; must conform to Identity, Zero, and Unary (division by self) properties

58  3A  :

    :a          a declaration of type "a"

59  3B  ;

60  3C  <

    a < b       is "greater than"; both types must match

    a<-b()      is an UCFS _actor_

61  3D  =

    a = b       an assignment

    a == b      A boolean equivalence

    a === b     A boolean AND type equivalence

62  3E  >

    a > b       is greater than; both types must match

63  3F  ?

    ?           null

64  40  @

91  5B  [

    [ x ]       an unnamed data structure (list); example has one value of "x"

    [[ x ]]     group of statements; example has statement "x" in it

92  5C  \

    a \         if at end of line, a continuation-of-line marker

    a\b         constant b found in library a

    a\b()       function or actor b found in library a

    c.a\b()     using function b in library a against object c

    c<-a\b()    using actor b in library a against object c

93  5D  ]

    [ x ]       an unnamed data structure (list); example has one value of "x"

    [[ x ]]     group of statements; example has statement "x" in it

94  5E  ^

    a ^ b       "to the power of"

95  5F  _

    _           void

    __A         a dunder; an identifier with a compiler-directed meaning

96  60  \`

    `+`         a compiler identifier escapement; example would be name name of the __PLUS function

123 7B  {

    { x = y }   a data structure (rolne); example has item named "x" having value "y"

    {{ x = y }} named statements; example has "section x" having statement "y"

124 7C  |

125 7D  }

    { x = y }   a data structure (rolne); example has item named "x" having value "y"

    {{ x = y }} named statements; example has "section x" having statement "y"

126 7E  ~


## non-latin unicode punctuation

Different human languages uses different punctuation marks for differing
purposes.

For example, chinese uses 「 」to quote text.

So, to a LIMITED extent, the language will allow aliases to have the same meaning.

And, to keep things consistent:

  * only allow these in the corresponding language
  * require escapements for the ones used in string notation; require them everywhere.

quotation examples:

```sulfur
#! sulfur 1.0 en

a = "one person said \"hi\" to me and the other said \「hello\」."  # notice the escapement for「 and 」, even in english quoting

# not allowed: a = 「one person said \"hi\" to me and the other said \「hello\」.」  # 「」not allowed in en files

```

```sulfur
#! sulfur 1.0 zh

a = "one person said \"hi\" to me and the other said \「hello\」."  # zh can use either "" or 「」because latin "" is universal

a = 「one person said \"hi\" to me and the other said \「hello\」.」 # still need to do escapement

```

Other punctuation is done the same way but without the escapement requirements. Actual allowed aliases depend on the language,
but they always map back to a universal latin symbol.

For example, a Chinese (hk) source file will alias `、` to be a latin comma `,` since they are both used for list separation. Behind-the-scenes, the parse considers them equivalent. Like the quotation restrictions, a `、` cannot be used as a comma in a English (en) source file (unless inside a string literal).

On the other hand, in Chinese,  `。` is NOT an alias for `.`. While they look similar, `。` means "full stop" and does not mean "subtending" or "fraction-part" which a latin period can _also_ mean.

alias examples:

```sulfur
#! sulfur 1.0 en

x = [ 1, 2, 3 ]
a = "start now。"   # you can mix languages in a string literal

# not allowed: x = [ 1、 2、 3 ]  # 、 is not an allowed alias in English (en)

```

```sulfur
#! sulfur 1.0 zh

x = [ 1, 2, 3 ]    # latin comma is universal
x = [ 1、 2、 3 ]    # allowed for zh marked source files; notice there is still a space after 、
x = [ 1、2、3 ]    # in Chinese (zh) 、aliases to a comma followed by a space, so excluding spaces can be allowed in this scenario.
x = 【 1、2、3 】     # also allowed in zh files. Spaces still needed in this scenario, however, as the parser expects them.

a = "start now。"   # you can mix languages in a string literal

```

```sulfur
#! sulfur 1.0 ar

x = y.detail﴾ 1، 2 ﴿    # an arabic ornate left/right parenthesis, params separated by arabic comma

a = "start now۔"   # you can mix languages in a string literal. Arabic full stop seen here.

```

LTR / RTL note: source files are in order of action. So in the above example, `y` is first followed
by `.detail` function that acts against it, regardless of the language.

Arabic being a RTL language means the _editor_ or _viewer_ is responsible for showing Arabic
content in RTL order. The source file is neither LTR or RTL. It is ordered by sequence of first to last.

Because of the limits of some OS systems, sulfur allows language-specific versions of the `.sulfur` source
extension. For example `.sulfur_ar` for arabic source files.