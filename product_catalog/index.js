import express from 'express';

const PRODUCTS = [
  {
    id: 1,
    title: "Flashlight",
    description: "High intensity flashlight, perfect for camping!",
    price: 2199,
  },
  {
    id: 2,
    title: "2 Person tent",
    description: "Dome shaped tent. Sleep comfortably with 2 people.",
    price: 1999,
  },
  {
    id: 3,
    title: "Pocket knife",
    description: "Durable pocket knife.",
    price: 2999,
  },
  {
    id: 4,
    title: "Magnesium fire starter",
    description: "Easily start fire using this stick made of magnesium!",
    price: 1499,
  },
]

const app = express();
const port = 80;

app.get('/products', (_, res) => {
  res.send({products: PRODUCTS});
});

app.get('/products/:id', (req, res) => {
  let product = PRODUCTS.find(product => product.id = req.params.id);

  if (product === undefined) {
    res.status(404)
    return
  }

  res.send({product});
});

app.listen(port, () => {
  console.log(`Server is running on port ${port}`);
});
