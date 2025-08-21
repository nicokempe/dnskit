package dnsutils

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"sync"
)

/*
Enumerate subdomains using a wordlist.

	@param domain base domain.
	@param wordlistPath path to subdomain wordlist.
	@param concurrency number of concurrent lookups.
	@param resolverAddr optional custom resolver.
	@returns discovered subdomains or an error.
*/
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
	lookupContext := context.Background()

	wordChan := make(chan string)
	var waitGroup sync.WaitGroup
	var resultMutex sync.Mutex

	for workerID := 0; workerID < concurrency; workerID++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			for word := range wordChan {
				candidateSubdomain := fmt.Sprintf("%s.%s", word, domain)
				if _, err := resolver.LookupIPAddr(lookupContext, candidateSubdomain); err == nil {
					resultMutex.Lock()
					discoveredSubdomains = append(discoveredSubdomains, candidateSubdomain)
					resultMutex.Unlock()
				}
			}
		}()
	}

	scanner := bufio.NewScanner(wordlistFile)
	for scanner.Scan() {
		wordChan <- scanner.Text()
	}
	close(wordChan)
	waitGroup.Wait()

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return discoveredSubdomains, nil
}
