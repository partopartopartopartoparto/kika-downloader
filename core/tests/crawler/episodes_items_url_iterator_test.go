package crawler

import (
	"fmt"
	"net/url"
	"rkl.io/kika-downloader/core/crawler"
	"rkl.io/kika-downloader/core/http"
	testConfig "rkl.io/kika-downloader/core/tests/config"
	"testing"
)

func TestItemsIteration(t *testing.T) {
	appContext, err := testConfig.InitTestContext(testConfig.TorSocksProxyURL)
	if err != nil {
		t.Error(err)
	}

	service, err := appContext.SafeGet("http_client")
	if err != nil {
		t.Error(err)
	}
	httpClient := service.(http.ClientInterface)

	testURL, err := url.Parse(testConfig.EpisodesItemsURL)
	if err != nil {
		t.Error(err)
	}

	service, err = appContext.SafeGet("episodes_items_url_iterator")
	if err != nil {
		t.Error(err)
	}

	iterator := service.(crawler.IteratorInterface)
	iterator.SetCrawlingURL(testURL)

	gotValidURL := false

	// validate every url received from iterator
	for rawURL := range iterator.Run() {

		if _, err := url.Parse(rawURL); err != nil {
			t.Error(err)
		}

		if _, err := httpClient.Get(rawURL); err != nil {
			t.Error(err)
		}

		fmt.Printf("validated item url: \"%s\"\n", rawURL)

		gotValidURL = true
	}

	if !gotValidURL {
		t.Error(fmt.Errorf("unable to get valid item url"))
	}
}
