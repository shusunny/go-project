package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch 返回 URL 的 body 内容，并且将在这个页面上找到的 URL 放到一个 slice 中。
	Fetch(url string) (body string, urls []string, err error)
}

// 创建map用来存放已抓取的网页, wg组用来管理线程
var m map[string]bool
var wg sync.WaitGroup

// TODO: 并行的抓取 URL。
// TODO: 不重复抓取页面。
// 这里用了辅助函数_crawl来完成上面的任务，并进行主要的抓取工作
func _crawl(url string, depth int, fetcher Fetcher, Results chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	if depth <= 0 {
		return
	}

	if exists := m[url]; exists {
		return
	}

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		Results <- fmt.Sprintf("not found: %s", url)
		return
	}
	m[url] = true

	Results <- fmt.Sprintf("found: %s %q", url, body)

	for _, u := range urls {
		wg.Add(1)
		go _crawl(u, depth-1, fetcher, Results, wg)
	}
	return
}

// Crawl 使用 fetcher 从某个 URL 开始递归的爬取页面，直到达到最大深度。
func Crawl(url string, depth int, fetcher Fetcher) {
	Results := make(chan string)
	wg.Add(1)
	go _crawl(url, depth, fetcher, Results, &wg)
	go func() {
		wg.Wait()
		close(Results)
	}()
	for i := range Results {
		fmt.Println(i)
	}
}

func main() {
	m = make(map[string]bool)
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher 是返回若干结果的 Fetcher。
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher 是填充后的 fakeFetcher。
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}