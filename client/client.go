package main

import (
	"fmt"
	"log"

	"golang.org/x/sync/errgroup"
)

func main() {

	var eg errgroup.Group

	for idx := 0; idx < 1050; idx++ {
		i := idx
		eg.Go(func() error {
			if i > 1024 {
				log.Println(fmt.Sprintf("request number :%d", i))
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		panic(err)
	}

}
