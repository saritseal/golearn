package parsers

import (
	"encoding/xml"
	"strings"

	"github.com/sirupsen/logrus"
)

// ParseXML This is a common method to parse any xml string
func ParseXML(xmlString string) map[string]interface{} {

	var result map[string]interface{}

	var decoder = xml.NewDecoder(strings.NewReader(xmlString))

	for {
		t, _ := decoder.Token()

		if t == nil {
			break
		}

		logrus.Println(t)

	}
	return result
}
