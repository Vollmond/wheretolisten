package streamingclient

import (
	"net/http"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func buildAlbum() Album {
	return Album{ArtistName: "Spiritual Front", Name: "Armageddon Gigolo"}
}

func TestBuildYandexQuery(t *testing.T) {
	album := buildAlbum()
	url := album.buildYandexQuery()
	if !strings.Contains(url, "text=Armageddon+Gigolo") {
		t.Errorf("url %s missing text param", url)
	}

	if !strings.Contains(url, "page=0") {
		t.Errorf("url %s missing text param", url)
	}

	if !strings.Contains(url, "type=album") {
		t.Errorf("url %s missing text param", url)
	}
}

func TestParseYandexResponse(t *testing.T) {
	album := buildAlbum()
	r, _ := recorder.New("fixtures/ya/existing")
	defer r.Stop()
	client := &http.Client{Transport: r}
	resp, err := client.Get(album.buildYandexQuery())
	if err != nil {
		t.Fail()
	}
	parsedResponse := parseYandexResponse(resp)

	if parsedResponse.Result.Albums.Results[0].Title != "Armageddon Gigolo" {
		t.Errorf("title supposed to be AG but its %s", parsedResponse.Result.Albums.Results[0].Title)
	}

	artistName := parsedResponse.Result.Albums.Results[0].Artists[0].Name

	t.Log(parsedResponse.Result.Albums.Results[0].Artists)
	if artistName != "Spiritual Front" {
		t.Errorf("artist supposed to be SF but its %s", artistName)
	}
}

func TestFindAlbumId(t *testing.T) {
	album := buildAlbum()
	r, _ := recorder.New("fixtures/ya/existing")
	defer r.Stop()

	client := &http.Client{Transport: r}
	resp, err := client.Get(album.buildYandexQuery())
	if err != nil {
		t.Fail()
	}

	id := parseYandexResponse(resp).findAlbumId(album)

	if id != 276007 {
		t.Errorf("Id supposed to be 276007 but its %d", id)
	}
}
