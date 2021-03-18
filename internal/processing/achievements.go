package processing

import (
	"github.com/karashiiro/godestone-gt/internal/site"
	"github.com/xivapi/godestone/v2/provider/models"
)

// AchievementsTable represents the reconstructed achievements table.
type AchievementsTable struct {
	Achievements []models.NamedEntity
}

// BuildAchievementsTable synthesizes achievement indices in each language
// into a reconstructed achievements table.
func BuildAchievementsTable(en *site.AchievementIndex, de *site.AchievementIndex, fr *site.AchievementIndex, ja *site.AchievementIndex) *AchievementsTable {
	achievements := make([]models.NamedEntity, len(en.Browse))
	for i := 0; i < len(achievements); i++ {
		enInfo := en.Browse[i]
		deInfo := de.Browse[i]
		frInfo := fr.Browse[i]
		jaInfo := ja.Browse[i]

		achievements[i] = models.NamedEntity{
			ID:   enInfo.ID,
			Name: enInfo.Name,

			NameEN: enInfo.Name,
			NameDE: deInfo.Name,
			NameFR: frInfo.Name,
			NameJA: jaInfo.Name,
		}
	}

	return &AchievementsTable{
		Achievements: achievements,
	}
}
