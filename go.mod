module example.com/hello

go 1.18

replace example.com/headless => ./headless/

replace example.com/zlibrary => ./zlibrary/

require (
	example.com/headless v0.0.0-00010101000000-000000000000
	example.com/myip v0.0.0-00010101000000-000000000000
	example.com/stock v0.0.0-00010101000000-000000000000
	example.com/zlibrary v0.0.0-00010101000000-000000000000
	github.com/gin-gonic/contrib v0.0.0-20201101042839-6a891bf89f19
	github.com/gin-gonic/gin v1.7.7
)

require (
	github.com/BurntSushi/toml v0.3.1 // indirect
	github.com/PuerkitoBio/goquery v1.8.0 // indirect
	github.com/andybalholm/cascadia v1.3.1 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.1.8 // indirect
	github.com/chromedp/cdproto v0.0.0-20220321060548-7bc2623472b3 // indirect
	github.com/chromedp/chromedp v0.8.0 // indirect
	github.com/chromedp/sysutil v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgrijalva/jwt-go v3.2.1-0.20180921172315-3af4c746e1c2+incompatible // indirect
	github.com/dustin/go-humanize v1.0.1-0.20200219035652-afde56e7acac // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/go-playground/locales v0.13.0 // indirect
	github.com/go-playground/universal-translator v0.17.0 // indirect
	github.com/go-playground/validator/v10 v10.4.1 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gobwas/httphead v0.1.0 // indirect
	github.com/gobwas/pool v0.2.1 // indirect
	github.com/gobwas/ws v1.1.0 // indirect
	github.com/gocolly/colly/v2 v2.1.0 // indirect
	github.com/golang/groupcache v0.0.0-20200121045136-8c9f03a8e57e // indirect
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/headzoo/surf v1.0.1 // indirect
	github.com/iancoleman/strcase v0.1.2 // indirect
	github.com/jinzhu/copier v0.0.0-20190924061706-b57f9002281a // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.2 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.10 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/leodido/go-urn v1.2.0 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/nu7hatch/gouuid v0.0.0-20131221200532-179d4d0c4d8d // indirect
	github.com/pilagod/gorm-cursor-paginator/v2 v2.2.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/root-gg/logger v0.0.0-20150501173826-5d9a47a35312 // indirect
	github.com/root-gg/plik v0.0.0-20220217080014-93cd25d42e4e // indirect
	github.com/root-gg/utils v0.0.0-20201113182430-4cef83d8cdcf // indirect
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/stretchr/testify v1.7.0 // indirect
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/ugorji/go/codec v1.1.7 // indirect
	github.com/xhit/go-str2duration/v2 v2.0.0 // indirect
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20210916014120-12bc252f5db8 // indirect
	golang.org/x/sys v0.0.0-20210615035016-665e8c7367d1 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/headzoo/surf.v1 v1.0.1 // indirect
	gopkg.in/yaml.v2 v2.2.8 // indirect
	gopkg.in/yaml.v3 v3.0.0-20200313102051-9f266ea9e77c // indirect
	gorm.io/gorm v1.22.3 // indirect
)

replace example.com/stock => ./stock/

replace example.com/myip => ./myip/
