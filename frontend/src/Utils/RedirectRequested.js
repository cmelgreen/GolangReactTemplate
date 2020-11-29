import React from "react"
import { Redirect } from "react-router-dom"

const redirectKey = "redirect"

const clearCookie = (cookie) => {
    document.cookie = cookie + "=''; max-age=-1"
}

const getCookieValue = (a) => {
    var b = document.cookie.match('(^|;)\\s*' + a + '\\s*=\\s*([^;]+)')
    return b ? b.pop() : ''
  }

const RedirectRequested = () => {
    const redirectTo = getCookieValue(redirectKey)
    clearCookie(redirectKey)

    if ( redirectTo ) {
        return <Redirect to={redirectTo} />
    }

    return null
}

export default RedirectRequested