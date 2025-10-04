# PixelPipeline Script (v1.0.0)

This is a custom language implementation with support for functions, variables, and various data types.

## Data Types

The language supports the following data types:

- `int`: Integer values
- `float`: Floating-point values
- `string`: Text values (enclosed in double quotes)
- `bool`: Boolean values (`true` or `false`)

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




## Functions

### `auto-contrast(img=- threshold=0.01 strength=1) ⮕ (result=-)`  
_Automatically adjusts the contrast of an image by stretching the histogram to use the full range of values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-contrast |
| `threshold` | `float64` | `0.01` |   |   | `%` | The percentage of pixels to ignore at both ends of the histogram (0-0.5) |
| `strength` | `float64` | `1` | `0` | `1` |   | How strongly to apply the contrast adjustment (0 = no change, 1 = full correction) |
| `⮕ result` | `any` | `-` |   |   |   | The contrast-adjusted image |
---

### `auto-levels(img=- lowPercentile=0.05 highPercentile=0.995 adjustAlpha=false) ⮕ (result=-)`  
_Automatically adjusts the contrast and brightness of an image by stretching the histogram to use the full range of values, ignoring outliers using percentiles_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-level |
| `lowPercentile` | `float64` | `0.05` |   |   | `%` | The lower percentile to ignore (e.g., 0.5) |
| `highPercentile` | `float64` | `0.995` |   |   | `%` | The upper percentile to ignore (e.g., 99.5) |
| `adjustAlpha` | `bool` | `false` |   |   |   | Whether to adjust alpha channel (false = preserve original alpha) |
| `⮕ result` | `any` | `-` |   |   |   | The auto-leveled image |
---

### `auto-tone(img=- levelsLow=0.005 levelsHigh=0.9995 whiteThresh=0.99 whiteStrength=0.8 contrastThresh=0.01 contrastStrength=0.8) ⮕ (result=-)`  
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
| `⮕ result` | `any` | `-` |   |   |   | The auto-toned image |
---

### `auto-white-balance(img=- threshold=0.95 strength=1) ⮕ (result=-)`  
_Automatically adjusts the white balance of an image by finding bright areas and making them neutral_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to auto-white-balance |
| `threshold` | `float64` | `0.95` |   |   | `%` | The brightness threshold to consider as white (0-1) |
| `strength` | `float64` | `1` | `0` | `1` |   | How strongly to apply the white balance (0 = no change, 1 = full correction) |
| `⮕ result` | `any` | `-` |   |   |   | The white-balanced image |
---

### `blend-average(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the average blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-color(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the color blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-color-burn(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the color burn blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-color-dodge(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the color dodge blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-contrast-negate(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the contrast negate blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-darken(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the darken blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-darker-color(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the darker color blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-difference(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the difference blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-divide(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the divide blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-erase(imgA=- imgB=-) ⮕ (result=-)`  
_Erases the bottom image wherever the top image is present (destination out)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-exclusion(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the exclusion blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-glow(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the glow blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-hard-light(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the hard light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-hard-mix(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the hard mix blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-hue(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the hue blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-lighten(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the lighten blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-lighter-color(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the lighter color blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-linear-light(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the linear light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-luminosity(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the luminosity blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-multiply(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the multiply blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-negation(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the negation blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-normal(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the normal blend mode (alpha compositing)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-overlay(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the overlay blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-pin-light(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the pin light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-reflect(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the reflect blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-saturation(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the saturation blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-screen(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the screen blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-soft-light(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the soft light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-subtract(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the subtract blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blend-vivid-light(imgA=- imgB=-) ⮕ (result=-)`  
_Blends the two images using the vivid light blend mode_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `imgA` | `*image.NRGBA64` | `-` |   |   |   | The bottom image |
| `imgB` | `*image.NRGBA64` | `-` |   |   |   | The top image |
| `⮕ result` | `any` | `-` |   |   |   | The blended image |
---

### `blur-box(img=- radius=1) ⮕ (result=-)`  
_Applies a box blur to an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `radius` | `int` | `1` | `1` | `10` |   | The blur radius (size of the box kernel) |
| `⮕ result` | `any` | `-` |   |   |   | The blurred image |
---

### `blur-gaussian(img=- radius=1) ⮕ (result=-)`  
_Applies a Gaussian blur to the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `radius` | `float64` | `1` | `1` | `10` |   | The blur radius (higher values create more blur) |
| `⮕ result` | `any` | `-` |   |   |   | The blurred image |
---

### `blur-motion(img=- length=5 angle=0) ⮕ (result=-)`  
_Applies a motion blur to an image along a specified angle._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `length` | `int` | `5` | `1` | `100` |   | The length of the motion blur (in pixels) |
| `angle` | `float64` | `0` | `0` | `360` |   | The angle of the motion blur (in degrees) |
| `⮕ result` | `any` | `-` |   |   |   | The blurred image |
---

### `blur-zoom(img=- strength=0.25 centerX=0.5 centerY=0.5) ⮕ (result=-)`  
_Applies a zoom blur effect to an image._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to blur |
| `strength` | `float64` | `0.25` | `0` | `1` |   | The strength of the blur effect (higher means more blur) |
| `centerX` | `float64` | `0.5` | `0` | `1` |   | X coordinate of the blur center (default: image center) |
| `centerY` | `float64` | `0.5` | `0` | `1` |   | Y coordinate of the blur center (default: image center) |
| `⮕ result` | `any` | `-` |   |   |   | The blurred image |
---

### `brightness(img=- factor=0) ⮕ (result=-)`  
_Changes the brightness of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to change brightness of |
| `factor` | `float64` | `0` | `0` | `2` |   | The change factor |
| `⮕ result` | `any` | `-` |   |   |   | The image with brightness changed |
---

### `chromatic-aberration(img=- amount=5) ⮕ (result=-)`  
_Creates a chromatic aberration effect by offsetting color channels_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply the effect to |
| `amount` | `float64` | `5` | `0` | `20` |   | The amount of color channel separation |
| `⮕ result` | `any` | `-` |   |   |   | The image with chromatic aberration |
---

### `clarity(img=- intensity=1 radius=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=-)`  
_Enhances local contrast while preserving overall image structure_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to enhance |
| `intensity` | `float64` | `1` | `0` | `1` |   | The intensity of the clarity effect |
| `radius` | `float64` | `1` | `0.1` | `2` |   | The radius of the filter in pixels (higher values affect larger areas) |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel |
| `⮕ result` | `any` | `-` |   |   |   | The enhanced image |
---

### `cmyk(c=0 m=0 y=0 k=0 alpha=1) ⮕ (result=-)`  
_Creates a color from CMYK values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `c` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The cyan component |
| `m` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The magenta component |
| `y` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The yellow component |
| `k` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The key (black) component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The alpha component |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `color-balance(img=- rFactor=1 gFactor=1 bFactor=1) ⮕ (result=-)`  
_Adjusts the balance of Red, Green, and Blue channels_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust color balance of |
| `rFactor` | `float64` | `1` | `0` | `2` |   | Red channel adjustment factor |
| `gFactor` | `float64` | `1` | `0` | `2` |   | Green channel adjustment factor |
| `bFactor` | `float64` | `1` | `0` | `2` |   | Blue channel adjustment factor |
| `⮕ result` | `any` | `-` |   |   |   | The color-balanced image |
---

### `colorize(img=- col=-) ⮕ (result=-)`  
_Colorizes the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to colorize |
| `col` | `color.RGBA64` | `-` |   |   |   | The color that determines the hue to use for colorization |
| `⮕ result` | `any` | `-` |   |   |   | The colorized image |
---

### `contrast(img=- factor=1) ⮕ (result=-)`  
_Adjusts the contrast of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust contrast of |
| `factor` | `float64` | `1` | `0` | `2` |   | The contrast factor (0 = gray, 1 = unchanged, 2 = maximum) |
| `⮕ result` | `any` | `-` |   |   |   | The contrast-adjusted image |
---

### `crop(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=-)`  
_Crops an image by specified percentages from each side_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `left` | `float64` | `0` |   |   |   | The percentage to crop from the left side (0-1) |
| `right` | `float64` | `0` |   |   |   | The percentage to crop from the right side (0-1) |
| `top` | `float64` | `0` |   |   |   | The percentage to crop from the top side (0-1) |
| `bottom` | `float64` | `0` |   |   |   | The percentage to crop from the bottom side (0-1) |
| `⮕ result` | `any` | `-` |   |   |   | The cropped image |
---

### `crop-circle(img=- radius=1 offsetX=0 offsetY=0) ⮕ (result=-)`  
_Crops an image using a circular mask. The circle is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the radius is a percentage (0-1) of half the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `radius` | `float64` | `1` | `0` | `1` |   | Radius as a percentage of half the min(width, height) |
| `offsetX` | `float64` | `0` | `-1` | `1` |   | Horizontal offset from image center (percentage of width, -1..1) |
| `offsetY` | `float64` | `0` | `-1` | `1` |   | Vertical offset from image center (percentage of height, -1..1) |
| `⮕ result` | `any` | `-` |   |   |   | The circularly cropped image (pixels outside the circle are transparent) |
---

### `crop-circle-px(img=- radius=1 offsetX=0 offsetY=0) ⮕ (result=-)`  
_Crops an image using a circular mask. The circle is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the radius is a percentage (0-1) of half the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `radius` | `float64` | `1` | `0` | `1` |   | Radius as a percentage of half the min(width, height) |
| `offsetX` | `int` | `0` |   |   |   | Horizontal offset from image center (pixels) |
| `offsetY` | `int` | `0` |   |   |   | Vertical offset from image center (pixels) |
| `⮕ result` | `any` | `-` |   |   |   | The circularly cropped image (pixels outside the circle are transparent) |
---

### `crop-px(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=-)`  
_Crops an image by specified amounts of pixels from each side_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `left` | `int` | `0` |   |   |   | The number of pixels to crop from the left side |
| `right` | `int` | `0` |   |   |   | The number of pixels to crop from the right side |
| `top` | `int` | `0` |   |   |   | The number of pixels to crop from the top side |
| `bottom` | `int` | `0` |   |   |   | The number of pixels to crop from the bottom side |
| `⮕ result` | `any` | `-` |   |   |   | The cropped image |
---

### `crop-square(img=- size=1 offsetX=0 offsetY=0) ⮕ (result=-)`  
_Crops an image using a square mask. The square is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the size is a percentage (0-1) of the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `size` | `float64` | `1` | `0` | `1` |   | Size as a percentage of the min(width, height) |
| `offsetX` | `float64` | `0` | `-1` | `1` |   | Horizontal offset from image center (percentage of width, -1..1) |
| `offsetY` | `float64` | `0` | `-1` | `1` |   | Vertical offset from image center (percentage of height, -1..1) |
| `⮕ result` | `any` | `-` |   |   |   | The square-cropped image (pixels outside the square are transparent) |
---

### `crop-square-px(img=- size=1 offsetX=0 offsetY=0) ⮕ (result=-)`  
_Crops an image using a square mask. The square is centered at (centerX&#43;offsetX, centerY&#43;offsetY) and the size is a percentage (0-1) of the minimum image dimension._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to crop |
| `size` | `float64` | `1` | `0` | `1` |   | Size as a percentage of the min(width, height) |
| `offsetX` | `int` | `0` |   |   |   | Horizontal offset from image center (pixels) |
| `offsetY` | `int` | `0` |   |   |   | Vertical offset from image center (pixels) |
| `⮕ result` | `any` | `-` |   |   |   | The square-cropped image (pixels outside the square are transparent) |
---

### `defisheye(img=- strength=1) ⮕ (result=-)`  
_Corrects fisheye lens distortion in an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to correct |
| `strength` | `float64` | `1` | `0` | `2` |   | The strength of the correction |
| `⮕ result` | `any` | `-` |   |   |   | The corrected image |
---

### `displace(img=- dMap=- amount=10) ⮕ (result=-)`  
_Displaces pixels based on the brightness of a displacement map_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to displace |
| `dMap` | `*image.NRGBA64` | `-` |   |   |   | The displacement map image |
| `amount` | `float64` | `10` | `0` | `50` |   | The amount of displacement |
| `⮕ result` | `any` | `-` |   |   |   | The displaced image |
---

### `edge-detect(img=-) ⮕ (result=-)`  
_Detects edges in the image using the Sobel operator_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to detect edges in |
| `⮕ result` | `any` | `-` |   |   |   | An image highlighting the edges |
---

### `enhance(img=- brightness=0 contrast=0 sharpening=1 rMin=0.75 rMax=1 gMin=0.75 gMax=1 bMin=0.75 bMax=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=-)`  
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
| `⮕ result` | `any` | `-` |   |   |   | The enhanceed image |
---

### `expand(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=-)`  
_Expands an image by adding transparent borders with specified percentage widths_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to expand |
| `left` | `float64` | `0` |   |   |   | The percentage to add to the left side (relative to original width) |
| `right` | `float64` | `0` |   |   |   | The percentage to add to the right side (relative to original width) |
| `top` | `float64` | `0` |   |   |   | The percentage to add to the top side (relative to original height) |
| `bottom` | `float64` | `0` |   |   |   | The percentage to add to the bottom side (relative to original height) |
| `⮕ result` | `any` | `-` |   |   |   | The expanded image |
---

### `expand-px(img=- left=0 right=0 top=0 bottom=0) ⮕ (result=-)`  
_Expands an image by adding transparent borders with specified pixel widths_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to expand |
| `left` | `int` | `0` |   |   |   | The number of pixels to add to the left side |
| `right` | `int` | `0` |   |   |   | The number of pixels to add to the right side |
| `top` | `int` | `0` |   |   |   | The number of pixels to add to the top side |
| `bottom` | `int` | `0` |   |   |   | The number of pixels to add to the bottom side |
| `⮕ result` | `any` | `-` |   |   |   | The expanded image |
---

### `exposure(img=- level=0) ⮕ (result=-)`  
_Adjusts the overall lightness or darkness of the image, simulating photographic exposure._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust exposure of |
| `level` | `float64` | `0` | `-2` | `2` |   | The exposure level adjustment (-2 = much darker, 0 = unchanged, 2 = much brighter) |
| `⮕ result` | `any` | `-` |   |   |   | The exposure-adjusted image |
---

### `fill(img=- col=-) ⮕ (result=-)`  
_Fills the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to fill |
| `col` | `color.RGBA64` | `-` |   |   |   | The fill color |
| `⮕ result` | `any` | `-` |   |   |   | The filled image |
---

### `fisheye(img=- strength=1) ⮕ (result=-)`  
_Applies a fisheye lens distortion effect to the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to distort |
| `strength` | `float64` | `1` | `0` | `2` |   | The strength of the fisheye effect |
| `⮕ result` | `any` | `-` |   |   |   | The distorted image |
---

### `flip-h(img=-) ⮕ (result=-)`  
_Flips an image horizontally (left to right)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to flip horizontally |
| `⮕ result` | `any` | `-` |   |   |   | The horizontally flipped image |
---

### `flip-v(img=-) ⮕ (result=-)`  
_Flips an image vertically (top to bottom)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to flip vertically |
| `⮕ result` | `any` | `-` |   |   |   | The vertically flipped image |
---

### `grayscale(img=-) ⮕ (result=-)`  
_Grayscales an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to grayscale |
| `⮕ result` | `any` | `-` |   |   |   | The grayscaled image |
---

### `highpass(img=- radius=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=-)`  
_Creates a high pass filter effect, resulting in a gray image with embossed edges_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply the high-pass filter to |
| `radius` | `float64` | `1` | `0.1` | `2` |   | The radius of the filter in pixels (higher values detect larger edges) |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel |
| `⮕ result` | `any` | `-` |   |   |   | The filtered image |
---

### `hsla(h=0 s=0.5 l=0.5 alpha=1) ⮕ (result=-)`  
_Creates a color from HSLA values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The color&#39;s hue |
| `s` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s saturation |
| `l` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s luminosity |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `hsv(h=0 s=0.5 v=0.5 alpha=1) ⮕ (result=-)`  
_Creates a color from HSV values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The color&#39;s hue |
| `s` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s saturation |
| `v` | `float64` | `0.5` | `0` | `1` | `&#34;%&#34;` | The color&#39;s value (brightness) |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `hue-rotate(img=- angle=0) ⮕ (result=-)`  
_Rotates the hue of image colors_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to rotate hue of |
| `angle` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The angle in degrees (0-360) |
| `⮕ result` | `any` | `-` |   |   |   | The hue-rotated image |
---

### `hwb(h=0 w=0 b=0 alpha=1) ⮕ (result=-)`  
_Creates a color from HWB (Hue, Whiteness, Blackness) values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The color&#39;s hue |
| `w` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The whiteness component |
| `b` | `float64` | `0` | `0` | `1` | `&#34;%&#34;` | The blackness component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `invert(img=-) ⮕ (result=-)`  
_Inverts an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to invert |
| `⮕ result` | `any` | `-` |   |   |   | The inverted image |
---

### `invert-hsl(img=- lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=-)`  
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
| `⮕ result` | `any` | `-` |   |   |   | The image with only pixels in all specified ranges visible (16-bit) |
---

### `lab(l=50 a=0 b=0 alpha=1) ⮕ (result=-)`  
_Creates a color from CIELAB values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `l` | `float64` | `50` | `0` | `100` | `&#34;&#34;` | The lightness component |
| `a` | `float64` | `0` | `-128` | `127` | `&#34;&#34;` | The green-red component |
| `b` | `float64` | `0` | `-128` | `127` | `&#34;&#34;` | The blue-yellow component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `lch(l=50 c=0 h=0 alpha=1) ⮕ (result=-)`  
_Creates a color from LCH (Lightness, Chroma, Hue) values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `l` | `float64` | `50` | `0` | `100` | `&#34;&#34;` | The lightness component |
| `c` | `float64` | `0` | `0` | `128` | `&#34;&#34;` | The chroma component |
| `h` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The hue component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `load(path="-") ⮕ (result=-)`  
_Loads an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `path` | `string` | `"-"` |   |   |   | - - Path to the image |
| `⮕ result` | `any` | `-` |   |   |   | The loaded image |
---

### `opacity(img=- amount=1) ⮕ (result=-)`  
_Adjusts the overall opacity/transparency of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust opacity of |
| `amount` | `float64` | `1` | `0` | `1` |   | The opacity amount (0 = fully transparent, 1 = unchanged) |
| `⮕ result` | `any` | `-` |   |   |   | The opacity-adjusted image |
---

### `pixelate(img=- size=8) ⮕ (result=-)`  
_Creates a pixelation effect by averaging colors in blocks_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to pixelate |
| `size` | `int` | `8` | `1` | `50` |   | The size of the pixel blocks |
| `⮕ result` | `any` | `-` |   |   |   | The pixelated image |
---

### `polar-to-rectangular(img=-) ⮕ (result=-)`  
_Converts a polar coordinate image to rectangular coordinates_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to transform |
| `⮕ result` | `any` | `-` |   |   |   | The transformed image |
---

### `posterize(img=- levels=4) ⮕ (result=-)`  
_Reduces the number of color levels in the image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to posterize |
| `levels` | `int` | `4` | `2` | `16` |   | Number of color levels per channel (2-16) |
| `⮕ result` | `any` | `-` |   |   |   | The posterized image |
---

### `rectangular-to-polar(img=-) ⮕ (result=-)`  
_Converts a rectangular coordinate image to polar coordinates_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to transform |
| `⮕ result` | `any` | `-` |   |   |   | The transformed image |
---

### `remove-brightness(img=- lowerBright=0.1 minBright=0.2 maxBright=0.8 upperBright=0.9) ⮕ (result=-)`  
_Removes pixels based on their brightness, making pixels inside the specified range transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerBright` | `float64` | `0.1` | `0` | `1` |   | The lower brightness threshold (below this becomes transparent) |
| `minBright` | `float64` | `0.2` | `0` | `1` |   | The minimum brightness for full opacity (fade from 0% to 100% between lowerBright and this) |
| `maxBright` | `float64` | `0.8` | `0` | `1` |   | The maximum brightness for full opacity (fade from 100% to 0% between this and upperBright) |
| `upperBright` | `float64` | `0.9` | `0` | `1` |   | The upper brightness threshold (above this becomes transparent) |
| `⮕ result` | `any` | `-` |   |   |   | The image with pixels in the specified brightness range removed (16-bit) |
---

### `remove-hsl(img=- lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=-)`  
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
| `⮕ result` | `any` | `-` |   |   |   | The image with only pixels in all specified ranges visible (16-bit) |
---

### `resize-max-mp(img=- mpMax=0) ⮕ (result=-)`  
_Resize an image to stay within a maximum amount of megapixels_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to resize |
| `mpMax` | `int` | `0` |   |   |   | The maximum amount of megapixels |
| `⮕ result` | `any` | `-` |   |   |   | The resized image |
---

### `rgba(r=0 g=0 b=0 alpha=1) ⮕ (result=-)`  
_Creates a color from RGBA values (8-bit per channel)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The red component |
| `g` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The green component |
| `b` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The blue component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The alpha component |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `rgba64(r=0 g=0 b=0 alpha=1) ⮕ (result=-)`  
_Creates a color from RGBA values (16-bit per channel)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `r` | `float64` | `0` | `0` | `65535` | `&#34;&#34;` | The red component |
| `g` | `float64` | `0` | `0` | `65535` | `&#34;&#34;` | The green component |
| `b` | `float64` | `0` | `0` | `65535` | `&#34;&#34;` | The blue component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The alpha component |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `rotate(img=- angle=0) ⮕ (result=-)`  
_Rotates an image around its center by a specified angle_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to rotate |
| `angle` | `float64` | `0` | `-360` | `360` |   | The rotation angle in degrees (positive = clockwise) |
| `⮕ result` | `any` | `-` |   |   |   | The rotated image |
---

### `rotate-hsl(img=- rotate=0 lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=-)`  
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
| `⮕ result` | `any` | `-` |   |   |   | The image with only pixels in all specified ranges visible (16-bit) |
---

### `saturation(img=- factor=1) ⮕ (result=-)`  
_Adjusts the color saturation of an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust saturation of |
| `factor` | `float64` | `1` | `0` | `2` |   | The saturation factor (0 = grayscale, 1 = unchanged, 2 = super saturated) |
| `⮕ result` | `any` | `-` |   |   |   | The saturation-adjusted image |
---

### `save(img=- path="-")`  
_Saves an image_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to save |
| `path` | `string` | `"-"` |   |   |   | - - Path where to save |
---

### `scale(img=- sx=0 sy=0) ⮕ (result=-)`  
_Scales an image by specified factors_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to scale |
| `sx` | `float64` | `0` |   |   |   | The horizontal scale factor |
| `sy` | `float64` | `0` |   |   |   | The vertical scale factor |
| `⮕ result` | `any` | `-` |   |   |   | The scaled image |
---

### `select-brightness(img=- lowerBright=0.1 minBright=0.2 maxBright=0.8 upperBright=0.9) ⮕ (result=-)`  
_Selects pixels based on their brightness, making pixels outside the specified range transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `lowerBright` | `float64` | `0.1` | `0` | `1` |   | The lower brightness threshold (below this becomes transparent) |
| `minBright` | `float64` | `0.2` | `0` | `1` |   | The minimum brightness for full opacity (fade from 0% to 100% between lowerBright and this) |
| `maxBright` | `float64` | `0.8` | `0` | `1` |   | The maximum brightness for full opacity (fade from 100% to 0% between this and upperBright) |
| `upperBright` | `float64` | `0.9` | `0` | `1` |   | The upper brightness threshold (above this becomes transparent) |
| `⮕ result` | `any` | `-` |   |   |   | The image with only pixels in the specified brightness range visible (16-bit) |
---

### `select-hsl(img=- lowerHue=0 minHue=30 maxHue=330 upperHue=360 lowerSat=0.1 minSat=0.2 maxSat=0.8 upperSat=0.9 lowerLum=0.1 minLum=0.2 maxLum=0.8 upperLum=0.9) ⮕ (result=-)`  
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
| `⮕ result` | `any` | `-` |   |   |   | The image with only pixels in all specified ranges visible (16-bit) |
---

### `select-hue(img=- hue=0 toleranceLeft=30 toleranceRight=30 softness=0.5 minSaturation=0.15) ⮕ (result=-)`  
_Selects a specific hue from the image and makes everything else transparent_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to process (16-bit) |
| `hue` | `float64` | `0` | `0` | `360` | `&#34;°&#34;` | The target hue to keep (in degrees) |
| `toleranceLeft` | `float64` | `30` | `0` | `180` | `&#34;°&#34;` | How much to include to the left (lower hue, in degrees) |
| `toleranceRight` | `float64` | `30` | `0` | `180` | `&#34;°&#34;` | How much to include to the right (higher hue, in degrees) |
| `softness` | `float64` | `0.5` | `0` | `1` |   | How soft the transition should be (0 = hard cut, 1 = very soft) |
| `minSaturation` | `float64` | `0.15` | `0` | `1` |   | The minimum saturation required to keep a pixel (smooth fade below) |
| `⮕ result` | `any` | `-` |   |   |   | The image with only the selected hue visible (16-bit) |
---

### `sepia(img=-) ⮕ (result=-)`  
_Changes the tone of an image to sepia_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to change to sepia tone |
| `⮕ result` | `any` | `-` |   |   |   | The sepia-toned image |
---

### `sharpen(img=- intensity=1 radius=1 rWeight=0.299 gWeight=0.587 bWeight=0.114) ⮕ (result=-)`  
_Sharpens an image using a highpass combined with vivid light blending_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to sharpen |
| `intensity` | `float64` | `1` | `0` | `1` |   | The intensity of the sharpening effect |
| `radius` | `float64` | `1` | `0.1` | `2` |   | The radius of the filter in pixels (higher values detect larger edges) |
| `rWeight` | `float64` | `0.299` | `0` | `1` |   | The weight of the red channel |
| `gWeight` | `float64` | `0.587` | `0` | `1` |   | The weight of the green channel |
| `bWeight` | `float64` | `0.114` | `0` | `1` |   | The weight of the blue channel |
| `⮕ result` | `any` | `-` |   |   |   | The sharpened image |
---

### `threshold(img=- level=0.5) ⮕ (result=-)`  
_Converts image to black and white based on a brightness threshold_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply thresholding to |
| `level` | `float64` | `0.5` | `0` | `1` |   | The brightness threshold |
| `⮕ result` | `any` | `-` |   |   |   | The thresholded (black and white) image |
---

### `transform(img=- dx=0 dy=0 angle=0 sx=0 sy=0) ⮕ (result=-)`  
_Applies translation, rotation, and scaling to an image in one operation_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to transform |
| `dx` | `float64` | `0` |   |   |   | The horizontal translation in pixels |
| `dy` | `float64` | `0` |   |   |   | The vertical translation in pixels |
| `angle` | `float64` | `0` |   |   |   | The rotation angle in degrees (clockwise) |
| `sx` | `float64` | `0` |   |   |   | The horizontal scale factor |
| `sy` | `float64` | `0` |   |   |   | The vertical scale factor |
| `⮕ result` | `any` | `-` |   |   |   | The transformed image |
---

### `translate(img=- dx=0 dy=0) ⮕ (result=-)`  
_Translates (moves) an image by a specified amount_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to translate |
| `dx` | `float64` | `0` |   |   |   | The horizontal translation amount in % (positive = right) |
| `dy` | `float64` | `0` |   |   |   | The vertical translation amount in % (positive = down) |
| `⮕ result` | `any` | `-` |   |   |   | The translated image |
---

### `vibrance(img=- factor=0) ⮕ (result=-)`  
_Adjusts the saturation of an image, protecting already saturated colors and skin tones._

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to adjust vibrance of |
| `factor` | `float64` | `0` | `-1` | `1` |   | The vibrance adjustment factor (-1 = less vibrant, 0 = unchanged, 1 = more vibrant) |
| `⮕ result` | `any` | `-` |   |   |   | The vibrance-adjusted image |
---

### `vignette(img=- strength=0.5 falloff=0.8) ⮕ (result=-)`  
_Adds a vignette effect (darkens/lightens edges)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `img` | `*image.NRGBA64` | `-` |   |   |   | The image to apply vignette to |
| `strength` | `float64` | `0.5` | `0` | `1` |   | Darkness/Lightness intensity (0 to 1) |
| `falloff` | `float64` | `0.8` | `0.1` | `2` |   | How quickly the effect fades (0.1 to 2.0) |
| `⮕ result` | `any` | `-` |   |   |   | The image with vignette effect |
---

### `xyz(x=0 y=0 z=0 alpha=1) ⮕ (result=-)`  
_Creates a color from CIE XYZ values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `x` | `float64` | `0` | `0` | `0.95047` | `&#34;&#34;` | The X component (red) |
| `y` | `float64` | `0` | `0` | `1` | `&#34;&#34;` | The Y component (green) |
| `z` | `float64` | `0` | `0` | `1.08883` | `&#34;&#34;` | The Z component (blue) |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `ycbcr(y=0 cb=128 cr=128 alpha=1) ⮕ (result=-)`  
_Creates a color from YCbCr values (digital video)_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `y` | `float64` | `0` | `0` | `255` | `&#34;&#34;` | The luminance component |
| `cb` | `float64` | `128` | `0` | `255` | `&#34;&#34;` | The blue-difference chroma component |
| `cr` | `float64` | `128` | `0` | `255` | `&#34;&#34;` | The red-difference chroma component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

### `yuv(y=0 u=0 v=0 alpha=1) ⮕ (result=-)`  
_Creates a color from YUV values_

| Name | Type | Default | Min | Max | Unit | Description |
|------|------|---------|-----|-----|------|-------------|
| `y` | `float64` | `0` | `0` | `1` | `&#34;&#34;` | The luminance component |
| `u` | `float64` | `0` | `-0.436` | `0.436` | `&#34;&#34;` | The U chrominance component |
| `v` | `float64` | `0` | `-0.615` | `0.615` | `&#34;&#34;` | The V chrominance component |
| `alpha` | `float64` | `1` | `0` | `1` | `&#34;%&#34;` | The color&#39;s alpha |
| `⮕ result` | `any` | `-` |   |   |   | The color as color.RGBA64 |
---

 
 