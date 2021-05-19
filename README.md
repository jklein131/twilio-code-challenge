### Twilio Sample API for delivery routing ### 

## Overview ## 
When an order is picked up, we may need to notify the customer that their order may be delayed. This is a CLI wrapper around a service that would recieve requests as part of the pickup process, and would notify the user if their delivery was going to be late. 

## Assumptions ## 
- I assumed a fixed rate of travel of 1 mile per minute. 

## Design Decisions ## 
Based on the design of this process, I designed the interfaces so that that it would be easy to translate into a web service. I also designed it using as much pure golang as possible, with the intent of making it easy to run, and show my knowledge of simple go features over libraries. The interfaces are meant to be exchangeable. In the test, we can exchange them for Mock, and in production we might also want to change the implimentation of for example our delivery estimation time service to use google maps instead. By using interfaces, we can easily switch those out. 

## Future Improvments ## 
- Dockerizing the code so that it can run without go installed. 
- Adding support for libphonenumber to validate phone numbers. 
- Convert into an HTTP service. 
- Adding additional parsing checks to make sure that bad data is not passed. 

## Usage ## 
```
go run . -o 123 -plat 27 -plong -80 -dlat 28 -dlong -79  -t "2020-05-19T15:45:00Z" -c 858-XXX-XXXX
```
