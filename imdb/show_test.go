package imdb

import (
	"fmt"
	"strings"

	_ "github.com/orchestrate-io/dvr"
)

func ExampleShow_ID() {
	movie := New(403358)
	fmt.Printf("id: %d\n", movie.ID())

	// Output:
	// id: 403358
}

func ExampleShow_Title() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	title, err := movie.Title()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("title: %s\n", title)

	// Output:
	// title: Nochnoy dozor
}

func ExampleShow_Title_episode() {
	episode := New(1577257) // Doctor Who s05e02
	defer episode.Free()

	title, err := episode.Title()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("title: %s\n", title)

	// Output:
	// title: The Beast Below
}

func ExampleShow_Title_series() {
	series := New(436992) // Doctor Who
	defer series.Free()

	title, err := series.Title()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("title: %s\n", title)

	// Output:
	// title: Doctor Who
}

func ExampleShow_Year() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	year, err := movie.Year()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("year: %d\n", year)

	// Output:
	// year: 2004
}

func ExampleShow_Year_episode() {
	episode := New(1577257) // Doctor Who s05e02
	defer episode.Free()

	year, err := episode.Year()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("year: %d\n", year)

	// Output:
	// year: 2010
}

func ExampleShow_Year_series() {
	series := New(436992) // Doctor Who
	defer series.Free()

	year, err := series.Year()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("year: %d\n", year)

	// Output:
	// year: 2005
}

func ExampleShow_OtherTitles() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	otherTitles, err := movie.OtherTitles()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	versions := []string{
		"UK",
		"Russia",
		"World-wide (English title)",
	}

	for _, version := range versions {
		fmt.Printf("%s: %s\n", version, otherTitles[version])
	}

	// Output:
	// UK: Night Watch
	// Russia: Ночной дозор
	// World-wide (English title): Night Watch
}

func ExampleShow_ReleaseDate() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	releaseDate, err := movie.ReleaseDate()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("release date: %s\n", releaseDate.Format("2006-01-02"))

	// Output:
	// release date: 2005-10-07
}

func ExampleShow_ReleaseDate_episode() {
	episode := New(1577257) // Doctor Who s05e02
	defer episode.Free()

	releaseDate, err := episode.ReleaseDate()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("release date: %s\n", releaseDate.Format("2006-01-02"))

	// Output:
	// release date: 2010-04-10
}

func ExampleShow_Tagline() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	tagline, err := movie.Tagline()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("tagline: %s\n", tagline)

	// Output:
	// tagline: All That Stands Between Light And Darkness Is The Night Watch.
}

func ExampleShow_Duration() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	duration, err := movie.Duration()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("duration: %s\n", duration)

	// Output:
	// duration: 1h54m0s
}

func ExampleShow_Plot() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	shortPlot, err := movie.Plot()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("short plot: %s\n", shortPlot)

	// Output:
	// short plot: A fantasy-thriller set in present-day Moscow where the respective forces that control daytime and nighttime do battle.
}

func ExampleShow_PlotMedium() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	mediumPlot, err := movie.PlotMedium()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	sentences := strings.Split(mediumPlot, ". ")

	fmt.Printf("medium plot sentence 3:\n%s\n", sentences[2])

	// Output:
	// medium plot sentence 3:
	// Ever since, the forces of light govern the day while the night belongs to their dark opponents
}

func ExampleShow_PlotLong() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	longPlot, err := movie.PlotLong()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	sentences := strings.Split(longPlot, ". ")

	fmt.Printf("long plot sentence 3:\n%s\n", sentences[2])

	// Output:
	// long plot sentence 3:
	// They are known as the Others and have co-existed with humans for as long as humanity has existed
}

func ExampleShow_PosterURL() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	posterURL, err := movie.PosterURL()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("poster url: %s\n", posterURL)

	// Output:
	// poster url: http://ia.media-imdb.com/images/M/MV5BMjE0Nzk0NDkyOV5BMl5BanBnXkFtZTcwMjkzOTkyMQ@@.jpg
}

func ExampleShow_Rating() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	rating, err := movie.Rating()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("rating: %.2g\n", rating)

	// Output:
	// rating: 6.5
}

func ExampleShow_Votes() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()

	votes, err := movie.Votes()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("thousands of votes: %d\n", votes/1000)

	// Output:
	// thousands of votes: 46
}

func ExampleShow_Type() {
	movie := New(403358) // Nochnoy Dozor (2004)
	defer movie.Free()
	episode := New(1577257) // Doctor Who s05e02
	defer episode.Free()
	series := New(436992)
	defer series.Free()

	movieType, err := movie.Type()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	episodeType, err := episode.Type()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	seriesType, err := series.Type()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("Nochnoy Dozor's type is %s\n", movieType)
	fmt.Printf("Doctor Who s05e02's type is %s\n", episodeType)
	fmt.Printf("Doctor Who's type is %s\n", seriesType)

	// Output:
	// Nochnoy Dozor's type is Movie
	// Doctor Who s05e02's type is Episode
	// Doctor Who's type is Series
}

func ExampleShow_SeasonEpisode() {
	episode := New(1577257) // Doctor Who s05e02
	defer episode.Free()

	seasonNumber, episodeNumber, err := episode.SeasonEpisode()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("s%02de%02d\n", seasonNumber, episodeNumber)

	// Output:
	// s05e02
}

func ExampleShow_Series() {
	episode := New(1577257) // Doctor Who s05e02
	defer episode.Free()

	series, err := episode.Series()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("id: %07d\n", series.ID())

	showType, err := series.Type()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("type: %s\n", showType)

	title, err := series.Title()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	fmt.Printf("title: %s\n", title)

	// Output:
	// id: 0436992
	// type: Series
	// title: Doctor Who
}

func ExampleShow_Seasons_testnumbers() {
	series := New(1286039) // Stargate Universe
	defer series.Free()

	seasons, err := series.Seasons()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	fmt.Printf("seasons:\n")
	for i := range seasons {
		number, err := seasons[i].Number()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		fmt.Printf("%d\n", number)
	}

	// Output:
	// seasons:
	// 1
	// 2
}

func ExampleShow_Seasons_episodes() {
	series := New(1286039) // Stargate Universe
	defer series.Free()

	seasons, err := series.Seasons()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	if len(seasons) < 2 {
		fmt.Printf("season 2 is missing :(\n")
	}

	season2 := seasons[1]

	episodes, err := season2.Episodes()
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	for i := range episodes {
		title, err := episodes[i].Title()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		// please note that i != episodeNumber, and possibly i+1 != episodeNumber
		seasonNumber, episodeNumber, err := episodes[i].SeasonEpisode()
		if err != nil {
			fmt.Printf("error: %s\n", err)
			return
		}

		fmt.Printf("s%02de%02d: %s\n", seasonNumber, episodeNumber, title)
	}

	// Output:
	// s02e01: Intervention
	// s02e02: Aftermath
	// s02e03: Awakening
	// s02e04: Pathogen
	// s02e05: Cloverdale
	// s02e06: Trial and Error
	// s02e07: The Greater Good
	// s02e08: Malice
	// s02e09: Visitation
	// s02e10: Resurgence
	// s02e11: Deliverance
	// s02e12: Twin Destinies
	// s02e13: Alliances
	// s02e14: Hope
	// s02e15: Seizure
	// s02e16: The Hunt
	// s02e17: Common Descent
	// s02e18: Epilogue
	// s02e19: Blockade
	// s02e20: Gauntlet
}
