package regex_test

import (
	"testing"

	"github.com/DaanV2/go-projects-launcher/pkg/config"
	"github.com/DaanV2/go-projects-launcher/pkg/regex"
)

func TestIsMatch(t *testing.T) {
	tests := []struct {
		name           string
		config         config.Matching
		item           string
		patterns       []string
		expectedResult bool
	}{
		{
			name:           "Case insensitive match",
			config:         config.Matching{CaseSensitive: false},
			item:           "HelloWorld",
			patterns:       []string{"helloworld"},
			expectedResult: true,
		},
		{
			name:           "Case sensitive match",
			config:         config.Matching{CaseSensitive: true},
			item:           "HelloWorld",
			patterns:       []string{"HelloWorld"},
			expectedResult: true,
		},
		{
			name:           "Case sensitive no match",
			config:         config.Matching{CaseSensitive: true},
			item:           "HelloWorld",
			patterns:       []string{"helloworld"},
			expectedResult: false,
		},
		{
			name:           "Multiple patterns match",
			config:         config.Matching{CaseSensitive: false},
			item:           "HelloWorld",
			patterns:       []string{"foo", "bar", "helloworld"},
			expectedResult: true,
		},
		{
			name:           "No patterns match",
			config:         config.Matching{CaseSensitive: false},
			item:           "HelloWorld",
			patterns:       []string{"foo", "bar"},
			expectedResult: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := regex.IsMatch(tt.config, tt.item, tt.patterns...)
			if result != tt.expectedResult {
				t.Errorf("expected %v, got %v", tt.expectedResult, result)
			}
		})
	}
}
