package godestonegt

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/karashiiro/godestone-gt/internal/processing"
	"github.com/karashiiro/godestone-gt/internal/site"
	"github.com/xivapi/godestone/v2/provider"
	"github.com/xivapi/godestone/v2/provider/models"
)

const baseUrl string = "https://www.garlandtools.org"

var languages = []string{"en", "de", "fr", "ja"}

type GarlandToolsProvider struct {
	client *http.Client
}

// New constructs a Garland Tools-backed data provider for use
// in a godestone parser.
func New() *GarlandToolsProvider {
	gt := GarlandToolsProvider{
		client: &http.Client{
			Timeout: time.Second * 30,
		},
	}

	return &gt
}

// WithHttpClient replaces the default-constructed HTTP client
// with a user-provided one.
func (gt *GarlandToolsProvider) WithHttpClient(client *http.Client) *GarlandToolsProvider {
	gt.client = client
	return gt
}

// AsProvider reinterprets the provided service as a DataProvider.
func (gt *GarlandToolsProvider) AsProvider() provider.DataProvider {
	return gt
}

func (gt *GarlandToolsProvider) request(url string) ([]byte, error) {
	res, err := gt.client.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (gt *GarlandToolsProvider) Achievement(name string) *models.NamedEntity {
	urlFmt := baseUrl + site.AchievementIndexPath

	indices := make(map[string]*site.AchievementIndex, 4)
	for _, lang := range languages {
		data, err := gt.request(fmt.Sprintf(urlFmt, lang))
		if err != nil {
			return nil
		}

		index := &site.AchievementIndex{}
		err = json.Unmarshal(data, index)
		if err != nil {
			return nil
		}

		indices[lang] = index
	}

	table := processing.BuildAchievementsTable(indices["en"], indices["de"], indices["fr"], indices["ja"])
	for _, entry := range table.Achievements {
		if entry.Name == name {
			return &entry
		}
	}

	return nil
}

func (gt *GarlandToolsProvider) ClassJob(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Deity(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) GrandCompany(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Item(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Minion(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Mount(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Race(name string) *models.GenderedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Reputation(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Title(name string) *models.TitleInternal {
	return nil
}

func (gt *GarlandToolsProvider) Town(name string) *models.NamedEntity {
	return nil
}

func (gt *GarlandToolsProvider) Tribe(name string) *models.GenderedEntity {
	return nil
}
