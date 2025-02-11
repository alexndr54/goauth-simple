package web

import (
	"encoding/json"
	"github.com/patrickmn/go-cache"
	"os"
	"time"
)

type FiberTemp struct {
	MetaData   Metadata
	PageTitle  string
	IsLoggedIn bool
	Optional   map[string]string
}

type Metadata struct {
	Title            string
	ShortDescription string
	Description      string
	Keywords         string
	Author           string
	FaviconPath      string
	LogoPath         string
}

func GetMetadata() Metadata {
	mycache := cache.New(5*time.Minute, 10*time.Minute)
	if cached, found := mycache.Get("metadata"); found {
		return cached.(Metadata)
	}

	data, _ := os.ReadFile("./metadata.json")

	var metadata Metadata
	_ = json.Unmarshal(data, &metadata)

	mycache.Set("metadata", metadata, cache.DefaultExpiration)
	return metadata
}
