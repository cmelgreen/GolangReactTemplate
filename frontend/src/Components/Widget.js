import React from "react"
import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import axios from "axios"

const api = "/api/post"

export default function WidgetContainer(props) {
    const id = useParams().widgetId
    const [text, setText] = useState('')

    useEffect(() => {
        readPost(id, setText)
      }, []);
    
    return <div> Widget {text} accessed via react-router </div>
}

const readPost = (id, action) => {
    axios.get(api + "?id=" + id).then((resp) => action(resp.data.content))
}