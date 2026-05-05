package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 60, Capacity: 82, Latency: 15, Risk: 8, Weight: 9}
	if got := Score(signal); got != 104 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 74, Capacity: 87, Latency: 22, Risk: 9, Weight: 10}
	if got := Score(signal); got != 104 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 84, Capacity: 86, Latency: 16, Risk: 18, Weight: 7}
	if got := Score(signal); got != 78 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
