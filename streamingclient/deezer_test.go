package streamingclient_test

import (
	"net/http"
	"testing"

	"github.com/dnaeon/go-vcr/recorder"
	"github.com/vollmond/wheretolisten/streamingclient"
)

func TestFindDeezer(t *testing.T) {
	album := streamingclient.Album{ArtistName: "Spiritual Front", Name: "Armageddon Gigolo"}

	r, _ := recorder.New("fixtures/deezer/existing")
	defer r.Stop()

	defaultClient := http.DefaultClient
	http.DefaultClient = &http.Client{Transport: r}

	link := album.FindDeezer()

	if link != "https://www.deezer.com/album/9323488" {
		t.Errorf("link supposed to be https://www.deezer.com/album/9323488 but its %s", link)
	}

	http.DefaultClient = defaultClient
}

func TestFindDeezerNotFound(t *testing.T) {
	album := streamingclient.Album{ArtistName: "Non Existing Garbage", Name: "Armageddon Gigolo"}

	r, _ := recorder.New("fixtures/deezer/nonexisting")
	defer r.Stop()

	defaultClient := http.DefaultClient
	http.DefaultClient = &http.Client{Transport: r}

	link := album.FindDeezer()

	if link != "" {
		t.Error("What")
	}

	http.DefaultClient = defaultClient
}
