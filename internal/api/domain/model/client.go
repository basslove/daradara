package model

import (
	"github.com/stretchr/stew/slice"
)

type ClientPlatformType string

const (
	ClientPlatformWeb     ClientPlatformType = "web"
	ClientPlatformIOS     ClientPlatformType = "ios"
	ClientPlatformAndroid ClientPlatformType = "android"
	ClientPlatformUnknown ClientPlatformType = "unknown"
)

type Client struct {
	Version       string
	SystemVersion string
	Platform      ClientPlatformType
	Device        string
}

func NewClient(version, systemVersion, platForm, device string) *Client {
	return &Client{
		Version:       version,
		SystemVersion: systemVersion,
		Platform:      NewClientPlatformType(platForm),
		Device:        device,
	}
}

func NewClientPlatformType(platform string) ClientPlatformType {
	switch platform {
	case "web":
		return ClientPlatformWeb
	case "ios":
		return ClientPlatformIOS
	case "android":
		return ClientPlatformAndroid
	case "":
		return ClientPlatformUnknown
	}
	return ClientPlatformUnknown
}

func (m *Client) IsWeb() bool {
	return slice.Contains([]ClientPlatformType{ClientPlatformWeb}, m.Platform)
}

func (m *Client) IsNativeApp() bool {
	return slice.Contains([]ClientPlatformType{ClientPlatformIOS, ClientPlatformAndroid}, m.Platform)
}
