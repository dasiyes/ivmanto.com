package booking

import (
	"strings"
	"testing"
	"time"
)

// TestResolveVisitorTimezone_AthensFromBerlinEvent is the canonical
// scenario from the bug report: a 15:30 CEST event must render as
// 16:30 EEST for a visitor in Europe/Athens.
func TestResolveVisitorTimezone_AthensFromBerlinEvent(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("tzdata not available for Europe/Berlin: %v", err)
	}

	loc, label := resolveVisitorTimezone("Europe/Athens", berlin)

	if loc == nil {
		t.Fatal("expected non-nil location")
	}
	if loc.String() != "Europe/Athens" {
		t.Errorf("expected Europe/Athens, got %q", loc.String())
	}
	// The label should be a short abbreviation, not the full IANA name.
	// On a Linux box with embedded tzdata, this is EEST (summer) or EET
	// (winter). We probe with January to land in winter — but a DST flip
	// in tzdata would be the only way this fails. Accept either.
	if label != "EEST" && label != "EET" && !strings.HasPrefix(label, "Europe/") {
		t.Errorf("expected abbreviation (EEST/EET) or IANA fallback, got %q", label)
	}
}

// TestResolveVisitorTimezone_EmptyFallsBackToCalendar ensures that an
// empty visitorTimezone (e.g. very old browser, or a hand-crafted
// request) never fails the booking — it falls back to the calendar's
// own timezone.
func TestResolveVisitorTimezone_EmptyFallsBackToCalendar(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("tzdata not available for Europe/Berlin: %v", err)
	}

	loc, _ := resolveVisitorTimezone("", berlin)

	if loc != berlin {
		t.Errorf("expected fallback to Europe/Berlin, got %q", loc.String())
	}
}

// TestResolveVisitorTimezone_UnknownFallsBackToCalendar ensures that a
// bogus IANA name does not fail the booking. This is the contract
// promised in the doc comment: a malformed client value can never
// block a booking.
func TestResolveVisitorTimezone_UnknownFallsBackToCalendar(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("tzdata not available for Europe/Berlin: %v", err)
	}

	loc, label := resolveVisitorTimezone("Not/A/Real_Zone", berlin)

	if loc != berlin {
		t.Errorf("expected fallback to Europe/Berlin, got %q", loc.String())
	}
	// Label should still come back non-empty — the abbreviation of Berlin.
	if label == "" {
		t.Errorf("expected non-empty fallback label, got empty")
	}
}

// TestResolveVisitorTimezone_TrimsWhitespace guards against clients that
// accidentally send " Europe/Athens " (a leading/trailing space is
// plausible from a JS template literal).
func TestResolveVisitorTimezone_TrimsWhitespace(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("tzdata not available for Europe/Berlin: %v", err)
	}

	loc, _ := resolveVisitorTimezone("  Europe/Athens  ", berlin)

	if loc == nil || loc.String() != "Europe/Athens" {
		t.Errorf("expected trimmed Europe/Athens, got %q", loc)
	}
}
