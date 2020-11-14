package fake

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func RandHeaders() http.Header {
	header := http.Header{}

	header.Set(`accept`, RandAccept())
	header.Set(`accept-language`, RandAcceptLanguage())

	if rand.Intn(1) == 0 {
		header.Set(`cache-control`, `no-cache`)
		header.Set(`pragma`, `no-cache`)
	}

	if rand.Intn(1) == 0 {
		header.Set(`dnt`, `1`)
	}

	header.Set(`upgrade-insecure-requests`, `1`)

	header.Set(`user-agent`, RandDesktopAgent())

	return header
}

func RandAccept() string {
	a := []string{
		`text/html`,
		`application/xhtml+xml`,
		`application/xml;q=` + fmt.Sprintf(`%.1f`, 0.1*float32(1+rand.Intn(8))),
		`image/webp`,
		`image/apng`,
		`*/*;q=` + fmt.Sprintf(`%.1f`, 0.1*float32(1+rand.Intn(5))),
	}

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return strings.Join(a, `,`)
}

func RandAcceptLanguage() string {
	a := []string{
		`ru-RU`,
		`ru;q=` + fmt.Sprintf(`%.1f`, 0.1*float32(1+rand.Intn(8))),
		`en-US;q=` + fmt.Sprintf(`%.1f`, 0.1*float32(1+rand.Intn(5))),
		`en;q=` + fmt.Sprintf(`%.1f`, 0.1*float32(1+rand.Intn(5))),
	}

	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })

	return strings.Join(a, `,`)
}

func RandDesktopAgent() string {
	osList := []string{
		`Macintosh; Intel Mac OS X 10_1#_#`,

		`Windows NT 5.1; Win64`,
		`Windows NT 6.1; Win64`,
		`Windows NT 6.2; Win64`,
		`Windows NT 6.3; Win64`,
		`Windows NT 10; Win64`,

		`X11; Linux x86_64`,
		`X11; Ubuntu; Linux x86_64`,
		`X11; Fedora; Linux x86_64`,
		`X11; OpenBSD amd64`,
	}

	fn := func(s string) string {
		r := make([]byte, len(s))

		for i, c := range s {
			if c != '#' {
				r[i] = byte(c)
			} else {
				r[i] = '1' + byte(rand.Intn(8))
			}
		}

		return string(r)
	}

	p := `Mozilla/5.0 (` + osList[rand.Intn(len(osList)-1)] + `) `

	switch rand.Intn(3) {
	default:
		return fn(p + `AppleWebKit/537.36 (KHTML, like Gecko) Chrome/##.0.####.### Safari/537.36`)
	case 1:
		return fn(p + `AppleWebKit/537.36 (KHTML, like Gecko) Chrome/##.0.####.### YaBrowser/##.##.####.##### Safari/537.36`)
	case 2:
		return fn(p + `AppleWebKit/537.36 (KHTML, like Gecko) Chrome/##.0.####.### Safari/537.36 OPR/##.0.####.##`)
	case 3:
		return fn(p + `Gecko/20100101 Firefox/##.#`)
	}
}

func RandBotAgent() string {
	ua := []string{
		`Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)`,
		`Mozilla/5.0 (compatible; YandexBot/3.0; +http://yandex.com/bots)`,
		`Mozilla/5.0 (compatible; Mail.RU_Bot/Fast/2.0)`,
		`StackRambler/2.0 (MSIE incompatible)`,
		`Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)`,
		`msnbot/1.1 (+http://search.msn.com/msnbot.htm)`,
		`Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)`,
	}

	return ua[rand.Intn(len(ua)-1)]
}
