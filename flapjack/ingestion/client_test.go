package ingestion

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetDefaultHostsUsesFlapjackIngestionHost(t *testing.T) {
	hosts := getDefaultHosts(US)
	if len(hosts) != 1 {
		t.Fatalf("expected exactly one default host, got %d", len(hosts))
	}

	scheme := reflect.ValueOf(hosts[0]).FieldByName("scheme").String()
	host := reflect.ValueOf(hosts[0]).FieldByName("host").String()
	if scheme != "https" {
		t.Fatalf("expected https scheme, got %q", scheme)
	}
	if host != "data.us.flapjack.io" {
		t.Fatalf("expected Flapjack ingestion host, got %q", host)
	}
	if strings.Contains(host, "algolia") {
		t.Fatalf("default ingestion host must not target Algolia: %q", host)
	}
}
