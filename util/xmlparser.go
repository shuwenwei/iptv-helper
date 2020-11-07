package util

import (
	"encoding/xml"
	"fmt"
)

type VideoUrlWrapper struct {
	XMLName xml.Name `xml:"LoadBalancing"`
	Result bool `xml:"Result"`
	Protocol string `xml:"Protocol"`
	StreamURL string `xml:"StreamURL"`
	Msg string `xml:"Msg"`
}

func ParseXmlToUrl(respPlain *[]byte) string {
	videoInfo := new(VideoUrlWrapper)
	err := xml.Unmarshal(*respPlain, &videoInfo)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	wrappedStreamURL := videoInfo.StreamURL
	return wrappedStreamURL
}
