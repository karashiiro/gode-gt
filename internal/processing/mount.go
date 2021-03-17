package processing

import "github.com/xivapi/godestone/v2/provider/models"

// MountTable represents the reconstructed mount table.
type MountTable struct {
	Mounts []models.NamedEntity
}
