const express = require('express');
const bodyParser = require('body-parser');
const calculadoraRoutes = require('./src/routes/calculadoraRoutes');

const app = express();

app.use(bodyParser.urlencoded({ extended: false }));
app.use(express.static('src/views'));
app.use('/calculadora', calculadoraRoutes);

const PORT = 3000;
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
