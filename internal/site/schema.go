package site

// AchievementIndex represents the structure of the achievement
// index, reduced to only the fields we need.
type AchievementIndex struct {
	Browse []struct {
		ID   uint32 `json:"i"`
		Name string `json:"n"`
	} `json:"browse"`
}
