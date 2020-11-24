import React from "react"

import { useParams } from "react-router-dom"

export default function WidgetContainer(props) {
    const id = useParams().widgetId

    return <div> Widget {id} accessed via react-router </div>
}