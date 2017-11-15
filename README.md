# gocrawler

Implemented features:

* [x] Static HTML parser
* [x] Retries
* [x] Parallel

Planned features:

* [ ] Maximum Depth

Missing features:

* [ ] Respect robots.txt
* [ ] Look for Sitemap in robots.txt
* [ ] Backoff Retries
* [ ] Render javascript (chrome headless?)
* [ ] Extract assets added by CSS

Demo:

```
# install
$ go get -u github.com/matthewrudy/gocrawler/...

# crawl http://tomblomfield.com
$ gocrawler git:(master) gocrawler
success: http://tomblomfield.com/
success: http://tomblomfield.com/about
success: http://tomblomfield.com/rss
success: http://tomblomfield.com/day/2015/12/13
...
http://tomblomfield.com/random
 - http://www.gravatar.com/avatar/c833be5582482777b51b8fc73e8b0586?s=128&d=identicon&r=PG
 - http://78.media.tumblr.com/ddebec46b60f554989f09682fc3d8e71/tumblr_inline_mtj697fPI11r5tr1m.jpg

http://tomblomfield.com/rss
```