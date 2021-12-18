package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "time"
   "strconv"
   "bytes"
   "encoding/json"

   "github.com/go-hexa/go-balance/internal/core"
)

func main() {

	client := http.Client{}

	//---------------------------------------------------
	log.Println("start LoadBalance data")
	for i:=0; i < 500; i++ {

		acc := "acc-" + strconv.Itoa(i)
		description := "description-"+ strconv.Itoa(i)
		balance := &core.Balance{
			Id:    strconv.Itoa(i),
			Account: acc,
			Amount: 1,
			DateBalance: time.Now(),
			Description: description,
		}

		payload := new(bytes.Buffer)
		json.NewEncoder(payload).Encode(balance)

		req_post , err := http.NewRequest("POST", "http://localhost:8901/add_balance", payload)
		if err != nil {
			log.Fatalln(err)
		}
		req_post.Header = http.Header{
			"Accept_Language": []string{"pt-BR"},
			"Authorization": []string{"Bearer cookie"},
			"Content-Type": []string{"application/json"},
		}

		_, err = client.Do(req_post)
		if err != nil {
		   log.Fatalln(err)
		}
	}
	log.Println("end LoadBalance data")
	//---------------------------------------------------

	log.Println("-----------------------------")
	req_get , err := http.NewRequest("GET", "http://localhost:8901/list_balance", nil)
	if err != nil {
    	log.Fatalln(err)
	}
	req_get.Header = http.Header{
		"Accept_Language": []string{"pt-BR"},
		"Authorization": []string{"Bearer cookie"},
	}

	for i:=0; i < 3600; i++ {
		resp, err := client.Do(req_get)
		defer resp.Body.Close()

		if err != nil {
		   log.Fatalln(err)
		}
	
		body, err := ioutil.ReadAll(resp.Body)
	
		if err != nil {
		   log.Fatalln(err)
		}

		sb := string(body)
		log.Printf(sb)
		log.Println("###########")

		time.Sleep(time.Second * time.Duration(1))
	}
	log.Println("-----------------------------")
}