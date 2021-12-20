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
	
	client := http.Client{}
	done := make(chan string)

	log.Println("-----------------------------")
	log.Println("Goroutine - Load data")
	go post(host, client, done)
	log.Println("End Load data")
	log.Println("-----------------------------")

	log.Println("-----------------------------")
	log.Println("Goroutine - Reading Data")
	go get(host, client, done)
	log.Println("End Reading Data")
	log.Println("-----------------------------")

	log.Println(<-done)
}

func post(host string, client http.Client, done chan string){
	for i:=0; i < 5; i++ {
		host_url := "http://" + host + "/add_balance"
		post_data(i, host_url, client)
		time.Sleep(time.Second * time.Duration(1))
	}
	done <- "END"
}

func get(host string, client http.Client, done chan string){
	for i:=0; i < 5; i++ {
		host_url := "http://" + host + "/list_balance"
		get_data(host_url, client)
		time.Sleep(time.Second * time.Duration(1))
	}
	done <- "END"
}

func post_data(i int ,host string, client http.Client){
	log.Println("POST DATA.....")
	acc := "acc-" + strconv.Itoa(i)
	description := "description-"+ strconv.Itoa(i) + " - UPDATED"
	balance := &core.Balance{
		Id:    strconv.Itoa(i),
		Account: acc,
		Amount: 1,
		DateBalance: time.Now(),
		Description: description,
	}

	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(balance)

	req_post , err := http.NewRequest("POST", host, payload)
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

func get_data(host_url string, client http.Client){
	log.Println("GET DATA.....")
	req_get , err := http.NewRequest("GET", host_url, nil)
	if err != nil {
		log.Println("Error http.NewRequest : ", err)
		panic(err)
	}
	req_get.Header = http.Header{
		"Accept_Language": []string{"pt-BR"},
		"Authorization": []string{"Bearer cookie"},
	}

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
}