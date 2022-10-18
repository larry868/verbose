# verbose go package 

verbose package generates formatted output only when verbose mode is turned on.

[![Go Reference](https://pkg.go.dev/badge/github.com/sunraylab/verbose.svg)](https://pkg.go.dev/github.com/sunraylab/verbose)

Every formatted output starts with a colored message type:
  - INFO: `>>info:` in cyan 
  - WARNING: `>>warn:` in orange 
  - ALERT: `>>alert:` in red 
  - TRACK: `>>track:` in green
  - DEBUG: `>>debug:` in yellow

Very easy to use like ``fmt.Println``, ``fmt.Printf``... but with `verbose.Println`, ``verbose.Printf``...

## Licence

[MIT License](LICENSE)