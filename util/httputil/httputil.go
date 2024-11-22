package httputil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	neturl "net/url"
	"time"

	"golang.org/x/net/proxy"
)

// Usage: connected := TestConnection("example.com", 80, 3 * time.Second, "")
func TestConnection(address string, port int, timeout time.Duration, proxyURL string) bool {
	// Combine the address and port into a single string
	target := fmt.Sprintf("%s:%d", address, port)
	// Create a dialer
	var dialer proxy.Dialer
	if proxyURL != "" {
		// Parse the proxy URL
		parsedProxyURL, err := url.Parse(proxyURL)
		if err != nil {
			fmt.Printf("Invalid proxy URL: %v\n", err)
			return false
		}
		// Create a proxy dialer
		dialer, err = proxy.FromURL(parsedProxyURL, proxy.Direct)
		if err != nil {
			fmt.Printf("Failed to create proxy dialer: %v\n", err)
			return false
		}
	} else {
		// Use a direct connection if no proxy is specified
		dialer = &net.Dialer{Timeout: timeout}
	}
	// Attempt to establish a connection using the dialer
	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
	// defer cancel()
	conn, err := dialer.Dial("tcp", target)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func GetString(url string, proxyURL string) (string, int, error) {
	resp, code, err := GetBytes(url, proxyURL)
	if err != nil {
		return "", 0, err
	}
	return string(resp), code, nil
}

func GetJSON(url string, proxyURL string) (string, int, error) {
	resp, code, err := GetBytes(url, proxyURL)
	if err != nil {
		return "", code, err
	}
	var prettyJSON bytes.Buffer
	err = json.Indent(&prettyJSON, resp, "", "  ")
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return "", 0, err
	}

	return prettyJSON.String(), code, nil
}

func GetBytes(url string, proxyURL string) ([]byte, int, error) {
	resp, code, err := get(url, proxyURL)
	if err != nil {
		return nil, 0, err
	}

	defer resp.Body.Close()
	body, readErr := io.ReadAll(resp.Body)
	if readErr != nil {
		return nil, code, readErr
	}
	return body, code, nil
}

func MeasureTime[T any](fn func(string, string) (T, int, error), url string, proxyURL string) (T, int, int64, error) {
	start := time.Now()
	result, code, err := fn(url, proxyURL)
	duration := time.Since(start).Milliseconds()
	return result, code, duration, err
}

func get(url string, proxyURL string) (*http.Response, int, error) {
	var transport *http.Transport
	if proxyURL != "" {
		proxy, err := neturl.Parse(proxyURL)
		if err != nil {
			return nil, 0, err
		}
		transport = &http.Transport{
			Proxy: http.ProxyURL(proxy),
		}
	} else {
		transport = &http.Transport{
			Proxy: nil,
		}
	}

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(url)
	if err != nil {
		return nil, 0, err
	}
	return resp, resp.StatusCode, nil
}
