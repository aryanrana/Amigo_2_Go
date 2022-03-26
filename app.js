
const express = require('express')
const { request } = require('http')
const app = express()
const port = 3000
const cmd = require("node-cmd")

app.get('/', (req, res) => {
  res.sendFile(__dirname+"/index.html")
})  

app.get('/api/:id', (req, res) => {
    const syncDir=cmd.runSync('go run main.go '+req.params.id);
    console.log((syncDir));
    res.json(syncDir.data.split("\r\n")[0]);
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})


