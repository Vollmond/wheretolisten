package streamingclient

import "sync"

type Album struct {
	ArtistName string
	Name       string
}

type AlbumFinder interface {
	FindDeezer() string
	FindYandex() string
	FindLinks() chan string
}

func (album Album) Links() []string {
	links := make([]string, 0, 5)
	var wg sync.WaitGroup
	wg.Add(2)

	go addLink(links, album.FindYandex(), &wg)
	go addLink(links, album.FindDeezer(), &wg)
	wg.Wait()
	return links
}

func addLink(links []string, link string, wg *sync.WaitGroup) {
	defer wg.Done()
	if link != "" {
		links = append(links, link)
	}
}
