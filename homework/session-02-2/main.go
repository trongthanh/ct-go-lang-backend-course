// You can edit this code!
// Click here and start typing.
package main

import (
	"context"
	"log"
	ad_listing "thanhtran-s02-2/clients/ad-listing"
)

func main() {

	// TODO #5 setup output for logger to write it to a file
	logger := log.Default()

	c := ad_listing.NewClient(ad_listing.BaseUrl, 3, logger)

	ads, err := c.GetAdByCate(context.TODO(), ad_listing.CatePty)
	if err != nil {
		panic("GetAdByCate " + err.Error())
	}

	logger.Printf("Number of Ads: %v", ads.Total)
}
