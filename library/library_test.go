package library

import (
	"testing"
	"time"

	"golang.org/x/text/language"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/orchestrate-io/dvr"
	"github.com/stretchr/testify/assert"
)

func TestSeries(t *testing.T) {
	plots := [...]string{
		`Lorem ipsum dolor sit amet`,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit.`,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Ut est arcu, tempor quis accumsan quis, imperdiet ut ex.
Pellentesque vel lobortis est. Vivamus lobortis eleifend dapibus.
Nullam laoreet ipsum sed massa ornare tristique in nec lorem.
In eleifend odio ac accumsan ultrices. Aenean lacinia vel risus quis mattis.
Donec suscipit pretium euismod. 
Etiam sed justo venenatis, interdum tortor quis, aliquet ipsum.
Vestibulum a facilisis lectus.
Fusce aliquam lectus vel vehicula consequat. 
Aenean venenatis, velit rhoncus scelerisque dictum, 
lorem neque auctor velit, id pretium dui dolor ut ex. Sed quis augue.`,
	}

	languages := Languages{
		NewLanguage(language.MustParseBase("en")),
		NewLanguage(language.MustParseBase("ru")),
	}

	lib, err := New("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	isin, err := lib.HasSeriesWithImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	if isin {
		t.Fatalf("Series already in the database?")
	}

	series, err := lib.GetSeriesByImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	series.Title = "title"
	series.Year = 2048
	series.OtherTitles = map[string]string{
		"foo": "bar",
		"bar": "baz",
	}
	series.Duration = Duration(3 * time.Minute)
	series.Plot = plots[0]
	series.PlotMedium = plots[1]
	series.PlotLong = plots[2]
	series.PosterURL = "http://example.com/foo.jpg"
	series.ImdbRating = 3.14
	series.ImdbVotes = 42
	series.Languages = languages

	err = lib.Save(series)
	if err != nil {
		t.Fatal(err)
	}

	isin, err = lib.HasSeriesWithImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	if !isin {
		t.Fatalf("Series not in the database?")
	}

	series2, err := lib.GetSeriesByImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	assert := assert.New(t)

	assert.Equal(999999, series2.ImdbID)
	assert.Equal("title", series2.Title)
	assert.Equal(2048, series2.Year)
	assert.Equal(
		map[string]string{"foo": "bar", "bar": "baz"},
		map[string]string(series2.OtherTitles),
	)
	assert.Equal(3*time.Minute, time.Duration(series2.Duration))
	assert.Equal(plots[0], series2.Plot)
	assert.Equal(plots[1], series2.PlotMedium)
	assert.Equal(plots[2], series2.PlotLong)
	assert.Equal("http://example.com/foo.jpg", series2.PosterURL)
	assert.InDelta(3.14, series2.ImdbRating, 0.0001)
	assert.Equal(42, series2.ImdbVotes)
	assert.Equal(languages, series2.Languages)
}

func TestShow(t *testing.T) {
	plots := [...]string{
		`Lorem ipsum dolor sit amet`,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit.`,
		`Lorem ipsum dolor sit amet, consectetur adipiscing elit. 
Ut est arcu, tempor quis accumsan quis, imperdiet ut ex.
Pellentesque vel lobortis est. Vivamus lobortis eleifend dapibus.
Nullam laoreet ipsum sed massa ornare tristique in nec lorem.
In eleifend odio ac accumsan ultrices. Aenean lacinia vel risus quis mattis.
Donec suscipit pretium euismod. 
Etiam sed justo venenatis, interdum tortor quis, aliquet ipsum.
Vestibulum a facilisis lectus.
Fusce aliquam lectus vel vehicula consequat. 
Aenean venenatis, velit rhoncus scelerisque dictum, 
lorem neque auctor velit, id pretium dui dolor ut ex. Sed quis augue.`,
	}

	languages := Languages{
		NewLanguage(language.MustParseBase("en")),
		NewLanguage(language.MustParseBase("ru")),
	}

	lib, err := New("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	isin, err := lib.HasShowWithImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	if isin {
		t.Fatalf("Movie already in the database?")
	}

	movie, err := lib.GetShowByImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	movie.Title = "title"
	movie.Year = 2048
	movie.OtherTitles = map[string]string{
		"foo": "bar",
		"bar": "baz",
	}
	movie.Duration = Duration(3 * time.Minute)
	movie.Plot = plots[0]
	movie.PlotMedium = plots[1]
	movie.PlotLong = plots[2]
	movie.PosterURL = "http://example.com/foo.jpg"
	movie.ImdbRating = 3.14
	movie.ImdbVotes = 42
	movie.Languages = languages
	movie.ReleaseDate = time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)

	movie.EpisodeData = nil
	movie.Tagline = "foo!"

	err = lib.Save(movie)
	if err != nil {
		t.Fatal(err)
	}

	isin, err = lib.HasShowWithImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	if !isin {
		t.Fatalf("Movie not in the database?")
	}

	movie2, err := lib.GetShowByImdbID(999999)
	if err != nil {
		t.Fatal(err)
	}

	assert := assert.New(t)

	assert.Equal(999999, movie2.ImdbID)
	assert.Equal("title", movie2.Title)
	assert.Equal(2048, movie2.Year)
	assert.Equal(
		map[string]string{"foo": "bar", "bar": "baz"},
		map[string]string(movie2.OtherTitles),
	)
	assert.Equal(3*time.Minute, time.Duration(movie2.Duration))
	assert.Equal(plots[0], movie2.Plot)
	assert.Equal(plots[1], movie2.PlotMedium)
	assert.Equal(plots[2], movie2.PlotLong)
	assert.Equal("http://example.com/foo.jpg", movie2.PosterURL)
	assert.InDelta(3.14, movie2.ImdbRating, 0.0001)
	assert.Equal(42, movie2.ImdbVotes)
	assert.Equal(languages, movie2.Languages)
	assert.Equal(
		time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC),
		movie2.ReleaseDate,
	)

	assert.Nil(movie2.EpisodeData)
	assert.Equal("foo!", movie2.Tagline)
}

func TestVideoFile(t *testing.T) {
	lib, err := New("sqlite3", ":memory:")
	if err != nil {
		t.Fatal(err)
	}

	isin, err := lib.HasFileWithPath("/foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	if isin {
		t.Fatalf("File already in the database?")
	}

	file, err := lib.GetFileByPath("/foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	file.Size = 98765432
	file.ResolutionX = 1920
	file.ResolutionY = 1080
	file.OsdbHash = 123456789
	file.Format = "h264"
	file.Duration = Duration(time.Minute * 20)

	file.LastPlayed = time.Date(2012, time.February, 10, 23, 15, 32, 5, time.UTC)
	file.LastPosition = Duration(time.Minute*12 + time.Second*38)

	err = lib.Save(file)
	if err != nil {
		t.Fatal(err)
	}

	isin, err = lib.HasFileWithPath("/foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	if !isin {
		t.Fatalf("File not in the database?")
	}

	file2, err := lib.GetFileByPath("/foo/bar")
	if err != nil {
		t.Fatal(err)
	}

	assert := assert.New(t)
	assert.Equal("/foo/bar", file2.Path)
	assert.Equal(uint64(98765432), file2.Size)
	assert.Equal(uint(1920), file2.ResolutionX)
	assert.Equal(uint(1080), file2.ResolutionY)
	assert.Equal(uint64(123456789), file2.OsdbHash)
	assert.Equal("h264", file2.Format)
	assert.Equal(time.Minute*20, time.Duration(file2.Duration))
	assert.Equal(
		file2.LastPlayed,
		time.Date(2012, time.February, 10, 23, 15, 32, 5, time.UTC),
	)
	assert.Equal(
		time.Duration(file2.LastPosition), time.Minute*12+time.Second*38,
	)
}