const express = require('express');
const app = express();
const axios = require('axios');

app.get("/", async (req, res) => {
    const data = await axios.get("http://localhost:3000/api/v1/anime")
    res.send(data.data)
})

app.listen(3001)