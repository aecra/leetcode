package problem2102

import "testing"

func TestSORTracker(t *testing.T) {
	obj := Constructor()
	metholds := []string{"add", "add", "get", "add", "get", "add", "get", "add", "get", "add", "get", "get"}
	params := []Sight{{"bradford", 2}, {"branford", 3}, {}, {"alps", 2}, {}, {"orland", 2}, {}, {"orlando", 3}, {}, {"alpine", 2}, {}, {}}
	expected := []string{"", "", "branford", "", "alps", "", "bradford", "", "bradford", "", "bradford", "orland"}
	for i, m := range metholds {
		if m == "add" {
			obj.Add(params[i].name, params[i].score)
		} else {
			if obj.Get() != expected[i] {
				t.Errorf("expected %s, but got %s", expected[i], obj.Get())
			}
		}
	}
}
