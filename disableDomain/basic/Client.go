package main

import (
	"fmt"
	"openApi-authentication/api/client"
	"openApi-authentication/common/auth"
	"os"
)

func main() {

	disableSingleDomainServiceRequest := client.DisableSingleDomainServiceRequest{}
	domainidset := os.Getenv("domainid")
	apiuri := fmt.Sprintf("/api/domain/%s/disable", domainidset)
	config := auth.BasicConfig{UserName: "{raypro}", ApiKey: "{98514TVVFKeNJ6dfQ8Ho6xBn5OTvph}", EndPoint: "api.cdnetworks.com", Uri: apiuri, Method: "PUT"}
	response := auth.Invoke(config, disableSingleDomainServiceRequest.String())
	fmt.Printf("response body is %#v\n", response)
}
