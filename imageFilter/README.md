# Image Filter tool

A CLI tool for applying various filters to images.
Works for both png and jpeg, but always produces a png for the filtered image.

## Installation

After downloading the source code, create a builds directory and then use make build to create a build version.

## Usage

You can use the following flags to use the program:

- -s (source): specify the path to the image you want to apply a filter to
- -f (filter): specify which filter you want to use. Following options are available:
  - boxBlur
  - gaussianBlur
  - edge
  - spot
  - invert
  - comic
  - heat
  - sort (VERY SLOW!)
  - pixel (Experimental)
  - basicKuwahara
  - generalKuwahara (Experimental)
- -t (threads): specify the number of threads that should be used when applying the filter. Omit to use maximum available threads.
- -h (help): shows a help message
- -c (convert): utility flag, used to create a png from a given jpeg image

## Filters

- blur: there are two different blur filters:
  - boxBlur: simple blur, breaks on bigger kernel sizes
  - gaussianBlur: slightly more expensive algorithm, but overall smoother
- edge: an edge filter, which weighs direct neighboring pixels twice as much as diagonal neighbors
- spot: creates a spot effect at the center of the image
- invert: simple color inversion
- comic: clamps intensities to a predefined palette of three colors
- heat: replicates a heat camera image effect
- sort: sorts pixels in each row, based on their intesity. Currently very slow, did not implement this in a smart way
- pixel: simple pixelation filter, which unfortunatly still produces black lines when running the filter concurrent
- basicKuwahara: the kuwahara filter gives an image an artsy look. Works best on realistic images
- generalKuwahara: a more sophisticated version, which I didn't manage to properly implemented yet

## Example

```bash
builds/imageFilter -s testImage.png -f gaussianBlur -t 12
```
