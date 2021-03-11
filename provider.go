package godestonegt

import (
	"io"
	"net/http"
	"time"

	"github.com/xivapi/godestone/v2/provider"
	"github.com/xivapi/godestone/v2/provider/models"
)

const baseUrl string = "https://www.garlandtools.org"

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
