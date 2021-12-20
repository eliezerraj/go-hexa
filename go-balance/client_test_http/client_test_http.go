package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "time"
   "strconv"
   "bytes"
   "encoding/json"
   "flag"
   "os"

   "github.com/go-hexa/go-balance/internal/core"
)

func main() {

	port := flag.String("port","","")
	flag.Parse()
	flag.VisitAll(func (f *flag.Flag) {
		if f.Value.String()=="" {
			log.Printf("A flag -%v não foi informado \n", f.Name )
			os.Exit(1)
		}
	})

	var host = "127.0.0.1:" + *port
	host_url := "http://" + host + "/add_balance"
	
	client := http.Client{}

	//---------------------------------------------------
	log.Println("start LoadBalance data")
	for i:=0; i < 100; i++ {

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

		req_post , err := http.NewRequest("POST", host_url, payload)
		if err != nil {
			log.Println("Error http.NewRequest : ", err)
			panic(err)
		}

		req_post.Header = http.Header{
			"Accept_Language": []string{"pt-BR"},
			"Authorization": []string{"Bearer cookie"},
			"Content-Type": []string{"application/json"},
		}

		_, err = client.Do(req_post)
		if err != nil {
		   log.Println("Error doing POST : ", err)
		   panic(err)
		}
	}
	log.Println("end LoadBalance data")
	//---------------------------------------------------

	log.Println("-----------------------------")
	host_url = "http://" + host + "/list_balance"

	req_get , err := http.NewRequest("GET", host_url, nil)
	if err != nil {
		log.Println("Error http.NewRequest : ", err)
		panic(err)
	}
	req_get.Header = http.Header{
		"Accept_Language": []string{"pt-BR"},
		"Authorization": []string{"Bearer cookie"},
	}

	for i:=0; i < 3600; i++ {
		resp, err := client.Do(req_get)
		defer resp.Body.Close()

		if err != nil {
			log.Println("Error doing GET : ", err)
			panic(err)
		}
	
		body, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Println("Error : ", err)
			panic(err)
		}

		sb := string(body)
		log.Printf(sb)
		log.Println("###########")

		time.Sleep(time.Second * time.Duration(1))
	}
	log.Println("-----------------------------")
}