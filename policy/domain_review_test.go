package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 50, Slack: 29, Drag: 27, Confidence: 71}
	if got := DomainReviewScore(item); got != 119 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "watch" {
		t.Fatalf("lane = %s", got)
	}
}
