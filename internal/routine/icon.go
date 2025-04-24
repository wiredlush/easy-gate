package routine

import (
	"encoding/base64"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/wiredlush/easy-gate/internal/config"
)

func (r *Routine) getIconData(service config.Service) template.URL {
	if strings.HasPrefix(service.Icon, "data:image") {
		return template.URL(service.Icon)
	}

	u, err := url.Parse(service.Icon)
	if err == nil && u.IsAbs() {
		return r.downloadIconFromURL(service.Icon)
	}

	u, err = url.Parse(service.URL)
	if err != nil {
		return ""
	}
	return r.downloadFavicon(fmt.Sprintf("%s://%s/%s", u.Scheme,
		u.Host, "favicon.ico"))
}

func (r *Routine) downloadIconFromURL(url string) template.URL {
	resp, err := r.Client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	mimeType := http.DetectContentType(respBytes)
	if !strings.HasPrefix(mimeType, "image/") {
		return ""
	}

	return template.URL(
		fmt.Sprintf(
			"data:%s;base64,%s", mimeType,
			base64.StdEncoding.EncodeToString(respBytes),
		),
	)
}

func (r *Routine) downloadFavicon(url string) template.URL {
	resp, err := r.Client.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return ""
	}

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	mimeType := http.DetectContentType(respBytes)
	if mimeType != "image/x-icon" {
		return ""
	}

	return template.URL(
		fmt.Sprintf(
			"data:image/x-icon;base64,%s",
			base64.StdEncoding.EncodeToString(respBytes),
		),
	)
}
