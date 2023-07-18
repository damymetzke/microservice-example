import express from 'express';

const app = express();
const port = 80;

app.get('/products', (_, res) => {
  res.send({message: "Hello from the microservice"});
});

app.get('/products/:id', (req, res) => {
  res.send({message: `Hello from the microservice, id = ${req.params.id}`});
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
