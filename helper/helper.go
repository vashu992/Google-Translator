package helper

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/vashu992/Google-Translator/constant"
	"github.com/vashu992/Google-Translator/model"
)

func Translate(obj model.Translate) ([]byte, error) {
	str2 := ""
	str2 = str2 + "q=" + obj.Text
	str2 = str2 + "&target=" + obj.TargetLanguage
	str2 = str2 + "&source=" + obj.SourceLanguage

	payload := strings.NewReader(str2)

	req, err := http.NewRequest(http.MethodPost, constant.TranslateURL, payload)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept-Encoding", "application/gzip")
	req.Header.Add("X-RapidAPI-Key", os.Getenv("GoogleApi"))
	req.Header.Add("X-RapidAPI-Host", "google-translate1.p.rapidapi.com")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}
