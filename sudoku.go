// 这是一个数独游戏的核心代码，其中包含了数独的验证、求解等功能。
package sudoku

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/jhunters/goassist/concurrent/syncx"
	"github.com/jhunters/goassist/conv"
)

type optional struct {
	x    int
	y    int
	opts []int
}

// SudokuX is a flexiable sudoku puzzle resolver
type SudokuX struct {
	// SudokuX puzzle content
	Puzzles [][]int

	Max int

	// 标记是否处理过， key=rowid+colid,  value=true表示已处理
	checked *syncx.Map[string, bool]

	// 标记 本次解题完成，可退出
	Exit *bool

	// 最大尝试次数
	tryCounter *int32

	mapping   map[int]int
	unmapping map[int]int

	boxWMax int
	boxHMax int
	xCount  int
	yCount  int
}

func NewSudokuX(max int) (*SudokuX, error) {
	if max > 1 && max < 10 && (isPerfectSquare(float64(max)) || max%2 == 0) {
		skd := &SudokuX{checked: syncx.NewMap[string, bool]()}
		skd.Exit = conv.ToPtr(false)
		skd.Max = max
		skd.Puzzles = Init2dimArray(max)
		skd.tryCounter = conv.ToPtr(int32(0))
		skd.mapping = make(map[int]int)
		skd.unmapping = make(map[int]int)
		for i := 1; i <= max; i++ {
			v := 1 << (i - 1)
			skd.mapping[v] = i
			skd.unmapping[i] = v
		}

		if max%2 == 0 {
			skd.boxWMax = max / 2
			skd.boxHMax = 2
			skd.xCount = max / 2
			skd.yCount = 2
		} else {
			sqrt := int(math.Sqrt(float64(max)))
			skd.boxWMax = sqrt
			skd.boxHMax = sqrt
			skd.xCount = sqrt
			skd.yCount = sqrt
		}

		return skd, nil
	}
	return nil, fmt.Errorf("invalid max number %d", max)
}

func isPerfectSquare(n float64) bool {
	// 计算平方根
	sqrt := math.Sqrt(n)

	// 判断平方根是否为正整数
	return sqrt == math.Floor(sqrt) && n > 0
}

// ResultIn 从原始数字导入
func (s *SudokuX) ResultIn(origin [][]int) {
	for i := 0; i < s.Max; i++ {
		for j := 0; j < s.Max; j++ {
			if origin[i][j] != 0 {
				s.Puzzles[i][j] = s.unmapping[origin[i][j]]
			}
		}
	}
}

// Copy a new SudokuX struct by has some pointer to Exit and tryCounter field
func (s *SudokuX) Copy() *SudokuX {
	mp := s.checked.Copy()
	puzzles := Init2dimArray(s.Max)
	for i := 0; i < s.Max; i++ {
		copy(puzzles[i], s.Puzzles[i])
	}
	ret := &SudokuX{Puzzles: puzzles, checked: mp, Max: s.Max, mapping: s.mapping,
		unmapping: s.unmapping, boxWMax: s.boxWMax, boxHMax: s.boxHMax, xCount: s.xCount, yCount: s.yCount}
	ret.Exit = s.Exit
	ret.tryCounter = s.tryCounter
	return ret
}

// Print SudokuX puzzle
func (s *SudokuX) Print() {
	PrintX(s.ResultOut(), s.Max, s.boxWMax, s.boxHMax)
}

// Exited return true if should exit loop
func (s *SudokuX) Exited() bool {
	return s.Exit != nil && *s.Exit
}

// Finished check all value is being wrote
func (s *SudokuX) Finished() bool {
	for _, v := range s.Puzzles {
		for _, v2 := range v {
			if v2 == 0 {
				return false
			}
		}
	}
	return true
}

// Success check SudokuX is well done
func (s *SudokuX) Success() bool {
	return s.validate(false)
}

// validate 判断是否是正确的数独解题
func (s *SudokuX) validate(ignoreZero bool) bool {
	// check line
	for _, v := range s.Puzzles {
		vv := v[:]
		ret := s.validateLine(vv, ignoreZero)
		if !ret {
			return ret
		}
	}

	// check column
	ret := s.validateCol(s.Puzzles[:], ignoreZero)
	if !ret {
		return ret
	}

	// check box
	for i := 0; i < s.Max; i = i + s.boxHMax {
		for j := 0; j < s.Max; j = j + s.boxWMax {
			if !s.validateBox(i, j) {
				return false
			}
		}
	}
	return true
}

// validateBox return true if target 3X3 box is finished
func (s *SudokuX) validateBox(x, y int) bool {
	var k int

	startx := x / s.boxHMax * s.boxHMax
	starty := y / s.boxWMax * s.boxWMax

	for i := startx; i < startx+s.boxHMax; i++ {
		for j := starty; j < starty+s.boxWMax; j++ {
			if k&s.Puzzles[i][j] != 0 {
				return false
			}
			k = k | s.Puzzles[i][j]
		}
	}

	return true
}

// validateLine return true if target row  finished
func (s *SudokuX) validateLine(line []int, ignoreZero bool) bool {
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
func (s *SudokuX) validateCol(col [][]int, ignoreZero bool) bool {

	sz := len(col)
	for i := 0; i < s.Max; i++ {
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

// ResultInFromString parse directly from string number
func (s *SudokuX) ResultInFromString(str string) error {
	total := int(math.Pow(float64(s.Max), 2))
	if len(str) != total {
		return fmt.Errorf("invalid sudoku puzzle by string. the length should be 81")
	}

	for x := 0; x < total; x++ {
		c := str[x]
		intVal, err := strconv.Atoi(string(c))
		if err != nil {
			return fmt.Errorf("with a ono integer value '%v'", c)
		}
		if intVal < 0 || intVal > 9 {
			return fmt.Errorf("with invalid integer value '%v', should between 0~9", c)
		}

		i := x / s.Max
		j := x % s.Max
		s.Puzzles[i][j] = s.unmapping[intVal]
	}
	return nil
}

// ResultOut 导出结果，为原始的数字
func (s *SudokuX) ResultOut() [][]int {
	var ret [][]int = make([][]int, s.Max)
	for i := 0; i < s.Max; i++ {
		ret[i] = make([]int, s.Max)
		for j := 0; j < s.Max; j++ {
			ret[i][j] = s.mapping[s.Puzzles[i][j]]
		}
	}
	return ret
}

// doSolve try to fill by only has one election number
func (s *SudokuX) doSolve() (bool, []optional) {
	if s.validate(false) { // if check no zero and finished
		return true, nil
	}

	if !s.validate(true) {
		return false, nil
	}

	optionals := make([]optional, 0)
	for x := 0; x < s.xCount; x++ {
		for y := 0; y < s.yCount; y++ {
			for e := 0; e < s.Max; e++ {

				if s.Exited() {
					break
				}

				i := x*s.boxHMax + e/s.boxWMax
				j := y*s.boxWMax + e%s.boxWMax
				// 找到空白格
				if s.Puzzles[i][j] == 0 {
					// 找出该格所在的格中，所有可能的解
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
func (s *SudokuX) Solve() (ok bool, count int32) {

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
		chanSudokuSolve := make(chan ChannelX)
		wg := new(sync.WaitGroup)

		for _, v := range ops.opts {
			if s.Exited() {
				break
			}
			wg.Add(1)

			go func(in *SudokuX, wg *sync.WaitGroup, rowID int, colID int, value int, c chan ChannelX) {
				defer wg.Done()
				in.Puzzles[rowID][colID] = value

				ret, count := in.Solve()
				atomic.AddInt32(in.tryCounter, count)
				c <- ChannelX{in, ret}

			}(s.Copy(), wg, ops.x, ops.y, v, chanSudokuSolve)

		}

		// wait for the threads to be done & close channel once all threads are done
		go func(wg *sync.WaitGroup, c chan ChannelX) {
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

type ChannelX struct {
	Intermediate *SudokuX
	Solved       bool
}

func (s *SudokuX) registerCheck(x, y int) {
	v := strconv.Itoa(x) + strconv.Itoa(y)
	s.checked.Store(v, true)
}

func (s *SudokuX) isCheck(x, y int) bool {
	v := strconv.Itoa(x) + strconv.Itoa(y)
	k, ok := s.checked.Load(v)
	return ok && k
}

// 求当前的空格在所在行中可填写的数字
func (s *SudokuX) getCandidatesInRow(row int) []int {
	found := 0
	ret := make([]int, s.Max)
	var k int
	for i := 0; i < s.Max; i++ {
		if s.Puzzles[row][i] != 0 {
			k = k | s.Puzzles[row][i]
		}
	}

	for key := range s.mapping {
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
func (s *SudokuX) getCandidatesInColumn(col int) []int {
	found := 0
	ret := make([]int, s.Max)
	var k int
	for i := 0; i < s.Max; i++ {
		if s.Puzzles[i][col] != 0 {
			k = k | s.Puzzles[i][col]
		}
	}

	for key := range s.mapping {
		if key&k != 0 {
			// exsit
			continue
		}
		ret[found] = key
		found++
	}
	return ret[:found]
}

// 求当前的空格在6格式中可填写的数字
func (s *SudokuX) getCandidatesInBox(x, y int) []int {
	found := 0
	ret := make([]int, s.Max)
	var k int

	// 计算x,y 属于哪一块 的格式中
	startx := x / s.boxHMax * s.boxHMax
	starty := y / s.boxWMax * s.boxWMax

	for i := startx; i < startx+s.boxHMax; i++ {
		for j := starty; j < starty+s.boxWMax; j++ {
			k = k | s.Puzzles[i][j]
		}
	}

	for key := range s.mapping {
		if key&k != 0 {
			// exsit
			continue
		}
		ret[found] = key
		found++
	}
	return ret[:found]
}
