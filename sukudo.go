// 这是一个数独游戏的核心代码，其中包含了数独的验证、求解等功能。
package sudoku

import (
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/jhunters/goassist/concurrent/syncx"
	"github.com/jhunters/goassist/conv"
)

const (

	// define digital mark
	One = 1 << iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
)

var (
	mapping   = map[int]int{One: 1, Two: 2, Three: 3, Four: 4, Five: 5, Six: 6, Seven: 7, Eight: 8, Nine: 9}
	unmapping = map[int]int{1: One, 2: Two, 3: Three, 4: Four, 5: Five, 6: Six, 7: Seven, 8: Eight, 9: Nine}
)

// 这是一个数独游戏的核心代码，其中包含了数独的验证、求解等功能。
// validateBox函数用于验证一个3x3的小方格是否符合数独规则，validateLine、validateLines、validateCol、validateBoxCol、validateBox2函数用于验证行、列、小方格是否符合数独规则。
// getBox、getBoxLines、getColLines、getBoxLinesScope函数用于获取数独中的小方格、行、列。ResultIn函数用于将原始的数独转换为程序中的数独，ResultOut函数用于将程序中的数独转换为原始的数独。
// doSolve函数用于求解数独。
type Sukudo struct {
	// sukudo puzzle content
	Puzzles [9][9]int

	// 标记是否处理过， key=rowid+colid,  value=true表示已处理
	checked *syncx.Map[string, bool]

	// 标记 本次解题完成，可退出
	Exit *bool

	// 最大尝试次数
	tryCounter *int32
}

func NewSukudo() *Sukudo {
	skd := &Sukudo{checked: syncx.NewMap[string, bool]()}
	skd.Exit = conv.ToPtr(false)
	skd.tryCounter = conv.ToPtr(int32(0))
	return skd
}

// Copy a new Sukudo struct by has some pointer to Exit and tryCounter field
func (s *Sukudo) Copy() *Sukudo {
	mp := s.checked.Copy()
	ret := &Sukudo{Puzzles: s.Puzzles, checked: mp}
	ret.Exit = s.Exit
	ret.tryCounter = s.tryCounter
	return ret
}

// Print sukudo puzzle
func (s *Sukudo) Print() {
	Print(s.ResultOut())
}

// Exited return true if should exit loop
func (s *Sukudo) Exited() bool {
	return s.Exit != nil && *s.Exit
}

// Finished check all value is being wrote
func (s *Sukudo) Finished() bool {
	for _, v := range s.Puzzles {
		for _, v2 := range v {
			if v2 == 0 {
				return false
			}
		}
	}
	return true
}

// Success check sukudo is well done
func (s *Sukudo) Success() bool {
	return s.validate(false)
}

// validate 判断是否是正确的数独解题
func (s *Sukudo) validate(ignoreZero bool) bool {
	// check line
	for _, v := range s.Puzzles {
		ret := validateLine(v, ignoreZero)
		if !ret {
			return ret
		}
	}

	// check column
	ret := validateCol(s.Puzzles[:], true)
	if !ret {
		return ret
	}

	// check box
	for i := 0; i < 9; i = i + 3 {
		for j := 0; j < 9; j = j + 3 {
			if !s.validateBox(i, j) {
				return false
			}
		}
	}
	return true
}

// validateBox return true if target 3X3 box is finished
func (s *Sukudo) validateBox(x, y int) bool {
	var k int

	startx := x / 3 * 3
	starty := y / 3 * 3

	for i := startx; i < startx+3; i++ {
		for j := starty; j < starty+3; j++ {
			if k&s.Puzzles[i][j] != 0 {
				return false
			}
			k = k | s.Puzzles[i][j]
		}
	}

	return true
}

// validateLine return true if target row  finished
func validateLine(line [9]int, ignoreZero bool) bool {
	var k int
	for _, v2 := range line {
		if v2 == 0 { // not finished
			if ignoreZero {
				continue
			}
			return false
		}
		if k&v2 != 0 {
			return false
		}
		k = k | v2
	}
	return true
}

// validateCol return true if target column finished
func validateCol(col [][9]int, ignoreZero bool) bool {

	sz := len(col)
	for i := 0; i < 9; i++ {
		var k int
		for j := 0; j < sz; j++ {
			if col[j][i] == 0 { // not finished
				if ignoreZero {
					continue
				}
				return false
			}

			if k&col[j][i] != 0 {
				return false
			}
			k = k | col[j][i]
		}
	}
	return true
}

// ResultIn 从原始数字导入
func (s *Sukudo) ResultIn(origin [9][9]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if origin[i][j] != 0 {
				s.Puzzles[i][j] = unmapping[origin[i][j]]
			}
		}
	}
}

// ResultOut 导出结果，为原始的数字
func (s *Sukudo) ResultOut() [9][9]int {
	var ret [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ret[i][j] = mapping[s.Puzzles[i][j]]
		}
	}
	return ret
}

// save alll optional number with target blank grid
type optional struct {
	x    int
	y    int
	opts []int
}

// doSolve try to fill by only has one election number
func (s *Sukudo) doSolve() (bool, []optional) {
	if !s.validate(true) {
		return false, nil
	}

	optionals := make([]optional, 0)
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			for e := 0; e < 9; e++ {

				if s.Exited() {
					break
				}

				i := x*3 + e/3
				j := y*3 + e%3

				// 找到空白格
				if s.Puzzles[i][j] == 0 {
					// 找出该格所在的9格中，所有可能的解
					boxsElection := s.getCandidatesInBox(i, j)
					if len(boxsElection) == 1 {
						s.Puzzles[i][j] = boxsElection[0]
						s.registerCheck(i, j)
						return s.doSolve()
					}

					// 找出所在该行中， 所有可能的解
					rowElection := s.getCandidatesInRow(i)
					if len(rowElection) == 1 {
						s.Puzzles[i][j] = rowElection[0]
						s.registerCheck(i, j)
						return s.doSolve()
					}

					// 找出所在该列中， 所有可能的解
					colElection := s.getCandidatesInColumn(j)
					if len(colElection) == 1 {
						s.Puzzles[i][j] = colElection[0]
						s.registerCheck(i, j)
						return s.doSolve()
					}

					if s.isCheck(i, j) {
						continue
					}

					// 保存所有可能的候选取交集
					electons := make([]int, 0)
					var k int
					// 根据当前9宫格， 取出全集可能性
					for _, v := range boxsElection {
						if k&v == 0 {
							k = k | v
						}
					}

					// 根据行取交集
					var lk int
					for _, v := range rowElection {
						if k&v != 0 {
							lk = lk | v
						}
					}

					// 根据行与列取交集
					for _, v := range colElection {
						if lk&v != 0 {
							electons = append(electons, v) // save result
						}
					}

					optionalData := &optional{x: i, y: j, opts: electons}
					optionals = append(optionals, *optionalData)
				} else {
					s.registerCheck(i, j)
				}
			}
		}
	}

	// solve data by all optional value
	foundUni := false
	for _, opt := range optionals {
		if len(opt.opts) == 1 {
			s.Puzzles[opt.x][opt.y] = opt.opts[0]
			// fmt.Println("found unique solution x=", opt.x, "y=", opt.y, " value=", mapping[opt.opts[0]])
			foundUni = true

			if s.Exited() {
				break
			}
		}
	}

	if foundUni {
		return s.doSolve() // redo solve
	}

	return true, optionals
}

// Solve to solve sudoku.  if ok is true means success.
func (s *Sukudo) Solve() (ok bool, count int32) {

	for {
		ok, optionals := s.doSolve() // if ok is true then continue to solve
		if !ok {
			return ok, *s.tryCounter
		}

		if s.Exited() {
			break
		}

		atomic.AddInt32(s.tryCounter, 1)

		if *s.tryCounter >= 100000 {
			break
		}

		// if all is fulfilled
		if len(optionals) == 0 {
			return true, *s.tryCounter
		}

		stepinfo := syncx.NewMap[int, optional]()
		for _, v := range optionals {
			stepinfo.Store(len(v.opts), v)
		}

		_, ops := stepinfo.MinKey(func(i1, i2 int) int {
			return i1 - i2
		})

		// Pick each eligible number, fill it and see if it works
		// Do this CONCURRENTLY to save time
		chanSudokuSolve := make(chan Channel)
		wg := new(sync.WaitGroup)

		for _, v := range ops.opts {
			if s.Exited() {
				break
			}
			wg.Add(1)

			go func(in *Sukudo, wg *sync.WaitGroup, rowID int, colID int, value int, c chan Channel) {
				defer wg.Done()
				in.Puzzles[rowID][colID] = value

				ret, count := in.Solve()
				atomic.AddInt32(in.tryCounter, count)
				c <- Channel{in, ret}

			}(s.Copy(), wg, ops.x, ops.y, v, chanSudokuSolve)

		}

		// wait for the threads to be done & close channel once all threads are done
		go func(wg *sync.WaitGroup, c chan Channel) {
			wg.Wait()
			close(c)
		}(wg, chanSudokuSolve)

		// collect the results and look for the right guess
		for r := range chanSudokuSolve {
			_solved := r.Solved

			if _solved {
				// 如果已求得解，则标记exit 为true, 让其它异步处理也及时退出
				*s.Exit = true
				s.Puzzles = r.Intermediate.Puzzles
				atomic.AddInt32(s.tryCounter, *r.Intermediate.tryCounter)
				return _solved, *s.tryCounter
			}
		}

	}

	return false, *s.tryCounter
}

type Channel struct {
	Intermediate *Sukudo
	Solved       bool
}

func (s *Sukudo) registerCheck(x, y int) {
	v := strconv.Itoa(x) + strconv.Itoa(y)
	s.checked.Store(v, true)
}

func (s *Sukudo) isCheck(x, y int) bool {
	v := strconv.Itoa(x) + strconv.Itoa(y)
	k, ok := s.checked.Load(v)
	return ok && k
}

// 求当前的空格在所在行中可填写的数字
func (s *Sukudo) getCandidatesInRow(row int) []int {
	found := 0
	ret := [9]int{}
	var k int
	for i := 0; i < 9; i++ {
		if s.Puzzles[row][i] != 0 {
			k = k | s.Puzzles[row][i]
		}
	}

	for key := range mapping {
		if key&k != 0 {
			// exsit
			continue
		}
		ret[found] = key
		found++
	}
	return ret[:found]
}

// 求当前的空格在所在列中可填写的数字
func (s *Sukudo) getCandidatesInColumn(col int) []int {
	found := 0
	ret := [9]int{}
	var k int
	for i := 0; i < 9; i++ {
		if s.Puzzles[i][col] != 0 {
			k = k | s.Puzzles[i][col]
		}
	}

	for key := range mapping {
		if key&k != 0 {
			// exsit
			continue
		}
		ret[found] = key
		found++
	}
	return ret[:found]
}

// 求当前的空格在9格式中可填写的数字
func (s *Sukudo) getCandidatesInBox(x, y int) []int {
	found := 0
	ret := [9]int{}
	var k int

	// 计算x,y 属于哪一块 3*3的格式中
	startx := x / 3 * 3
	starty := y / 3 * 3

	for i := startx; i < startx+3; i++ {
		for j := starty; j < starty+3; j++ {
			k = k | s.Puzzles[i][j]
		}
	}

	for key := range mapping {
		if key&k != 0 {
			// exsit
			continue
		}
		ret[found] = key
		found++
	}
	return ret[:found]
}
