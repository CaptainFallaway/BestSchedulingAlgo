package graphing

type ColumnInterface interface {
	Col(renderable Renderable, colSpan ...int) ColumnInterface
}

type renderableItem struct {
	Renderable Renderable
	Bounds     componentBounds
	Dimensions CompDimensions
}

type colItem struct {
	colSpan    int
	renderable Renderable
}

type row struct {
	rowSpan       int
	totalColSpans int
	cols          []colItem
}

type Layout struct {
	Items      []renderableItem
	ItemsCount int

	rows          []row
	totalRowSpans int
}

// twidth and theight just stand for terminal width / height
func (l *Layout) CalcSizes(twidth, theight int) {
	rowParts := theight / l.totalRowSpans

	l.Items = make([]renderableItem, 0) // Reset the items

	rowOffset := 0
	for _, row := range l.rows {
		colParts := twidth / row.totalColSpans
		rowHeight := rowParts * row.rowSpan

		colOffset := 0
		for _, col := range row.cols {
			colWidth := colParts * col.colSpan

			l.Items = append(l.Items, renderableItem{
				col.renderable,
				componentBounds{
					Width:   colWidth + 1,
					Height:  rowHeight + 1,
					OffsetX: colOffset,
					OffsetY: rowOffset,
				},
				CompDimensions{
					Width:  colWidth,
					Height: rowHeight,
				},
			})

			colOffset += colWidth
		}

		rowOffset += rowHeight
	}
}

func (l *Layout) Row(rowSpan ...int) ColumnInterface {
	span := getSpan(rowSpan...)

	l.totalRowSpans += span

	l.rows = append(l.rows, row{
		rowSpan: span,
	})

	return ColumnAdder{l, len(l.rows) - 1}
}

type ColumnAdder struct {
	l       *Layout
	rowIndx int
}

// colSpan is the number of columns this renderable should span, default is 1 if not provided
func (ca ColumnAdder) Col(renderable Renderable, colSpan ...int) ColumnInterface {
	if ca.rowIndx < 0 {
		panic("rowIndx cannot be less than 0")
	}

	span := getSpan(colSpan...)

	temp := colItem{
		span,
		renderable,
	}

	ca.l.rows[ca.rowIndx].cols = append(ca.l.rows[ca.rowIndx].cols, temp)
	ca.l.rows[ca.rowIndx].totalColSpans += span
	ca.l.ItemsCount++

	return ca
}
