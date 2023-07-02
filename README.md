# logo2go

Convert Logo program to Go.

go run logo2go.go example.logo > example.go

## Logo

http://people.eecs.berkeley.edu/~bh/downloads/ucblogo.pdf

### Logo Commands

| Command | Description | Usage     |
| ------- | ----------- | --------- |
| FD      | Forward     | FD <step> |
| BK      | Backward    | BK <step> |
| RT      | Right       | RT <deg>  |
| LT      | Left        | LT <deg>  |
| PU      | Pen Up      | PU        |
| PD      | Pen Down    | PD        |

### Example

```logo
FD 100
RT 90
FD 100
RT 90
FD 100
RT 90
FD 100
RT 90
```
