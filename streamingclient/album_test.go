package streamingclient_test

import (
	"fmt"
	"testing"

	"github.com/vollmond/wheretolisten/streamingclient"
)

type testAlbum struct {
	streamingclient.Album
}

func (album testAlbum) FindYandex() string {
	return "https://music.yandex.ru/album/276007"
}

func (album testAlbum) FindDeezer() string {
	fmt.Print("YAAAAAY")
	return "https://www.deezer.com/album/9323488"
}

func TestFindLinks(t *testing.T) {
	// album := testAlbum{streamingclient.Album{Name: "Spiritual Front", ArtistName: "Armageddon Gigolo"}}
	// links := album.Links()
	// if len(links) == 0 {
	// 	t.Error("links are empty!")
	// }
	// for _, link := range album.Links() {
	// 	if link != album.FindYandex() || link != album.FindDeezer() {
	// 		t.Errorf("wrong link for album - %s", link)
	// 	}
	// }
}
