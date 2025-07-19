package main

import (
	"fmt"

	"github.com/lamned/leetcode-scraper-service/internal/lc"
)


func main() {
	client := lc.CreateLeetcodeScraper()
	slug_channel := client.Get_Slugs()

	for slug := range slug_channel {
		fmt.Println(slug.Slug)
	}
}
