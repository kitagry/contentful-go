package contentful

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLocaleItem_MarshalJSON(t *testing.T) {
	tests := map[string]struct {
		data     LocaleItem[string]
		expected string
	}{
		"locale item": {
			data:     LocaleItem[string]{Map: map[string]string{"en-US": "hoge"}},
			expected: `{"en-US":"hoge"}`,
		},
		"not locale item": {
			data:     LocaleItem[string]{Item: toPtr("hoge")},
			expected: `"hoge"`,
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			assertions := assert.New(t)

			b, err := json.Marshal(tt.data)
			assertions.NoError(err)

			assertions.Equal(tt.expected, string(b))
		})
	}
}

func TestLocaleItem_UnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		data         string
		expectedMap  map[string]string
		expectedItem *string
	}{
		"locale item": {
			data:        `{"en-US": "hoge"}`,
			expectedMap: map[string]string{"en-US": "hoge"},
		},
		"not locale item": {
			data:         `"hoge"`,
			expectedItem: toPtr("hoge"),
		},
	}

	for n, tt := range tests {
		t.Run(n, func(t *testing.T) {
			assertions := assert.New(t)

			var localeItem LocaleItem[string]
			err := json.Unmarshal([]byte(tt.data), &localeItem)
			assertions.NoError(err)

			assertions.Equal(tt.expectedMap, localeItem.Map)
			assertions.Equal(tt.expectedItem, localeItem.Item)
		})
	}
}

func toPtr[T any](t T) *T {
	return &t
}

func TestNewCollection(t *testing.T) {
	setup()
	defer teardown()
}
