package data

import (
	"encoding/json"
	"fmt"
)

type Exhibition struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

var list = []Exhibition{
	{
		Title:       "Life in Ancient Greek",
		Description: "Uncover the world of ancient Greece through the sculptures, tools, and jewelry found in ruins from over 2000 years ago that have been unearthed through modern science and technology.",
		Image:       "ancient-greece.png",
	},
	{
		Title:       "Aristotle: Life and Legacy",
		Description: "Explore the life and legacy of the great philosopher Aristotle, one of the most influential thinkers of all time. Through rare artifacts and ancient texts, learn about his ideas on ethics, politics, and metaphysics that have shaped the world for centuries.",
		Image:       "aristotle.png",
	},
	{
		Title:       "Chameleon: Colorful Adaptations",
		Description: "Discover the amazing world of chameleons and their incredible ability to change color. Through interactive displays and live chameleon exhibits, learn about the science behind their color changing and how they use it to communicate and camouflage in their environments.",
		Image:       "colorful-adaptations.png",
	},
	{
		Title:       "Sea Monsters: Myth and Reality",
		Description: "Dive into the world of sea monsters and explore the myths and legends that have captured our imaginations for centuries. Through fossils, ancient maps, and interactive displays, discover the truth behind the stories and learn about the real-life creatures that inhabit our oceans.",
		Image:       "sea-monsters.png",
	},
}

func (exhibition Exhibition) String() string {
	json, err := json.MarshalIndent(exhibition, "", "  ")

	if err != nil {
		return fmt.Sprintf("unable to encode exhibition as json %v\n", err)
	}

	return string(json)
}

func GetLatestExhibition() Exhibition {
	latest := len(list)
	return list[latest-1]
}

func GetAllExhibitions() []Exhibition {
	return list
}
