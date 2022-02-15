# Exercises - Ninja Level 8

## Hands-on exercise #1

Starting [with this code](https://go.dev/play/p/_fQUGm9Utvl), marshal the `[]user` to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.
solution: https://play.golang.org/p/8BK6PXj3aG

## Hands-on exercise #2

Starting [with this code](https://go.dev/play/p/Ic-EWXs7m9B), unmarshal the JSON into a Go data structure. Hint:
https://mholt.github.io/json-to-go/
code:

- code setup (just fyi, not needed for exercise):
  - https://play.golang.org/p/nWPP5Z2Q4e
- solution:
  - https://play.golang.org/p/r8oSG8DIPR

## Hands-on exercise #3

Starting [with this code](https://go.dev/play/p/HBiH2ZYYdSG), encode to JSON the `[]user` sending the results to Stdout.

- Hint: you will need to use `json.NewEncoder(os.Stdout).Encode(v interface{})`
- solution: https://play.golang.org/p/ql90D1OQqd

## Hands-on exercise #4

Starting [with this code](https://go.dev/play/p/pUE8ga9Jm6X), sort the `[]int` and `[]string` for each person.

- solution: https://play.golang.org/p/jz_llY1aPp

## Hands-on exercise #5

Starting [with this code](https://go.dev/play/p/HBiH2ZYYdSG), sort the `[]user` by

- age
- last
  Also sort each `[]string` “Sayings” for each user
- print everything in a way that is pleasant
  solution: https://play.golang.org/p/8RKkdtLl6w
