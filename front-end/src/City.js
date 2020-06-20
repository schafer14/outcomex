import React, {useState, useEffect} from 'react';
import {
  Card,
  CardMedia,
  CardHeader,
  CardContent,
  Button,
  CircularProgress,
  Typography,
  CardActions,
} from '@material-ui/core';
import {chain, round, map, capitalize} from 'lodash';
import Forecast from './ForecastImg';
import * as axios from 'axios';

// City is a component that is responsible for displaying the weather of a given city. 
// As the application grows this component could also take responsiblity for caching 
// the data and minimize the number of requests made. Another feature that would live
// in this component is the ability to to automatic refreshing after a given amount of time. 
//
// The problem with this component right now is that it is both fetching data from the api _and_
// displaying that data. These two responsiblities should be separated into their own components.
//
// Another issue is there is a lot of boilerplate code (such as displaying problems and the loading state).
// Since this code is only used once in the application I did not abstract it to a separate component,
// but as soon as another component used the API the loading and problem if conditionals would be moved
// out of this component.
export default function City({name, isShowing}) {

  const [isLoading, setIsLoading] = useState(true);
  const [data, setData] = useState();
  const [problem, setProblem] = useState();
  const pointsInWeatherImage = 8;

  const fetchData = async () => {
    setIsLoading(true);
    // Would probably not hardcode this in real life.
    try {
      const resp = await axios.get(`http://localhost:3000/weather/${name}`);
      setData(resp.data);
    } catch (e) {
      setProblem(e.response.data);
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    if (isShowing) {
      fetchData();
    }
  }, [isShowing]);

  if (!isShowing) {
    return <></>;
  }

  if (isLoading) {
    return (
      <Card style={{margin: 20}}>
        <CardHeader title={capitalize(name)} />
        <CardContent>
          <CircularProgress />
        </CardContent>
      </Card>
    );
  }

  if (problem) {
    return (
      <Card style={{margin: 20}}>
        <CardHeader title={capitalize(name)} />
        <CardContent>
          <Typography color="error">{problem}</Typography>
        </CardContent>
      </Card>
    );
  }

  const current = data[0];
  const forecast = chain(data).map('temp').take(pointsInWeatherImage).value();

  return (
    <Card style={{margin: 20}}>
      <CardHeader title={capitalize(name)} />
      <CardMedia>
        <Forecast points={forecast} />
      </CardMedia>
      <CardContent>
        <Typography variant="h6">Current</Typography>
        {map(current.description, description => (
          <Typography key={description} variant="p" color="textSecondary">
            {capitalize(description)}
          </Typography>
        ))}
        <Typography variant="p" component="div" color="textSecondary">
          Temperature: {round(current.temp - 272.15)}&deg;C
        </Typography>
        <Typography variant="p" component="div" color="textSecondary">
          Feels like: {round(current.feelsLike - 272.15)}&deg;C
        </Typography>
        <Typography variant="p" component="div" color="textSecondary">
          Wind: {current.windSpeed} m/s{' '}
          {degreesToDirection(current.windDirection)}
        </Typography>
      </CardContent>
      <CardActions>
        <Button onClick={() => fetchData()}>Reload</Button>
      </CardActions>
    </Card>
  );
}

function degreesToDirection(degrees) {
  if (degrees < 22.5) {
    return 'N';
  } else if (degrees < 22.5 + 45) {
    return 'NW';
  } else if (degrees < 22.5 + 45 * 2) {
    return 'W';
  } else if (degrees < 22.5 + 45 * 3) {
    return 'SW';
  } else if (degrees < 22.5 + 45 * 4) {
    return 'S';
  } else if (degrees < 22.5 + 45 * 5) {
    return 'SE';
  } else if (degrees < 22.5 + 45 * 6) {
    return 'E';
  } else if (degrees < 22.5 + 45 * 7) {
    return 'NE';
  } else {
    return 'N';
  }
}
