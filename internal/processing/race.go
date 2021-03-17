package processing

import "github.com/xivapi/godestone/v2/provider/models"

// RaceTable represents the reconstructed race table.
type RaceTable struct {
	Races []models.GenderedEntity
}
