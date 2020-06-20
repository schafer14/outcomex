import React from 'react';

import {map} from 'lodash';

export default function Forecast({points}) {
  return (
    <svg
      xmlns="http://www.w3.org/2000/svg"
      xmlnsXlink="http://www.w3.org/1999/xlink"
      width="100%"
      height="100px"
      viewBox={`0 0 ${points.length * 10} 40`}>
      <text x="1" y="35" fontSize="6" textAnchor="middle">0</text>
      <text x="1" y="5" fontSize="6" textAnchor="middle">40</text>

      {map(points, (point, index) => {
        if (index <= points.length - 2) {
          return (
            <line
              x1={index * 10}
              y1={40 - (point - 272.15)}
              x2={(index + 1) * 10}
              y2={40 - (points[index + 1] - 272.15)}
              stroke="black"
            />
          );
        }
      })}
    </svg>
  );
}
