exports.calcular = (num1, num2, operacao) => {
    switch (operacao) {
        case '+':
            return num1 + num2;
        case '-':
            return num1 - num2;
        case '*':
            return num1 * num2;
        case '/':
            return num2 !== 0 ? num1 / num2 : 'Erro: Divisão por zero';
        default:
            return null;
    }
};

exports.getOperador = (operacao) => {
    switch (operacao) {
        case '+':
            return 'adição';
        case '-':
            return 'subtração';
        case '*':
            return 'multiplicação';
        case '/':
            return 'divisão';
        default:
            return '';
    }
};
