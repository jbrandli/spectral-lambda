package data

type band struct {
	start      int64
	end        int64
	resolution int64
}

func (b *band) characterizeSignal(spike []int) signal {
	return signal{}
}
