# Oddzy

Oddzy is a partial implementation of bookmakers racing pages using live racing data. 


## Why?

I wanted a small project to use while learning Go and microservices. [Matched betting](https://en.wikipedia.org/wiki/Matched_betting) has been a hobby of mine for a couple of years so I decided to try and recreate something that i'm very familiar with. It also gave me an opportunity to try a number of products that I hadn't got around to using.


<a href="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/race_card.png"><img src="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/race_card.png" width="400"></a> <a href="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/race_schedule.png"><img src="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/race_schedule.png" width="400"></a>


## Tech Stack
Overkill for a project of this size, but this was done for the experience rather than being practical.
* Go microservices using the [go-micro](https://micro.mu/) framework and toolkit.
* Front end using VueJS and the Bulma framework for styling
* MongoDB for storing racing data
* Consul for service discovery
* Logging using Elasticsearch, Logstash and Kibana
* Metrics using Prometheus and Grafana
* NGINX to serve the website and as a reverse proxy to access Kibana & Grafana
* Terraform for remote deployment, docker compose for local development
* Hosted on AWS using ECS
* Socket.IO for sending live updates to the browser
* NATS as a broker for publishing changes to racing entities


## Front End

The website is a single page application using Vue 2. The CSS framework Bulma is used for styling. 

Most of the data on the Racing schedule and Race pages should be kept up to date depending on the source data. As no pricing data is scraped the values shown are randomly generated placeholder values.

The date and race type (horse racing / harness / greyhounds) filters work. The race countdown timers are updated live and results are displayed when the races complete.

My career to date has been almost entirely back-end development and my goal with this was just to make something that looked reasonable and made use of the API I had created.


## Microservices

These are written in Go using the go-micro framework with Consul used for service discovery.

#### [Racing](/services/srv/racing)
Backend service managing racing entites such as Meetings, Races and Selections. Uses a MongoDB database for storage.

#### [Race Scraper](/services/srv/race-scraper)
Backend service which periodically polls an external data source for the latest data on upcoming races.

The frequency that a race is polled is determined by how close it is to starting. A race that is about to start is polled every 30 seconds, while a race that is days away might only be polled every hour. Only racing data is scraped (race status, results, scratchings etc), not pricing data.


#### [Racing API](/services/api/racing)
API service called by the front end. Currently only has two methods:
* RaceCard - retrieves all of the information needed to display a single race
* Schedule - Returns the data needed to display a race schedule

These can be accessed at \
https://api.example.com/racing/schedule?date=yyyy-mm-dd \
https://api.example.com/racing/racecard?race_id=id

#### [Price Updater](/services/web/price-updater)
A very basic Socket IO server written in Node.js for generating test pricing data. 


## Logging

The microservices use Elasticsearch, Logstash and Kibana for logging. These are deployed to their own containers on ECS.
Kibana can be externally via http://internal.example.com/kibana although this is password protected.

<a href="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/kibana.png"><img src="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/kibana.png" width="400"></a>


## Metrics
The microservices use Prometheus, StatsD and Grafana for tracking metrics. These are deployed to their own containers on ECS.
All services have their timings and success/failure tracked as well as a few other miscellaneous metrics like the frequency that races are scraped. 
Granfana can be externally via http://internal.example.com/grafana although this is password protected.

<a href="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/grafana.png"><img src="https://raw.githubusercontent.com/krozlink/oddzy/master/docs/grafana.png" width="400"></a>

## Deployment
For development the entire application can be run locally using docker compose.

For remote deployment Terraform is used to create the stack on AWS. This includes:
* ECS cluster and tasks
* Route53 records
* Network resources - VPC, subnets, security groups etc
* Application load balancer using a HTTPS listener
* EFS volume for persistant storage
* Optional EC2 jumpbox for remote debugging
* AWS Systems Manager Document used push updates to the website remotely

Both local and remote environments can be created with a single command in the makefile

The terraform files can be found in the [deploy directory](/deploy/terraform)