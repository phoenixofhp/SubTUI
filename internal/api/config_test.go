package api

import (
	"testing"

	"github.com/pelletier/go-toml/v2"
)

// Existing config files do not contain new sections. LoadConfig starts from the
// embedded defaults and applies the user config on top, so the defaults here
// are what users without a [cache] section get.
func TestDefaultConfigProvidesCacheDefaults(t *testing.T) {
	var config Config
	if err := toml.Unmarshal(defaultConfig, &config); err != nil {
		t.Fatalf("decode default config: %v", err)
	}

	cache := config.Cache.CoverArt
	if !cache.Enabled {
		t.Error("cover art cache is disabled by default")
	}
	if cache.Dir != "" {
		t.Errorf("default cover art cache dir = %q, want an empty value (user cache dir)", cache.Dir)
	}
	if cache.MaxSizeMB != 256 {
		t.Errorf("default cover art cache size = %d MiB, want 256", cache.MaxSizeMB)
	}
}

func TestUserConfigOverridesCacheDefaults(t *testing.T) {
	var config Config
	if err := toml.Unmarshal(defaultConfig, &config); err != nil {
		t.Fatalf("decode default config: %v", err)
	}

	userConfig := []byte(`
[cache.cover_art]
enabled = false
dir = '/tmp/covers'
max_size_mb = 8
`)

	if err := toml.Unmarshal(userConfig, &config); err != nil {
		t.Fatalf("decode user config: %v", err)
	}

	cache := config.Cache.CoverArt
	if cache.Enabled {
		t.Error("user override did not disable the cache")
	}
	if cache.Dir != "/tmp/covers" {
		t.Errorf("dir = %q, want /tmp/covers", cache.Dir)
	}
	if cache.MaxSizeMB != 8 {
		t.Errorf("max size = %d MiB, want 8", cache.MaxSizeMB)
	}
}
