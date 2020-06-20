import React, {useState} from 'react';
import {Button, Paper, Grid, Typography} from '@material-ui/core';
import {map} from 'lodash';
import update from 'immutability-helper';
import City from './City';

function App() {
  const [cities, setCities] = useState([
    {name: 'sydney', isShowing: true},
    {name: 'adelaide', isShowing: false},
    {name: 'melbourne', isShowing: false},
    {name: 'covid', isShowing: true},
  ]);

  return (
    <Grid
      style={{height: '100%'}}
      container
      alignContent="center"
      justify="center">
      <Paper style={{padding: '40px'}}>
        <Grid direction="row" container>
          <Grid item style={{padding: 20}}>
            <Typography variant="h5">Cities</Typography>
            {map(cities, (city, key) => {
              return (
                <Typography
                  key={city.name}
                  onClick={() => {
                    setCities(
                      update(cities, {
                        [key]: {isShowing: {$set: !city.isShowing}},
                      }),
                    );
                  }}>
                  <Button
                    color="primary"
                    variant={city.isShowing ? 'contained' : 'outlined'}
                    style={{margin: 10}}>
                    {city.name}
                  </Button>
                </Typography>
              );
            })}
          </Grid>
          <Grid item style={{padding: 20}}>
            <Typography variant="h5">Weather</Typography>
            <Grid container direction="row">
              {map(cities, (city, key) => {
                return (
                  <City
                    key={city.name}
                    name={city.name}
                    isShowing={city.isShowing}
                  />
                );
              })}
            </Grid>
          </Grid>
        </Grid>
      </Paper>
    </Grid>
  );
}

export default App;
