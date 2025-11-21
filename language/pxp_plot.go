package language

import (
	"fmt"
	"image"
	"image/color"

	"github.com/toxyl/math"
)

// @Name: plot-data
// @Desc: Renders a chart from CSV data by plotting selected columns with specified colors
// @Param:      width   - - -   Chart width in pixels
// @Param:      height  - - -   Chart height in pixels
// @Param:      data    - - -   2D array of data from CSV
// @Param:      columns - - -   Array of column indices to plot
// @Param:      colors  - - -   Array of colors for each column
// @Returns:    result  - - -   The chart image
func plotData(width, height int, data [][]float64, columns []any, colors []any) (*image.NRGBA64, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be greater than 0")
	}

	if len(data) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	if len(columns) != len(colors) {
		return nil, fmt.Errorf("columns and colors arrays must have the same length")
	}

	if len(columns) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	columnsInt := make([]int, len(columns))
	for i, v := range columns {
		if intVal, ok := v.(int); ok {
			columnsInt[i] = intVal
		} else if floatVal, ok := v.(float64); ok {
			columnsInt[i] = int(floatVal)
		} else {
			return nil, fmt.Errorf("invalid column index type at position %d", i)
		}
	}

	colorsRGBA64 := make([]color.RGBA64, len(colors))
	for i, v := range colors {
		if col, ok := v.(color.RGBA64); ok {
			colorsRGBA64[i] = col
		} else {
			return nil, fmt.Errorf("invalid color type at position %d", i)
		}
	}

	numDatapoints := len(data)
	if numDatapoints == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	img := IC(width, height, color.RGBA64{0, 0, 0, 0})
	xScale := float64(width) / float64(numDatapoints)

	for colIdx, colIndex := range columnsInt {
		if colIndex < 0 {
			continue
		}

		colColor := colorsRGBA64[colIdx]
		nrgbaColor := color.NRGBA64{R: colColor.R, G: colColor.G, B: colColor.B, A: colColor.A}

		for i := 0; i < numDatapoints; i++ {
			if colIndex >= len(data[i]) {
				continue
			}

			yValue := data[i][colIndex]
			if math.IsNaN(yValue) {
				continue
			}

			yNormalized := math.Clamp(yValue, 0.0, 1.0)
			yPixel := height - 1 - int(yNormalized*float64(height))
			xPixel := int(float64(i) * xScale)

			if xPixel < 0 {
				xPixel = 0
			}
			if xPixel >= width {
				xPixel = width - 1
			}
			if yPixel < 0 {
				yPixel = 0
			}
			if yPixel >= height {
				yPixel = height - 1
			}

			img.Set(xPixel, yPixel, nrgbaColor)
		}
	}

	return img, nil
}

// @Name: plot-data-compact
// @Desc: Renders a chart from CSV data by plotting selected columns with specified colors. The series will be normalized to 0..1 based on their minium and maximimum value.
// @Param:      width   - - -   Chart width in pixels
// @Param:      height  - - -   Chart height in pixels
// @Param:      data    - - -   2D array of data from CSV
// @Param:      columns - - -   Array of column indices to plot
// @Param:      colors  - - -   Array of colors for each column
// @Returns:    result  - - -   The chart image
func plotDataCompact(width, height int, data [][]float64, columns []any, colors []any) (*image.NRGBA64, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be greater than 0")
	}

	if len(data) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	if len(columns) != len(colors) {
		return nil, fmt.Errorf("columns and colors arrays must have the same length")
	}

	if len(columns) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	columnsInt := make([]int, len(columns))
	for i, v := range columns {
		if intVal, ok := v.(int); ok {
			columnsInt[i] = intVal
		} else if floatVal, ok := v.(float64); ok {
			columnsInt[i] = int(floatVal)
		} else {
			return nil, fmt.Errorf("invalid column index type at position %d", i)
		}
	}

	colorsRGBA64 := make([]color.RGBA64, len(colors))
	for i, v := range colors {
		if col, ok := v.(color.RGBA64); ok {
			colorsRGBA64[i] = col
		} else {
			return nil, fmt.Errorf("invalid color type at position %d", i)
		}
	}

	numDatapoints := len(data)
	if numDatapoints == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	img := IC(width, height, color.RGBA64{0, 0, 0, 0})
	xScale := float64(width) / float64(numDatapoints)

	for colIdx, colIndex := range columnsInt {
		if colIndex < 0 {
			continue
		}

		colColor := colorsRGBA64[colIdx]
		nrgbaColor := color.NRGBA64{R: colColor.R, G: colColor.G, B: colColor.B, A: colColor.A}

		// Find the minimum and maximum values for the selected column index
		minValue := math.Inf(1)
		maxValue := math.Inf(-1)
		for i := 0; i < numDatapoints; i++ {
			if colIndex >= len(data[i]) {
				continue
			}
			val := data[i][colIndex]
			if math.IsNaN(val) {
				continue
			}
			if val < minValue {
				minValue = val
			}
			if val > maxValue {
				maxValue = val
			}
		}
		// Avoid division by zero if data is constant
		normalize := func(v float64) float64 {
			if maxValue == minValue {
				return 0.5
			}
			return (v - minValue) / (maxValue - minValue)
		}

		for i := 0; i < numDatapoints; i++ {
			if colIndex >= len(data[i]) {
				continue
			}

			yValue := normalize(data[i][colIndex])
			if math.IsNaN(yValue) {
				continue
			}

			yNormalized := math.Clamp(yValue, 0.0, 1.0)
			yPixel := height - 1 - int(yNormalized*float64(height))
			xPixel := int(float64(i) * xScale)

			if xPixel < 0 {
				xPixel = 0
			}
			if xPixel >= width {
				xPixel = width - 1
			}
			if yPixel < 0 {
				yPixel = 0
			}
			if yPixel >= height {
				yPixel = height - 1
			}

			img.Set(xPixel, yPixel, nrgbaColor)
		}
	}

	return img, nil
}

// @Name: plot-series
// @Desc: Renders a series from CSV data by plotting a single column with colors determined by value using color stops
// @Param:      width   - - -   Chart width in pixels
// @Param:      height  - - -   Chart height in pixels
// @Param:      data    - - -   2D array of data from CSV
// @Param:      column  - - -   Column index to plot
// @Param:      min     - - -   Minimum value for color mapping
// @Param:      max     - - -   Maximum value for color mapping
// @Param:      stops   - - -   Color stops as [][]any where each stop is [threshold, hue, saturation, lightness, alpha]; additional fields are ignored
// @Param:      invertY false   Whether flip the y-axis when plotting
// @Returns:    result  - - -   The chart image
func plotSeries(width, height int, data [][]float64, column int, min, max float64, stops [][]any, invertY bool) (*image.NRGBA64, error) {
	if width <= 0 || height <= 0 {
		return nil, fmt.Errorf("width and height must be greater than 0")
	}

	if len(data) == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	if column < 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	numDatapoints := len(data)
	if numDatapoints == 0 {
		return IC(width, height, color.RGBA64{0, 0, 0, 0}), nil
	}

	img := IC(width, height, color.RGBA64{0, 0, 0, 0})
	xScale := float64(width) / float64(numDatapoints)

	for i := 0; i < numDatapoints; i++ {
		if column >= len(data[i]) {
			continue
		}

		value := data[i][column]
		if math.IsNaN(value) {
			continue
		}

		col, err := mapColor(value, min, max, stops)
		if err != nil {
			return nil, fmt.Errorf("failed to map color for data point %d: %w", i, err)
		}
		nrgbaColor := color.NRGBA64{R: col.R, G: col.G, B: col.B, A: col.A}

		valueNormalized := (value - min) / (max - min)
		if max == min {
			valueNormalized = 0.5
		}
		valueNormalized = math.Max(0.0, math.Min(1.0, valueNormalized))
		if invertY {
			valueNormalized = 1.0 - valueNormalized
		}
		yPixel := height - 1 - int(valueNormalized*float64(height))
		xPixel := int(float64(i) * xScale)

		if xPixel < 0 {
			xPixel = 0
		}
		if xPixel >= width {
			xPixel = width - 1
		}
		if yPixel < 0 {
			yPixel = 0
		}
		if yPixel >= height {
			yPixel = height - 1
		}

		img.Set(xPixel, yPixel, nrgbaColor)
	}

	return img, nil
}
