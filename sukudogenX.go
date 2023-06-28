package sudoku

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/jhunters/goassist/arrayutil"
)

// SukudoGen generate sudoku puzzle
type SukudoGenX struct {
	MissCount  int
	max        int
	boxWMax    int
	boxHMax    int
	xCount     int
	yCount     int
	totalCount int
}

func NewSukudoGenX(max, misscount int) (*SukudoGenX, error) {
	if max > 1 && max < 10 && (isPerfectSquare(float64(max)) || max%2 == 0) {
		sg := &SukudoGenX{max: max, MissCount: misscount}
		sg.totalCount = int(math.Pow(float64(sg.max), 2))
		if sg.MissCount > sg.totalCount-5 {
			fmt.Println("too big miss count value:", sg.MissCount)
			sg.totalCount = sg.totalCount - 5
		}
		if max%2 == 0 {
			sg.boxWMax = max / 2
			sg.boxHMax = 2
			sg.xCount = max / 2
			sg.yCount = 2
		} else {
			sqrt := int(math.Sqrt(float64(max)))
			sg.boxWMax = sqrt
			sg.boxHMax = sqrt
			sg.xCount = sqrt
			sg.yCount = sqrt
		}

		return sg, nil
	}
	return nil, fmt.Errorf("invalid max number %d", max)
}

func (sg *SukudoGenX) genRandArray() []int {
	ret := make([]int, sg.max)
	for i := 1; i <= sg.max; i++ {
		ret[i-1] = i
	}

	r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	arrayutil.ShuffleRandom(ret, r)

	return ret
}

func (sg *SukudoGenX) fillBox(result [][]int, fillby []int, rowB, ColB int) [][]int {
	for i := 0; i < sg.max; i++ {
		rowid := i/sg.boxWMax + rowB
		colid := i%sg.boxWMax + ColB
		result[rowid][colid] = fillby[i]
	}
	return result
}

// GenSukudo generate sudoku puzzle and solved result by random way
func (sg *SukudoGenX) GenSukudo() (result [][]int, answer [][]int, err error) {
	result = Init2dimArray(sg.max)

	beginX := 0
	beginY := 0
	for i := 0; i < sg.yCount; i++ {
		box1 := sg.genRandArray()
		result = sg.fillBox(result, box1, beginX, beginY)
		time.Sleep(10 * time.Millisecond)
		beginX = beginX + sg.boxHMax
		beginY = beginY + sg.boxWMax
	}

	sdk, err := NewSukudoX(sg.max)
	sdk.ResultIn(result)
	sdk.Solve()
	result = sdk.ResultOut()
	copy(answer, result)

	// 替换需要空缺的个数
	if sg.MissCount > 0 {
		count := sg.totalCount
		missedPos := make([]int, count)
		for i := 0; i < count; i++ {
			missedPos[i] = i
		}
		r := rand.New(rand.NewSource(int64(time.Now().UnixNano())))
		arrayutil.ShuffleRandom(missedPos, r)

		i := 0
		for _, v := range missedPos {
			rowid := v % sg.max
			colid := v / sg.max
			result[rowid][colid] = 0
			i++
			if sg.MissCount <= i {
				break
			}
		}
	}

	return result, answer, err
}
