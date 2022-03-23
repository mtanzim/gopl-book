package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

//!+main
type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}

func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byYear []*Track

func (x byYear) Len() int           { return len(x) }
func (x byYear) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byYear) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type byTitle []*Track

func (x byTitle) Len() int           { return len(x) }
func (x byTitle) Less(i, j int) bool { return x[i].Title < x[j].Title }
func (x byTitle) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

func main() {

	fmt.Println("\nReverse(byArtist):")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	fmt.Println("\nbyYear:")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom:")
	sortCustom(tracks)

	fmt.Println("byArtist:")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom Dynamic: Artist")
	sortCustomDynamic(tracks, false, false, false, true)
	printTracks(tracks)

	fmt.Println("byTitle:")
	sort.Sort(byTitle(tracks))
	printTracks(tracks)

	fmt.Println("\nCustom Dynamic: Title")
	sortCustomDynamic(tracks, true, false, false, false)
	printTracks(tracks)

	fmt.Println("\nCustom:")
	sortCustom(tracks)
	printTracks(tracks)

	fmt.Println("\nCustom Dynamic: Title, Year, Length")
	sortCustomDynamic(tracks, true, true, true, false)
	printTracks(tracks)

}

// TODO: doesn't really work
func sortCustomDynamic(tracks []*Track, sortTitle, sortYear, sortLength, sortArtist bool) {

	sort.Sort(customSort{tracks, func(x, y *Track) bool {

		return (sortTitle && x.Title != y.Title && x.Title < y.Title) || (sortYear && x.Year != y.Year && x.Year < y.Year) || (sortLength && x.Length != y.Length && x.Length < y.Length) || (sortArtist && x.Artist != y.Artist && x.Artist < y.Artist)
	}})
}

func sortCustom(tracks []*Track) {
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
