package house

import "strings"

const testVersion = 1

// Verse takes 2 strings, []string and returns a string with verses of the song
func Verse(subject string, relPhrase []string, nounPhrase string) string {

	if len(relPhrase) <= 0 {
		return embed(subject, nounPhrase)
	}
	return Verse(embed(subject, relPhrase[0]), relPhrase[1:], nounPhrase)
}

// Song takes no arguments, and returns a string with all verses
func Song() string {
	firstLine := "This is"
	lastLine := "the house that Jack built."
	var verses = make([]string, 0)
	for i := 0; i <= len(lyrics); i++ {
		verse := Verse(firstLine, reverse(lyrics[:i]), lastLine)
		verses = append(verses, verse)
	}

	return strings.Join(verses, "\n\n")
}

func embed(relPhrase, nounPhrase string) string {
	return relPhrase + " " + nounPhrase
}

func reverse(items []string) []string {
	l := len(items)
	var reversed = make([]string, l)
	for i, v := range items {
		reversed[l-i-1] = v
	}
	return reversed
}

var lyrics = []string{
	"the malt\nthat lay in",
	"the rat\nthat ate",
	"the cat\nthat killed",
	"the dog\nthat worried",
	"the cow with the crumpled horn\nthat tossed",
	"the maiden all forlorn\nthat milked",
	"the man all tattered and torn\nthat kissed",
	"the priest all shaven and shorn\nthat married",
	"the rooster that crowed in the morn\nthat woke",
	"the farmer sowing his corn\nthat kept",
	"the horse and the hound and the horn\nthat belonged to",
}
