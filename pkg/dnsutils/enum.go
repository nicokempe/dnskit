package dnsutils

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
)

// Enumerate performs subdomain enumeration using a wordlist. Lookups are
// executed concurrently according to the provided concurrency level and may use
// a custom DNS resolver.
func Enumerate(domain, wordlistPath string, concurrency int, resolverAddr string) ([]string, error) {
	var subdomains []string

	if wordlistPath == "" {
		return nil, fmt.Errorf("no wordlist provided")
	}

	if concurrency < 1 {
		concurrency = 1
	}

	file, err := os.Open(wordlistPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	r := getResolver(resolverAddr)
	ctx := context.Background()

	jobs := make(chan string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for sub := range jobs {
				candidate := fmt.Sprintf("%s.%s", sub, domain)
				if _, err := r.LookupIPAddr(ctx, candidate); err == nil {
					mu.Lock()
					subdomains = append(subdomains, candidate)
					mu.Unlock()
				}
			}
		}()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		jobs <- scanner.Text()
	}
	close(jobs)
	wg.Wait()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return subdomains, nil
}
