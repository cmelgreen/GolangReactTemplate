import React from "react"
import { Redirect } from "react-router-dom"

const getCookieValue = (a) => {
    var b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)')
    return b ? b.pop() : ''
  }

const RedirectIfRequested = () => {
    const redirectTo = getCookieValue("redirect")

    if ( redirectTo !== "" ) {
        document.cookie = "redirect=''; max-age=-1"
        return <Redirect to={redirectTo} />
    }

    return null
}

export default RedirectIfRequested