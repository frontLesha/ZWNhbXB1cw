package zachetka

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/pejman-hkh/gdp/gdp"
)

func GetZachetka(client *http.Client) {
	resp, err := client.Get("https://ecampus.ncfu.ru/details/zachetka")
	if err != nil {
		log.Fatal(err)
	}
	page, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc := gdp.Default(string(page))
	el := doc.Find("[type='text/javascript']")
	elHTML := el.Html()
	formatEl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.Split(elHTML, "var viewModel = ")[1], "</script>", ""), ";", ""), "JSON.parse", ""), `""`, `"`), "(", ""), ")", ""), `Модуль "`, "")
	re := regexp.MustCompile(`\s+`)

	resultData := re.ReplaceAllString(formatEl, " ")
	fmt.Println(resultData)
}
