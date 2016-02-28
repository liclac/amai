package ffxiv

import (
	"fmt"
	"strconv"
	"strings"
	"regexp"
)

type ConfusedByMarkupError string

func (e ConfusedByMarkupError) Error() string {
	return string(e)
}

func normalizeServerName(name string) string {
	name = strings.TrimSpace(name)
	return strings.TrimSuffix(strings.TrimPrefix(name, "("), ")")
}

func parseEorzeanDate(s string) (sun int, moon int, err error) {
	eorzeanDateRegex := regexp.MustCompile(`(\d+)(?:st|nd|rd|th) sun of the (\d+)(?:st|nd|rd|th) (astral|umbral) moon`)
	matches := eorzeanDateRegex.FindStringSubmatch(strings.ToLower(s))
	
	if len(matches) != 4 {
		return 0, 0, ConfusedByMarkupError(fmt.Sprintf("Can't parse Eorzean date: %s", s))
	}
	
	sun, err = strconv.Atoi(matches[1])
	if err != nil {
		return 0, 0, err
	}
	
	moon, err = strconv.Atoi(matches[2])
	if err != nil {
		return 0, 0, err
	}
	
	moon = moon + (moon - 1)
	if matches[3] == "umbral" {
		moon += 1
	}
	
	return sun, moon, nil
}
