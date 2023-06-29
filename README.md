<p align="center">
Sudoku is a Go programming language software development kit (SDK) used for generating and solving Sudoku puzzles. It offers great flexibility by supporting various grid sizes such as 9x9, 6x6, 4x4, 8x8, and more.
</p>

[![Go Report Card](https://goreportcard.com/badge/github.com/jhunters/sudoku)](https://goreportcard.com/report/github.com/jhunters/sudoku)
[![Build Status](https://github.com/jhunters/sudoku/actions/workflows/go.yml/badge.svg)](https://github.com/jhunters/sudoku/actions/workflows/go.yml)
[![codecov](https://codecov.io/gh/jhunters/sudoku/branch/main/graph/badge.svg)](https://codecov.io/gh/jhunters/sudoku)
[![Releases](https://img.shields.io/github/release/jhunters/sudoku/all.svg?style=flat-square)](https://github.com/jhunters/sudoku/releases)
[![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/jhunters/sudoku)
[![LICENSE](https://img.shields.io/github/license/jhunters/sudoku.svg?style=flat-square)](https://github.com/jhunters/sudoku/blob/main/LICENSE)


# Go Required Version
need go 1.18

# Install
go get github.com/jhunters/sudoku


# Qiuck start

### Do solve sudoku puzzle
```go
	skd, _ := sudoku.NewSukudoX(9)
	origin := [][]int{{6, 0, 0, 0, 0, 2, 0, 0, 0}, {0, 0, 1, 0, 0, 7, 0, 0, 2},
		{0, 3, 4, 9, 0, 0, 0, 0, 0}, {8, 6, 0, 0, 5, 0, 0, 4, 0}, {1, 0, 0, 0, 0, 0, 0, 0, 6},
		{0, 0, 9, 7, 0, 0, 8, 0, 5}, {0, 0, 0, 0, 2, 0, 9, 6, 0}, {0, 0, 0, 0, 0, 1, 0, 0, 4}, 
        {4, 0, 0, 0, 0, 5, 0, 8, 0},
	}

	skd.ResultIn(origin)

	result, _ := skd.Solve()
	fmt.Println(result)

	skd.Print()
```

output result as follow:

```shell
6 7 8 | 5 3 2 | 4 9 1 
9 5 1 | 8 4 7 | 6 3 2 
2 3 4 | 9 1 6 | 7 5 8 
----- + ----- + ----- 
8 6 7 | 1 5 3 | 2 4 9 
1 4 5 | 2 8 9 | 3 7 6 
3 2 9 | 7 6 4 | 8 1 5 
----- + ----- + ----- 
5 1 3 | 4 2 8 | 9 6 7 
7 8 6 | 3 9 1 | 5 2 4 
4 9 2 | 6 7 5 | 1 8 3
```


### generate a random sudoku puzzle

```go
    sg := &sudoku.SukudoGenX{9, 40}
    result, _ := sg.GenSukudo()

    sudoku.PrintX(result, 9, 3, 3)
```
output result will be like:

```shell
4 2 0 | 0 0 0 | 0 5 7
0 0 0 | 0 4 7 | 0 0 8
8 0 0 | 0 0 0 | 4 0 0
----- + ----- + -----
3 0 0 | 5 0 9 | 1 2 6
2 0 0 | 0 0 4 | 9 8 5
9 0 5 | 1 6 0 | 7 4 3
----- + ----- + -----
0 0 2 | 0 7 0 | 0 3 0
7 3 0 | 4 5 0 | 0 0 9
5 0 1 | 6 2 0 | 0 7 4
```

## License
sudoku is [Apache 2.0 licensed](./LICENSE).
