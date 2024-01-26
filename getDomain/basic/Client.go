package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
)

type DomainSummary struct {
	DomainID         int    `xml:"domain-id"`
	DomainName       string `xml:"domain-name"`
	ServiceType      int    `xml:"service-type"`
	CName            string `xml:"cname"`
	Status           string `xml:"status"`
	CdnServiceStatus bool   `xml:"cdn-service-status"`
	Enabled          bool   `xml:"enabled"`
}

type DomainList struct {
	XMLName         xml.Name        `xml:"domain-list"`
	DomainSummaries []DomainSummary `xml:"domain-summary"`
}

func main() {

	queryApiDomainListServiceRequest := client.QueryApiDomainListServiceRequest{}

	config := auth.BasicConfig{UserName: "raypro", ApiKey: "98514TVVFKeNJ6dfQ8Ho6xBn5OTvph", EndPoint: "api.cdnetworks.com", Uri: "/api/domain", Method: "GET"}
	response := auth.Invoke(config, queryApiDomainListServiceRequest.String())

	xmlData := response
	var domainList DomainList
	err := xml.Unmarshal([]byte(xmlData), &domainList)
	if err != nil {
		fmt.Println("解析 XML 失败:", err)
		return
	}

	// 将结构体转换为 JSON
	jsonData, err := json.Marshal(domainList)
	if err != nil {
		fmt.Println("转换为 JSON 失败:", err)
		return
	}

	// 打印 JSON 数据
	fmt.Println(string(jsonData))

}
