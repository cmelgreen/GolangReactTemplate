import React, { useState } from 'react';
import ReactDOM from 'react-dom';
import "./index.css"

function Component() {
    const [data, setData] = useState("")

    const apiUrl = '../api';
    fetch(apiUrl)
      .then((response) => response.text())
      .then((data) => setData(data));

    return <React.Fragment><div id="loadedText">{data}</div><a href="/dynamic">Click to see dynamicly rendered script</a></React.Fragment>;

};

var element = <Component />;

ReactDOM.render(element, document.getElementById("root"))
