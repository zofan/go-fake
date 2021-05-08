package fake

import "math/rand"

func UTM() (source, medium, campaign string) {
	sourceList := []string{`google`, `yandex`, `vk`, `facebook`, `targetmail`, `instagram`, `youtube`, `twitter`}
	mediumList := []string{`cpc`, `email`, `banner`, `article`, `cpm`, `smm`, `target`, `referral`, `retargeting`}
	campaignList := []string{`promo`, `discount`, `sale`, Guid()}

	return sourceList[rand.Intn(len(sourceList)-1)],
		mediumList[rand.Intn(len(mediumList)-1)],
		campaignList[rand.Intn(len(campaignList)-1)]
}

func Referrer() string {
	switch rand.Intn(3) {
	default:
		return ReferrerSearch()
	case 1:
		return ReferrerShortURL()
	case 2:
		return ReferrerSocial()
	case 3:
		return ReferrerQuestion()
	}
}

func ReferrerSearch() string {
	refs := []string{
		`https://www.google.com/`,
		`https://yandex.ru/`,
		`https://search.yahoo.com/`,
		`https://go.mail.ru/`,
		`https://www.bing.com/`,
		`https://search.aol.com/`,
		`https://www.baidu.com/`,
	}

	return refs[rand.Intn(len(refs)-1)]
}

func ReferrerShortURL() string {
	refs := []string{
		`https://bit.ly/`,
		`https://bit.do/`,
		`https://clck.ru/`,
		`https://tinyurl.com/`,
		`https://to.click/`,
		`https://u.to/`,
		`https://vk.cc/`,
	}

	return refs[rand.Intn(len(refs)-1)]
}

func ReferrerSocial() string {
	refs := []string{
		`https://www.facebook.com/`,
		`https://ok.ru/`,
		`https://vk.com/`,
		`https://www.instagram.com/`,
		`https://twitter.com/`,
		`https://www.pinterest.ru/`,
		`https://www.reddit.com/`,
		`https://www.tumblr.com/`,
		`https://www.flickr.com/`,
	}

	return refs[rand.Intn(len(refs)-1)]
}

func ReferrerQuestion() string {
	refs := []string{
		`https://ask.fm/`,
		`https://yandex.ru/q/`,
		`http://vorum.ru/`,
		`https://qna.habr.com/`,
		`https://stackoverflow.com/`,
	}

	return refs[rand.Intn(len(refs)-1)]
}
