import React from "react"
import { Redirect } from "react-router-dom"

const redirectKey = "redirect"

const clearCookie = (cookie) => {
    document.cookie = "{" + cookie + "}=''; max-age=-1"
}

const getCookieValue = (a) => {
    var b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)')
    return b ? b.pop() : ''
  }

const RedirectIfRequested = () => {
    const redirectTo = getCookieValue(redirectKey)

    if ( redirectTo ) {
        clearCookie(redirectKey)
        return <Redirect to={redirectTo} />
    }

    return null
}

export default RedirectIfRequested