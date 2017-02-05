package streamingclient

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type deezerApiData struct {
	Data []struct {
		Link string `json:"link"`
	} `json:"data"`
}

func (album Album) FindDeezer() string {
	rawData, err := http.Get(album.buildDeezerQuery())
	if err != nil {
		log.Fatal(err)
		return ""
	}
	defer rawData.Body.Close()
	return parseDeezerResponse(rawData).findLink()
}

func (apiAlbum deezerApiData) findLink() string {
	if len(apiAlbum.Data) == 0 {
		return ""
	}
	return apiAlbum.Data[0].Link
}

// TODO: remove double code
func parseDeezerResponse(resp *http.Response) *deezerApiData {
	data := new(deezerApiData)
	json.NewDecoder(resp.Body).Decode(&data)
	return data
}

func (album Album) buildDeezerQuery() string {
	apiUrl, _ := url.Parse("https://api.deezer.com/search/album")
	params := url.Values{}
	params.Add("q", fmt.Sprintf("artist:\"%s\" album:\"%s\"", album.ArtistName, album.Name))
	apiUrl.RawQuery = params.Encode()
	return apiUrl.String()
}
