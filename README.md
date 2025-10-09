<h1>MailScrapper</h1>
    <p>MailScrapper is a Go script for extracting emails from a web page.</p>

<hr>
<h2>How to Use</h2>
<p>To use MailScrapper, run the script in the terminal:</p>

    mailscrapper https://exemple.com
    
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
