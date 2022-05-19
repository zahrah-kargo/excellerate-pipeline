const express = require('express')
const handler = require('./src/handler/hello_world.js')
const app = express()
const port = 8080

app.get('/', (req, res) => {
  res.send(handler.helloWorld())
})

app.listen(port, () => {
  console.log(`Example app listening on port ${port}`)
})