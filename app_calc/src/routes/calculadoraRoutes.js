const express = require('express');
const router = express.Router();
const calculadoraController = require('../controllers/calculadoraController');

router.post('/resultado', calculadoraController.calcular);

module.exports = router;
