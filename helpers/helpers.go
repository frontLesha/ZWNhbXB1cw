package helpers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"testreq/models"

	"github.com/pejman-hkh/gdp/gdp"
)

func GetUserProperties(client *http.Client) models.UserProperties {
	var res interface{}
	var result models.UserProperties
	resp, err := client.Get("https://ecampus.ncfu.ru/studies")
	if err != nil {
		log.Fatal(err)
	}
	page, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc := gdp.Default(string(page))
	el := doc.Find("[type='text/javascript']")
	elStr := el.Html()
	formatEl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(elStr, `<script type="text/javascript">`, ""), "</script>", ""), "window.iams.coreRequire(function(core){", ""), "$.extend(core.res, ", "")
	formatEl = strings.ReplaceAll(strings.Split(formatEl, "var viewModel = ")[1], ";", "")
	err = json.Unmarshal([]byte(formatEl), &res)
	if err != nil {
		log.Fatal(err)
	}
	userId := res.(map[string]interface{})["specialities"].([]interface{})[0].(map[string]interface{})["Id"].(float64)
	userIdStr := fmt.Sprintf("%0.f", userId)
	result.UserId = userIdStr
	terms := res.(map[string]interface{})["specialities"].([]interface{})[0].(map[string]interface{})["AcademicYears"].([]interface{})
	for _, v := range terms {
		formatV := v.(map[string]interface{})["Terms"].([]interface{})
		for _, sem := range formatV {
			fSem := sem.(map[string]interface{})
			term := models.Term{TermId: fmt.Sprintf("%0.f", fSem["Id"].(float64)), TermNum: fSem["Name"].(string)}
			result.Terms = append(result.Terms, term)
		}
	}
	result.RaspID, result.TargetType = getRaspId(client)
	return result
}

func getRaspId(client *http.Client) (string, string) {
	var res interface{}
	resp, err := client.Get("https://ecampus.ncfu.ru/schedule/my/student")
	if err != nil {
		log.Fatal(err)
	}
	page, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc := gdp.Default(string(page))
	el := doc.Find("[type='text/javascript']")
	elStr := el.Html()
	formatEl := strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(elStr, `<script type="text/javascript">`, ""), "</script>", ""), "window.iams.coreRequire(function(core){", ""), "$.extend(core.res, ", ""), "JSON.parse", ""), `"\"`, `"`), "(", ""), ")", "")
	formatEl = strings.ReplaceAll(strings.Split(formatEl, "var viewModel = ")[1], ";", "")
	err = json.Unmarshal([]byte(formatEl), &res)
	if err != nil {
		log.Fatal(err)
	}
	raspId := res.(map[string]interface{})["Model"].(map[string]interface{})["Id"].(float64)
	targetType := res.(map[string]interface{})["Model"].(map[string]interface{})["Type"].(float64)
	raspIdStr := fmt.Sprintf("%0.f", raspId)
	targetTypeStr := fmt.Sprintf("%0.f", targetType)
	return raspIdStr, targetTypeStr
}
