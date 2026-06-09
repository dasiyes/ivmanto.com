package ical

import (
	"strings"
	"testing"
	"time"
)

// TestGenerate_EmitsXWrTimezoneWhenProvided verifies that the visitor's IANA
// timezone, when supplied, is included as an X-WR-TIMEZONE header at the
// VCALENDAR level. DTSTART/DTEND must remain UTC (Z-suffixed) regardless,
// because the iCal spec is ambiguous on floating-vs-UTC and modern clients
// resolve the Z-suffixed UTC fields to local time correctly.
func TestGenerate_EmitsXWrTimezoneWhenProvided(t *testing.T) {
	start := time.Date(2026, 6, 15, 13, 30, 0, 0, time.UTC) // 15:30 CEST, 16:30 EEST
	end := time.Date(2026, 6, 15, 14, 0, 0, 0, time.UTC)

	out := Generate(EventDetails{
		UID:         "test-uid@ivmanto.com",
		StartTime:   start,
		EndTime:     end,
		Summary:     "Consultation: Marina",
		Description: "Test booking",
		Location:    "https://meet.google.com/abc-defg-hij",
		Name:        "Marina",
		Email:       "marina@example.com",
		Timezone:    "Europe/Athens",
	})

	if !strings.Contains(out, "X-WR-TIMEZONE:Europe/Athens\r\n") {
		t.Errorf("expected X-WR-TIMEZONE:Europe/Athens header, got:\n%s", out)
	}
	if !strings.Contains(out, "DTSTART:20260615T133000Z\r\n") {
		t.Errorf("expected DTSTART in UTC Z format, got:\n%s", out)
	}
	if !strings.Contains(out, "DTEND:20260615T140000Z\r\n") {
		t.Errorf("expected DTEND in UTC Z format, got:\n%s", out)
	}
}

// TestGenerate_OmitsXWrTimezoneWhenAbsent verifies that no X-WR-TIMEZONE
// header is emitted when the visitor's timezone is empty, so the .ics
// attachment remains compatible with all clients when no TZ info is
// available (e.g. legacy events that predate this feature).
func TestGenerate_OmitsXWrTimezoneWhenAbsent(t *testing.T) {
	start := time.Date(2026, 6, 15, 13, 30, 0, 0, time.UTC)
	end := time.Date(2026, 6, 15, 14, 0, 0, 0, time.UTC)

	out := Generate(EventDetails{
		UID:       "test-uid@ivmanto.com",
		StartTime: start,
		EndTime:   end,
		Summary:   "Consultation",
		Location:  "https://meet.google.com/abc-defg-hij",
		Name:      "Visitor",
		Email:     "visitor@example.com",
		// Timezone deliberately left empty
	})

	if strings.Contains(out, "X-WR-TIMEZONE") {
		t.Errorf("expected no X-WR-TIMEZONE header, got:\n%s", out)
	}
}

// TestGenerate_EscapesIANACharacters guards against IANA names that
// happen to contain semicolons or commas (none do today, but a future
// tzdata addition is plausible). iCal property values must escape these.
func TestGenerate_EscapesIANACharacters(t *testing.T) {
	start := time.Date(2026, 6, 15, 13, 30, 0, 0, time.UTC)
	end := time.Date(2026, 6, 15, 14, 0, 0, 0, time.UTC)

	out := Generate(EventDetails{
		UID:       "test-uid@ivmanto.com",
		StartTime: start,
		EndTime:   end,
		Summary:   "Consultation",
		Location:  "https://meet.google.com/abc-defg-hij",
		Name:      "Visitor",
		Email:     "visitor@example.com",
		Timezone:  "Fake/Zone;with,commas",
	})

	if !strings.Contains(out, `X-WR-TIMEZONE:Fake/Zone\;with\,commas`) {
		t.Errorf("expected escaped IANA value, got:\n%s", out)
	}
}
