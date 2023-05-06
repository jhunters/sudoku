package sudoku

import (
	"math/rand"
	"time"

	"github.com/jhunters/goassist/arrayutil"
)

// SukudoGen generate sudoku puzzle
type SukudoGen struct {
	MissCount int
}

func genRandArray() []int {
	ret := make([]int, 9)
	for i := 1; i < 10; i++ {
		ret[i-1] = i
	}

	r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	arrayutil.ShuffleRandom(ret, r)

	return ret
}

func fillBox(result [9][9]int, fillby []int, rowB, ColB int) [9][9]int {
	for i := 0; i < 9; i++ {
		rowid := i%3 + rowB
		colid := i/3 + ColB
		result[rowid][colid] = fillby[i]
	}
	return result
}

// GenSukudo generate sudoku puzzle and solved result by random way
func (sg *SukudoGen) GenSukudo() (result [9][9]int, answer [9][9]int) {
	box1 := genRandArray()
	result = fillBox(result, box1, 0, 0)
	time.Sleep(time.Millisecond)
	box2 := genRandArray()
	result = fillBox(result, box2, 3, 3)
	time.Sleep(time.Millisecond)
	box3 := genRandArray()
	result = fillBox(result, box3, 6, 6)

	sdk := NewSukudo()
	sdk.ResultIn(result)
	sdk.Solve()

	result = sdk.ResultOut()
	answer = result

	// 替换需要空缺的个数
	if sg.MissCount > 0 {
		missedPos := make([]int, 81)
		for i := 0; i < 81; i++ {
			missedPos[i] = i
		}
		r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
		arrayutil.ShuffleRandom(missedPos, r)

		i := 0
		for _, v := range missedPos {
			rowid := v % 9
			colid := v / 9
			result[rowid][colid] = 0
			i++
			if sg.MissCount <= i {
				break
			}
		}
	}

	return result, answer
}
