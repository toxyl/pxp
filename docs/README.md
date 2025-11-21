# PixelPipeline Script (v1.0.0)

This is a custom language implementation with support for functions, variables, and various data types.

## Data Types

The language supports the following data types:

- `int`: Integer values
- `float`: Floating-point values
- `string`: Text values (enclosed in double quotes)
- `bool`: Boolean values (`true` or `false`)
- `*image.NRGBA64`: 16-bit NRGBA image from `image`
- `color.RGBA64`: 16-bit RGBA color from `image/color`
- `Point`: Point with X and Y coordinates
- `Rect`: Reactangle with X1, Y, X2 and Y2, W and H properties

## Syntax

### Comments
Comments start and end with `#`. Linebreaks are treated as part of the comment. In comments `#` can be escaped with `\`.

### String Literals
Strings start and end with `"`. Linebreaks are treated as part of the string. In strings `"` can be escaped with `\`.

### Argument References
Script arguments can be referenced using `$1`, `$2`, etc.

### Variables

Variables can be declared and assigned using the `:` operator:

`myVar: 42`

`text: "Hello World"`

### Functions

Functions are called using the syntax `functionName(arg1 arg2 ...)`.

Arguments can be passed by position or by name.
You must either use positional arguments or named arguments, mixing is not allowed.
All arguments have defaults.

### For Loops

For loops iterate over slices or matrices using the syntax:

```
for listName[indexVar itemVar]
    # body statements #
done
```

The `done` keyword marks the end of the loop body.




## Functions

### `C(centerX=- centerY=- radius=-) ⮕ (result=)`  
_Creates a new circle with the given radius at P(x|y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `centerX` | `float64` | `-` |   |   |   | The center of the circle on the x-axis |
| `centerY` | `float64` | `-` |   |   |   | The center of the circle on the y-axis |
| `radius` | `float64` | `-` |   |   |   | The radius of the circle |
| `⮕ result` | `error` |   |   |   |   | - - - A new circle |
---

### `E(centerX=- centerY=- radiusX=- radiusY=-) ⮕ (result=)`  
_Creates a new ellipse with the given radius at P(x|y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `centerX` | `float64` | `-` |   |   |   | The center of the ellipse on the x-axis |
| `centerY` | `float64` | `-` |   |   |   | The center of the ellipse on the y-axis |
| `radiusX` | `float64` | `-` |   |   |   | The radius of the ellipse on the x-axis |
| `radiusY` | `float64` | `-` |   |   |   | The radius of the ellipse on the y-axis |
| `⮕ result` | `error` |   |   |   |   | - - - A new ellipse |
---

### `Erx(e=-) ⮕ (result=)`  
_Returns the x-component of the radius of an ellipse._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `e` | `Ellipse` | `-` |   |   |   | The ellipse to return the x-component of the radius of |
| `⮕ result` | `error` |   |   |   |   | - - - the x-component of the radius of e |
---

### `Ery(e=-) ⮕ (result=)`  
_Returns the y-component of the radius of an ellipse._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `e` | `Ellipse` | `-` |   |   |   | The ellipse to return the y-component of the radius of |
| `⮕ result` | `error` |   |   |   |   | - - - the y-component of the radius of e |
---

### `Ex(e=-) ⮕ (result=)`  
_Returns the center x-coordinate of an ellipse._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `e` | `Ellipse` | `-` |   |   |   | The ellipse to return the center x-coordinate of |
| `⮕ result` | `error` |   |   |   |   | - - - the center x-coordinate of e |
---

### `Ey(e=-) ⮕ (result=)`  
_Returns the center y-coordinate of an ellipse._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `e` | `Ellipse` | `-` |   |   |   | The ellipse to return the center y-coordinate of |
| `⮕ result` | `error` |   |   |   |   | - - - the center y-coordinate of e |
---

### `FS(color=-) ⮕ (result=)`  
_Creates a new fill style._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `color` | `color.RGBA64` | `-` |   |   |   | The fill color |
| `⮕ result` | `error` |   |   |   |   | - - - A new fill style |
---

### `I(w=- h=-) ⮕ (result=)`  
_Creates a new transparent image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `w` | `int` | `-` |   |   |   | The width of the image |
| `h` | `int` | `-` |   |   |   | The height of the image |
| `⮕ result` | `error` |   |   |   |   | - - - The new image |
---

### `IC(w=- h=- cFill=-) ⮕ (result=)`  
_Creates a new image with the given color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `w` | `int` | `-` |   |   |   | The width of the image |
| `h` | `int` | `-` |   |   |   | The height of the image |
| `cFill` | `color.RGBA64` | `-` |   |   |   | The fill color |
| `⮕ result` | `error` |   |   |   |   | - - - The new image |
---

### `Ih(img=-) ⮕ (result=)`  
_Returns the height of an image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to return the height of |
| `⮕ result` | `error` |   |   |   |   | - - - The height of img |
---

### `Ir(img=-) ⮕ (result=)`  
_Returns the aspect ratio of the given image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to return the aspect ratio of |
| `⮕ result` | `error` |   |   |   |   | - - - Aspect ratio of the image |
---

### `It(img=- dt=-) ⮕ (result=)`  
_Translates the given image by expanding/cropping the left &#43; top borders._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to translate |
| `dt` | `Point` | `-` |   |   |   | The translation to apply |
| `⮕ result` | `error` |   |   |   |   | - - - The new image |
---

### `Iw(img=-) ⮕ (result=)`  
_Returns the width of an image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to return the width of |
| `⮕ result` | `error` |   |   |   |   | - - - The width of img |
---

### `LS(color=- thickness=1) ⮕ (result=)`  
_Creates a new line style._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `color` | `color.RGBA64` | `-` |   |   |   | The line color |
| `thickness` | `float64` | `1` |   |   | `1` | The line thickness |
| `⮕ result` | `error` |   |   |   |   | - - - A new line style |
---

### `P(x=0 y=0) ⮕ (result=)`  
_Creates a new point at P(x|y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `0` |   |   |   | The start position on the x-axis |
| `y` | `float64` | `0` |   |   |   | The start position on the y-axis |
| `⮕ result` | `error` |   |   |   |   | - - - A new point |
---

### `Px(p=-) ⮕ (result=)`  
_Returns the x-coordinate of a point._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `p` | `Point` | `-` |   |   |   | The point to return the x-coordinate of |
| `⮕ result` | `error` |   |   |   |   | - - - The x-coordinate of p |
---

### `Py(p=-) ⮕ (result=)`  
_Returns the y-coordinate of a point._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `p` | `Point` | `-` |   |   |   | The point to return the y-coordinate of |
| `⮕ result` | `error` |   |   |   |   | - - - The y-coordinate of p |
---

### `R(x=- y=- w=- h=-) ⮕ (result=)`  
_Creates a new rectangle with the given dimensions at P(x|y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The upper-left corner of the rectangle |
| `y` | `float64` | `-` |   |   |   | The upper-left corner of the rectangle |
| `w` | `float64` | `-` |   |   |   | The width of the rectangle |
| `h` | `float64` | `-` |   |   |   | The width of the rectangle |
| `⮕ result` | `error` |   |   |   |   | - - - A new rectangle |
---

### `Rh(r=-) ⮕ (result=)`  
_Returns the height of a rect._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `Rect` | `-` |   |   |   | The rect to return the height of |
| `⮕ result` | `error` |   |   |   |   | - - - The height of r |
---

### `Rw(r=-) ⮕ (result=)`  
_Returns the width of a rect._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `Rect` | `-` |   |   |   | The rect to return the width of |
| `⮕ result` | `error` |   |   |   |   | - - - The width of r |
---

### `Rx(r=-) ⮕ (result=)`  
_Returns the x-coordinate of a rect._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `Rect` | `-` |   |   |   | The rect to return the x-coordinate of |
| `⮕ result` | `error` |   |   |   |   | - - - The x-coordinate of r |
---

### `Ry(r=-) ⮕ (result=)`  
_Returns the y-coordinate of a rect._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `Rect` | `-` |   |   |   | The rect to return the y-coordinate of |
| `⮕ result` | `error` |   |   |   |   | - - - The y-coordinate of r |
---

### `SI(img=- r=-) ⮕ (result=)`  
_Copies an area from a source image and returns it as a new image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The source image |
| `r` | `Rect` | `-` |   |   |   | The selection to copy |
| `⮕ result` | `error` |   |   |   |   | - - - The new image |
---

### `T(style=- text="-") ⮕ (result=)`  
_Creates a new text._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `style` | `TextStyle` | `-` |   |   |   | The text style to use |
| `text` | `string` | `"-"` |   |   |   | - - The text to print |
| `⮕ result` | `error` |   |   |   |   | - - - A new text |
---

### `TS(color=- size=10 family="mono") ⮕ (result=)`  
_Creates a new font style._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `color` | `color.RGBA64` | `-` |   |   |   | The font color |
| `size` | `float64` | `10` |   |   | `1` | The font size |
| `family` | `string` | `"mono"` |   |   |   | The font family |
| `⮕ result` | `error` |   |   |   |   | - - - A new font style |
---

### `V(x=0 y=0 z=0) ⮕ (result=)`  
_Creates a new Vector from x, y and z._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `0` |   |   |   | The x-component |
| `y` | `float64` | `0` |   |   |   | The y-component |
| `z` | `float64` | `0` |   |   |   | The z-component |
| `⮕ result` | `error` |   |   |   |   | - - - A new vector |
---

### `Vx(v=-) ⮕ (result=)`  
_Returns the x-component of a vector._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `v` | `Vector` | `-` |   |   |   | The vector to return the x-component of |
| `⮕ result` | `error` |   |   |   |   | - - - The x-component of v |
---

### `Vy(v=-) ⮕ (result=)`  
_Returns the y-component of a vector._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `v` | `Vector` | `-` |   |   |   | The vector to return the y-component of |
| `⮕ result` | `error` |   |   |   |   | - - - The y-component of vp |
---

### `Vz(v=-) ⮕ (result=)`  
_Returns the z-component of a vector._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `v` | `Vector` | `-` |   |   |   | The vector to return the z-component of |
| `⮕ result` | `error` |   |   |   |   | - - - The z-component of vp |
---

### `abs(x=-) ⮕ (result=)`  
_Returns the absolute value of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The absolute value of x |
---

### `acos(radians=-) ⮕ (result=)`  
_calculates the arccosine (inverse cosine) of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - angle in radians between 0 and PI |
---

### `acosh(x=-) ⮕ (result=)`  
_calculates the inverse hyperbolic cosine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - inverse hyperbolic cosine value (x &gt;= 1) |
---

### `acoth(x=-) ⮕ (result=)`  
_calculates the inverse hyperbolic cotangent of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - inverse hyperbolic cotangent value (|x| &gt; 1) |
---

### `acsch(x=-) ⮕ (result=)`  
_calculates the inverse hyperbolic cosecant of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - inverse hyperbolic cosecant value (x != 0) |
---

### `add(a=- b=-) ⮕ (result=)`  
_Adds the two numbers_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `a` | `float64` | `-` |   |   |   | The first number |
| `b` | `float64` | `-` |   |   |   | The second number |
| `⮕ result` | `error` |   |   |   |   | - - - a&#43;b |
---

### `add-n(a=- b=- n=-) ⮕ (result=)`  
_Multiplies b by n and adds the result to a_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `a` | `float64` | `-` |   |   |   | The first number |
| `b` | `float64` | `-` |   |   |   | The second number |
| `n` | `float64` | `-` |   |   |   | The multiplier for the second number |
| `⮕ result` | `error` |   |   |   |   | - - - a &#43; (n * b) |
---

### `adjacent-of-triangle(hypotenuse=- opposite=- alpha=-) ⮕ (result=)`  
_Calculates adjacent side from hypotenuse, opposite and alpha angle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `alpha` | `float64` | `-` |   |   |   | The alpha angle |
| `⮕ result` | `error` |   |   |   |   | - - - The adjacent side length |
---

### `angle-between(x1=- y1=- x2=- y2=-) ⮕ (result=)`  
_Calculates angle between two points_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x1` | `float64` | `-` |   |   |   | The x coordinate of the first point |
| `y1` | `float64` | `-` |   |   |   | The y coordinate of the first point |
| `x2` | `float64` | `-` |   |   |   | The x coordinate of the second point |
| `y2` | `float64` | `-` |   |   |   | The y coordinate of the second point |
| `⮕ result` | `error` |   |   |   |   | - - - The angle in radians |
---

### `angle-difference(angle1=- angle2=-) ⮕ (result=)`  
_calculates the smallest difference between two angles_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `angle1` | `float64` | `-` |   |   |   | The first angle in radians |
| `angle2` | `float64` | `-` |   |   |   | The second angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - smallest angle difference in radians |
---

### `angle-difference-degrees(angle1=- angle2=-) ⮕ (result=)`  
_calculates the smallest difference between two angles in degrees_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `angle1` | `float64` | `-` |   |   |   | The first angle in degrees |
| `angle2` | `float64` | `-` |   |   |   | The second angle in degrees |
| `⮕ result` | `error` |   |   |   |   | - - - smallest angle difference in degrees |
---

### `asech(x=-) ⮕ (result=)`  
_calculates the inverse hyperbolic secant of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - inverse hyperbolic secant value (0 &lt; x &lt;= 1) |
---

### `asin(radians=-) ⮕ (result=)`  
_calculates the arcsine (inverse sine) of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - angle in radians between -PI/2 and PI/2 |
---

### `asinh(x=-) ⮕ (result=)`  
_calculates the inverse hyperbolic sine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - inverse hyperbolic sine value |
---

### `atan(radians=-) ⮕ (result=)`  
_calculates the arctangent (inverse tangent) of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - angle in radians between -PI/2 and PI/2 |
---

### `atanh(x=-) ⮕ (result=)`  
_calculates the inverse hyperbolic tangent of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - inverse hyperbolic tangent value (|x| &lt; 1) |
---

### `auto-contrast(img=- threshold=0.01 strength=1) ⮕ (result=)`  
_Automatically adjusts the contrast of an image by stretching the histogram to use the full range of values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-contrast |
| `threshold` | `float64` | `0.01` |   |   | `%` | The percentage of pixels to ignore at both ends of the histogram (0-0.5) |
| `strength` | `float64` | `1` | `0` | `1` |   | How strongly to apply the contrast adjustment (0 = no change, 1 = full correction) |
| `⮕ result` | `error` |   |   |   |   | - - - The contrast-adjusted image |
---

### `auto-levels(img=- lowPercentile=0.05 highPercentile=0.995 adjustAlpha=false) ⮕ (result=)`  
_Automatically adjusts the contrast and brightness of an image by stretching the histogram to use the full range of values, ignoring outliers using percentiles_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-level |
| `lowPercentile` | `float64` | `0.05` |   |   | `%` | The lower percentile to ignore (e.g., 0.5) |
| `highPercentile` | `float64` | `0.995` |   |   | `%` | The upper percentile to ignore (e.g., 99.5) |
| `adjustAlpha` | `bool` | `false` |   |   |   | Whether to adjust alpha channel (false = preserve original alpha) |
| `⮕ result` | `error` |   |   |   |   | - - - The auto-leveled image |
---

### `auto-tone(img=- levelsLow=0.005 levelsHigh=0.9995 whiteThresh=0.99 whiteStrength=0.8 contrastThresh=0.01 contrastStrength=0.8) ⮕ (result=)`  
_Automatically enhances the image by applying auto-levels, auto-white-balance, and auto-contrast in sequence_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-tone |
| `levelsLow` | `float64` | `0.005` |   |   | `%` | Lower percentile for auto-levels (e.g., 0.5) |
| `levelsHigh` | `float64` | `0.9995` |   |   | `%` | Upper percentile for auto-levels (e.g., 99.5) |
| `whiteThresh` | `float64` | `0.99` |   |   | `%` | Brightness threshold for auto-white-balance (0-1) |
| `whiteStrength` | `float64` | `0.8` | `0` | `1` |   | Strength for auto-white-balance (0-1) |
| `contrastThresh` | `float64` | `0.01` |   |   | `%` | Percentile for auto-contrast (0-0.5) |
| `contrastStrength` | `float64` | `0.8` | `0` | `1` |   | Strength for auto-contrast (0-1) |
| `⮕ result` | `error` |   |   |   |   | - - - The auto-toned image |
---

### `auto-white-balance(img=- threshold=0.95 strength=1) ⮕ (result=)`  
_Automatically adjusts the white balance of an image by finding bright areas and making them neutral_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-white-balance |
| `threshold` | `float64` | `0.95` |   |   | `%` | The brightness threshold to consider as white (0-1) |
| `strength` | `float64` | `1` | `0` | `1` |   | How strongly to apply the white balance (0 = no change, 1 = full correction) |
| `⮕ result` | `error` |   |   |   |   | - - - The white-balanced image |
---

### `blend(imgA=- imgB=- mode="normal") ⮕ (result=)`  
_Blends the two images using the given blendmode (defaults to normal)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `mode` | `string` | `"normal"` |   |   |   | The blendmode name |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-average(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the average blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-color(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the color blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-color-burn(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the color burn blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-color-dodge(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the color dodge blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-contrast-negate(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the contrast negate blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-darken(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the darken blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-darker-color(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the darker color blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-difference(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the difference blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-divide(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the divide blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-erase(imgA=- imgB=-) ⮕ (result=)`  
_Erases the bottom image wherever the top image is present (destination out)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-exclusion(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the exclusion blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-glow(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the glow blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-hard-light(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the hard light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-hard-mix(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the hard mix blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-hue(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the hue blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-lighten(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the lighten blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-lighter-color(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the lighter color blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-linear-light(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the linear light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-luminosity(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the luminosity blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-multiply(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the multiply blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-negation(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the negation blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-normal(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the normal blend mode (alpha compositing)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-overlay(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the overlay blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-pin-light(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the pin light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-reflect(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the reflect blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-saturation(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the saturation blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-screen(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the screen blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-soft-light(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the soft light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-subtract(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the subtract blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blend-vivid-light(imgA=- imgB=-) ⮕ (result=)`  
_Blends the two images using the vivid light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `error` |   |   |   |   | - - - The blended image |
---

### `blur-box(img=- radius=1) ⮕ (result=)`  
_Applies a box blur to an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `radius` | `int` | `1` | `1` | `10` |   | The blur radius (size of the box kernel) |
| `⮕ result` | `error` |   |   |   |   | - - - The blurred image |
---

### `blur-gaussian(img=- radius=1) ⮕ (result=)`  
_Applies a Gaussian blur to the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `radius` | `float64` | `1` | `1` | `10` |   | The blur radius (higher values create more blur) |
| `⮕ result` | `error` |   |   |   |   | - - - The blurred image |
---

### `blur-motion(img=- length=5 angle=0) ⮕ (result=)`  
_Applies a motion blur to an image along a specified angle._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `length` | `int` | `5` | `1` | `100` |   | The length of the motion blur (in pixels) |
| `angle` | `float64` | `0` | `0` | `360` |   | The angle of the motion blur (in degrees) |
| `⮕ result` | `error` |   |   |   |   | - - - The blurred image |
---

### `blur-zoom(img=- strength=0.25 centerX=0.5 centerY=0.5) ⮕ (result=)`  
_Applies a zoom blur effect to an image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `strength` | `float64` | `0.25` | `0` | `1` |   | The strength of the blur effect (higher means more blur) |
| `centerX` | `float64` | `0.5` | `0` | `1` |   | X coordinate of the blur center (default: image center) |
| `centerY` | `float64` | `0.5` | `0` | `1` |   | Y coordinate of the blur center (default: image center) |
| `⮕ result` | `error` |   |   |   |   | - - - The blurred image |
---

### `border(img=- style=-) ⮕ (result=)`  
_Draws a border around the image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw border around |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the border |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `box(img=- styleBorder=- styleFill=-) ⮕ (result=)`  
_Fills the image with the given background color and then draws a border around it._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw border around |
| `styleBorder` | `LineStyle` | `-` |   |   |   | The thickness and color of the border |
| `styleFill` | `FillStyle` | `-` |   |   |   | The color of the fill |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `brightness(img=- factor=0) ⮕ (result=)`  
_Changes the brightness of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to change brightness of |
| `factor` | `float64` | `0` | `0` | `2` |   | The change factor |
| `⮕ result` | `error` |   |   |   |   | - - - The image with brightness changed |
---

### `ceil(x=-) ⮕ (result=)`  
_Returns the smallest integer greater than or equal to x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The smallest integer greater than or equal to x |
---

### `chord(x=-) ⮕ (result=)`  
_calculates the chord of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - chord value (2 * sin(x/2)) |
---

### `chromatic-aberration(img=- amount=5) ⮕ (result=)`  
_Creates a chromatic aberration effect by offsetting color channels_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply the effect to |
| `amount` | `float64` | `5` | `0` | `20` |   | The amount of color channel separation |
| `⮕ result` | `error` |   |   |   |   | - - - The image with chromatic aberration |
---

### `circumference-of-a_circle(radius=-) ⮕ (result=)`  
_Calculates circumference from radius_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radius` | `float64` | `-` |   |   |   | The radius of the circle |
| `⮕ result` | `error` |   |   |   |   | - - - The circumference length |
---

### `clarity(img=- intensity=1 radius=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=)`  
_Enhances local contrast while preserving overall image structure_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to enhance |
| `intensity` | `float64` | `1` | `0` | `1` |   | The intensity of the clarity effect |
| `radius` | `float64` | `1` | `0.1` | `2` |   | The radius of the filter in pixels (higher values affect larger areas) |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel |
| `⮕ result` | `error` |   |   |   |   | - - - The enhanced image |
---

### `cmyk(c=0 m=0 y=0 k=0 alpha=1) ⮕ (result=)`  
_Creates a color from CMYK values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `c` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The cyan component |
| `m` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The magenta component |
| `y` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The yellow component |
| `k` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The key (black) component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The alpha component |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `color-balance(img=- rFactor=1 gFactor=1 bFactor=1) ⮕ (result=)`  
_Adjusts the balance of Red, Green, and Blue channels_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust color balance of |
| `rFactor` | `float64` | `1` | `0` | `2` |   | Red channel adjustment factor |
| `gFactor` | `float64` | `1` | `0` | `2` |   | Green channel adjustment factor |
| `bFactor` | `float64` | `1` | `0` | `2` |   | Blue channel adjustment factor |
| `⮕ result` | `error` |   |   |   |   | - - - The color-balanced image |
---

### `colorize(img=- col=-) ⮕ (result=)`  
_Colorizes the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to colorize |
| `col` | `color.RGBA64` | `-` |   |   |   | The color that determines the hue to use for colorization |
| `⮕ result` | `error` |   |   |   |   | - - - The colorized image |
---

### `contrast(img=- factor=1) ⮕ (result=)`  
_Adjusts the contrast of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust contrast of |
| `factor` | `float64` | `1` | `0` | `2` |   | The contrast factor (0 = gray, 1 = unchanged, 2 = maximum) |
| `⮕ result` | `error` |   |   |   |   | - - - The contrast-adjusted image |
---

### `cos(radians=-) ⮕ (result=)`  
_calculates the cosine of an angle in radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - cosine value between -1 and 1 |
---

### `cos-of-triangle(adjacent=- hypotenuse=-) ⮕ (result=)`  
_calculates cosine using adjacent and hypotenuse sides_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `⮕ result` | `error` |   |   |   |   | - - - cosine value (adjacent/hypotenuse) |
---

### `cos2(x=-) ⮕ (result=)`  
_calculates the square of cosine (cos²(x))_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - squared cosine value (cos(x)²) |
---

### `cosec(radians=-) ⮕ (result=)`  
_calculates the cosecant of an angle in radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - cosecant value (1/sin) |
---

### `cosec-of-triangle(hypotenuse=- opposite=-) ⮕ (result=)`  
_calculates cosecant using hypotenuse and opposite sides_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `⮕ result` | `error` |   |   |   |   | - - - cosecant value (hypotenuse/opposite) |
---

### `cosh(x=-) ⮕ (result=)`  
_calculates the hyperbolic cosine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - hyperbolic cosine value (always positive) |
---

### `cot(radians=-) ⮕ (result=)`  
_calculates the cotangent of an angle in radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - cotangent value (1/tan) |
---

### `cot-of-triangle(adjacent=- opposite=-) ⮕ (result=)`  
_calculates cotangent using adjacent and opposite sides_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `⮕ result` | `error` |   |   |   |   | - - - cotangent value (adjacent/opposite) |
---

### `coth(x=-) ⮕ (result=)`  
_calculates the hyperbolic cotangent of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - hyperbolic cotangent value (1/tanh) |
---

### `covercos(x=-) ⮕ (result=)`  
_calculates the coversed cosine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - coversed cosine value (1 &#43; sin(x)) |
---

### `coversin(x=-) ⮕ (result=)`  
_calculates the coversed sine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - coversed sine value (1 - sin(x)) |
---

### `crop(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=)`  
_Crops an image by specified percentages from each side_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `left` | `float64` | `0` |   |   |   | The percentage to crop from the left side (0-1) |
| `right` | `float64` | `0` |   |   |   | The percentage to crop from the right side (0-1) |
| `top` | `float64` | `0` |   |   |   | The percentage to crop from the top side (0-1) |
| `bottom` | `float64` | `0` |   |   |   | The percentage to crop from the bottom side (0-1) |
| `⮕ result` | `error` |   |   |   |   | - - - The cropped image |
---

### `crop-circle(img=- radius=1 offsetX=0 offsetY=0) ⮕ (result=)`  
_Crops an image using a circular mask. The circle is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the radius is a percentage (0-1) of half the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `radius` | `float64` | `1` | `0` | `1` |   | Radius as a percentage of half the min(width, height) |
| `offsetX` | `float64` | `0` | `-1` | `1` |   | Horizontal offset from image center (percentage of width, -1..1) |
| `offsetY` | `float64` | `0` | `-1` | `1` |   | Vertical offset from image center (percentage of height, -1..1) |
| `⮕ result` | `error` |   |   |   |   | - - - The circularly cropped image (pixels outside the circle are transparent) |
---

### `crop-circle-px(img=- radius=1 offsetX=0 offsetY=0) ⮕ (result=)`  
_Crops an image using a circular mask. The circle is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the radius is a percentage (0-1) of half the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `radius` | `float64` | `1` | `0` | `1` |   | Radius as a percentage of half the min(width, height) |
| `offsetX` | `int` | `0` |   |   |   | Horizontal offset from image center (pixels) |
| `offsetY` | `int` | `0` |   |   |   | Vertical offset from image center (pixels) |
| `⮕ result` | `error` |   |   |   |   | - - - The circularly cropped image (pixels outside the circle are transparent) |
---

### `crop-px(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=)`  
_Crops an image by specified amounts of pixels from each side_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `left` | `int` | `0` |   |   |   | The number of pixels to crop from the left side |
| `right` | `int` | `0` |   |   |   | The number of pixels to crop from the right side |
| `top` | `int` | `0` |   |   |   | The number of pixels to crop from the top side |
| `bottom` | `int` | `0` |   |   |   | The number of pixels to crop from the bottom side |
| `⮕ result` | `error` |   |   |   |   | - - - The cropped image |
---

### `crop-square(img=- size=1 offsetX=0 offsetY=0) ⮕ (result=)`  
_Crops an image using a square mask. The square is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the size is a percentage (0-1) of the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `size` | `float64` | `1` | `0` | `1` |   | Size as a percentage of the min(width, height) |
| `offsetX` | `float64` | `0` | `-1` | `1` |   | Horizontal offset from image center (percentage of width, -1..1) |
| `offsetY` | `float64` | `0` | `-1` | `1` |   | Vertical offset from image center (percentage of height, -1..1) |
| `⮕ result` | `error` |   |   |   |   | - - - The square-cropped image (pixels outside the square are transparent) |
---

### `crop-square-px(img=- size=1 offsetX=0 offsetY=0) ⮕ (result=)`  
_Crops an image using a square mask. The square is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the size is a percentage (0-1) of the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `size` | `float64` | `1` | `0` | `1` |   | Size as a percentage of the min(width, height) |
| `offsetX` | `int` | `0` |   |   |   | Horizontal offset from image center (pixels) |
| `offsetY` | `int` | `0` |   |   |   | Vertical offset from image center (pixels) |
| `⮕ result` | `error` |   |   |   |   | - - - The square-cropped image (pixels outside the square are transparent) |
---

### `csch(x=-) ⮕ (result=)`  
_calculates the hyperbolic cosecant of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - hyperbolic cosecant value (1/sinh) |
---

### `defisheye(img=- strength=1) ⮕ (result=)`  
_Corrects fisheye lens distortion in an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to correct |
| `strength` | `float64` | `1` | `0` | `2` |   | The strength of the correction |
| `⮕ result` | `error` |   |   |   |   | - - - The corrected image |
---

### `degrees2radians(degrees=-) ⮕ (result=)`  
_converts degrees to radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `degrees` | `float64` | `-` |   |   |   | The angle in degrees |
| `⮕ result` | `error` |   |   |   |   | - - - angle in radians |
---

### `delta(x=- y=-) ⮕ (result=)`  
_Returns the delta between x and y_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The x value |
| `y` | `float64` | `-` |   |   |   | The y value |
| `⮕ result` | `error` |   |   |   |   | - - - The delta between x and y |
---

### `displace(img=- dMap=- amount=10) ⮕ (result=)`  
_Displaces pixels based on the brightness of a displacement map_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to displace |
| `dMap` | `*image.NRGBA64` | `-` |   |   |   | The displacement map image |
| `amount` | `float64` | `10` | `0` | `50` |   | The amount of displacement |
| `⮕ result` | `error` |   |   |   |   | - - - The displaced image |
---

### `distance-between(x1=- y1=- x2=- y2=-) ⮕ (result=)`  
_Calculates distance between two points_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x1` | `float64` | `-` |   |   |   | The x coordinate of the first point |
| `y1` | `float64` | `-` |   |   |   | The y coordinate of the first point |
| `x2` | `float64` | `-` |   |   |   | The x coordinate of the second point |
| `y2` | `float64` | `-` |   |   |   | The y coordinate of the second point |
| `⮕ result` | `error` |   |   |   |   | - - - The distance between the points |
---

### `div(a=- b=-) ⮕ (result=)`  
_Divides the two numbers_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `a` | `float64` | `-` |   |   |   | The first number |
| `b` | `float64` | `-` |   |   |   | The second number |
| `⮕ result` | `error` |   |   |   |   | - - - a/b |
---

### `draw-circle(img=- c=- style=-) ⮕ (result=)`  
_Draws a circle._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `c` | `Ellipse` | `-` |   |   |   | The circle to draw |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-circle-px(img=- c=- style=-) ⮕ (result=)`  
_Draws a circle._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `c` | `Ellipse` | `-` |   |   |   | The circle to draw |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-ellipse(img=- e=- style=-) ⮕ (result=)`  
_Draws an ellipse._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `e` | `Ellipse` | `-` |   |   |   | The ellipse to draw |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-ellipse-px(img=- e=- style=-) ⮕ (result=)`  
_Draws an ellipse._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `e` | `Ellipse` | `-` |   |   |   | The ellipse to draw |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-grid(img=- r=- rows=- cols=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The area to draw the grid on (relative) |
| `rows` | `int` | `-` |   |   |   | The number of rows |
| `cols` | `int` | `-` |   |   |   | The number of cols |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-grid-h(img=- r=- rows=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The area to draw the grid on (relative) |
| `rows` | `int` | `-` |   |   |   | The number of rows |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-grid-h-px(img=- r=- rows=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The area to draw the grid on |
| `rows` | `int` | `-` |   |   |   | The number of rows |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-grid-px(img=- r=- rows=- cols=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The area to draw the grid on |
| `rows` | `int` | `-` |   |   |   | The number of rows |
| `cols` | `int` | `-` |   |   |   | The number of cols |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-grid-v(img=- r=- cols=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The area to draw the grid on (relative) |
| `cols` | `int` | `-` |   |   |   | The number of cols |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-grid-v-px(img=- r=- cols=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The area to draw the grid on |
| `cols` | `int` | `-` |   |   |   | The number of cols |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-line(img=- p1=- p2=- style=-) ⮕ (result=)`  
_Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `p1` | `Point` | `-` |   |   |   | The start position (relative) |
| `p2` | `Point` | `-` |   |   |   | The end position (relative) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-line-h(img=- y=- x1=- x2=- style=-) ⮕ (result=)`  
_Draws a line from P(x1|y) to P(x2|y) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `y` | `float64` | `-` |   |   |   | The position on the y-axis (relative) |
| `x1` | `float64` | `-` |   |   |   | The start position on the x-axis (relative) |
| `x2` | `float64` | `-` |   |   |   | The end position on the x-axis (relative) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-line-h-px(img=- y=- x1=- x2=- style=-) ⮕ (result=)`  
_Draws a line from P(x1|y) to P(x2|y) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `y` | `float64` | `-` |   |   |   | The position on the y-axis |
| `x1` | `float64` | `-` |   |   |   | The start position on the x-axis |
| `x2` | `float64` | `-` |   |   |   | The end position on the x-axis |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-line-px(img=- p1=- p2=- style=-) ⮕ (result=)`  
_Draws a line from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `p1` | `Point` | `-` |   |   |   | The start position |
| `p2` | `Point` | `-` |   |   |   | The end position |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-line-v(img=- x=- y1=- y2=- style=-) ⮕ (result=)`  
_Draws a line from P(x|y1) to P(x|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `x` | `float64` | `-` |   |   |   | The position on the x-axis (relative) |
| `y1` | `float64` | `-` |   |   |   | The start position on the y-axis (relative) |
| `y2` | `float64` | `-` |   |   |   | The end position on the y-axis (relative) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-line-v-px(img=- x=- y1=- y2=- style=-) ⮕ (result=)`  
_Draws a line from P(x|y1) to P(x|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `x` | `float64` | `-` |   |   |   | The position on the x-axis |
| `y1` | `float64` | `-` |   |   |   | The start position on the y-axis |
| `y2` | `float64` | `-` |   |   |   | The end position on the y-axis |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-rect(img=- r=- style=-) ⮕ (result=)`  
_Draws a rectangle at position (x,y) with the given width and height._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The rectangle to draw (relative) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-rect-px(img=- r=- style=-) ⮕ (result=)`  
_Draws a rectangle at position (x,y) with the given width and height._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `r` | `Rect` | `-` |   |   |   | The rectangle to draw (absolute) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-square(img=- s=- style=-) ⮕ (result=)`  
_Draws a square at position (x,y) with the given size._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `s` | `Rect` | `-` |   |   |   | The square to draw (relative) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-square-px(img=- s=- style=-) ⮕ (result=)`  
_Draws a square at position (x,y) with the given size._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `s` | `Rect` | `-` |   |   |   | The square to draw (absolute) |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the line |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-text(img=- p=- t=-) ⮕ (result=)`  
_Draws a text at position (x,y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `p` | `Point` | `-` |   |   |   | The upper-left coordinate of the text |
| `t` | `Text` | `-` |   |   |   | The text to draw |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-text-outline(img=- p=- t=- outline=-) ⮕ (result=)`  
_Draws only the outline of text at position (x,y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `p` | `Point` | `-` |   |   |   | The upper-left coordinate of the text |
| `t` | `Text` | `-` |   |   |   | The text to outline |
| `outline` | `LineStyle` | `-` |   |   |   | The outline style (thickness and color) |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-text-outline-px(img=- p=- t=- outline=-) ⮕ (result=)`  
_Draws only the outline of text at position (x,y)._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `p` | `Point` | `-` |   |   |   | The upper-left coordinate of the text |
| `t` | `Text` | `-` |   |   |   | The text to outline |
| `outline` | `LineStyle` | `-` |   |   |   | The outline style (thickness and color) |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `draw-text-px(img=- p=- t=-) ⮕ (result=)`  
_Draws text at position (x,y) with the given style using TrueType fonts._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `p` | `Point` | `-` |   |   |   | The upper-left coordinate of the text |
| `t` | `Text` | `-` |   |   |   | The text to draw |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `edge-detect(img=-) ⮕ (result=)`  
_Detects edges in the image using the Sobel operator_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to detect edges in |
| `⮕ result` | `error` |   |   |   |   | - - - An image highlighting the edges |
---

### `enhance(img=- brightness=0 contrast=0 sharpening=1 rMin=0.75 rMax=1 gMin=0.75 gMax=1 bMin=0.75 bMax=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=)`  
_Enhances colors and sharpness of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to enhance |
| `brightness` | `float64` | `0` | `-1` | `1` |   | The brightness adjustment of the image |
| `contrast` | `float64` | `0` | `-1` | `1` |   | The contrast adjustment of the image |
| `sharpening` | `float64` | `1` | `0` | `5` |   | The sharpening intensity in pixels (higher values detect larger edges) |
| `rMin` | `float64` | `0.75` | `0` | `1` |   | The minimum intensity of the red channel |
| `rMax` | `float64` | `1` | `0` | `1` |   | The maximum intensity of the red channel |
| `gMin` | `float64` | `0.75` | `0` | `1` |   | The minimum intensity of the green channel |
| `gMax` | `float64` | `1` | `0` | `1` |   | The maximum intensity of the green channel |
| `bMin` | `float64` | `0.75` | `0` | `1` |   | The minimum intensity of the blue channel |
| `bMax` | `float64` | `1` | `0` | `1` |   | The maximum intensity of the blue channel |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel (used for sharpening) |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel (used for sharpening) |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel (used for sharpening) |
| `⮕ result` | `error` |   |   |   |   | - - - The enhanceed image |
---

### `excsc(x=-) ⮕ (result=)`  
_calculates the excosecant of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - excosecant value (cosec(x) - 1) |
---

### `expand(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=)`  
_Expands an image by adding transparent borders with specified percentage widths_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to expand |
| `left` | `float64` | `0` |   |   |   | The percentage to add to the left side (relative to original width) |
| `right` | `float64` | `0` |   |   |   | The percentage to add to the right side (relative to original width) |
| `top` | `float64` | `0` |   |   |   | The percentage to add to the top side (relative to original height) |
| `bottom` | `float64` | `0` |   |   |   | The percentage to add to the bottom side (relative to original height) |
| `⮕ result` | `error` |   |   |   |   | - - - The expanded image |
---

### `expand-px(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=)`  
_Expands an image by adding transparent borders with specified pixel widths_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to expand |
| `left` | `int` | `0` |   |   |   | The number of pixels to add to the left side |
| `right` | `int` | `0` |   |   |   | The number of pixels to add to the right side |
| `top` | `int` | `0` |   |   |   | The number of pixels to add to the top side |
| `bottom` | `int` | `0` |   |   |   | The number of pixels to add to the bottom side |
| `⮕ result` | `error` |   |   |   |   | - - - The expanded image |
---

### `exposure(img=- level=0) ⮕ (result=)`  
_Adjusts the overall lightness or darkness of the image, simulating photographic exposure._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust exposure of |
| `level` | `float64` | `0` | `-2` | `2` |   | The exposure level adjustment (-2 = much darker, 0 = unchanged, 2 = much brighter) |
| `⮕ result` | `error` |   |   |   |   | - - - The exposure-adjusted image |
---

### `exsec(x=-) ⮕ (result=)`  
_calculates the exsecant of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - exsecant value (sec(x) - 1) |
---

### `fibonacci(nth=-) ⮕ (result=)`  
_Calculates the nth fibonacci number using 1-based indexing with memoization_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `nth` | `float64` | `-` |   |   |   | The nth fibonacci number to calculate |
| `⮕ result` | `error` |   |   |   |   | - - - The nth fibonacci number |
---

### `fill(img=- style=-) ⮕ (result=)`  
_Fills the given image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to fill |
| `style` | `FillStyle` | `-` |   |   |   | The color of the fill |
| `⮕ result` | `error` |   |   |   |   | - - - The filled image |
---

### `first-csv-row(data=-) ⮕ (result=)`  
_Returns the first row of CSV data_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `data` | `[][]float64` | `-` |   |   |   | Data to return first row of |
| `⮕ result` | `error` |   |   |   |   | - - - A slice with the data |
---

### `first-csv-rows(data=- n=1) ⮕ (result=)`  
_Returns the first `n` rows of CSV data_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `data` | `[][]float64` | `-` |   |   |   | Data to return first `n` rows of |
| `n` | `int` | `1` |   |   | `1` | Number rows to return |
| `⮕ result` | `error` |   |   |   |   | - - - A slice with the data |
---

### `fisheye(img=- strength=1) ⮕ (result=)`  
_Applies a fisheye lens distortion effect to the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to distort |
| `strength` | `float64` | `1` | `0` | `2` |   | The strength of the fisheye effect |
| `⮕ result` | `error` |   |   |   |   | - - - The distorted image |
---

### `flip-h(img=-) ⮕ (result=)`  
_Flips an image horizontally (left to right)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to flip horizontally |
| `⮕ result` | `error` |   |   |   |   | - - - The horizontally flipped image |
---

### `flip-v(img=-) ⮕ (result=)`  
_Flips an image vertically (top to bottom)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to flip vertically |
| `⮕ result` | `error` |   |   |   |   | - - - The vertically flipped image |
---

### `floor(x=-) ⮕ (result=)`  
_Returns the largest integer less than or equal to x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The largest integer less than or equal to x |
---

### `grads2radians(grads=-) ⮕ (result=)`  
_converts grads to radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `grads` | `float64` | `-` |   |   |   | The angle in grads |
| `⮕ result` | `error` |   |   |   |   | - - - angle in radians |
---

### `grayscale(img=-) ⮕ (result=)`  
_Grayscales an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to grayscale |
| `⮕ result` | `error` |   |   |   |   | - - - The grayscaled image |
---

### `grid(img=- rows=- cols=- style=-) ⮕ (result=)`  
_Draws a grid from P(x1|y1) to P(x2|y2) with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `rows` | `int` | `-` |   |   |   | The number of rows |
| `cols` | `int` | `-` |   |   |   | The number of cols |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the border |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `grid-h(img=- rows=- style=-) ⮕ (result=)`  
_Draws horizontal grid lines on the image with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `rows` | `int` | `-` |   |   |   | The number of rows |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the border |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `grid-v(img=- cols=- style=-) ⮕ (result=)`  
_Draws vertical grid lines on the image with the given thickness and color._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to draw to |
| `cols` | `int` | `-` |   |   |   | The number of columns |
| `style` | `LineStyle` | `-` |   |   |   | The thickness and color of the border |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `group(img=- title="-" colText=- colFill=- colBorder=- padding=3 fillAlphaHeader=0.9 fillAlphaBody=0.8 borderThickness=1 borderAlphaHeader=0.95 borderAlphaBody=0.95) ⮕ (result=)`  
_Generates the given group with the given styles._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to wrap in a group |
| `title` | `string` | `"-"` |   |   |   | - - The title of the group |
| `colText` | `color.RGBA64` | `-` |   |   |   | The color of the title |
| `colFill` | `color.RGBA64` | `-` |   |   |   | The color of the group |
| `colBorder` | `color.RGBA64` | `-` |   |   |   | The color of the border |
| `padding` | `float64` | `3` |   |   | `0` | The padding for the image to wrap |
| `fillAlphaHeader` | `float64` | `0.9` |   |   | `0` | The alpha of header fill |
| `fillAlphaBody` | `float64` | `0.8` |   |   | `0` | The alpha of body fill |
| `borderThickness` | `float64` | `1` |   |   | `0` | The thickness of the border |
| `borderAlphaHeader` | `float64` | `0.95` |   |   | `0` | The alpha of header border |
| `borderAlphaBody` | `float64` | `0.95` |   |   | `0` | The alpha of body border |
| `⮕ result` | `error` |   |   |   |   | - - - The group wrapping the input image |
---

### `haversin(x=-) ⮕ (result=)`  
_calculates the haversine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - haversine value (0.5 * (1 - cos(x))) |
---

### `highpass(img=- radius=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=)`  
_Creates a high pass filter effect, resulting in a gray image with embossed edges_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply the high-pass filter to |
| `radius` | `float64` | `1` | `0.1` | `2` |   | The radius of the filter in pixels (higher values detect larger edges) |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel |
| `⮕ result` | `error` |   |   |   |   | - - - The filtered image |
---

### `hsla(h=0 s=0.5 l=0.5 alpha=1) ⮕ (result=)`  
_Creates a color from HSLA values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The color&#39;s hue |
| `s` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s saturation |
| `l` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s luminosity |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `hsv(h=0 s=0.5 v=0.5 alpha=1) ⮕ (result=)`  
_Creates a color from HSV values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The color&#39;s hue |
| `s` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s saturation |
| `v` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s value (brightness) |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `hue-rotate(img=- angle=0) ⮕ (result=)`  
_Rotates the hue of image colors_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to rotate hue of |
| `angle` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The angle in degrees (0-360) |
| `⮕ result` | `error` |   |   |   |   | - - - The hue-rotated image |
---

### `hwb(h=0 w=0 b=0 alpha=1) ⮕ (result=)`  
_Creates a color from HWB (Hue, Whiteness, Blackness) values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The color&#39;s hue |
| `w` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The whiteness component |
| `b` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The blackness component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `hypotenuse-of-triangle(adjacent=- opposite=- gamma=-) ⮕ (result=)`  
_Calculates hypotenuse from adjacent, opposite and gamma angle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `gamma` | `float64` | `-` |   |   |   | The gamma angle |
| `⮕ result` | `error` |   |   |   |   | - - - The hypotenuse length |
---

### `invert(img=-) ⮕ (result=)`  
_Inverts an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to invert |
| `⮕ result` | `error` |   |   |   |   | - - - The inverted image |
---

### `invert-hsl(img=- lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=)`  
_Inverts pixels based on their hue, saturation, and luminance, inverting pixels inside the specified ranges_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerHue` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The lower hue threshold (below this becomes transparent) |
| `minHue` | `float64` | `30` | `0` | `360` | `&#34;°&#34;` | The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this) |
| `maxHue` | `float64` | `330` | `0` | `360` | `&#34;°&#34;` | The maximum hue for full opacity (fade from 100% to 0% between this and upperHue) |
| `upperHue` | `float64` | `360` | `0` | `360` | `&#34;°&#34;` | The upper hue threshold (above this becomes transparent) |
| `lowerSat` | `float64` | `0.1` | `0` | `1` |   | The lower saturation threshold (below this becomes transparent) |
| `minSat` | `float64` | `0.2` | `0` | `1` |   | The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this) |
| `maxSat` | `float64` | `0.8` | `0` | `1` |   | The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat) |
| `upperSat` | `float64` | `0.9` | `0` | `1` |   | The upper saturation threshold (above this becomes transparent) |
| `lowerLum` | `float64` | `0.1` | `0` | `1` |   | The lower luminance threshold (below this becomes transparent) |
| `minLum` | `float64` | `0.2` | `0` | `1` |   | The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this) |
| `maxLum` | `float64` | `0.8` | `0` | `1` |   | The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum) |
| `upperLum` | `float64` | `0.9` | `0` | `1` |   | The upper luminance threshold (above this becomes transparent) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with only pixels in all specified ranges visible (16-bit) |
---

### `lab(l=50 a=0 b=0 alpha=1) ⮕ (result=)`  
_Creates a color from CIELAB values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `l` | `float64` | `50` | `0` | `100` | `&#34;&#34;` | The lightness component |
| `a` | `float64` | `0` | `-128` | `127` | `&#34;&#34;` | The green-red component |
| `b` | `float64` | `0` | `-128` | `127` | `&#34;&#34;` | The blue-yellow component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `last-csv-row(data=-) ⮕ (result=)`  
_Returns the last row of CSV data_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `data` | `[][]float64` | `-` |   |   |   | Data to return last row of |
| `⮕ result` | `error` |   |   |   |   | - - - A slice with the data |
---

### `last-csv-rows(data=- n=1) ⮕ (result=)`  
_Returns the last `n` rows of CSV data_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `data` | `[][]float64` | `-` |   |   |   | Data to return last `n` rows of |
| `n` | `int` | `1` |   |   | `1` | Number rows to return |
| `⮕ result` | `error` |   |   |   |   | - - - A slice with the data |
---

### `lch(l=50 c=0 h=0 alpha=1) ⮕ (result=)`  
_Creates a color from LCH (Lightness, Chroma, Hue) values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `l` | `float64` | `50` | `0` | `100` | `&#34;&#34;` | The lightness component |
| `c` | `float64` | `0` | `0` | `128` | `&#34;&#34;` | The chroma component |
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The hue component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `len(v=-) ⮕ (result=)`  
_Returns the length of the given value._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `v` | `any` | `-` |   |   |   | The value to get the length of |
| `⮕ result` | `error` |   |   |   |   | - - - The length of v |
---

### `lerp-angle(angle1=- angle2=- t=-) ⮕ (result=)`  
_linearly interpolates between two angles in radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `angle1` | `float64` | `-` |   |   |   | The first angle in radians |
| `angle2` | `float64` | `-` |   |   |   | The second angle in radians |
| `t` | `float64` | `-` |   |   |   | The interpolation factor (0-1) |
| `⮕ result` | `error` |   |   |   |   | - - - interpolated angle in radians |
---

### `lerp-angle-degrees(angle1=- angle2=- t=-) ⮕ (result=)`  
_linearly interpolates between two angles in degrees_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `angle1` | `float64` | `-` |   |   |   | The first angle in degrees |
| `angle2` | `float64` | `-` |   |   |   | The second angle in degrees |
| `t` | `float64` | `-` |   |   |   | The interpolation factor (0-1) |
| `⮕ result` | `error` |   |   |   |   | - - - interpolated angle in degrees |
---

### `load(path="-") ⮕ (result=)`  
_Loads an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `path` | `string` | `"-"` |   |   |   | - - Path to the image |
| `⮕ result` | `error` |   |   |   |   | - - - The loaded image |
---

### `load-csv(path="-" sep="\t" hasHeader=true) ⮕ (result=)`  
_Loads a CSV file_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `path` | `string` | `"-"` |   |   |   | - - Path to the CSV file |
| `sep` | `string` | `"\t"` |   |   |   | The separator to split columns with |
| `hasHeader` | `bool` | `true` |   |   |   | Whether the first row is a header row |
| `⮕ result` | `error` |   |   |   |   | - - - A 2D slice with the data |
---

### `load-csv-column(path="-" index=- sep="\t" hasHeader=true) ⮕ (result=)`  
_Loads a column from a CSV file_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `path` | `string` | `"-"` |   |   |   | - - Path to the CSV file |
| `index` | `int` | `-` |   |   |   | The index of the column to retrieve |
| `sep` | `string` | `"\t"` |   |   |   | The separator to split columns with |
| `hasHeader` | `bool` | `true` |   |   |   | Whether the first row is a header row |
| `⮕ result` | `error` |   |   |   |   | - - - A slice with the data |
---

### `load-csv-row(path="-" index=- sep="\t" hasHeader=true) ⮕ (result=)`  
_Loads a row from a CSV file_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `path` | `string` | `"-"` |   |   |   | - - Path to the CSV file |
| `index` | `int` | `-` |   |   |   | The index of the row to retrieve |
| `sep` | `string` | `"\t"` |   |   |   | The separator to split columns with |
| `hasHeader` | `bool` | `true` |   |   |   | Whether the first row is a header row |
| `⮕ result` | `error` |   |   |   |   | - - - A slice with the data |
---

### `log(x=-) ⮕ (result=)`  
_Returns the natural logarithm of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The natural logarithm of x |
---

### `lowercase(str="-") ⮕ (result=)`  
_Lowercases a string_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `str` | `string` | `"-"` |   |   |   | - - The string to lowercase |
| `⮕ result` | `error` |   |   |   |   | - - - The lowercased string |
---

### `map-color(value=0 min=0 max=1 stops=-) ⮕ (result=)`  
_Maps a value to a color using color stops with HSLA interpolation_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `value` | `float64` | `0` |   |   | `&#34;&#34;` | The value to map to a color |
| `min` | `float64` | `0` |   |   | `&#34;&#34;` | Minimum value of the range |
| `max` | `float64` | `1` |   |   | `&#34;&#34;` | Maximum value of the range |
| `stops` | `[][]any` | `-` |   |   | `&#34;&#34;` | Color stops as [][]any where each stop is [threshold, hue, saturation, lightness, alpha]; threshold is a raw value between min and max |
| `⮕ result` | `error` |   |   |   |   | - - - The interpolated color |
---

### `max(x=- y=-) ⮕ (result=)`  
_Returns the maximum value of x and y_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The x value |
| `y` | `float64` | `-` |   |   |   | The y value |
| `⮕ result` | `error` |   |   |   |   | - - - The maximum value of x and y |
---

### `min(x=- y=-) ⮕ (result=)`  
_Returns the minimum value of x and y_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The x value |
| `y` | `float64` | `-` |   |   |   | The y value |
| `⮕ result` | `error` |   |   |   |   | - - - The minimum value of x and y |
---

### `mul(a=- b=-) ⮕ (result=)`  
_Multiplies the two numbers_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `a` | `float64` | `-` |   |   |   | The first number |
| `b` | `float64` | `-` |   |   |   | The second number |
| `⮕ result` | `error` |   |   |   |   | - - - a*b |
---

### `normalize-angle(radians=-) ⮕ (result=)`  
_normalizes an angle to [0, 2π)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - normalized angle in radians |
---

### `normalize-angle-degrees(degrees=-) ⮕ (result=)`  
_normalizes an angle to [0, 360)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `degrees` | `float64` | `-` |   |   |   | The angle in degrees |
| `⮕ result` | `error` |   |   |   |   | - - - normalized angle in degrees |
---

### `opacity(img=- amount=1) ⮕ (result=)`  
_Adjusts the overall opacity/transparency of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust opacity of |
| `amount` | `float64` | `1` | `0` | `1` |   | The opacity amount (0 = fully transparent, 1 = unchanged) |
| `⮕ result` | `error` |   |   |   |   | - - - The opacity-adjusted image |
---

### `opposite-of-triangle(hypotenuse=- adjacent=- beta=-) ⮕ (result=)`  
_Calculates opposite side from hypotenuse, adjacent and beta angle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `beta` | `float64` | `-` |   |   |   | The beta angle |
| `⮕ result` | `error` |   |   |   |   | - - - The opposite side length |
---

### `or(value1=- value2=-) ⮕ (result=)`  
_Returns either value1 or value2 randomly_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `value1` | `float64` | `-` |   |   |   | The first value |
| `value2` | `float64` | `-` |   |   |   | The second value |
| `⮕ result` | `error` |   |   |   |   | - - - One of the two input values randomly |
---

### `pixelate(img=- size=8) ⮕ (result=)`  
_Creates a pixelation effect by averaging colors in blocks_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to pixelate |
| `size` | `int` | `8` | `1` | `50` |   | The size of the pixel blocks |
| `⮕ result` | `error` |   |   |   |   | - - - The pixelated image |
---

### `plot-data(width=- height=- data=- columns=- colors=-) ⮕ (result=)`  
_Renders a chart from CSV data by plotting selected columns with specified colors_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `width` | `int` | `-` |   |   |   | Chart width in pixels |
| `height` | `int` | `-` |   |   |   | Chart height in pixels |
| `data` | `[][]float64` | `-` |   |   |   | 2D array of data from CSV |
| `columns` | `[]any` | `-` |   |   |   | Array of column indices to plot |
| `colors` | `[]any` | `-` |   |   |   | Array of colors for each column |
| `⮕ result` | `error` |   |   |   |   | - - - The chart image |
---

### `plot-data-compact(width=- height=- data=- columns=- colors=-) ⮕ (result=)`  
_Renders a chart from CSV data by plotting selected columns with specified colors. The series will be normalized to 0..1 based on their minium and maximimum value._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `width` | `int` | `-` |   |   |   | Chart width in pixels |
| `height` | `int` | `-` |   |   |   | Chart height in pixels |
| `data` | `[][]float64` | `-` |   |   |   | 2D array of data from CSV |
| `columns` | `[]any` | `-` |   |   |   | Array of column indices to plot |
| `colors` | `[]any` | `-` |   |   |   | Array of colors for each column |
| `⮕ result` | `error` |   |   |   |   | - - - The chart image |
---

### `plot-series(width=- height=- data=- column=- min=- max=- stops=- invertY=false) ⮕ (result=)`  
_Renders a series from CSV data by plotting a single column with colors determined by value using color stops_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `width` | `int` | `-` |   |   |   | Chart width in pixels |
| `height` | `int` | `-` |   |   |   | Chart height in pixels |
| `data` | `[][]float64` | `-` |   |   |   | 2D array of data from CSV |
| `column` | `int` | `-` |   |   |   | Column index to plot |
| `min` | `float64` | `-` |   |   |   | Minimum value for color mapping |
| `max` | `float64` | `-` |   |   |   | Maximum value for color mapping |
| `stops` | `[][]any` | `-` |   |   |   | Color stops as [][]any where each stop is [threshold, hue, saturation, lightness, alpha]; additional fields are ignored |
| `invertY` | `bool` | `false` |   |   |   | Whether flip the y-axis when plotting |
| `⮕ result` | `error` |   |   |   |   | - - - The chart image |
---

### `polar-to-rectangular(img=-) ⮕ (result=)`  
_Converts a polar coordinate image to rectangular coordinates_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to transform |
| `⮕ result` | `error` |   |   |   |   | - - - The transformed image |
---

### `posterize(img=- levels=4) ⮕ (result=)`  
_Reduces the number of color levels in the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to posterize |
| `levels` | `int` | `4` | `2` | `16` |   | Number of color levels per channel (2-16) |
| `⮕ result` | `error` |   |   |   |   | - - - The posterized image |
---

### `pow(base=- n=-) ⮕ (result=)`  
_Calculates base raised to the power of n, using lookup tables for integer bases when possible_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `base` | `float64` | `-` |   |   |   | The base value |
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - Base^n |
---

### `pow10(n=-) ⮕ (result=)`  
_Calculates 10 raised to the power of n_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - 10^n |
---

### `pow12(n=-) ⮕ (result=)`  
_Calculates 12 raised to the power of n_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - 12^n |
---

### `pow16(n=-) ⮕ (result=)`  
_Calculates 16 raised to the power of n_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - 16^n |
---

### `pow2(n=-) ⮕ (result=)`  
_Calculates 2 raised to the power of n_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - 2^n |
---

### `pow4(n=-) ⮕ (result=)`  
_Calculates 4 raised to the power of n_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - 4^n |
---

### `pow8(n=-) ⮕ (result=)`  
_Calculates 8 raised to the power of n_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `n` | `float64` | `-` |   |   |   | The exponent |
| `⮕ result` | `error` |   |   |   |   | - - - 8^n |
---

### `printf(str="-" a=-) ⮕ (result=)`  
_Prints formatted strings to the console_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `str` | `string` | `"-"` |   |   |   | - - The format string |
| `a` | `[]any` | `-` |   |   |   | A slice with the arguments |
| `⮕ result` | `error` |   |   |   |   | - - - Number of bytes written |
---

### `printfln(str="-" a=-) ⮕ (result=)`  
_Prints formatted strings to the console_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `str` | `string` | `"-"` |   |   |   | - - The format string |
| `a` | `[]any` | `-` |   |   |   | A slice with the arguments |
| `⮕ result` | `error` |   |   |   |   | - - - Number of bytes written |
---

### `radians-of-triangle(adjacent=- opposite=- hypotenuse=-) ⮕ (result=)`  
_calculates angle in radians using all three sides of a triangle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `⮕ result` | `error` |   |   |   |   | - - - angle in radians between adjacent and opposite sides |
---

### `radians2degrees(radians=-) ⮕ (result=)`  
_converts radians to degrees_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - angle in degrees |
---

### `radians2grads(radians=-) ⮕ (result=)`  
_converts radians to grads_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - angle in grads |
---

### `random-range(min=- max=-) ⮕ (result=)`  
_Returns a random number between min and max_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `min` | `float64` | `-` |   |   |   | The minimum value |
| `max` | `float64` | `-` |   |   |   | The maximum value |
| `⮕ result` | `error` |   |   |   |   | - - - A random float64 value between min and max, with NaN handling |
---

### `rectangular-to-polar(img=-) ⮕ (result=)`  
_Converts a rectangular coordinate image to polar coordinates_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to transform |
| `⮕ result` | `error` |   |   |   |   | - - - The transformed image |
---

### `remove-brightness(img=- lowerBright=0.1 minBright=0.2 maxBright=0.8 upperBright=0.9) ⮕ (result=)`  
_Removes pixels based on their brightness, making pixels inside the specified range transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerBright` | `float64` | `0.1` | `0` | `1` |   | The lower brightness threshold (below this becomes transparent) |
| `minBright` | `float64` | `0.2` | `0` | `1` |   | The minimum brightness for full opacity (fade from 0% to 100% between lowerBright and this) |
| `maxBright` | `float64` | `0.8` | `0` | `1` |   | The maximum brightness for full opacity (fade from 100% to 0% between this and upperBright) |
| `upperBright` | `float64` | `0.9` | `0` | `1` |   | The upper brightness threshold (above this becomes transparent) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with pixels in the specified brightness range removed (16-bit) |
---

### `remove-hsl(img=- lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=)`  
_Removes pixels based on their hue, saturation, and luminance, making pixels inside the specified ranges transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerHue` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The lower hue threshold (below this becomes transparent) |
| `minHue` | `float64` | `30` | `0` | `360` | `&#34;°&#34;` | The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this) |
| `maxHue` | `float64` | `330` | `0` | `360` | `&#34;°&#34;` | The maximum hue for full opacity (fade from 100% to 0% between this and upperHue) |
| `upperHue` | `float64` | `360` | `0` | `360` | `&#34;°&#34;` | The upper hue threshold (above this becomes transparent) |
| `lowerSat` | `float64` | `0.1` | `0` | `1` |   | The lower saturation threshold (below this becomes transparent) |
| `minSat` | `float64` | `0.2` | `0` | `1` |   | The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this) |
| `maxSat` | `float64` | `0.8` | `0` | `1` |   | The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat) |
| `upperSat` | `float64` | `0.9` | `0` | `1` |   | The upper saturation threshold (above this becomes transparent) |
| `lowerLum` | `float64` | `0.1` | `0` | `1` |   | The lower luminance threshold (below this becomes transparent) |
| `minLum` | `float64` | `0.2` | `0` | `1` |   | The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this) |
| `maxLum` | `float64` | `0.8` | `0` | `1` |   | The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum) |
| `upperLum` | `float64` | `0.9` | `0` | `1` |   | The upper luminance threshold (above this becomes transparent) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with only pixels in all specified ranges visible (16-bit) |
---

### `resize-fit(img=- maxW=0 maxH=0) ⮕ (result=)`  
_Resize an image to fit within a bounding box while preserving aspect ratio_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to resize |
| `maxW` | `int` | `0` |   |   |   | The maximum width (pixels) |
| `maxH` | `int` | `0` |   |   |   | The maximum height (pixels) |
| `⮕ result` | `error` |   |   |   |   | - - - The resized image |
---

### `resize-max-mp(img=- mpMax=0) ⮕ (result=)`  
_Resize an image to stay within a maximum amount of megapixels_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to resize |
| `mpMax` | `int` | `0` |   |   |   | The maximum amount of megapixels |
| `⮕ result` | `error` |   |   |   |   | - - - The resized image |
---

### `rgba(r=0 g=0 b=0 alpha=1) ⮕ (result=)`  
_Creates a color from RGBA values (8-bit per channel)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The red component |
| `g` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The green component |
| `b` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The blue component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The alpha component |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `rgba64(r=0 g=0 b=0 alpha=1) ⮕ (result=)`  
_Creates a color from RGBA values (16-bit per channel)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `float64` | `0` | `0` | `65535` | `&#34;&#34;` | The red component |
| `g` | `float64` | `0` | `0` | `65535` | `&#34;&#34;` | The green component |
| `b` | `float64` | `0` | `0` | `65535` | `&#34;&#34;` | The blue component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The alpha component |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `rotate(img=- angle=0) ⮕ (result=)`  
_Rotates an image around its center by a specified angle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to rotate |
| `angle` | `float64` | `0` | `-360` | `360` |   | The rotation angle in degrees (positive = clockwise) |
| `⮕ result` | `error` |   |   |   |   | - - - The rotated image |
---

### `rotate-hsl(img=- rotate=0 lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=)`  
_Rotates the hue of pixels based on their hue, saturation, and luminance_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `rotate` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The lower hue threshold (below this becomes transparent) |
| `lowerHue` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The lower hue threshold (below this becomes transparent) |
| `minHue` | `float64` | `30` | `0` | `360` | `&#34;°&#34;` | The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this) |
| `maxHue` | `float64` | `330` | `0` | `360` | `&#34;°&#34;` | The maximum hue for full opacity (fade from 100% to 0% between this and upperHue) |
| `upperHue` | `float64` | `360` | `0` | `360` | `&#34;°&#34;` | The upper hue threshold (above this becomes transparent) |
| `lowerSat` | `float64` | `0.1` | `0` | `1` |   | The lower saturation threshold (below this becomes transparent) |
| `minSat` | `float64` | `0.2` | `0` | `1` |   | The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this) |
| `maxSat` | `float64` | `0.8` | `0` | `1` |   | The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat) |
| `upperSat` | `float64` | `0.9` | `0` | `1` |   | The upper saturation threshold (above this becomes transparent) |
| `lowerLum` | `float64` | `0.1` | `0` | `1` |   | The lower luminance threshold (below this becomes transparent) |
| `minLum` | `float64` | `0.2` | `0` | `1` |   | The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this) |
| `maxLum` | `float64` | `0.8` | `0` | `1` |   | The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum) |
| `upperLum` | `float64` | `0.9` | `0` | `1` |   | The upper luminance threshold (above this becomes transparent) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with only pixels in all specified ranges visible (16-bit) |
---

### `round(x=-) ⮕ (result=)`  
_Returns the nearest integer to x, rounding to even on ties_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The nearest integer to x |
---

### `saturation(img=- factor=1) ⮕ (result=)`  
_Adjusts the color saturation of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust saturation of |
| `factor` | `float64` | `1` | `0` | `2` |   | The saturation factor (0 = grayscale, 1 = unchanged, 2 = super saturated) |
| `⮕ result` | `error` |   |   |   |   | - - - The saturation-adjusted image |
---

### `save(img=- path="-")`  
_Saves an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to save |
| `path` | `string` | `"-"` |   |   |   | - - Path where to save |
---

### `scale(img=- sx=0 sy=0) ⮕ (result=)`  
_Scales an image by specified factors_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to scale |
| `sx` | `float64` | `0` |   |   |   | The horizontal scale factor |
| `sy` | `float64` | `0` |   |   |   | The vertical scale factor |
| `⮕ result` | `error` |   |   |   |   | - - - The scaled image |
---

### `sec(radians=-) ⮕ (result=)`  
_calculates the secant of an angle in radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - secant value (1/cos) |
---

### `sec-of-triangle(hypotenuse=- adjacent=-) ⮕ (result=)`  
_calculates secant using hypotenuse and adjacent sides_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `⮕ result` | `error` |   |   |   |   | - - - secant value (hypotenuse/adjacent) |
---

### `sech(x=-) ⮕ (result=)`  
_calculates the hyperbolic secant of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - hyperbolic secant value (1/cosh) |
---

### `select-brightness(img=- lowerBright=0.1 minBright=0.2 maxBright=0.8 upperBright=0.9) ⮕ (result=)`  
_Selects pixels based on their brightness, making pixels outside the specified range transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerBright` | `float64` | `0.1` | `0` | `1` |   | The lower brightness threshold (below this becomes transparent) |
| `minBright` | `float64` | `0.2` | `0` | `1` |   | The minimum brightness for full opacity (fade from 0% to 100% between lowerBright and this) |
| `maxBright` | `float64` | `0.8` | `0` | `1` |   | The maximum brightness for full opacity (fade from 100% to 0% between this and upperBright) |
| `upperBright` | `float64` | `0.9` | `0` | `1` |   | The upper brightness threshold (above this becomes transparent) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with only pixels in the specified brightness range visible (16-bit) |
---

### `select-hsl(img=- lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=)`  
_Selects pixels based on their hue, saturation, and luminance, making pixels outside the specified ranges transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerHue` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The lower hue threshold (below this becomes transparent) |
| `minHue` | `float64` | `30` | `0` | `360` | `&#34;°&#34;` | The minimum hue for full opacity (fade from 0% to 100% between lowerHue and this) |
| `maxHue` | `float64` | `330` | `0` | `360` | `&#34;°&#34;` | The maximum hue for full opacity (fade from 100% to 0% between this and upperHue) |
| `upperHue` | `float64` | `360` | `0` | `360` | `&#34;°&#34;` | The upper hue threshold (above this becomes transparent) |
| `lowerSat` | `float64` | `0.1` | `0` | `1` |   | The lower saturation threshold (below this becomes transparent) |
| `minSat` | `float64` | `0.2` | `0` | `1` |   | The minimum saturation for full opacity (fade from 0% to 100% between lowerSat and this) |
| `maxSat` | `float64` | `0.8` | `0` | `1` |   | The maximum saturation for full opacity (fade from 100% to 0% between this and upperSat) |
| `upperSat` | `float64` | `0.9` | `0` | `1` |   | The upper saturation threshold (above this becomes transparent) |
| `lowerLum` | `float64` | `0.1` | `0` | `1` |   | The lower luminance threshold (below this becomes transparent) |
| `minLum` | `float64` | `0.2` | `0` | `1` |   | The minimum luminance for full opacity (fade from 0% to 100% between lowerLum and this) |
| `maxLum` | `float64` | `0.8` | `0` | `1` |   | The maximum luminance for full opacity (fade from 100% to 0% between this and upperLum) |
| `upperLum` | `float64` | `0.9` | `0` | `1` |   | The upper luminance threshold (above this becomes transparent) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with only pixels in all specified ranges visible (16-bit) |
---

### `select-hue(img=- hue=0 toleranceLeft=30 toleranceRight=30 softness=0.5 minSaturation=0.15) ⮕ (result=)`  
_Selects a specific hue from the image and makes everything else transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `hue` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The target hue to keep (in degrees) |
| `toleranceLeft` | `float64` | `30` | `0` | `180` | `&#34;°&#34;` | How much to include to the left (lower hue, in degrees) |
| `toleranceRight` | `float64` | `30` | `0` | `180` | `&#34;°&#34;` | How much to include to the right (higher hue, in degrees) |
| `softness` | `float64` | `0.5` | `0` | `1` |   | How soft the transition should be (0 = hard cut, 1 = very soft) |
| `minSaturation` | `float64` | `0.15` | `0` | `1` |   | The minimum saturation required to keep a pixel (smooth fade below) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with only the selected hue visible (16-bit) |
---

### `sepia(img=-) ⮕ (result=)`  
_Changes the tone of an image to sepia_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to change to sepia tone |
| `⮕ result` | `error` |   |   |   |   | - - - The sepia-toned image |
---

### `set-alpha(c=- alpha=1) ⮕ (result=)`  
_Sets the alpha channel of a color_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `c` | `color.RGBA64` | `-` |   |   | `&#34;&#34;` | The color to modify |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The new alpha value |
| `⮕ result` | `error` |   |   |   |   | - - - The color with new alpha |
---

### `sharpen(img=- intensity=1 radius=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=)`  
_Sharpens an image using a highpass combined with vivid light blending_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to sharpen |
| `intensity` | `float64` | `1` | `0` | `1` |   | The intensity of the sharpening effect |
| `radius` | `float64` | `1` | `0.1` | `2` |   | The radius of the filter in pixels (higher values detect larger edges) |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel |
| `⮕ result` | `error` |   |   |   |   | - - - The sharpened image |
---

### `sin(radians=-) ⮕ (result=)`  
_calculates the sine of an angle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - sine value between -1 and 1 |
---

### `sin-of-triangle(opposite=- hypotenuse=-) ⮕ (result=)`  
_calculates sine using opposite and hypotenuse sides_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `hypotenuse` | `float64` | `-` |   |   |   | The hypotenuse length |
| `⮕ result` | `error` |   |   |   |   | - - - sine value (opposite/hypotenuse) |
---

### `sin2(x=-) ⮕ (result=)`  
_calculates the square of sine (sin²(x))_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - squared sine value (sin(x)²) |
---

### `sinc(x=-) ⮕ (result=)`  
_calculates the sinc function (sin(x)/x)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - sinc value (sin(x)/x, with sinc(0) = 1) |
---

### `sinh(x=-) ⮕ (result=)`  
_calculates the hyperbolic sine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - hyperbolic sine value |
---

### `slope(x1=- y1=- x2=- y2=-) ⮕ (result=)`  
_Calculates the slope between two points_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x1` | `float64` | `-` |   |   |   | The x coordinate of the first point |
| `y1` | `float64` | `-` |   |   |   | The y coordinate of the first point |
| `x2` | `float64` | `-` |   |   |   | The x coordinate of the second point |
| `y2` | `float64` | `-` |   |   |   | The y coordinate of the second point |
| `⮕ result` | `error` |   |   |   |   | - - - The slope value |
---

### `sprintf(str="-" a=-) ⮕ (result=)`  
_Creates formatted strings_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `str` | `string` | `"-"` |   |   |   | - - The format string |
| `a` | `[]any` | `-` |   |   |   | A slice with the arguments |
| `⮕ result` | `error` |   |   |   |   | - - - The formatted string |
---

### `sqrt(x=-) ⮕ (result=)`  
_Returns the square root of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The square root of x |
---

### `square(x=-) ⮕ (result=)`  
_Calculates the square of a number_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - The square of x |
---

### `sub(a=- b=-) ⮕ (result=)`  
_Subtracts the two numbers_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `a` | `float64` | `-` |   |   |   | The first number |
| `b` | `float64` | `-` |   |   |   | The second number |
| `⮕ result` | `error` |   |   |   |   | - - - a-b |
---

### `sub-n(a=- b=- n=-) ⮕ (result=)`  
_Multiplies b by n and subtracts the result from a_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `a` | `float64` | `-` |   |   |   | The first number |
| `b` | `float64` | `-` |   |   |   | The second number |
| `n` | `float64` | `-` |   |   |   | The multiplier for the second number |
| `⮕ result` | `error` |   |   |   |   | - - - a - (n * b) |
---

### `tan(radians=-) ⮕ (result=)`  
_calculates the tangent of an angle in radians_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `radians` | `float64` | `-` |   |   |   | The angle in radians |
| `⮕ result` | `error` |   |   |   |   | - - - tangent value (unbounded) |
---

### `tan-of-slope(m=-) ⮕ (result=)`  
_Calculates the angle from a slope value_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `m` | `float64` | `-` |   |   |   | The slope value |
| `⮕ result` | `error` |   |   |   |   | - - - The angle in radians |
---

### `tan-of-triangle(opposite=- adjacent=-) ⮕ (result=)`  
_calculates tangent using opposite and adjacent sides_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `opposite` | `float64` | `-` |   |   |   | The opposite side length |
| `adjacent` | `float64` | `-` |   |   |   | The adjacent side length |
| `⮕ result` | `error` |   |   |   |   | - - - tangent value (opposite/adjacent) |
---

### `tan2(x=-) ⮕ (result=)`  
_calculates the square of tangent (tan²(x))_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - squared tangent value (tan(x)²) |
---

### `tanh(x=-) ⮕ (result=)`  
_calculates the hyperbolic tangent of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - hyperbolic tangent value between -1 and 1 |
---

### `text(t="-" style=- outline=-) ⮕ (result=)`  
_Generates the given text with the given styles._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `t` | `string` | `"-"` |   |   |   | - - The text to generate |
| `style` | `TextStyle` | `-` |   |   |   | The text style (font, size, color) |
| `outline` | `LineStyle` | `-` |   |   |   | The thickness and color of the outline |
| `⮕ result` | `error` |   |   |   |   | - - - The resulting image |
---

### `threshold(img=- level=0.5) ⮕ (result=)`  
_Converts image to black and white based on a brightness threshold_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply thresholding to |
| `level` | `float64` | `0.5` | `0` | `1` |   | The brightness threshold |
| `⮕ result` | `error` |   |   |   |   | - - - The thresholded (black and white) image |
---

### `time(layout="2006-01-02 15:04:05") ⮕ (result=)`  
_Returns the current time according to the given layout_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `layout` | `string` | `"2006-01-02 15:04:05"` |   |   |   | The layout |
| `⮕ result` | `error` |   |   |   |   | - - - The current time as string |
---

### `transform(img=- dx=0 dy=0 angle=0 sx=0 sy=0) ⮕ (result=)`  
_Applies translation, rotation, and scaling to an image in one operation_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to transform |
| `dx` | `float64` | `0` |   |   |   | The horizontal translation in pixels |
| `dy` | `float64` | `0` |   |   |   | The vertical translation in pixels |
| `angle` | `float64` | `0` |   |   |   | The rotation angle in degrees (clockwise) |
| `sx` | `float64` | `0` |   |   |   | The horizontal scale factor |
| `sy` | `float64` | `0` |   |   |   | The vertical scale factor |
| `⮕ result` | `error` |   |   |   |   | - - - The transformed image |
---

### `translate(img=- dx=0 dy=0) ⮕ (result=)`  
_Translates (moves) an image by a specified amount_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to translate |
| `dx` | `float64` | `0` |   |   |   | The horizontal translation amount in % (positive = right) |
| `dy` | `float64` | `0` |   |   |   | The vertical translation amount in % (positive = down) |
| `⮕ result` | `error` |   |   |   |   | - - - The translated image |
---

### `uppercase(str="-") ⮕ (result=)`  
_Uppercases a string_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `str` | `string` | `"-"` |   |   |   | - - The string to uppercase |
| `⮕ result` | `error` |   |   |   |   | - - - The uppercased string |
---

### `vercos(x=-) ⮕ (result=)`  
_calculates the versed cosine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - versed cosine value (1 &#43; cos(x)) |
---

### `versin(x=-) ⮕ (result=)`  
_calculates the versed sine of x_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `-` |   |   |   | The input value |
| `⮕ result` | `error` |   |   |   |   | - - - versed sine value (1 - cos(x)) |
---

### `vibrance(img=- factor=0) ⮕ (result=)`  
_Adjusts the saturation of an image, protecting already saturated colors and skin tones._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust vibrance of |
| `factor` | `float64` | `0` | `-1` | `1` |   | The vibrance adjustment factor (-1 = less vibrant, 0 = unchanged, 1 = more vibrant) |
| `⮕ result` | `error` |   |   |   |   | - - - The vibrance-adjusted image |
---

### `vignette(img=- strength=0.5 falloff=0.8) ⮕ (result=)`  
_Adds a vignette effect (darkens/lightens edges)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply vignette to |
| `strength` | `float64` | `0.5` | `0` | `1` |   | Darkness/Lightness intensity (0 to 1) |
| `falloff` | `float64` | `0.8` | `0.1` | `2` |   | How quickly the effect fades (0.1 to 2.0) |
| `⮕ result` | `error` |   |   |   |   | - - - The image with vignette effect |
---

### `xyz(x=0 y=0 z=0 alpha=1) ⮕ (result=)`  
_Creates a color from CIE XYZ values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `0` | `0` | `0.95047` | `&#34;&#34;` | The X component (red) |
| `y` | `float64` | `0` | `0` | `1` | `&#34;&#34;` | The Y component (green) |
| `z` | `float64` | `0` | `0` | `1.08883` | `&#34;&#34;` | The Z component (blue) |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `ycbcr(y=0 cb=128 cr=128 alpha=1) ⮕ (result=)`  
_Creates a color from YCbCr values (digital video)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `y` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The luminance component |
| `cb` | `float64` | `128` | `0` | `255` | `&#34;&#34;` | The blue-difference chroma component |
| `cr` | `float64` | `128` | `0` | `255` | `&#34;&#34;` | The red-difference chroma component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

### `yuv(y=0 u=0 v=0 alpha=1) ⮕ (result=)`  
_Creates a color from YUV values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `y` | `float64` | `0` | `0` | `1` | `&#34;&#34;` | The luminance component |
| `u` | `float64` | `0` | `-0.436` | `0.436` | `&#34;&#34;` | The U chrominance component |
| `v` | `float64` | `0` | `-0.615` | `0.615` | `&#34;&#34;` | The V chrominance component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `error` |   |   |   |   | - - - The color |
---

 
 