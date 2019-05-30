## News aggregator service

### How to install
`$ go get -u github.com/codecure/news-aggregator/...`

### How to run aggregator
`$ aggregator -feed https://echo.tochka.com/feeds/topics/ru/ -rule "date=UpdatedParsed&content=Content" -db /path/to/db/dir`
`$ aggregator -feed https://habr.com/ru/rss/all/all/?fl=ru -rule "date=PublishedParsed&content=Description" -db /path/to/db/dir`

### How to run aggregator web service
`$ aggregator-server -db /path/to/db/dir`
