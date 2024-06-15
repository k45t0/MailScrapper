package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"

	"github.com/fatih/color"
)

var (
	emailRegex       = `[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}`
	outputLock       sync.Mutex
	processedDomains sync.Map
	verbose          bool
	quiet            bool
)

func extractEmailsFromHTML(htmlContent string) []string {
	re := regexp.MustCompile(emailRegex)
	matches := re.FindAllString(htmlContent, -1)
	emailSet := make(map[string]struct{})
	for _, email := range matches {
		emailSet[email] = struct{}{}
	}
	emails := make([]string, 0, len(emailSet))
	for email := range emailSet {
		emails = append(emails, email)
	}
	return emails
}

func filterEmails(emails []string) []string {
	validEmails := []string{}
	for _, email := range emails {
		if !strings.HasSuffix(email, ".png") &&
			!strings.HasSuffix(email, ".gif") &&
			!strings.HasSuffix(email, ".int") &&
			!strings.HasSuffix(email, ".js") &&
			!strings.HasSuffix(email, ".jpg") &&
			!strings.HasSuffix(email, ".jpeg") &&
			!strings.HasSuffix(email, ".webp") &&
			!strings.HasSuffix(email, ".ld") {
			validEmails = append(validEmails, email)
		}
	}
	return validEmails
}

func saveEmailsToFile(emails []string, outputFile string) {
	outputLock.Lock()
	defer outputLock.Unlock()

	file, err := os.OpenFile(outputFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, email := range emails {
		_, err := writer.WriteString(email + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
	writer.Flush()
}

func processURL(link string, outputFile string) {
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "http://" + link
	}
	u, err := url.Parse(link)
	if err != nil || u.Host == "" {
		return
	}
	domain := u.Host
	if _, loaded := processedDomains.LoadOrStore(domain, true); loaded {
		return
	}

	resp, err := http.Get(link)
	if err != nil || resp.StatusCode >= 400 {
		if !quiet {
			color.Red("[-] %s", domain)
		}
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		if !quiet {
			color.Red("[-] %s", domain)
		}
		return
	}

	emails := extractEmailsFromHTML(string(body))
	emails = filterEmails(emails)
	if len(emails) > 0 {
		saveEmailsToFile(emails, outputFile)
		color.Green("[+] %s | %d", domain, len(emails))
		if verbose {
			for _, email := range emails {
				color.Cyan("    %s", email)
			}
		}
	} else {
		if !quiet {
			color.Red("[-] %s", domain)
		}
	}
}

func processFile(filePath string, threadsCount int, outputFile string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	links := make(chan string, threadsCount)
	var wg sync.WaitGroup

	for i := 0; i < threadsCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range links {
				processURL(link, outputFile)
			}
		}()
	}

	for scanner.Scan() {
		link := strings.TrimSpace(scanner.Text())
		links <- link
	}
	close(links)

	wg.Wait()
}

func main() {
	listFlag := flag.String("l", "", "Path to the file containing URLs")
	domainFlag := flag.String("d", "", "Single domain to process")
	threadsFlag := flag.Int("t", 10, "Number of threads to use")
	outputFlag := flag.String("o", "emails.txt", "Output file to save emails")
	verboseFlag := flag.Bool("v", false, "Enable verbose output")
	quietFlag := flag.Bool("q", false, "Suppress output for domains with no emails found")

	flag.Parse()

	verbose = *verboseFlag
	quiet = *quietFlag

	if *listFlag == "" && *domainFlag == "" {
		fmt.Println("Either -l (list file) or -d (domain) must be specified")
		flag.Usage()
		return
	}

	if *listFlag != "" {
		processFile(*listFlag, *threadsFlag, *outputFlag)
	} else if *domainFlag != "" {
		links := make(chan string, *threadsFlag)
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range links {
				processURL(link, *outputFlag)
			}
		}()
		links <- *domainFlag
		close(links)
		wg.Wait()
	}
}
