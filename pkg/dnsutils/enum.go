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
	var discoveredSubdomains []string

	if wordlistPath == "" {
		return nil, fmt.Errorf("no wordlist provided")
	}

	if concurrency < 1 {
		concurrency = 1
	}

	wordlistFile, err := os.Open(wordlistPath)
	if err != nil {
		return nil, err
	}
	defer wordlistFile.Close()

	resolver := getResolver(resolverAddr)
	lookupCtx := context.Background()

	wordQueue := make(chan string)
	var waitGroup sync.WaitGroup
	var resultMutex sync.Mutex

	for workerID := 0; workerID < concurrency; workerID++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for word := range wordQueue {
				candidateSubdomain := fmt.Sprintf("%s.%s", word, domain)
				if _, err := resolver.LookupIPAddr(lookupCtx, candidateSubdomain); err == nil {
					resultMutex.Lock()
					discoveredSubdomains = append(discoveredSubdomains, candidateSubdomain)
					resultMutex.Unlock()
				}
			}
		}()
	}

	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		wordQueue <- scanner.Text()
	}
	close(wordQueue)
	waitGroup.Wait()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return discoveredSubdomains, nil
}
