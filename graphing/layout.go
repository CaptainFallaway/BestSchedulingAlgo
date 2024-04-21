package graphing

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

// Helper function since i do the exact same thing in two places
func getSpan(span ...int) int {
	if len(span) >= 1 {
		return span[0]
	}

	return 1
}

func (l *Layout) AddRow(rowSpan ...int) {
	span := getSpan(rowSpan...)

	l.totalRowSpans += span

	l.rows = append(l.rows, row{
		rowSpan: span,
	})
}

// The default spans are 1, 1 (row, col) is hos they should be parsed as parameters
func (l *Layout) AddRenderable(renderable Renderable, rowIndx int, colSpan ...int) {
	if rowIndx < 0 {
		panic("rowIndx cannot be less than 0")
	}

	// Make a helper function for this
	span := getSpan(colSpan...)

	temp := colItem{
		span,
		renderable,
	}

	l.rows[rowIndx].cols = append(l.rows[rowIndx].cols, temp)
	l.rows[rowIndx].totalColSpans += span
	l.ItemsCount++
}
