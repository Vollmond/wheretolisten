package streamingclient

type Album struct {
	ArtistName string
	Name       string
}

// func (album Album) Links() []string {
// 	links := make([]string, 0, 5)
// 	var wg sync.WaitGroup
// 	wg.Add(2)
//
// 	go addLink(links, album.FindYandex(), &wg)
// 	go addLink(links, album.FindDeezer(), &wg)
// 	wg.Wait()
// 	return links
// }
//
// func addLink(links []string, link string, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	if link != "" {
// 		links = append(links, link)
// 	}
// }

func (album Album) Links() []string {
	links := make(chan string, 2)

	go func() { links <- album.FindYandex() }()
	go func() { links <- album.FindDeezer() }()

	resultLinks := make([]string, 0, 5)

	counter := 0
	for counter < 2 {
		link := <-links
		if link != "" {
			resultLinks = append(resultLinks, link)
		}
		counter++
	}

	return resultLinks
}
