package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {

	singleRequest()

}

func singleRequest() {
	httpClient := &http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			TLSHandshakeTimeout:   10 * time.Second,
			ResponseHeaderTimeout: 10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
		Timeout: 3 * time.Second,
	}
	resp, err := httpClient.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("resp", resp)
		panic(err)
	}
	defer resp.Body.Close()
	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	log.Println(string(byteArray))
}

func bulkRequest() {
	var eg errgroup.Group

	for idx := 0; idx < 1050; idx++ {
		i := idx
		eg.Go(func() error {
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			byteArray, err := io.ReadAll(resp.Body)
			if err != nil {
				return err
			}
			log.Println(fmt.Sprintf("request number :%d, content: %s", i, string(byteArray)))
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		panic(err)
	}
}
