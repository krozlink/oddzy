package main

import (
	"fmt"
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"log"
	"time"
)

const (
	meetingDataURL = "https://www.odds.com.au/api/web/public/Racing/getUpcomingRaces/?sport=%s&date=%s"
	raceDataURL    = "https://www.odds.com.au/api/web/public/Odds/getOddsComparisonCacheable/?eventId=%s&includeTAB=1&includeOdds=1&arrangeOdds=0&betType=FixedWin&includeTote=true&allowGet=true"
)

type meetingData struct {
	meetings    []meeting.Meeting
	races       []race
	selections  []selection
	competitors []competitor
}

func main() {

	intData := readInternalMeetingData()
	extData := scrapeExternalMeetingData()
	eventIds := getMissingEvents(intData, extData)
	scrapeMissingRaceData(eventIds)

	startRaceUpdater()

	srv := micro.NewService(
		micro.Name("oddzy.services.race-scraper"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterScraperServiceHandler(srv.Server(), &service{})
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func readInternalMeetingData() {
	// read all meetings, races, selections and competitors for yesterday, today, tomorrow, and the following day
}

func scrapeMissingRaceData() {

}

func scrapeExternalMeetingData() {
	log.Println("Scraping meeting data")

	for i, url := range meetingUrls() {
		if i > 0 {
			<-time.After(1 * time.Second)
			scrapeMeeting(url)
		}
	}
}

func startRaceUpdater() {

}

func scrapeMeeting(url string) {
	// scrape the url
}

func meetingUrls() []string {
	dates := []string{
		time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 0).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 1).Format("2006-01-02"),
		time.Now().AddDate(0, 0, 2).Format("2006-01-02"),
	}

	types := []string{"horse-racing", "greyhounds", "harness"}

	urls := make([]string, 0, 12)

	for _, d := range dates {
		for _, t := range types {
			u := fmt.Sprintf(meetingDataURL, t, d)
			urls = append(urls, u)
		}
	}

	return urls
}
