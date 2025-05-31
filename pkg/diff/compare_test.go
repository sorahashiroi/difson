package diff

import (
	"os"
	"reflect"
	"testing"
)

func TestCompareJSONBrief(t *testing.T) {
	tests := []struct {
		name string
		a, b map[string]interface{}
		want bool
	}{
		{
			name: "equal",
			a:    map[string]interface{}{"name": "Jane", "age": 30},
			b:    map[string]interface{}{"name": "Jane", "age": 30},
			want: true,
		},
		{
			name: "not equal",
			a:    map[string]interface{}{"name": "Jane", "age": 30},
			b:    map[string]interface{}{"name": "Jane", "age": 31},
			want: false,
		},
	}

	for _, td := range tests {
		t.Run(td.name, func(t *testing.T) {
			got := CompareJSONBrief(td.a, td.b)
			if got != td.want {
				t.Errorf("CompareJSONBrief() = %v, want %v", got, td.want)
			}
		})
	}
}

func TestGetDiff(t *testing.T) {
	tests := []struct {
		name           string
		givenA, givenB map[string]interface{}
		want           map[string]interface{}
	}{
		{
			name:   "basic diff",
			givenA: map[string]interface{}{"age": 30},
			givenB: map[string]interface{}{"age": 31},
			want: map[string]interface{}{
				"age": map[string]interface{}{
					"from": 30,
					"to":   31,
				},
			},
		},
		{
			name:   "added key",
			givenA: map[string]interface{}{},
			givenB: map[string]interface{}{"city": "Tokyo"},
			want: map[string]interface{}{
				"city": map[string]interface{}{
					"from": "<missing>",
					"to":   "Tokyo",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetDiff(tt.givenA, tt.givenB)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDiff() = %v; want %v", got, tt.want)
			}
		})
	}
}

func TestLoadJSON(t *testing.T) {
	fileContent := `{"name": "Alice", "age": 30}`
	tmpFile, err := os.CreateTemp("", "testjson")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString(fileContent)
	tmpFile.Close()

	got, err := LoadJSON(tmpFile.Name())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	want := map[string]interface{}{"name": "Alice", "age": float64(30)} // JSON数値はfloat64

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
