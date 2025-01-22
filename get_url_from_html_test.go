package main

import (
	"reflect"
	"testing"
)

func TestGetUrlsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
		<a href="/path/one">
			<span>example.dev</span>
		</a>
		<a href="https://other.com/path/one">
			<span>Example.dev</span>
		</a>
	</body>
</html>
`,
			expected: []string{"https://blog.example.dev/path/one", "https://other.com/path/one"},
		},
		{
			name:     "single absolute URL",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
			<a href="https://other.com/path/one">
				<span>Something</span>
			</a>
	</body>			
</html>
`,
			expected: []string{"https://other.com/path/one"},
		},
		{
			name:     "single relative URL",
			inputURL: "https://blog.example.dev",
			inputBody: `
<html>
	<body>
			<a href="/path/one">
				<span>Something</span>
			</a>
	</body>			
</html>
`,
			expected: []string{"https://blog.example.dev/path/one"},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputBody, tc.inputURL)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("expected URLs: %v, got: %v", tc.expected, actual)
			}
		})
	}
}
