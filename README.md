<h1>MailScrapper</h1>
    <p>MailScrapper é um script em Go para extrair e-mails de uma página da web.</p>

<h2>Instalação</h2>
            git clone https://github.com/k45t0/MailScrapper.git
            
            cd MailScrapper
            
            go mod tidy
            
            go build mailscrapper.go
            
            cp mailscrapper /usr/local/bin

<hr>
<h2>Modo de Uso</h2>
<p>Para usar o MailScrapper, execute o script no terminal:</p>

    mailscrapper https://exemplo.com
    
    mailscrapper -l urls.txt -t 10 -o output.txt -v -q

    subfinder -d domain.txt -o outputdomain.txt | mailscrapper -l outputdomain.txt -t 100 -o outputmail.txt -v -q

<hr>
<h2>Help</h2>

        Usage of mailscrapper:
          -d string
            	Single domain to process
          -l string
            	Path to the file containing URLs
          -o string
            	Output file to save emails (default "emails.txt")
          -q	Suppress output for domains with no emails found
          -t int
            	Number of threads to use (default 10)
          -v	Enable verbose output

</body>
</html>
