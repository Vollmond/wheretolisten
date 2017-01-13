package streamingclient_test

import (
	"net/http"
	"testing"

	"github.com/vollmond/wheretolisten/streamingclient"

	"github.com/dnaeon/go-vcr/recorder"
)

func TestFind(t *testing.T) {
	album := streamingclient.Album{ArtistName: "Spiritual Front", Name: "Armageddon Gigolo"}
	r, _ := recorder.New("fixtures/ya/found")
	defer r.Stop()

	defaultClient := http.DefaultClient
	http.DefaultClient = &http.Client{Transport: r}
	url := album.FindYandex()

	if url != "https://music.yandex.ru/album/276007" {
		t.Errorf("Wrong url, %s", url)
	}
	http.DefaultClient = defaultClient
}

func TestNotFound(t *testing.T) {
	album := streamingclient.Album{ArtistName: "iddqd", Name: "idkfa"}

	r, _ := recorder.New("fixtures/ya/nonexisting")
	defer r.Stop()

	defaultClient := http.DefaultClient
	http.DefaultClient = &http.Client{Transport: r}
	url := album.FindYandex()

	if url != "" {
		t.Errorf("Wrong url, %s", url)
	}
	http.DefaultClient = defaultClient
}
