package handle

import (
	"covid_sumary/model"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type CaseData struct {
	Data []model.Case
}

func NewCaseData() *CaseData {
	return &CaseData{}
}
func (cd *CaseData) FetchDataForm(url string) string {
	// Disable security check for a client.
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}
	resp, err := client.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	// fmt.Println(sb)

	return sb
}

func (cd *CaseData) ConvertToModel(s string) *CaseData {
	err := json.Unmarshal([]byte(s), cd)
	if err != nil {
		log.Fatalln(err)
	}
	return cd
}
