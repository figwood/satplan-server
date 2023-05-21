# satplan-server
Satplan-server is a backend application written in Go that provides APIs for mission planning. The following techniques are utilized:
* gin
* gorm

At present, only rectangular areas are supported. However, rtree in sqlite has the potential to support other types of areas.

## Database
Satplan-server share databases with [calpath](https://github.com/figwh/calpath)

## APIs
API|type|description
----|----|----
/sattree|GET|get all the satellites and sensors
/satplan|POST|misson planning
/sat/all|GET|get all satellites
/sat/add|POST|create new satellite
/sat/:id|GET|get satellite by id
/sat/update/:id|PUT|update satellite
/sat/:id|DELETE|delete satellite
/sat/tle/update|PUT|update tles
/sat/tle/cal|POST|recalculation data
/sen/all|GET|get all sensors
/sen/add|GET|get all sensors
/sen/bysat|GET|get sensors by satellite
/sen/:id|GET|get sensor by id
/sen/update/:id|PUT|update sensor
/sen/:id|DELETE|delete sensor

For further information, please refer to syscfg/router.go.

## Auto update TLEs
The Satplan-server will update the Two-Line Elements (TLEs) and recalculate data every day at 00:00 AM UTC.

## Build and run
### Prerequisites
Ensure that you have installed all the necessary prerequisites on your development machine:
* Go 1.15+

Execute the following commands to build and run.
```bash
go mod tidy
CGO_ENABLED=1 go build 
./satplan-server
```
The server will be launched at http://localhost:8080.
