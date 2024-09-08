package wfi

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/pejman-hkh/gdp/gdp"
)

func GetWiFiInfo(client *http.Client) {
	resp, err := client.Get("https://ecampus.ncfu.ru/DomainAccountInfo")
	if err != nil {
		log.Fatal(err)
	}
	page, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc := gdp.Default(string(page))
	el := doc.Find("[class='form-control-static']")
	elStr := strings.Split(strings.Split(el.Html(), "<b>")[1], "</b>")[0]
	fmt.Println(elStr)

}
