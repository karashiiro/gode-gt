package processing

import (
	"github.com/karashiiro/godestone-gt/internal/site"
	"github.com/xivapi/godestone/v2/provider/models"
)

// ClassJobTable represents the reconstructed ClassJob table.
type ClassJobTable struct {
	ClassJobs []models.NamedEntity
}

// BuildClassJobTable synthesizes core data stores in each language
// into a reconstructed ClassJob table.
func BuildClassJobTable(en *site.CoreData, de *site.CoreData, fr *site.CoreData, ja *site.CoreData) *ClassJobTable {
	classJobs := make([]models.NamedEntity, len(en.Jobs))
	for i := 0; i < len(classJobs); i++ {
		enInfo := en.Jobs[i]
		deInfo := de.Jobs[i]
		frInfo := fr.Jobs[i]
		jaInfo := ja.Jobs[i]

		classJobs[i] = models.NamedEntity{
			ID:   enInfo.ID,
			Name: enInfo.Name,

			NameEN: enInfo.Name,
			NameDE: deInfo.Name,
			NameFR: frInfo.Name,
			NameJA: jaInfo.Name,
		}
	}

	return &ClassJobTable{
		ClassJobs: classJobs,
	}
}
