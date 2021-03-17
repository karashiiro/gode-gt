package godestonegt

import (
	"encoding/json"
	"errors"
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
const notFound = "no entity matching the criteria was found"

var languages = []string{"en", "de", "fr", "ja"}

type GarlandToolsProvider struct {
	client *http.Client

	achievementsTable *processing.AchievementsTable
	classJobTable     *processing.ClassJobTable
	deityTable        *processing.DeityTable
	grandCompanyTable *processing.GrandCompanyTable
	itemTable         *processing.ItemTable
	minionTable       *processing.MinionTable
	mountTable        *processing.MountTable
	raceTable         *processing.RaceTable
	reputationTable   *processing.ReputationTable
	titleTable        *processing.TitleTable
	townTable         *processing.TownTable
	tribeTable        *processing.TribeTable
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

func (gt *GarlandToolsProvider) getAchievementsTable() (*processing.AchievementsTable, error) {
	if gt.achievementsTable == nil {
		urlFmt := baseUrl + site.AchievementIndexPath

		indices := make(map[string]*site.AchievementIndex, 4)
		for _, lang := range languages {
			data, err := gt.request(fmt.Sprintf(urlFmt, lang))
			if err != nil {
				return nil, err
			}

			index := &site.AchievementIndex{}
			err = json.Unmarshal(data, index)
			if err != nil {
				return nil, err
			}

			indices[lang] = index
		}

		gt.achievementsTable = processing.BuildAchievementsTable(indices["en"], indices["de"], indices["fr"], indices["ja"])
	}
	return gt.achievementsTable, nil
}

func (gt *GarlandToolsProvider) getClassJobTable() (*processing.ClassJobTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getDeityTable() (*processing.DeityTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getGrandCompanyTable() (*processing.GrandCompanyTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getItemTable() (*processing.ItemTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getMinionTable() (*processing.MinionTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getMountTable() (*processing.MountTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getRaceTable() (*processing.RaceTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getReputationTable() (*processing.ReputationTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getTitleTable() (*processing.TitleTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getTownTable() (*processing.TownTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) getTribeTable() (*processing.TribeTable, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Achievement(name string) (*models.NamedEntity, error) {
	table, err := gt.getAchievementsTable()
	if err != nil {
		return nil, err
	}

	for _, entry := range table.Achievements {
		if entry.Name == name {
			return &entry, nil
		}
	}

	return nil, errors.New(notFound)
}

func (gt *GarlandToolsProvider) ClassJob(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Deity(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) GrandCompany(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Item(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Minion(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Mount(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Race(name string) (*models.GenderedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Reputation(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Title(name string) (*models.TitleInternal, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Town(name string) (*models.NamedEntity, error) {
	return nil, nil
}

func (gt *GarlandToolsProvider) Tribe(name string) (*models.GenderedEntity, error) {
	return nil, nil
}
