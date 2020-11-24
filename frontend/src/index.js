import React from 'react';
import ReactDOM from 'react-dom';
import { BrowserRouter as Router, Route } from "react-router-dom";

import  ParseRoutes from "./Utils/ParseRoutes"
import { routes } from "./routes.json"

function AppContainer() {
  const routeProps = ParseRoutes(routes)
  
  return (
    <Router>
     {routeProps.map((props, i) => (<Route key={i} {...props} /> ))}
    </Router>
  )
}

ReactDOM.render(<AppContainer />, document.getElementById("root"))
