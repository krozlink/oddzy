const server = require('http').createServer();

const io = require('socket.io')(8882, {
  path: '/',
  serveClient: false,
  cookie: false
});
io.set('origins', '*:*');

const priceNamespace = io.of('/prices');

const testWinPrices = [2.5, 4.6, 5.5, 7.5, 101, 101, 9.5, 51, 41,  26, 41, 151];
const testPlacePrices = [1.29, 1.73, 1.83, 2.2, 13, 13, 2.5, 8, 6, 4.33, 6, 21];

var interval = setInterval(() => {

    // Randomly pick whether to update a win or a place price
    const isWin = Math.round(Math.random()) == 1;

    // Pick a random runner between 1 and 12 (runners > 12 will not have their prices updated)
    const barrier = Math.round(Math.random() * (12 - 1) + 1);

    let price = isWin ? testWinPrices[barrier - 1] : testPlacePrices[barrier-1];

    // change the price between -10% and +10%
    let changeAmount = (Math.random() * 0.20) - 0.10; // random number between -0.10 and +0.10

    price = price * (1 + changeAmount);
    
    // round the price to something that looks slightly more realistic
    if(price > 25) {
      price = (Math.round(price / 5) * 5).toFixed(0);
    } else if(price > 10) {
      price = Math.round(price).toFixed(0);
    } else if (price > 5) {
      price = (Math.round(price / 0.2) * 0.2).toFixed(2);
    } else if( price > 2) {
      price = (Math.round(price / 0.1) * 0.1).toFixed(2);
    } else {
      price = price.toFixed();
    }

    const data = {
      selection_id: barrier,
      type: isWin? 'win': 'place',
      price
    }

    priceNamespace.emit('updatePrice', data);
}, 5000);