package parser

import (
	"testing"

	"github.com/severity1/claude-agent-sdk-go/internal/shared"
)

// TestParseMessageRateLimitEvent verifies that the parser does not error on
// `rate_limit_event` telemetry frames emitted by recent Claude CLI versions.
// They are surfaced as SystemMessage{Subtype: "rate_limit_event"} so consumers
// can read the payload via Data without the iteration aborting.
func TestParseMessageRateLimitEvent(t *testing.T) {
	p := New()
	data := map[string]any{
		"type": "rate_limit_event",
		"rate_limit": map[string]any{
			"resets_at":  "2026-04-28T12:00:00Z",
			"percentage": 0.8,
		},
	}

	msg, err := p.ParseMessage(data)
	if err != nil {
		t.Fatalf("ParseMessage returned error for rate_limit_event: %v", err)
	}
	if msg == nil {
		t.Fatalf("ParseMessage returned nil message for rate_limit_event")
	}
	sys, ok := msg.(*shared.SystemMessage)
	if !ok {
		t.Fatalf("expected *SystemMessage, got %T", msg)
	}
	if sys.Subtype != "rate_limit_event" {
		t.Fatalf("expected Subtype=rate_limit_event, got %q", sys.Subtype)
	}
	if sys.Data["type"] != "rate_limit_event" {
		t.Fatalf("expected Data to preserve original payload, got %v", sys.Data)
	}
}
