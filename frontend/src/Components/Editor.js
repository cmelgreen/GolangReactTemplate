import React from "react"

import { useState } from "react"

import axios from "axios"

const api = "/api/post"

export default function EditorContainer(props) {
    const [title, setTitle] = useState('Title')
    const [content, setContent] = useState('I was just thinking...')

    return <Editor title={title} setTitle={setTitle} content={content} setContent={setContent} />
}

function Editor(props) {
    
    return (
        <div className="editor-container">
            Create Post:<br/>
            <textarea id="title-editor" onChange={(e) => props.setTitle(e.target.value)}/><br/>
            <textarea id="post-editor"  onChange={(e) => props.setContent(e.target.value)} /><br/>
            <button className="editor-post-button" onClick={() => createPost(props.title, props.content)}>Post</button>
        </div>
    )
}

const createPost = (title, content) => {
    console.log("creating post", title, content)
    axios.post(api, {id: title, content: content})
}