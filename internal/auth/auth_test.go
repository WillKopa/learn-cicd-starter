package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	token := "frodo"
	prefix := "ApiKey "

	header1 := http.Header{}
	header1.Set("Authorization", prefix+token)

	no_prefix_header := http.Header{}
	no_prefix_header.Set("Authorization", token)

	tests := []struct {
		name           string
		expected_token string
		header         http.Header
		expect_error   bool
	}{
		{
			name:           "Header with correct syntax",
			expected_token: token,
			header:         header1,
			expect_error:   false,
		},
		{
			name:           "Header with no auth",
			expected_token: "",
			header:         http.Header{},
			expect_error:   true,
		},
		{
			name:           "Header without API key",
			expected_token: "",
			header:         no_prefix_header,
			expect_error:   true,
		},
	}

	for _, token_test := range tests {
		t.Run(token_test.name, func(t *testing.T) {
			result, err := GetAPIKey(token_test.header)
			if (err != nil) && !token_test.expect_error {
				t.Errorf("error calling GetAPIKey %v", err)
			}
			if result != token_test.expected_token {
				t.Errorf("error calling GetAPIKey. Got %s, but expected %s", result, token_test.expected_token)
			}
		})
	}
}
