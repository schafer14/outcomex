import React from 'react';
import City from './City';

test('City should initially enter a loading state', () => {

  // Arrange

  // Act
  const myCity = new City('sydney')

  // Assert
  expect(myCity.isLoading).toBe(true);
  expect(myCity.isShowing).toBe(true);

});

test('City name should be capitalied for displaying', () => {

  // Arrange

  // Act
  const myCity = new City('sydney')

  // Assert
  expect(myCity.name).toBe("Sydney");

});

test('Toggle is showing should change the showing status', () => {

  // Arrange
  const myCity = new City('sydney')

  // Act
  myCity.toggleIsShowing()

  // Assert
  expect(myCity.isShowing).toBe(false);

});
