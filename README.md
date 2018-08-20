# Oddzy

Oddzy is a partial implementation of bookmakers racing pages using live racing data. 


## Why?

I wanted a small project to use while learning Go and microservices. [Matched betting](https://en.wikipedia.org/wiki/Matched_betting) has been a hobby of mine for a couple of years so I decided to try and recreate something that i'm very familiar with.

<Image placeholder>


## Tech Stack
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


## Microservices

These are written in Go using the go-micro framework and Consul for service discovery.

#### Racing
Backend service managing racing entites such as Meetings, Races and Selections. Uses a MongoDB database for storage.

#### Race Scraper
Backend service which periodically polls an external data source for the latest data on upcoming races.

The frequency that a race is polled is determined by how close it is to starting. A race that is about to start is polled every 30 seconds, while a race that is days away might only be polled every hour. Only racing data is scraped (race status, results, scratchings etc), not pricing data.

#### Racing API
API service called by the front end. Currently only has two methods:
* RaceCard - retrieves all of the information needed to display a single race
* Schedule - Returns the data needed to display a race schedule

#### Price Updater
A very basic Socket IO server written in Node.js for generating test pricing data. 


## Logging

The microservices use Elasticsearch, Logstash, Kibana for logging. These are deployed to their own containers on ECS.
Kibana can be externally via internal.oddzy.xyz/kibana although this is password protected.

## Metrics
The microservices use Prometheus, StatsD and Grafana for tracking metrics. All services have their timings and success/failure tracked as well as a few other useful metrics like the number of race updates scraped. 

## Deployment
For development the entire application can be run locally using docker compose.

For remote deployment terraform is used to create the stack on AWS. This includes:
* Route53 records
* Network resources - VPC, subnets, security groups etc
* Application load balancer and auto scaling group
* ECS cluster and tasks
* EFS volume for persistant storage
* Optional EC2 jumpbox for remote debugging

Both local and remote environments can be created with a single command in the makefile

