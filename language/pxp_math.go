package language

import (
	"math/rand/v2"
	"sync"

	"github.com/toxyl/math"
)

const (
	PRECOMPUTE = 100
)

var (
	fibonacciValues = func() []float64 {
		vals := make([]float64, PRECOMPUTE)
		vals[0], vals[1] = 1, 1
		for i := 2; i < PRECOMPUTE; i++ {
			vals[i] = vals[i-1] + vals[i-2]
		}
		return vals
	}()
	mutex = sync.RWMutex{}
	lt    = [][64]float64{} // Lookup table
)

func genPowNTable(n uint) [64]float64 {
	table := [64]float64{}
	for i := range 64 {
		table[i] = math.Pow(float64(n), float64(i))
	}
	return table
}

// @Name: add
// @Desc: Adds the two numbers
// @Param:      a       - -   -   The first number
// @Param:      b       - -   -   The second number
// @Returns:    result  - -   -   a+b
func add(a, b float64) (float64, error) { return a + b, nil }

// @Name: add-n
// @Desc: Multiplies b by n and adds the result to a
// @Param:      a       - -   -   The first number
// @Param:      b       - -   -   The second number
// @Param:      n       - -   -   The multiplier for the second number
// @Returns:    result  - -   -   a + (n * b)
func addN(a, b, n float64) (float64, error) { return a + (n * b), nil }

// @Name: sub
// @Desc: Subtracts the two numbers
// @Param:      a       - -   -   The first number
// @Param:      b       - -   -   The second number
// @Returns:    result  - -   -   a-b
func sub(a, b float64) (float64, error) { return a - b, nil }

// @Name: sub-n
// @Desc: Multiplies b by n and subtracts the result from a
// @Param:      a       - -   -   The first number
// @Param:      b       - -   -   The second number
// @Param:      n       - -   -   The multiplier for the second number
// @Returns:    result  - -   -   a - (n * b)
func subN(a, b, n float64) (float64, error) { return a - (n * b), nil }

// @Name: mul
// @Desc: Multiplies the two numbers
// @Param:      a       - -   -   The first number
// @Param:      b       - -   -   The second number
// @Returns:    result  - -   -   a*b
func mul(a, b float64) (float64, error) { return a * b, nil }

// @Name: div
// @Desc: Divides the two numbers
// @Param:      a       - -   -   The first number
// @Param:      b       - -   -   The second number
// @Returns:    result  - -   -   a/b
func div(a, b float64) (float64, error) {
	if b == 0 {
		return 0, nil // avoid division-by-zero error
	}
	return a / b, nil
}

// @Name: fibonacci
// @Desc: Calculates the nth fibonacci number using 1-based indexing with memoization
// @Param:      nth     - -   -   The nth fibonacci number to calculate
// @Returns:    result  - -   -   The nth fibonacci number
func fibonacci(nth float64) (float64, error) {
	n := int(nth)
	// Handle edge cases
	if n <= 0 {
		return 0, nil
	}
	if n == 1 || n == 2 {
		return 1, nil
	}

	mutex.Lock()
	defer mutex.Unlock()

	// Ensure we have enough values in the slice
	for len(fibonacciValues) < n {
		fibonacciValues = append(fibonacciValues, 0)
	}

	// If we already calculated this value, return it
	if fibonacciValues[n-1] != 0 {
		return fibonacciValues[n-1], nil
	}

	// Calculate all values up to n if needed
	for i := 3; i <= n; i++ {
		if fibonacciValues[i-1] == 0 {
			fibonacciValues[i-1] = fibonacciValues[i-2] + fibonacciValues[i-3]
		}
	}

	return fibonacciValues[n-1], nil
}

// @Name: floor
// @Desc: Returns the largest integer less than or equal to x
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The largest integer less than or equal to x
func floor(x float64) (float64, error) {
	return math.Floor(x), nil
}

// @Name: ceil
// @Desc: Returns the smallest integer greater than or equal to x
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The smallest integer greater than or equal to x
func ceil(x float64) (float64, error) {
	return math.Ceil(x), nil
}

// @Name: round
// @Desc: Returns the nearest integer to x, rounding to even on ties
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The nearest integer to x
func round(x float64) (float64, error) {
	return math.Round(x), nil
}

// @Name: min
// @Desc: Returns the minimum value of x and y
// @Param:      x       - -   -   The x value
// @Param:      y       - -   -   The y value
// @Returns:    result  - -   -   The minimum value of x and y
func min(x, y float64) (float64, error) {
	return math.Min(x, y), nil
}

// @Name: max
// @Desc: Returns the maximum value of x and y
// @Param:      x       - -   -   The x value
// @Param:      y       - -   -   The y value
// @Returns:    result  - -   -   The maximum value of x and y
func max(x, y float64) (float64, error) {
	return math.Max(x, y), nil
}

// @Name: delta
// @Desc: Returns the delta between x and y
// @Param:      x       - -   -   The x value
// @Param:      y       - -   -   The y value
// @Returns:    result  - -   -   The delta between x and y
func delta(x, y float64) (float64, error) {
	return math.Max(x, y) - math.Min(x, y), nil
}

// @Name: abs
// @Desc: Returns the absolute value of x
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The absolute value of x
func abs(x float64) (float64, error) {
	return math.Abs(x), nil
}

// @Name: slope
// @Desc: Calculates the slope between two points
// @Param:      x1      - -   -   The x coordinate of the first point
// @Param:      y1      - -   -   The y coordinate of the first point
// @Param:      x2      - -   -   The x coordinate of the second point
// @Param:      y2      - -   -   The y coordinate of the second point
// @Returns:    result  - -   -   The slope value
func slope(x1, y1, x2, y2 float64) (float64, error) {
	return (y2 - y1) / (x2 - x1), nil
}

// @Name: tan-of-slope
// @Desc: Calculates the angle from a slope value
// @Param:      m       - -   -   The slope value
// @Returns:    result  - -   -   The angle in radians
func tanOfSlope(m float64) (float64, error) {
	return math.Atan(m), nil
}

// @Name: hypotenuse-of-triangle
// @Desc: Calculates hypotenuse from adjacent, opposite and gamma angle
// @Param:      adjacent - -   -   The adjacent side length
// @Param:      opposite - -   -   The opposite side length
// @Param:      gamma    - -   -   The gamma angle
// @Returns:    result   - -   -   The hypotenuse length
func hypotenuseOfTriangle(adjacent, opposite, gamma float64) (float64, error) {
	return math.Sqrt((adjacent * adjacent) + (opposite * opposite) - 2*adjacent*opposite*math.Cos(gamma)), nil
}

// @Name: adjacent-of-triangle
// @Desc: Calculates adjacent side from hypotenuse, opposite and alpha angle
// @Param:      hypotenuse - - -   The hypotenuse length
// @Param:      opposite   - - -   The opposite side length
// @Param:      alpha      - - -   The alpha angle
// @Returns:    result     - - -   The adjacent side length
func adjacentOfTriangle(hypotenuse, opposite, alpha float64) (float64, error) {
	return math.Sqrt((opposite * opposite) + (hypotenuse * hypotenuse) - 2*opposite*hypotenuse*math.Cos(alpha)), nil
}

// @Name: opposite-of-triangle
// @Desc: Calculates opposite side from hypotenuse, adjacent and beta angle
// @Param:      hypotenuse - - -   The hypotenuse length
// @Param:      adjacent   - - -   The adjacent side length
// @Param:      beta       - - -   The beta angle
// @Returns:    result     - - -   The opposite side length
func oppositeOfTriangle(hypotenuse, adjacent, beta float64) (float64, error) {
	return math.Sqrt((hypotenuse * hypotenuse) + (adjacent * adjacent) - 2*hypotenuse*adjacent*math.Cos(beta)), nil
}

// @Name: circumference-of-a_circle
// @Desc: Calculates circumference from radius
// @Param:      radius  - -   -   The radius of the circle
// @Returns:    result  - -   -   The circumference length
func circumferenceOfACircle(radius float64) (float64, error) {
	return 2 * math.Pi * radius, nil
}

// @Name: distance-between
// @Desc: Calculates distance between two points
// @Param:      x1      - -   -   The x coordinate of the first point
// @Param:      y1      - -   -   The y coordinate of the first point
// @Param:      x2      - -   -   The x coordinate of the second point
// @Param:      y2      - -   -   The y coordinate of the second point
// @Returns:    result  - -   -   The distance between the points
func distanceBetween(x1, y1, x2, y2 float64) (float64, error) {
	dx := x2 - x1
	dy := y2 - y1
	return math.Sqrt(dx*dx + dy*dy), nil
}

// @Name: angle-between
// @Desc: Calculates angle between two points
// @Param:      x1      - -   -   The x coordinate of the first point
// @Param:      y1      - -   -   The y coordinate of the first point
// @Param:      x2      - -   -   The x coordinate of the second point
// @Param:      y2      - -   -   The y coordinate of the second point
// @Returns:    result  - -   -   The angle in radians
func angleBetween(x1, y1, x2, y2 float64) (float64, error) {
	return math.Atan2(y2-y1, x2-x1), nil
}

// @Name: square
// @Desc: Calculates the square of a number
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The square of x
func square(x float64) (float64, error) {
	return x * x, nil
}

// @Name: pow
// @Desc: Calculates base raised to the power of n, using lookup tables for integer bases when possible
// @Param:      base    - -   -   The base value
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   Base^n
func pow(base, n float64) (float64, error) {
	if base != float64(int(base)) {
		return math.Pow(base, n), nil // this is a float value, so we can't use a lookup

	}
	// Ensure the lookup table has enough entries
	for len(lt) <= int(base) {
		lt = append(lt, genPowNTable(uint(len(lt))))
	}

	if n == float64(int(n)) && int(n) > -64 && n < 64 {
		// this is a whole number we can use with a lookup table
		v2 := lt[int(base)][uint(math.Abs(n))]
		if n < 0 {
			return float64(1.0 / float64(v2)), nil
		}
		return float64(v2), nil
	}
	// this is not a whole number or outside of range of a lookup table
	return math.Pow(float64(base), n), nil
}

// @Name: pow2
// @Desc: Calculates 2 raised to the power of n
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   2^n
func pow2(n float64) (float64, error) { return pow(2, n) }

// @Name: pow4
// @Desc: Calculates 4 raised to the power of n
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   4^n
func pow4(n float64) (float64, error) { return pow(4, n) }

// @Name: pow8
// @Desc: Calculates 8 raised to the power of n
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   8^n
func pow8(n float64) (float64, error) { return pow(8, n) }

// @Name: pow10
// @Desc: Calculates 10 raised to the power of n
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   10^n
func pow10(n float64) (float64, error) { return pow(10, n) }

// @Name: pow12
// @Desc: Calculates 12 raised to the power of n
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   12^n
func pow12(n float64) (float64, error) { return pow(12, n) }

// @Name: pow16
// @Desc: Calculates 16 raised to the power of n
// @Param:      n       - -   -   The exponent
// @Returns:    result  - -   -   16^n
func pow16(n float64) (float64, error) { return pow(16, n) }

// @Name: sqrt
// @Desc: Returns the square root of x
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The square root of x
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return math.NaN[float64](), nil
	}
	return math.Sqrt(x), nil
}

// @Name: log
// @Desc: Returns the natural logarithm of x
// @Param:      x       - -   -   The input value
// @Returns:    result  - -   -   The natural logarithm of x
func log(x float64) (float64, error) {
	if x <= 0 {
		return math.NaN[float64](), nil
	}
	return math.Log(x), nil
}

// @Name: random-range
// @Desc: Returns a random number between min and max
// @Param:      min     - -   -   The minimum value
// @Param:      max     - -   -   The maximum value
// @Returns:    result  - -   -   A random float64 value between min and max, with NaN handling
func randomRange(min, max float64) (float64, error) {
	if min != min { // Check for NaN
		min = 0
	}
	if max != max { // Check for NaN
		max = 1
	}
	return min + rand.Float64()*(max-min), nil
}

// @Name: or
// @Desc: Returns either value1 or value2 randomly
// @Param:      value1  - -   -   The first value
// @Param:      value2  - -   -   The second value
// @Returns:    result  - -   -   One of the two input values randomly
func randomOr(value1, value2 float64) (float64, error) {
	if rand.IntN(2) == 1 {
		return value1, nil
	}
	return value2, nil
}

// @Name: degrees2radians
// @Desc: converts degrees to radians
// @Param:      degrees  - -   -   The angle in degrees
// @Returns:    result   - -   -   angle in radians
func degrees2Radians(degrees float64) (float64, error) { return degrees * (math.Pi / 180), nil }

// @Name: grads2radians
// @Desc: converts grads to radians
// @Param:      grads    - -   -   The angle in grads
// @Returns:    result   - -   -   angle in radians
func grads2Radians(grads float64) (float64, error) { return grads * (math.Pi / 200), nil }

// @Name: radians2degrees
// @Desc: converts radians to degrees
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   angle in degrees
func radians2Degrees(radians float64) (float64, error) { return radians * (180 / math.Pi), nil }

// @Name: radians2grads
// @Desc: converts radians to grads
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   angle in grads
func radians2Grads(radians float64) (float64, error) { return radians * (200 / math.Pi), nil }

// @Name: normalize-angle
// @Desc: normalizes an angle to [0, 2π)
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   normalized angle in radians
func normalizeAngle(radians float64) (float64, error) {
	for radians < 0 {
		radians += 2 * math.Pi
	}
	for radians >= 2*math.Pi {
		radians -= 2 * math.Pi
	}
	return radians, nil
}

// @Name: normalize-angle-degrees
// @Desc: normalizes an angle to [0, 360)
// @Param:      degrees  - -   -   The angle in degrees
// @Returns:    result   - -   -   normalized angle in degrees
func normalizeAngleDegrees(degrees float64) (float64, error) {
	for degrees < 0 {
		degrees += 360
	}
	for degrees >= 360 {
		degrees -= 360
	}
	return degrees, nil
}

// @Name: angle-difference
// @Desc: calculates the smallest difference between two angles
// @Param:      angle1  - -   -   The first angle in radians
// @Param:      angle2  - -   -   The second angle in radians
// @Returns:    result   - -   -   smallest angle difference in radians
func angleDifference(angle1, angle2 float64) (float64, error) {
	diff := math.Abs(angle1 - angle2)
	if diff > math.Pi {
		diff = 2*math.Pi - diff
	}
	return diff, nil
}

// @Name: angle-difference-degrees
// @Desc: calculates the smallest difference between two angles in degrees
// @Param:      angle1  - -   -   The first angle in degrees
// @Param:      angle2  - -   -   The second angle in degrees
// @Returns:    result   - -   -   smallest angle difference in degrees
func angleDifferenceDegrees(angle1, angle2 float64) (float64, error) {
	diff := math.Abs(angle1 - angle2)
	if diff > 180 {
		diff = 360 - diff
	}
	return diff, nil
}

// @Name: sin
// @Desc: calculates the sine of an angle
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   sine value between -1 and 1
func sin(radians float64) (float64, error) { return math.Sin(radians), nil }

// @Name: asin
// @Desc: calculates the arcsine (inverse sine) of x
// @Param:      radians  - -   -   The input value
// @Returns:    result  - -   -   angle in radians between -PI/2 and PI/2
func asin(radians float64) (float64, error) { return math.Asin(radians), nil }

// @Name: cos
// @Desc: calculates the cosine of an angle in radians
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   cosine value between -1 and 1
func cos(radians float64) (float64, error) { return math.Cos(radians), nil }

// @Name: acos
// @Desc: calculates the arccosine (inverse cosine) of x
// @Param:      radians  - -   -   The input value
// @Returns:    result   - -   -   angle in radians between 0 and PI
func acos(radians float64) (float64, error) { return math.Acos(radians), nil }

// @Name: tan
// @Desc: calculates the tangent of an angle in radians
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   tangent value (unbounded)
func tan(radians float64) (float64, error) { return math.Tan(radians), nil }

// @Name: atan
// @Desc: calculates the arctangent (inverse tangent) of x
// @Param:      radians  - -   -   The input value
// @Returns:    result   - -   -   angle in radians between -PI/2 and PI/2
func atan(radians float64) (float64, error) { return math.Atan(radians), nil }

// @Name: sec
// @Desc: calculates the secant of an angle in radians
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   secant value (1/cos)
func sec(radians float64) (float64, error) { return 1 / math.Cos(radians), nil }

// @Name: cosec
// @Desc: calculates the cosecant of an angle in radians
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   cosecant value (1/sin)
func cosec(radians float64) (float64, error) { return 1 / math.Sin(radians), nil }

// @Name: cot
// @Desc: calculates the cotangent of an angle in radians
// @Param:      radians  - -   -   The angle in radians
// @Returns:    result   - -   -   cotangent value (1/tan)
func cot(radians float64) (float64, error) { return 1 / math.Tan(radians), nil }

// @Name: sinh
// @Desc: calculates the hyperbolic sine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   hyperbolic sine value
func sinh(x float64) (float64, error) { return math.Sinh(x), nil }

// @Name: cosh
// @Desc: calculates the hyperbolic cosine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   hyperbolic cosine value (always positive)
func cosh(x float64) (float64, error) { return math.Cosh(x), nil }

// @Name: tanh
// @Desc: calculates the hyperbolic tangent of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   hyperbolic tangent value between -1 and 1
func tanh(x float64) (float64, error) { return math.Tanh(x), nil }

// @Name: sech
// @Desc: calculates the hyperbolic secant of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   hyperbolic secant value (1/cosh)
func sech(x float64) (float64, error) { return 1 / math.Cosh(x), nil }

// @Name: csch
// @Desc: calculates the hyperbolic cosecant of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   hyperbolic cosecant value (1/sinh)
func csch(x float64) (float64, error) { return 1 / math.Sinh(x), nil }

// @Name: coth
// @Desc: calculates the hyperbolic cotangent of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   hyperbolic cotangent value (1/tanh)
func coth(x float64) (float64, error) { return 1 / math.Tanh(x), nil }

// @Name: asinh
// @Desc: calculates the inverse hyperbolic sine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   inverse hyperbolic sine value
func asinh(x float64) (float64, error) { return math.Asinh(x), nil }

// @Name: acosh
// @Desc: calculates the inverse hyperbolic cosine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   inverse hyperbolic cosine value (x >= 1)
func acosh(x float64) (float64, error) { return math.Acosh(x), nil }

// @Name: atanh
// @Desc: calculates the inverse hyperbolic tangent of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   inverse hyperbolic tangent value (|x| < 1)
func atanh(x float64) (float64, error) { return math.Atanh(x), nil }

// @Name: asech
// @Desc: calculates the inverse hyperbolic secant of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   inverse hyperbolic secant value (0 < x <= 1)
func asech(x float64) (float64, error) { return math.Acosh(1 / x), nil }

// @Name: acsch
// @Desc: calculates the inverse hyperbolic cosecant of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   inverse hyperbolic cosecant value (x != 0)
func acsch(x float64) (float64, error) { return math.Asinh(1 / x), nil }

// @Name: acoth
// @Desc: calculates the inverse hyperbolic cotangent of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   inverse hyperbolic cotangent value (|x| > 1)
func acoth(x float64) (float64, error) { return 0.5 * math.Log((x+1)/(x-1)), nil }

// @Name: versin
// @Desc: calculates the versed sine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   versed sine value (1 - cos(x))
func versin(x float64) (float64, error) { return 1 - math.Cos(x), nil }

// @Name: vercos
// @Desc: calculates the versed cosine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   versed cosine value (1 + cos(x))
func vercos(x float64) (float64, error) { return 1 + math.Cos(x), nil }

// @Name: coversin
// @Desc: calculates the coversed sine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   coversed sine value (1 - sin(x))
func coversin(x float64) (float64, error) { return 1 - math.Sin(x), nil }

// @Name: covercos
// @Desc: calculates the coversed cosine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   coversed cosine value (1 + sin(x))
func covercos(x float64) (float64, error) { return 1 + math.Sin(x), nil }

// @Name: haversin
// @Desc: calculates the haversine of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   haversine value (0.5 * (1 - cos(x)))
func haversin(x float64) (float64, error) { return 0.5 * (1 - math.Cos(x)), nil }

// @Name: exsec
// @Desc: calculates the exsecant of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   exsecant value (sec(x) - 1)
func exsec(x float64) (float64, error) { s, _ := sec(x); return s - 1, nil }

// @Name: excsc
// @Desc: calculates the excosecant of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   excosecant value (cosec(x) - 1)
func excsc(x float64) (float64, error) { c, _ := cosec(x); return c - 1, nil }

// @Name: chord
// @Desc: calculates the chord of x
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   chord value (2 * sin(x/2))
func chord(x float64) (float64, error) { return 2 * math.Sin(x/2), nil }

// @Name: sin-of-triangle
// @Desc: calculates sine using opposite and hypotenuse sides
// @Param:      opposite   - - -   The opposite side length
// @Param:      hypotenuse - - -   The hypotenuse length
// @Returns:    result     - - -   sine value (opposite/hypotenuse)
func sinOfTriangle(opposite, hypotenuse float64) (float64, error) { return opposite / hypotenuse, nil }

// @Name: cos-of-triangle
// @Desc: calculates cosine using adjacent and hypotenuse sides
// @Param:      adjacent   - - -   The adjacent side length
// @Param:      hypotenuse - - -   The hypotenuse length
// @Returns:    result     - - -   cosine value (adjacent/hypotenuse)
func cosOfTriangle(adjacent, hypotenuse float64) (float64, error) { return adjacent / hypotenuse, nil }

// @Name: tan-of-triangle
// @Desc: calculates tangent using opposite and adjacent sides
// @Param:      opposite - - -   The opposite side length
// @Param:      adjacent - - -   The adjacent side length
// @Returns:    result   - - -   tangent value (opposite/adjacent)
func tanOfTriangle(opposite, adjacent float64) (float64, error) { return opposite / adjacent, nil }

// @Name: sec-of-triangle
// @Desc: calculates secant using hypotenuse and adjacent sides
// @Param:      hypotenuse - - -   The hypotenuse length
// @Param:      adjacent   - - -   The adjacent side length
// @Returns:    result     - - -   secant value (hypotenuse/adjacent)
func secOfTriangle(hypotenuse, adjacent float64) (float64, error) { return hypotenuse / adjacent, nil }

// @Name: cosec-of-triangle
// @Desc: calculates cosecant using hypotenuse and opposite sides
// @Param:      hypotenuse - - -   The hypotenuse length
// @Param:      opposite   - - -   The opposite side length
// @Returns:    result     - - -   cosecant value (hypotenuse/opposite)
func cosecOfTriangle(hypotenuse, opposite float64) (float64, error) {
	return hypotenuse / opposite, nil
}

// @Name: cot-of-triangle
// @Desc: calculates cotangent using adjacent and opposite sides
// @Param:      adjacent - - -   The adjacent side length
// @Param:      opposite - - -   The opposite side length
// @Returns:    result   - - -   cotangent value (adjacent/opposite)
func cotOfTriangle(adjacent, opposite float64) (float64, error) { return adjacent / opposite, nil }

// @Name: radians-of-triangle
// @Desc: calculates angle in radians using all three sides of a triangle
// @Param:      adjacent   - - -   The adjacent side length
// @Param:      opposite   - - -   The opposite side length
// @Param:      hypotenuse - - -   The hypotenuse length
// @Returns:    result     - - -   angle in radians between adjacent and opposite sides
func radiansOfTriangle(adjacent, opposite, hypotenuse float64) (float64, error) {
	return acos(((adjacent * adjacent) + (opposite * opposite) - (hypotenuse * hypotenuse)) / (2 * adjacent * opposite))
}

// @Name: sin2
// @Desc: calculates the square of sine (sin²(x))
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   squared sine value (sin(x)²)
func sin2(x float64) (float64, error) {
	sin := math.Sin(x)
	return sin * sin, nil
}

// @Name: cos2
// @Desc: calculates the square of cosine (cos²(x))
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   squared cosine value (cos(x)²)
func cos2(x float64) (float64, error) {
	cos := math.Cos(x)
	return cos * cos, nil
}

// @Name: tan2
// @Desc: calculates the square of tangent (tan²(x))
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   squared tangent value (tan(x)²)
func tan2(x float64) (float64, error) {
	tan := math.Tan(x)
	return tan * tan, nil
}

// @Name: sinc
// @Desc: calculates the sinc function (sin(x)/x)
// @Param:      x       - -   -   The input value
// @Returns:    result   - -   -   sinc value (sin(x)/x, with sinc(0) = 1)
func sinc(x float64) (float64, error) {
	if math.Abs(x) < 1e-10 {
		return 1.0, nil
	}
	return math.Sin(x) / x, nil
}

// @Name: lerp-angle
// @Desc: linearly interpolates between two angles in radians
// @Param:      angle1  - -   -   The first angle in radians
// @Param:      angle2  - -   -   The second angle in radians
// @Param:      t       - -   -   The interpolation factor (0-1)
// @Returns:    result   - -   -   interpolated angle in radians
func lerpAngle(angle1, angle2, t float64) (float64, error) {
	// Normalize both angles first
	angle1, _ = normalizeAngle(angle1)
	angle2, _ = normalizeAngle(angle2)

	// Handle angle wrapping by finding the shortest path
	diff := angle2 - angle1
	if diff > math.Pi {
		diff -= 2 * math.Pi
	} else if diff < -math.Pi {
		diff += 2 * math.Pi
	}

	result := angle1 + diff*t
	return normalizeAngle(result)
}

// @Name: lerp-angle-degrees
// @Desc: linearly interpolates between two angles in degrees
// @Param:      angle1  - -   -   The first angle in degrees
// @Param:      angle2  - -   -   The second angle in degrees
// @Param:      t       - -   -   The interpolation factor (0-1)
// @Returns:    result   - -   -   interpolated angle in degrees
func lerpAngleDegrees(angle1, angle2, t float64) (float64, error) {
	// Normalize both angles first
	angle1, _ = normalizeAngleDegrees(angle1)
	angle2, _ = normalizeAngleDegrees(angle2)

	// Handle angle wrapping by finding the shortest path
	diff := angle2 - angle1
	if diff > 180 {
		diff -= 360
	} else if diff < -180 {
		diff += 360
	}

	result := angle1 + diff*t
	return normalizeAngleDegrees(result)
}
