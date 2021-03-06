package services

import (
	"github.com/sarulabs/di"
	"kika-downloader/commands"
	"kika-downloader/contract"
	"kika-downloader/crawler"
)

// AssignEpisodesOverviewItemsUrlIterator
func AssignFetchAllCommandHandler(builder *di.Builder) error {
	builder.AddDefinition(di.Definition{
		Name:  "command_handler.fetch_all",
		Scope: di.App,
		Build: func(ctx di.Context) (interface{}, error) {
			service, err := ctx.SafeGet("episodes_overview_url_iterator")
			if err != nil {
				return nil, err
			}
			episodesPageIterator := service.(crawler.IteratorInterface)

			service, err = ctx.SafeGet("episodes_items_url_iterator")
			if err != nil {
				return nil, err
			}
			pageItemsIterator := service.(crawler.IteratorInterface)

			service, err = ctx.SafeGet("video_extractor")
			if err != nil {
				return nil, err
			}
			videoExtractor := service.(contract.VideoExtractorInterface)

			service, err = ctx.SafeGet("video_downloader")
			if err != nil {
				return nil, err
			}
			videoDownloader := service.(contract.VideoDownloaderInterface)

			return commands.NewFetchAllHandler(
				episodesPageIterator,
				pageItemsIterator,
				videoExtractor,
				videoDownloader,
			).(contract.CommandHandlerInterface), nil
		},
	})

	return nil
}
