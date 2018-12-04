package parser

import (
	"io/ioutil"
	"net/http"
	"strings"
)

var titleTags = []string{"h1", "title"}
var contentTags = []string{"body"}

func ParseFromUrl(url string) (string, string, []string) {

	var message = []string{}
	var title = ""
	var content = ""

	response, err := http.Get(url)

	if err != nil {
		message = append(message, err.Error())
		return title, content, message
	}

	defer response.Body.Close()

	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	articleHasTitle, title := getTagInnerText(titleTags, pageContent)
	articleHasContent, content := getTagInnerText(contentTags, pageContent)

	if articleHasTitle == false {
		message = append(message, "could not read title")
	}

	if articleHasContent == false {
		message = append(message, "could not read content")
	}

	return title, content, message

}

func getTagInnerText(tags []string, pageContent string) (bool, string) {

	var isTagMatched = false
	var matchedTagContent = ""
	for _, tag := range tags {

		formattedTagStart := formatStartTag(tag)
		formattedTagEnd := formatEndTag(tag)

		if isTagExists(formattedTagStart, pageContent) == true {

			isTagMatched = true

			tagStartIndex := strings.Index(pageContent, formattedTagStart)
			tagEndIndex := strings.Index(pageContent, formattedTagEnd)
			splitted := strings.Split(pageContent[tagStartIndex:tagEndIndex], ">")

			slice := splitted[1:len(splitted)]
			matchedTagContent = strings.Join(slice, ">")
			break
		}

	}

	return isTagMatched, matchedTagContent
}

func formatStartTag(tag string) string {
	return "<" + tag
}

func formatEndTag(tag string) string {
	return "</" + tag + ">"
}

func isTagExists(tag string, content string) bool {
	contentLower := strings.ToLower(content)
	return strings.Contains(contentLower, tag)
}
