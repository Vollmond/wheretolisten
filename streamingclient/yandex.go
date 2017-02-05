package streamingclient

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

type yandexApiData struct {
	Result struct {
		Albums struct {
			Results []struct {
				ID      int    `json:"id"`
				Title   string `json:"title"`
				Artists []struct {
					Name string `json:"name"`
				} `json:"artists"`
			} `json:"results"`
		} `json:"albums"`
	} `json:"result"`
}

func (album Album) FindYandex() string {
	response, _ := http.Get(album.buildYandexQuery())
	albumID := parseYandexResponse(response).findAlbumId(album)
	var url string
	if albumID != 0 {
		// fmt.Println(albumID)
		url = fmt.Sprintf("https://music.yandex.ru/album/%d", albumID)
	}
	return url
}

func (apiData *yandexApiData) findAlbumId(album Album) int {
	albumsList := apiData.Result.Albums.Results
	var albumId int
	for i := 0; i < len(albumsList); i++ {
		if strings.EqualFold(albumsList[i].Title, album.Name) && strings.EqualFold(albumsList[i].Artists[0].Name, album.ArtistName) {
			albumId = albumsList[i].ID
		}
	}
	return albumId
}

func parseYandexResponse(resp *http.Response) *yandexApiData {
	data := new(yandexApiData)
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func (album Album) buildYandexQuery() string {
	apiUrl, _ := url.Parse("https://api.music.yandex.net/search")
	params := url.Values{}
	params.Add("page", "0")
	params.Add("text", album.Name)
	params.Add("type", "album")
	apiUrl.RawQuery = params.Encode()
	return apiUrl.String()
}
