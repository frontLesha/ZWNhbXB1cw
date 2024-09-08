package auth

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"

	"github.com/pejman-hkh/gdp/gdp"
)

func getCaptchaRequest(client *http.Client) ([]byte, []*http.Cookie) {
	res, err := client.Get("https://ecampus.ncfu.ru/Captcha/Captcha")
	if err != nil {
		panic(err)
	}
	ontent, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return ontent, res.Cookies()
}

func loginRequest(client *http.Client, token, login, pswd string, code string) []*http.Cookie {
	var param = url.Values{}
	param.Set("__RequestVerificationToken", token)
	param.Set("Login", login)
	param.Set("Password", pswd)
	param.Set("Code", code)
	param.Set("RememberMe", "true")
	var payload = bytes.NewBufferString(param.Encode())
	req, err := http.NewRequest("POST", "https://ecampus.ncfu.ru/Account/Login", payload)
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	return res.Cookies()
}

func getLoginPage() (string, []*http.Cookie) {
	res, err := http.Get("https://ecampus.ncfu.ru/account/login")
	if err != nil {
		log.Fatal(err)
	}
	content, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	return string(content), res.Cookies()
}

func saveSession(sessions map[string]string) {
	sessionsByte, err := json.Marshal(sessions)
	if err != nil {
		panic(err)
	}
	os.WriteFile("sesions.json", sessionsByte, 0644)
}
func readSessions() map[string]string {
	var sessionsPool = make(map[string]string)
	sessionsByte, err := os.ReadFile("sesions.json")
	if err != nil {
		return sessionsPool
	}
	err = json.Unmarshal(sessionsByte, &sessionsPool)
	if err != nil {
		panic(err)
	}
	return sessionsPool
}

func Login(login, password string) *http.Client {
	u, err := url.Parse("https://ecampus.ncfu.ru/")
	if err != nil {
		panic(err)
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client := http.Client{
		Jar: jar,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	sessionsPool := readSessions()
	if sessionsPool[login] == "" {
		pageHtml, cookies := getLoginPage()
		client.Jar.SetCookies(u, cookies)

		doc := gdp.Default(pageHtml)
		el := doc.Find("[name='__RequestVerificationToken']")
		elStr := el.Html()
		authToken := strings.Split(strings.Split(elStr, `value="`)[1], `"`)[0]

		capcha, captchaCookies := getCaptchaRequest(&client)
		os.WriteFile("captcha.png", capcha, 0644)
		code := ReadInTerminal("enter captcha ")
		client.Jar.SetCookies(u, captchaCookies)
		cook := loginRequest(&client, authToken, login, password, code)
		sessionsPool[login] = cook[0].Value
		saveSession(sessionsPool)
	}
	authCookie := http.Cookie{Name: "ecampus", Value: sessionsPool[login]}
	client.Jar.SetCookies(u, []*http.Cookie{&authCookie})
	return &client
}

func ReadInTerminal(startText string) string {
	fmt.Print("\n" + startText + ": ")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(text, "\n", ""), "\r", ""), " ", "")
}
