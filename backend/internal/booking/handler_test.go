package booking

import (
	"testing"
	"time"
)

// TestResolveVisitorTimezone_AthensFromBerlinEvent is the canonical
// scenario from the bug report. The visitor is in Europe/Athens, the
// event is in the calendar's own Europe/Berlin zone. The helper
// must resolve Athens.
func TestResolveVisitorTimezone_AthensFromBerlinEvent(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("tzdata not available for Europe/Berlin: %v", err)
	}

	loc := resolveVisitorTimezone("Europe/Athens", berlin)

	if loc == nil {
		t.Fatal("expected non-nil location")
	}
	if loc.String() != "Europe/Athens" {
		t.Errorf("expected Europe/Athens, got %q", loc.String())
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

	loc := resolveVisitorTimezone("", berlin)

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

	loc := resolveVisitorTimezone("Not/A/Real_Zone", berlin)

	if loc != berlin {
		t.Errorf("expected fallback to Europe/Berlin, got %q", loc.String())
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

	loc := resolveVisitorTimezone("  Europe/Athens  ", berlin)

	if loc == nil || loc.String() != "Europe/Athens" {
		t.Errorf("expected trimmed Europe/Athens, got %q", loc)
	}
}

// TestVisitorLabelFollowsDST is the regression test for the bug caught
// in PR review: the helper must NOT embed a fixed-probe abbreviation
// (Jan 1 always returns EET for Athens even when the booking is in
// June, which is EEST). The label is computed at the call site from
// the event's actual start time, so:
//   - a 15:30 CEST event in June labelled for an Athens visitor = EEST
//   - a 15:30 CET  event in January labelled for an Athens visitor = EET
func TestVisitorLabelFollowsDST(t *testing.T) {
	berlin, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		t.Skipf("tzdata not available for Europe/Berlin: %v", err)
	}

	// June event: 15:30 CEST = 13:30 UTC
	summerStart, err := time.Parse(time.RFC3339, "2026-06-15T15:30:00+02:00")
	if err != nil {
		t.Fatalf("could not parse summer event time: %v", err)
	}
	// January event: 15:30 CET = 14:30 UTC
	winterStart, err := time.Parse(time.RFC3339, "2026-01-15T15:30:00+01:00")
	if err != nil {
		t.Fatalf("could not parse winter event time: %v", err)
	}

	visitorLoc := resolveVisitorTimezone("Europe/Athens", berlin)

	summerLabel := summerStart.In(visitorLoc).Format("MST")
	winterLabel := winterStart.In(visitorLoc).Format("MST")

	if summerLabel != "EEST" {
		t.Errorf("summer Athens label: expected EEST, got %q", summerLabel)
	}
	if winterLabel != "EET" {
		t.Errorf("winter Athens label: expected EET, got %q", winterLabel)
	}

	// Sanity: the same instants in the Berlin zone for the admin-facing
	// label should also follow DST (CEST in summer, CET in winter).
	adminLoc := resolveVisitorTimezone("", berlin)
	if got := summerStart.In(adminLoc).Format("MST"); got != "CEST" {
		t.Errorf("summer Berlin label: expected CEST, got %q", got)
	}
	if got := winterStart.In(adminLoc).Format("MST"); got != "CET" {
		t.Errorf("winter Berlin label: expected CET, got %q", got)
	}
}
