package goconfluence

import (
	"net/url"
)

// DownloadAttachment takes a fully qualified download url and
// and just downloads it. It will also handle the redirect and the
// token insertion via the underlying API.
func (a *API) DownloadAttachment(download string) ([]byte, error) {
	ep, err := url.ParseRequestURI(a.confluenceBase.String() + download)
	if err != nil {
		return nil, err
	}
	return a.SendDownloadRequest(ep, "GET")
}
