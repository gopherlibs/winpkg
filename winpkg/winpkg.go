package winpkg

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
)

// The root XML element returned from the Chocolatey Search API.
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Entries []Entry  `xml:"entry"`
}

// The Chocolatey Search API can return zero or more `entry` items.
// Each entry represents a package.
type Entry struct {
	XMLName xml.Name `xml:"entry"`
	Title   string   `xml:"title"`
	Version string   `xml:"properties>Version"`
}

func GetVersion(pkg string) (string, error) {

	resp, err := http.Get("https://community.chocolatey.org/api/v2/Search()?$filter=IsLatestVersion&$skip=0&$top=1&searchTerm=%27" + pkg + "%27&targetFramework=%27%27&includePrerelease=false")
	if err != nil {
		return "", errors.New("Failed to reach Chocolately API.")
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.New("Failed to read in data from API response.")
	}

	var searchResult Feed

	err = xml.Unmarshal(data, &searchResult)
	if err != nil {
		return "", errors.New("Failed to unmarshal API response.")
	}

	return searchResult.Entries[0].Version, nil
}
