package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name       string
		headerKey  string
		headerVal  string
		wantAPIKey string
		wantErr    bool
	}{
		{
			name:       "Valid API Key",
			headerKey:  "Authorization",
			headerVal:  "ApiKey 123456",
			wantAPIKey: "123456",
			wantErr:    false,
		},
		{
			name:       "Invalid API Key",
			headerKey:  "Authorization",
			headerVal:  "badAPIKey",
			wantAPIKey: "",
			wantErr:    true,
		},
		{
			name:       "No Auth Header",
			headerKey:  "",
			headerVal:  "",
			wantAPIKey: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		header := http.Header{}
		header.Set(tt.headerKey, tt.headerVal)
		t.Run(tt.name, func(t *testing.T) {
			gotAPIKey, err := GetAPIKey(header)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotAPIKey != tt.wantAPIKey {
				t.Errorf("GetAPIKey() gotAPIKey = %v, want %v", gotAPIKey, tt.wantAPIKey)
			}
		})
	}
}
