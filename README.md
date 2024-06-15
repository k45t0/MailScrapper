<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>MailScrapper</title>
    <style>
        body {
            font-family: Arial, sans-serif;
            line-height: 1.6;
            max-width: 800px;
            margin: 0 auto;
            padding: 20px;
        }
        h1, h2, h3 {
            color: #333;
        }
        code {
            background-color: #f4f4f4;
            padding: 5px;
            border-radius: 4px;
            font-size: 0.9em;
        }
    </style>
</head>
<body>
    <h1>MailScrapper</h1>
    <p>MailScrapper é um script em Go para extrair e-mails de uma página da web.</p>
    
    <h2>Instalação</h2>
    <p>Para utilizar o MailScrapper, siga os passos abaixo:</p>
    <ol>
        <li>Clone o repositório:</li>
        <code>
            git clone https://github.com/k45t0/MailScrapper.git
        </code>
        <li>Navegue até o diretório do projeto:</li>
        <code>
            cd MailScrapper
        </code>
        <li>Opcional: Instale dependências do projeto (se necessário):</li>
        <code>
            go mod tidy
        </code>
        <li>Compile o script (se necessário):</li>
        <code>
            go build mailscrapper.go
        </code>
    </ol>

    <h2>Modo de Uso</h2>
    <p>Para usar o MailScrapper, execute o script no terminal:</p>
    <code>
        ./mailscrapper
    </code>

    <h2>Exemplo</h2>
    <p>Suponha que você deseje extrair e-mails de uma página web:</p>
    <code>
        ./mailscrapper https://exemplo.com
    </code>

    <h2>Contribuições</h2>
    <p>Contribuições são bem-vindas! Sinta-se à vontade para enviar pull requests.</p>

    <h2>Licença</h2>
    <p>Este projeto está licenciado sob a Licença MIT - veja o arquivo <code>LICENSE</code> para mais detalhes.</p>

    <hr>
    <p>Este README foi criado para fins demonstrativos. Adaptado e personalizado conforme necessário.</p>
</body>
</html>
