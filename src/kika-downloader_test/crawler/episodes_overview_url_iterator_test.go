package crawler

import (
	"fmt"
	"kika-downloader/config"
	"kika-downloader/crawler"
	"kika-downloader/http"
	testConfig "kika-downloader_test/config"
	"net/url"
	"testing"
)

func TestPageIteration(t *testing.T) {
	appContext, err := config.InitApp(testConfig.TorSocksProxyURL)
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("http_client")
	if err != nil {
		t.Error(err)
	}
	httpClient := service.(http.ClientInterface)

	service, err = appContext.SafeGet("episodes_overview_url_iterator")
	if err != nil {
		t.Error(err)
	}
	iterator := service.(crawler.IteratorInterface)

	testUrl, err := url.Parse(testConfig.EpisodesOverviewURL)
	if err != nil {
		t.Error(err)
	}

	iterator.SetCrawlingURL(testUrl)

	gotValidURL := false

	// validate every url received from iterator
	for rawURL := range iterator.Run() {
		if _, err := url.Parse(rawURL); err != nil {
			t.Error(err)
		}

		if _, err := httpClient.Get(rawURL); err != nil {
			t.Error(err)
		}

		fmt.Printf("validated overview url: \"%s\"\n", rawURL)

		gotValidURL = true
	}

	if !gotValidURL {
		t.Error(fmt.Errorf("unable to retrieve a valid page url"))
	}
}
