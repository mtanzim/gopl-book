package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"sync"
)

var count int
var mu sync.Mutex

type LisParams struct {
	cycles  float64
	res     float64
	size    float64
	nframes int
	delay   int
}

func NewLisParams() LisParams {
	params := LisParams{
		cycles:  5,     // number of complete x oscillator revolutions
		res:     0.001, // angular resolution
		size:    100,   // image canvas covers [-size..+size]
		nframes: 64,    // number of animation frames
		delay:   8,     // delay between frames in 10ms units
	}
	return params
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	http.HandleFunc("/lis", func(rw http.ResponseWriter, r *http.Request) {

		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		lissajous(rw, NewLisParams())

	})
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, _ *http.Request) {
	mu.Lock()
	fmt.Fprintf(w, "Count = %d\n", count)
	mu.Unlock()
}

var palette = []color.Color{color.White, color.Black}

const (
	whiteIndex = 0 // first color in palette
	blackIndex = 1 // next color in palette
)

func lissajous(out io.Writer, params LisParams) {
	cycles := params.cycles   // number of complete x oscillator revolutions
	res := params.res         // angular resolution
	size := params.size       // image canvas covers [-size..+size]
	nframes := params.nframes // number of animation frames
	delay := params.delay     // delay between frames in 10ms units

	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, int(2*size+1), int(2*size+1))
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(int(size)+int(x*size+0.5), int(size)+int(y*size+0.5),
				blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
