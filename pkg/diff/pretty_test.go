package diff

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrintPretty(t *testing.T) {
	// 差分データ（入力）
	diffData := map[string]interface{}{
		"age": map[string]interface{}{
			"from": 30,
			"to":   31,
		},
		"name": map[string]interface{}{
			"from": "Alice",
			"to":   "Bob",
		},
	}

	// 標準出力をキャプチャ
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	PrintPretty(diffData)

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)

	output := buf.String()

	// 整形されたJSONを含んでいるかの基本的な確認
	expectedStrings := []string{
		`"age"`, `"from": 30`, `"to": 31`,
		`"name"`, `"from": "Alice"`, `"to": "Bob"`,
		"\n  ", // インデントがあるか
	}

	for _, str := range expectedStrings {
		if !strings.Contains(output, str) {
			t.Errorf("expected output to contain %q, but it did not. Full output:\n%s", str, output)
		}
	}
}
