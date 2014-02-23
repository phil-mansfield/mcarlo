package mcrand

const (
	tauswortheDigitsRandomized = 15

	tauswortheSeqLen = 9689
	tauswortheFirstOffset = 2444
	tauswortheSecondOffset = 4187
)

type tauswortheGenerator struct {
	seq []float64
	leader, firstFollower, secondFollower int
}

func (gen *tauswortheGenerator) Init(seed uint64) {
	gen.seq = make([]float64, tauswortheSeqLen)

	digitGen := New(GoRand, seed)

	f := 0.5
	for digit := 0; digit < tauswortheDigitsRandomized; digit++ {
		for i, _ := range gen.seq {
			gen.seq[i] = digitGen.Uniform(0, f)
		}
		f /= 2.0
	}

	gen.leader = 0
	gen.firstFollower = tauswortheFirstOffset
	gen.secondFollower = tauswortheSecondOffset
}

func (gen *tauswortheGenerator) Next() float64 {
	next := gen.seq[gen.firstFollower] - gen.seq[gen.secondFollower]
	if next < 0 { next += 1.0 }
	gen.seq[gen.leader] = next

	if gen.leader == 0 { gen.leader = len(gen.seq) }
	if gen.firstFollower == 0 { gen.firstFollower = len(gen.seq) }
	if gen.secondFollower == 0 { gen.secondFollower = len(gen.seq) }

	gen.leader--
	gen.firstFollower--
	gen.secondFollower--

	return next
}
