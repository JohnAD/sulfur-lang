# 2024 plan for language

possible alt name: cloth

final decision: this is a TRANSPILER. It outputs to C.
Specifically clang under LLVM; which is under the Apache 2.0 license.

The compiler is being bootstraped with Nim, which also produces C code.

## STEP 1: write a simple console app: Colossal Cave Adventure

  See https://gitlab.com/esr/open-adventure
  the framework only uses /usr/lib/libc.so
  https://sourceware.org/glibc/wiki/HomePage
  https://en.wikipedia.org/wiki/C_standard_library

## STEP 2: write a simple web framework: TODO for Everyone

  Server-Side web framework
  CSS-to-style embedding something
  each page is self-contained
  Mustache templating
  API can be anything: Flask, FeathersJS, whatever

## STEP 3: write an API for the TODO app

  Now the API uses the lang also.
  Shared compile-time API spec of some kind.

Only work on this project for 2 hours per week. No more.
