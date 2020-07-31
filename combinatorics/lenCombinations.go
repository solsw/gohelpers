package combinatorics

type round struct {
	index  int
	inner  *round
	outer  *round
	action func(*round)
}

var (
	_arrayLength int
	_resultFull  [][]int
	_resultPart  []int
)

func makeInnermostRound() *round {
	return &round{action: func(selfRound *round) {
		if selfRound.outer != nil {
			selfRound.index = selfRound.outer.index + 1
		} else {
			selfRound.index = 0
		}
		for ; selfRound.index < _arrayLength; selfRound.index++ {
			_resultFull = append(_resultFull, append(_resultPart, selfRound.index))
		}
		_resultPart = nil
	}}
}

func makeRound(delta int) *round {
	return &round{action: func(selfRound *round) {
		var resultPartPrev = make([]int, len(_resultPart))
		copy(resultPartPrev, _resultPart)
		if selfRound.outer != nil {
			selfRound.index = selfRound.outer.index + 1
		} else {
			selfRound.index = 0
		}
		for ; selfRound.index < _arrayLength-delta; selfRound.index++ {
			_resultPart = append(resultPartPrev, selfRound.index)
			if selfRound.inner != nil {
				selfRound.inner.action(selfRound.inner)
			}
		}
	}}
}

// LenCombinations returns all combinations with length 'combLen' of the array indexes [0..arrLen-1].
// (E.g. for the array with length 4 the following combinations of length 2 are returned: 01 02 03 12 13 23.)
func LenCombinations(arrLen, combLen int) [][]int {
	if arrLen <= 0 || combLen <= 0 || arrLen < combLen {
		return nil
	}
	_arrayLength = arrLen
	// clear possible previous content of results
	_resultFull = nil
	_resultPart = nil
	currentRound := makeInnermostRound()
	for i := 1; i < combLen; i++ {
		newRound := makeRound(i)
		currentRound.outer = newRound
		newRound.inner = currentRound
		currentRound = newRound
	}
	currentRound.action(currentRound)
	return _resultFull
}
