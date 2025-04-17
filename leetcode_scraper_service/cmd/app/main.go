package main

import (
	"fmt"

	"github.com/lamned/leetcode-scraper-service/internal/lc"
)


func main() {
	slug_channel := lc.Get_Slugs()

	for slug := range slug_channel {
		fmt.Println(slug.Slug)
	}
}
