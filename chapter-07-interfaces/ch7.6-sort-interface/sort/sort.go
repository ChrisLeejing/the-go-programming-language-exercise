package main

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	fmt.Println("==============================byArtist==================================")
	sort.Sort(byArtist(tracks))
	printTracks(tracks)

	fmt.Println("==============================byArtist Reverse==================================")
	reverse := sort.Reverse(byArtist(tracks))
	sort.Sort(reverse)
	printTracks(tracks)

	fmt.Println("==============================byYear==================================")
	sort.Sort(byYear(tracks))
	printTracks(tracks)

	fmt.Println("==============================byYear Reverse==================================")
	reverse = sort.Reverse(byYear(tracks))
	sort.Sort(reverse)
	printTracks(tracks)

	fmt.Println("==============================Title Year Length==================================")
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
	printTracks(tracks)
	// ==============================byArtist==================================
	// Title       Artist          Album              Year     Length
	// -------     -------         -------            -------  -------
	// Go Ahead    Alicia Keys     As I Am            2007     4m36s
	// Go          Delilah         From the Roots Up  2012     3m38s
	// Ready 2 Go  Martin Solveig  Smash              2011     4m24s
	// Go          Moby            Moby               1992     3m37s
	// ==============================byArtist Reverse==================================
	// Title       Artist          Album              Year     Length
	// -------     -------         -------            -------  -------
	// Go          Moby            Moby               1992     3m37s
	// Ready 2 Go  Martin Solveig  Smash              2011     4m24s
	// Go          Delilah         From the Roots Up  2012     3m38s
	// Go Ahead    Alicia Keys     As I Am            2007     4m36s
	// ==============================byYear==================================
	// Title       Artist          Album              Year     Length
	// -------     -------         -------            -------  -------
	// Go          Moby            Moby               1992     3m37s
	// Go Ahead    Alicia Keys     As I Am            2007     4m36s
	// Ready 2 Go  Martin Solveig  Smash              2011     4m24s
	// Go          Delilah         From the Roots Up  2012     3m38s
	// ==============================byYear Reverse==================================
	// Title       Artist          Album              Year     Length
	// -------     -------         -------            -------  -------
	// Go          Delilah         From the Roots Up  2012     3m38s
	// Ready 2 Go  Martin Solveig  Smash              2011     4m24s
	// Go Ahead    Alicia Keys     As I Am            2007     4m36s
	// Go          Moby            Moby               1992     3m37s
	// ==============================Title Year Length==================================
	// Title       Artist          Album              Year     Length
	// -------     -------         -------            -------  -------
	// Go          Moby            Moby               1992     3m37s
	// Go          Delilah         From the Roots Up  2012     3m38s
	// Go Ahead    Alicia Keys     As I Am            2007     4m36s
	// Ready 2 Go  Martin Solveig  Smash              2011     4m24s
	//
	// Process finished with exit code 0
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

func length(s string) time.Duration {
	duration, err := time.ParseDuration(s) // 解析: 1 year = 365 days
	if err != nil {
		panic(err)
	}
	return duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func printTracks(tracks []*Track) {
	const format string = "%v\t%v\t%v\t%v\t%v\t\n"
	w := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(w, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(w, format, "-------", "-------", "-------", "-------", "-------")
	for _, t := range tracks {
		fmt.Fprintf(w, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	w.Flush()
}

type byArtist []*Track

func (a byArtist) Len() int {
	return len(a)
}

func (a byArtist) Less(i, j int) bool {
	return a[i].Artist < a[j].Artist
}

func (a byArtist) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

type byYear []*Track

func (y byYear) Len() int {
	return len(y)
}

func (y byYear) Less(i, j int) bool {
	return y[i].Year < y[j].Year
}

func (y byYear) Swap(i, j int) {
	y[i], y[j] = y[j], y[i]
}

// 自定义排序结构体: Title -> Year -> Length
type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (c customSort) Len() int {
	return len(c.t)
}

func (c customSort) Less(i, j int) bool {
	return c.less(c.t[i], c.t[j])
}

func (c customSort) Swap(i, j int) {
	c.t[i], c.t[j] = c.t[j], c.t[i]
}
