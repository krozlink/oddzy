package main

import (
	"encoding/json"
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
	meetings    []racing.Meeting
	races       []racing.Race
	selections  []racing.Selection
	competitors []racing.Competitor
}

var (
	status   string
	meetings map[string]string
	done     = make(chan bool)
)

func main() {
	registerService()

	addMissingRaceData()

	monitorUpcomingRaces()

	<-done
}

func registerService() {
	status = "INITIALISING"
	srv := micro.NewService(
		micro.Name("oddzy.services.race-scraper"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterScraperServiceHandler(srv.Server(), &service{})

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalln(err)
		}
	}()
}

func addMissingRaceData() {
	status = "SETUP"

	// intData := readInternalMeetingData()
	// extData := scrapeExternalMeetingData()
	// eventIds := getMissingEvents(intData, extData)
}

func readInternalMeetingData() {
	// ListMeetingsByDate(date_range)
	// ListRacesByMeetingDate(date_range)
}

func scrapeExternalMeetingData() <-chan bool {

	result := make(chan bool)
	go func() {
		log.Println("Scraping meeting data")

		for i, url := range calendarUrls() {
			if i > 0 {
				<-time.After(1 * time.Second)
				scrapeCalendar(url)
			}
		}

		result <- true
	}()

	return result
}

func startRaceUpdater() {

}

func scrapeCalendar(url string) (*raceCalendar, error) {
	encodedResponse := getResponse(url)
	odds := oddsResponse{}
	json.Unmarshal(encodedResponse, odds)
	calendar := &raceCalendar{}
	err := json.Unmarshal([]byte(odds.r), calendar)

	if err != nil {
		fmt.Println("Unable to decode response into race calendar")
		fmt.Println(odds.r)
		return nil, err
	}

	return calendar, nil
}

func calendarUrls() []string {
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

func monitorUpcomingRaces() {
	status = "RACE_MONITORING"

}
