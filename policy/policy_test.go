package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 60, Capacity: 82, Latency: 15, Risk: 8, Weight: 9}, wantScore: 104, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 74, Capacity: 87, Latency: 22, Risk: 9, Weight: 10}, wantScore: 104, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 84, Capacity: 86, Latency: 16, Risk: 18, Weight: 7}, wantScore: 78, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
