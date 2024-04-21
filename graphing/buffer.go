package graphing

type Buffer struct {
	BuffArr []TermPixel
	width   int
	height  int
}

func NewBuffer(w, h int) *Buffer {
	return &Buffer{
		BuffArr: make([]TermPixel, w*h),
		width:   w,
		height:  h,
	}
}

func (b *Buffer) Set(tp TermPixel) {
	indx := ((tp.Y - 1) * b.width) + (tp.X - 1)
	b.BuffArr[indx] = tp
}

func (b *Buffer) Get(x, y int) TermPixel {
	indx := ((y - 1) * b.width) + (x - 1)
	return b.BuffArr[indx]
}
