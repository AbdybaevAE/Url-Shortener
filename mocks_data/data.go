package mocks_data

import pbLink "github.com/abdybaevae/url-shortener/proto"

const BasicLink = "https://google.com"

const BasicLinkShortenedKey = "some_key"

const InvalidLink = "https://https://google.com/httpss"

var BasicShoretenLinkReq = &pbLink.ShortenReq{
	Link: BasicLink,
}
var InvalidShoretenLinkReq = &pbLink.ShortenReq{
	Link: InvalidLink,
}
var BasicShoretenLinkRes = &pbLink.ShortenRes{
	ShortLink: BasicLinkShortenedKey,
}
