package scrap

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/itpourya/queuingBot/internal/serializers"
)

var SEARCH_API = "https://apigw.paziresh24.com/seapi/v1/search/zahedan/ophthalmology"

func GetDoctorDetail() {
	var responseSerializer serializers.Root

	req, err := http.NewRequest("GET", SEARCH_API, nil)
	if err != nil {
		log.Fatal("Faild Create Request", err.Error())
	}

	client := http.Client{}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/130.0.6723.70 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Faild Send Request", err.Error())
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Println("Faild Read Request Body", err.Error())
	}

	err = json.Unmarshal(body, &responseSerializer)
	if err != nil {
		log.Println("Warning Unmarshal Json Response", err.Error())
	}

	for _, doctor := range responseSerializer.Search.Result {
		fmt.Println(doctor.Slug)
	}
}
