package main

import (
   "io/ioutil"
   "log"
   "net/http"
   "time"
   "bytes"
   "encoding/json"
   "flag"
   "os"

   "github.com/go-hexa/go-balance/dummy_data"
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

	var host = "a13602eb146cf4786af99f26bc7a9db8-2126738867.us-east-2.elb.amazonaws.com" + ":" + *port 
	//var host = "127.0.0.1:" + *port
	
	client := http.Client{}
	done := make(chan string)

	log.Println("-----------------------------")
	log.Println("Goroutine - POST data")
	go post(host, client, done)
	log.Println("End POST data")
	log.Println("-----------------------------")

	log.Println("-----------------------------")
	log.Println("Goroutine - GET Data")
	go get(host, client, done)
	log.Println("End GET Data")
	log.Println("-----------------------------")

	log.Println(<-done)
}

func get(host string, client http.Client, done chan string){
	for i:=0; i < 3600; i++ {
		host_url := "http://" + host + "/list_balance"
		get_data(host_url, client)
		time.Sleep(time.Millisecond * time.Duration(1000))
	}
	done <- "END"
}

func post(host string, client http.Client, done chan string){
	for a:=0; a < 3600; a++ {
		for i:=0; i < 50; i++ {
			host_url := "http://" + host + "/add_balance"
			post_data(i, host_url, client)
			time.Sleep(time.Millisecond * time.Duration(150))
		}
		time.Sleep(time.Millisecond * time.Duration(1000))
	}
	done <- "END"
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
		"jwt": []string{"cookie"},
	}

	req_get.Close = true
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

	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("Erro Close :",err)
		}
	}()

	log.Println("###########")
}

func post_data(i int ,host string, client http.Client){
	log.Println("POST DATA.....")

	balance := dummy_data.NewBalance(i)
	payload := new(bytes.Buffer)
	json.NewEncoder(payload).Encode(balance)

	req_post , err := http.NewRequest("POST", host, payload)
	if err != nil {
		log.Println("Error http.NewRequest : ", err)
		panic(err)
	}

	req_post.Header = http.Header{
		"Accept_Language": []string{"pt-BR"},
		"jwt": []string{"cookie"},
		"Content-Type": []string{"application/json"},
	}

	resp, err := client.Do(req_post)
	defer resp.Body.Close()
	if err != nil {
	   log.Println("Error doing POST : ", err)
	   panic(err)
	}

	log.Println("StatusCode : ", resp.StatusCode )
	
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Println("Erro Close :",err)
		}
	}()

}
