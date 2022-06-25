import React, { useState } from 'react';
import axios from "axios"



function Test() {

    const [content, setContent] = useState("offline")

  async function TestInit() {
        await axios.post("http://localhost:8080/getData", {
            status: "online",
            port: 8080,
            project: "NALYA",

        }).then((res) => {
            console.log("RESPONSE =>", res)
            setContent(res.data)
        })
    }
    

    return (
        <h3 onClick={TestInit}>server: {content}</h3>
    )
    
}

export default Test