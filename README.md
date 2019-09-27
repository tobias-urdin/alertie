# WARNING! Alertie was never finished

# Alertie
## What is Alertie?
Alerting system with an API service to manage resources and trigger alerts.

Alertie is written in Go.

## Goals
Goal is to to be an alert endpoints for your alarms with support for already know systems/API:s such as:
* PagerDuty
* VictorOps
* Grafana webhook

Alertie will then be in charge of recieving your alarms and sending out the alerts to the proper place/person using the proper method.

All parts of Alertie should have scalability and high availability in mind.

## Design
* MySQL server for storing resources and alerts.
* nsq.io message queue system, nsqd + nsqlookup daemons.
* API service to manage all resources and trigger alerts, alert endpoints queues a message in NSQ, API service is horizontally scalable.
* Worker service consumes messages from NSQ and makes decisions on who to contact and how and sends out the alerts. Worker service is horizontally scalable.

nsq.io was chosen for it's possibility to scale thus making every single part of Alertie scalable with a high availbility design model.

## Usage
Create your config file, see the example in the config section below.

`make depend-update`
`make build`

Starting the api:
`./bin/alertie-api -config alertie.ini`

Starting the worker:
`./bin/alertie-worker -config alertie.ini`

## Config

```
[database]
# MySQL connection string
connection = myuser:mypassword@tcp(127.0.0.1:3306)/alertie?parseTime=true

[nsq]
# NSQ lookups, can be multiple with a colon separated list
lookups = localhost:4161
#lookups = 10.0.0.11:4161,10.0.0.12:4161
```

## TODO
* Finish PoC
* AngularJS/React or similar front-end web interface
* Determine exact things that Alertie should support and what it should not support
* PagerDuty API support
* VictorOps API support
* Grafana webhook support
* Write scripts/agents for monitoring systems to use Alertie to alerting
* More alerters
* Metric support such as statsd for alerting metrics

## Contributions
Any contributions are welcome.

## Makefile

Build
`make build`

Install deps with glide
`make depend`

Update deps with glide
`make depend-update`

Run tests
`make test`

Cleanup
`make clean`

Check formatting
`make fmt`

Auto fix format
`make fmtfix`
