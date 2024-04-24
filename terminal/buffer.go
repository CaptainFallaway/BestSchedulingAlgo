package terminal

type buffer struct {
	BuffArr []termPixel
	width   int
	height  int
}

func NewBuffer(w, h int) *buffer {
	return &buffer{
		BuffArr: make([]termPixel, w*h),
		width:   w,
		height:  h,
	}
}

func (b *buffer) Set(tp termPixel) {
	indx := ((tp.Y - 1) * b.width) + (tp.X - 1)
	b.BuffArr[indx] = tp
}

func (b *buffer) Get(x, y int) termPixel {
	indx := ((y - 1) * b.width) + (x - 1)
	return b.BuffArr[indx]
}
