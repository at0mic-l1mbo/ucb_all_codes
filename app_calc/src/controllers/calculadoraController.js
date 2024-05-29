const calculadora = require('../calculadora/calculadora');

exports.calcular = (req, res) => {
    const { num1, num2, operacao } = req.body;

    if (isNaN(num1) || isNaN(num2)) {
        res.send(`
            <script>
                alert('Por favor, insira valores numéricos.');
                window.location.href = "/";
            </script>
        `);
    } else if (operacao === '/' && parseFloat(num2) === 0) {
        res.send(`
            <script>
                alert('Divisão por zero não é permitida.');
                window.location.href = "/";
            </script>
        `);
    } else {
        const resultado = calculadora.calcular(parseFloat(num1), parseFloat(num2), operacao);
        res.send(`
            <html>
                <body>
                    <h1>Resultado</h1>
                    <p>O resultado da operação ${calculadora.getOperador(operacao)} entre ${num1} e ${num2} é ${resultado}</p>
                    <a href="/">Voltar</a>
                </body>
            </html>
        `);
    }
};
