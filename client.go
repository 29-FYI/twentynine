package twentynine

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	TwentyNine       = 29
	TwentyNineApiUrl = "https://api.29.fyi"
)

func PostLink(link Link) (err error) {
	buf := new(bytes.Buffer)
	if err = json.NewEncoder(buf).Encode(link); err != nil {
		return
	}

	r, err := http.Post(TwentyNineApiUrl, "application/json", buf)
	if err != nil {
		return
	}

	if r.StatusCode != http.StatusCreated {
		msgBuf, _ := ioutil.ReadAll(r.Body)
		err = Error{
			Code:    r.StatusCode,
			Message: string(msgBuf),
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&link)
	return
}

func GetLinks() (links []Link, err error) {
	r, err := http.Get(TwentyNineApiUrl)
	if err != nil {
		return
	}

	if r.StatusCode != http.StatusOK {
		msgBuf, _ := ioutil.ReadAll(r.Body)
		err = Error{
			Code:    r.StatusCode,
			Message: string(msgBuf),
		}
		return
	}

	err = json.NewDecoder(r.Body).Decode(&links)
	return
}
