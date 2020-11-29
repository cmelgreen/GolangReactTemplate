import React from "react"

import RedirectIfRequested from "../Utils/RedirectRequested.js"

export default function MainPageContainer(props) {
    const redirect = RedirectIfRequested()

    return redirect ? redirect : <MainPage /> 
}

function MainPage(props) {
    return (
        <div>Main Page</div>
    )
}