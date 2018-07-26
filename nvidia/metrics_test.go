package nvidia

import "testing"

func TestMetrics_NewMetrics(t *testing.T) {
	m := NewMetrics()
	output, _ := m.Get("test", "query")
	for _, o := range output {
		if o != nil {
			t.Errorf("output has to be nil.")
		}
	}
}
