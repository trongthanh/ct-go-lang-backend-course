// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	ad_listing "thanhtran-s02-2/clients/ad-listing"
)

func main() {
	// TODO #5 setup output for logger to write it to a file ✅
	logger := log.Default()

	file, err:= os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err == nil {
		logger.SetOutput(file)
	}
	defer file.Close()

	c := ad_listing.NewClient(
		ad_listing.WithBaseUrl(ad_listing.BaseUrl),
		ad_listing.WithRetryTimes(3),
		ad_listing.WithLogger(logger),
	)

	getAdListing := func(cat string) {
		ads, err := c.GetAdByCate(context.TODO(), cat)
		if err != nil {
			panic("GetAdByCate " + err.Error())
		}
		logger.Printf("Number of Ads: %v", ads.Total)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		getAdListing(ad_listing.CateVeh)
		fmt.Println("getAdListing done cate veh")
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		getAdListing(ad_listing.CatePty)
		fmt.Println("getAdListing done cate pty")
		wg.Done()
	}()

	wg.Wait()
}


