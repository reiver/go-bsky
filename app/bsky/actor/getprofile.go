package actor

import (
	gourl "net/url"

	"github.com/reiver/go-xrpc"

	"github.com/reiver/go-bsky/internal/config"
)

func GetProfile(dst any, actor string) error {

	const nsid string = "app.bsky.actor.getProfile"

	var query string = "actor=" + gourl.QueryEscape(actor)

	var xrpcURL xrpc.URL = xrpc.URL{
		Host: config.Host,
		NSID: nsid,
		Query: query,
	}

	return xrpc.Query(dst, xrpcURL.String())
}
