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

func (album Album) FindLinks() chan string {
	links := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)

	go addLink(&links, album.FindDeezer(), &wg)
	go addLink(&links, album.FindYandex(), &wg)
	wg.Wait()
	return links
}

func addLink(links *chan string, link string, wg *sync.WaitGroup) {
	defer wg.Done()
	if link != "" {
		*links <- link
	}
}
