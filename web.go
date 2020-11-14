package fake

import "math/rand"

func RandUTM() (source, medium, campaign string) {
	sourceList := []string{`google`, `yandex`, `vk`, `facebook`, `targetmail`, `instagram`, `youtube`, `twitter`}
	mediumList := []string{`cpc`, `email`, `banner`, `article`, `cpm`, `smm`, `target`, `referral`, `retargeting`}
	campaignList := []string{`promo`, `discount`, `sale`, RandGuid()}

	return sourceList[rand.Intn(len(sourceList)-1)],
		mediumList[rand.Intn(len(mediumList)-1)],
		campaignList[rand.Intn(len(campaignList)-1)]
}

func RandReferrer() string {
	refs := []string{
		`https://www.google.com/`,
		`https://yandex.ru/`,
		`https://search.yahoo.com/`,
		`https://go.mail.ru/`,
		`https://www.bing.com/`,
		`https://www.ask.com/`,
		`https://search.aol.com/`,
		`http://www.baidu.com/`,
	}

	return refs[rand.Intn(len(refs)-1)]
}
