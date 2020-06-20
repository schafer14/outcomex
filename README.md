# Outcomex Tech Challenge

## Building

For both testing and serving the back end requires a API Key to be passed in. For testing the 

### Development

### Testing

### Production

### Docs

## Application Structure

## Dependencies

- https://openweathermap.org/api
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
