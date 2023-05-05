# sulfter: append vs extend

```sulfer
x = 12
a = [1, 2, 3]
b = [7]

# change in place:
a << append(x)  # a is changed and is now [1, 2, 3, 12]
a << extend(y)  # a is changed and is now [1, 2, 3, 12, 7]

# new list
c = a & b    # a is left alone, c becomes [1, 2, 3, 12, 7, 7]
d = a & [x]  # a is left alone, d becomes [1, 2, 3, 12, 7, 12]
e = a.appendWith(x) # a is left alone, e becomes [1, 2, 3, 12, 7, 12]
```

please note that `+` is never used with a list. That is because addition
is Commutative. eg. a + b == b + a . But `&` does not have that restriction.

idea: properties follow initial symbols even when combined.
So, `+#` must also be Commutative, not matter what that function does.

the only exception is when the last character is `)`, `]`, or `}`.

so, `(***` must be a grouping identifier, regardless of what it actually does.
And, it must have a corresponding `***)`.

might also include support for unicode indicated with Pi and Pf (initial and final punctuation)

or a subset:

LEFT CORNER BRACKET (U+300C, Ps): 「
RIGHT CORNER BRACKET (U+300D, Pe): 」

though, really, these should be reserved for quotation mark aliases as they are
used in Chinese.

