package yaml

import (
	"bytes"
	"strings"
	"testing"
)

func TestChineseTextRoundTrip(t *testing.T) {
	input := `
meta:
  title: 中文论证图
statements:
  source_confirmed: 原始来源已经核验。
assumptions: [source_confirmed]
`

	graph, err := Import(strings.NewReader(input))
	if err != nil {
		t.Fatalf("import Chinese YAML: %v", err)
	}

	statement, ok := graph.Statements["source_confirmed"]
	if !ok {
		t.Fatal("Chinese YAML statement was not imported")
	}
	if statement.Text != "原始来源已经核验。" {
		t.Fatalf("statement text = %q", statement.Text)
	}

	var output bytes.Buffer
	Export(&output, graph)
	if !strings.Contains(output.String(), "原始来源已经核验。") {
		t.Fatalf("export lost Chinese text:\n%s", output.String())
	}

	if _, err := Import(bytes.NewReader(output.Bytes())); err != nil {
		t.Fatalf("re-import exported Chinese YAML: %v", err)
	}
}
