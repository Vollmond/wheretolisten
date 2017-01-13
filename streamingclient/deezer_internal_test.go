package streamingclient

import (
	"net/http"
	"strings"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestBuildDeezerQuery(t *testing.T) {
	album := buildAlbum()
	url := album.buildDeezerQuery()
	if !strings.Contains(url, "album%3A%22Armageddon+Gigolo%22") {
		t.Errorf("url %s missing album param", url)
	}

	if !strings.Contains(url, "artist%3A%22Spiritual+Front%22") {
		t.Errorf("url %s missing artist param", url)
	}
}

func TestFindLink(t *testing.T) {
	album := buildAlbum()
	r, err := recorder.New("fixtures/deezer/existing")
	if err != nil {
		t.Error(err.Error())
	}
	defer r.Stop()

	client := &http.Client{Transport: r}
	resp, err := client.Get(album.buildDeezerQuery())
	if err != nil {
		t.Error(err.Error())
	}

	link := parseDeezerResponse(resp).findLink()

	if link != "https://www.deezer.com/album/9323488" {
		t.Errorf("Link supposed to be 9323488 but its %s", link)
	}
}
