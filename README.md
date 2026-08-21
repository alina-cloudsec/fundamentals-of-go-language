# Fundamentals of Go Language

This is my practice space for learning Go. I write and run small Go
programs here almost every day. This helps me actually learn the
language, not just read about it.

Each folder is one thing I was learning at that time.

---

## What is inside

- **03 blank identifier** > Practicing the blank identifier and a
  function that returns coordinates.
- **array value semantics** > Learning how arrays copy their values in
  Go, instead of sharing them like slices do.
- **hello github** > My very first Go file. Just getting the basics
  working.
- **infinite loop with break** > Practicing infinite loops and how to
  stop them the right way.
- **pointers value modification** > Using pointers to actually change
  a value, instead of just copying it.
- **short variable declarations** > Practicing the short way to make
  variables in Go.
- **slice range loop** > Looping through a slice and printing values
  based on simple rules.
- **switch fallthrough** > Practicing switch statements that fall
  through to the next case.
- **switch with no condition** > Learning that a Go switch does not
  need a variable to check against.
- **variadic functions** > Practicing functions that can take any
  number of inputs.

  ---

## Why I am doing this

Go is the language used to build most cloud native tools. Meshery,
Kubernetes, HAMi, and Flatcar Linux are all built using Go, or use it
in a big part of their systems. I wanted a strong base in Go before
trying to read or work on any of these real project codebases.

---

## Where Go is used

Go is used a lot in cloud and cybersecurity tools today. Here is where
it shows up the most:

- **Cloud infrastructure tools** > Kubernetes, Docker, and Meshery are
  all built using Go. Most tools that manage cloud servers and
  containers are written in Go.
- **Command line tools** > Go is a common choice for building CLI
  tools, because it compiles into one single file that runs fast and
  does not need extra software installed to work.
- **Network and security tools** > Many modern security scanners and
  network tools are written in Go, because it handles many tasks at
  once well, which matters when checking lots of servers or
  connections quickly.
- **Backend servers** > Go is also used to build the backend of many
  websites and apps, because it is fast and simple to read.

Mostly, Go is used for building tools, not for writing malware or
attacking systems. It is the language of the people who build and
protect cloud systems, which is exactly why I am learning it.

---

## Status

I am still adding new folders as I learn more. In the future, I plan
to build bigger, full sized projects with Go, not just small practice
files.

---
```bash
>  **Note:** Honestly, I didn't know that Go has such strict rules for curly brackets `{`. Coming from a C and C++ background, I am used to writing code in that format, so I didn't pay much attention to it at first and just wrote it that way. But when compiling, I got syntax errors, which led me to research and discover Go's strict styling rules. Over the next one or two days, I will be fixing the formatting across all my codes.

