package race

import (
	"strconv"
	"strings"

	"github.com/ironarachne/world/pkg/words"
)

// SimplifiedRace is a simplified version of a race for display
type SimplifiedRace struct {
	Name        string `json:"name"`
	Adjective   string `json:"adjective"`
	Description string `json:"description"`
}

// Simplify returns a simplified version of a race
func (race Race) Simplify() SimplifiedRace {
	return SimplifiedRace{
		Name:        race.Name,
		Adjective:   race.Adjective,
		Description: race.Describe(),
	}
}

// Describe returns a description of a race
func (race Race) Describe() string {
	description := ""

	description += strings.Title(race.PluralName) + " are " + race.SizeCategory.Name + ", with men averaging " + heightToString(race.getAverageHeight("male")) + " tall and "
	description += "weighing " + strconv.Itoa(race.getAverageWeight("male")) + " pounds, and women averaging " + heightToString(race.getAverageHeight("female")) + " tall and "
	description += "weighing " + strconv.Itoa(race.getAverageWeight("female")) + " pounds. "

	if len(race.Appearance.UniqueTraits) > 0 {
		uniqueTraits := words.CombinePhrases(race.Appearance.UniqueTraits)
		description += strings.Title(race.PluralName) + " " + uniqueTraits + ". "
	}

	return description
}
