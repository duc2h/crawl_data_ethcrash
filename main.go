package main

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/hoangduc02011998/crawl_ethcrash/database"
)

func getValueFromEthCrashIO(db *sql.DB, job <-chan int, results chan<- int) {

	// for i := 0; i < 117000; i++ {
	for j := range job {
		var id = 2130785
		var link = "https://www.ethercrash.io/game/"
		id += j

		link = link + strconv.Itoa(id)
		doc, _ := goquery.NewDocument(link)

		doc.Find(".pt-3").Each(func(i int, s *goquery.Selection) {
			valueS := between(s.Find(".lb-title").Text(), "@", "x")
			date_betS := between(s.Find(".text-muted").Text(), "on ", "GMT+0200")

			value, _ := strconv.ParseFloat(valueS, 0)

			dateaa := strings.Split(date_betS, " ")
			timee := dateaa[2] + " " + dateaa[1] + " " + dateaa[3] + " " + dateaa[4]
			date_time, _ := time.Parse("2 Jan 2006 15:04:05", timee)

			//Lưu dữ liệu bài viết vào DB
			insPost, _ := db.Prepare("INSERT INTO crawl_2129900 (value, date_bet, id_bet) VALUES(?, ?, ?)")
			insPost.Exec(value, date_time, id)

		})

		results <- j
	}
	//}

	fmt.Println("Finished")
}

func between(value string, a string, b string) string {
	// Get substring between two strings.
	posFirst := strings.Index(value, a)
	if posFirst == -1 {
		return ""
	}
	posLast := strings.Index(value, b)
	if posLast == -1 {
		return ""
	}
	posFirstAdjusted := posFirst + len(a)
	if posFirstAdjusted >= posLast {
		return ""
	}
	return value[posFirstAdjusted:posLast]
}

func before(value string, a string) string {
	// Get substring before a string.
	pos := strings.Index(value, a)
	if pos == -1 {
		return ""
	}
	return value[0:pos]
}

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}

func main() {
	db := database.DBConn()

	// Lưu dữ liệu bài viết vào DB
	//insPost, _ := db.Prepare("INSERT INTO crawl_ethcrash (value, date_bet) VALUES(?, ?)")

	//insPost.Exec(1.43, time.Now())
	//fmt.Printf("Inserted: %q\n", content)
	//getValueFromEthCrashIO(db)
	//fmt.Println(db)

	const numJobs = 110000
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	for w := 1; w <= 2; w++ {
		go getValueFromEthCrashIO(db, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
}
