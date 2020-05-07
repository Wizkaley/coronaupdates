package utils

import (
	"io/ioutil"
	"log"
	"net/http"
)

// GetData gets the data from the api
func GetData() (body []byte) {

	url := "https://coronavirus-tracker-india-covid-19.p.rapidapi.com/api/getStatewise"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "coronavirus-tracker-india-covid-19.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "c00a03fa2dmsh6bd2d46a085b6bap1cbcaejsn2db4322306d3")

	res, err := http.DefaultClient.Do(req)

	CheckError(err)

	defer res.Body.Close()
	body, err = ioutil.ReadAll(res.Body)
	CheckError(err)

	//fmt.Println(res)
	//fmt.Println(string(body))

	return body
}

// CheckError utility for checking error
func CheckError(err error) {
	if err != nil {
		log.Print("Error Occured::", err.Error())
		return
	}
}
