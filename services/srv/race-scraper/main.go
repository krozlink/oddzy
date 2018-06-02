package main

import (
	proto "github.com/krozlink/oddzy/services/srv/race-scraper/proto"
	// racing "github.com/krozlink/oddzy/services/srv/racing/proto"
	micro "github.com/micro/go-micro"
	_ "github.com/micro/go-plugins/registry/consul"
	"log"
)

var process scrapeProcess

type scrapeProcess struct {
	status string
	done   chan bool
	http   handler
}

func main() {

	process = newScrapeProcess()
	registerProcessMonitor()

	start()

	<-process.done
}

func newScrapeProcess() scrapeProcess {
	return scrapeProcess{
		status: "INITIALISING",
		done:   make(chan bool),
		http:   handler{},
	}
}

func registerProcessMonitor() {
	monitor := newMonitor(&process)

	srv := micro.NewService(
		micro.Name("oddzy.services.race-scraper"),
		micro.Version("latest"),
	)

	srv.Init()

	proto.RegisterMonitorServiceHandler(srv.Server(), monitor)

	go func() {
		if err := srv.Run(); err != nil {
			log.Fatalln(err)
		}
	}()
}

func start() {
	// STATUS - SETUP

	// read all internal meeting data for scraping period (yesterday to 2 days from now)
	//		ListMeetingsByDate
	//		ListRacesByMeetingDate

	// read race calendars from scraping period
	//		ScrapeCalendar

	// for each external meeting
	// 		if the meeting doesnt exist internally (look up source id), create meeting id and add to new meetings list

	// for each race
	// 		if the race doesnt exist internally (look up source id) create a race id and add to new races list

	// update new meeting list with their race ids
	// call AddMeetings with new meetings list
	// call AddRaces with  new races list (leave last_updated as null)

	// Add all existing races with a null last_updated as well as all new races to a missingEvents queue

	// STATUS - RACE_CREATION
	// For each item on the missingEvents queue
	//	Scrape the race
	// 	Update the race - UpdateRace

	// STATUS - RACE MONITORING
	// Process 1
	//		Monitor time until race and race status and push required races to queue if they are not already on there
	//		Only needs to run at most every 30 seconds. It can determine its own sleep duration depending on upcoming races
	// Process 2 - Periodically take items off the queue and scrape them. Ensure minimum interval to avoid overloading server
}
