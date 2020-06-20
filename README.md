# Outcomex Tech Challenge

![Demo Image](./example.png)

## Building

### Using Docker

The only important thing is to provide your openweathermap api key as a flag (or environment variable) to the program. You may also specify a custome host for the API using `--api-host=0.0.0.0:12345`

```
docker build -t outcomex .
docker run -it -p 3000:3000 outcomex --api-key="123"
```

You can then view the application in your browser at localhost:3000.

### Testing

For testing make sure you have a openweathermap api key available as an environment variable as `$API_Key` then run `go test ./...`. 

### Docs

## Application Structure

A good entrypoint to get started is with Docker, however you can also use go modd to run the app in development mode and the application will restart as you make changes. 

The internal directory is for go source that could become it's own go module, the code in this folder is completely domain agnostic and could be turned into it's own module if required. 

The cmd directory contains a set of commands to facilitate using this application. Currently the only command is api which runs the server; however, if there were migrations or an admin tool in this project it would also live in this directory. Code in this directory is specific to this application and should not be copied or used in other applications.

The front-end directory contains the source code required to run a browser application. It is mostly boiler plate create react app code. The three main files are in src/ and are `App.js` which houses the main application structure. `City.js` which contains view logic for displaying a city card, and `ForecastImg.js` which contains logic for creating a SVG (image) of a forecast.

Note: City.js is currently doing two things. It is responsible for fetching data from the server _and_ it is responsible for displaying that data. This is not a great design and should be seprated into two compoments for the two tasks. 

## Dependencies

### https://openweathermap.org/api

This is the API used for fetching the weather forecasts. An API key is required to run either the tests or the server for this application.

- Cli arg parser <3
- Testify
- Material UI
- Modd (Kinda)
- Axios
- Lodash
- Immutability Helper
- Create React App

## Improvements 

- Use multiple APIs for extra reliablity
- Map on frontend
- Weather picture for current weather (image of rain)
- SVG needs labeling along the x axis!!!
- SVG needs to be easier to read
- Current temperature typography needs to be easier to read.
- Makefile

## Design decisions

### Testing

- Unmanaged out of bound process
  - Mocks not integration
- Separate module for testing
- Classic style unit tests
- Adding a test from upstream API
- Just end to end testing because I would need to mock out the API to do proper unit testing.

### Front-end

- No router
- Simple state management 
- No Elm :(
- No TS because no Elm :(
- Repurposed logo and Favicon because reasons
- No immutable JS
- No weather cacheing
- Not real time
- No design (because that should be done sperate to implementation)
- Not tested (time)
  - Needs snapshot test
- City component needs to be refactored
- Loading and problem can be moved somewhere else
- Better responsiveness on smaller screens
- SVG responsive to temperature bounds for better viewing experience.
- Need to wrap fetchData in useCallback hook in City.js

### Back-end 

- Default HTTP router
- No metrics :(
- No tracing :(
- All configuration passed in by env or cmd line args
- Use of internal folders
- Only current weather
- Closed world assumption for available cities
- No API docs
- Single API
- Cors conf
- Health Check endpoint
- Default HTTP Router (anything can use it)
- Graceful shutdown 
- No support for SSL
- In the handleWeather function most of the boiler plate would be a separate function, but since I'm only doing it once that abstraction doesn't make sense.
- Middlewares (in particular logging)

### Git 

- Single commit because no issue number
- Git branching (master/tag, gitflow)
