package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func main() {

	var eg errgroup.Group

	for idx := 0; idx < 1050; idx++ {
		i := idx
		eg.Go(func() error {
			resp, err := http.Get("http://localhost:8080")
			if err != nil {
				return err
			}
			defer resp.Body.Close()
			byteArray, err := ioutil.ReadAll(resp.Body)
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
