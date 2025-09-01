package ical

import (
	"bytes"
	"fmt"
	"net/textproto"
	"strings"
	"time"
)

// EventDetails contains all the necessary information to generate an iCalendar file.
type EventDetails struct {
	UID         string
	StartTime   time.Time
	EndTime     time.Time
	Summary     string
	Description string
	Location    string // Typically the Google Meet link
	Name        string // Client's name
	Email       string // Client's email
}

// Attachment represents an email attachment.
type Attachment struct {
	Headers textproto.MIMEHeader
	Body    []byte
}

// timeToUTCiCalFormat converts a time.Time to the iCalendar UTC format (YYYYMMDDTHHMMSSZ).
func timeToUTCiCalFormat(t time.Time) string {
	return t.UTC().Format("20060102T150405Z")
}

// escapeString escapes characters in a string according to iCalendar specs.
func escapeString(s string) string {
	s = strings.ReplaceAll(s, "\\", "\\\\")
	s = strings.ReplaceAll(s, ";", "\\;")
	s = strings.ReplaceAll(s, ",", "\\,")
	s = strings.ReplaceAll(s, "\n", "\\n")
	return s
}

// Generate creates an iCalendar (.ics) file content as a string.
func Generate(details EventDetails) string {
	var b bytes.Buffer

	b.WriteString("BEGIN:VCALENDAR\r\n")
	b.WriteString("VERSION:2.0\r\n")
	b.WriteString("PRODID:-//ivmanto.com//Booking Service//EN\r\n")
	b.WriteString("CALSCALE:GREGORIAN\r\n")
	b.WriteString("METHOD:REQUEST\r\n")
	b.WriteString("BEGIN:VEVENT\r\n")
	b.WriteString(fmt.Sprintf("UID:%s\r\n", details.UID))
	b.WriteString(fmt.Sprintf("DTSTAMP:%s\r\n", timeToUTCiCalFormat(time.Now())))
	b.WriteString(fmt.Sprintf("DTSTART:%s\r\n", timeToUTCiCalFormat(details.StartTime)))
	b.WriteString(fmt.Sprintf("DTEND:%s\r\n", timeToUTCiCalFormat(details.EndTime)))
	b.WriteString(fmt.Sprintf("SUMMARY:%s\r\n", escapeString(details.Summary)))
	b.WriteString(fmt.Sprintf("DESCRIPTION:%s\r\n", escapeString(details.Description)))
	b.WriteString(fmt.Sprintf("LOCATION:%s\r\n", escapeString(details.Location)))
	b.WriteString(fmt.Sprintf("ORGANIZER;CN=IVMANTO:mailto:no-reply@ivmanto.com\r\n")) // Using a generic organizer
	b.WriteString(fmt.Sprintf("ATTENDEE;CN=%s;ROLE=REQ-PARTICIPANT;PARTSTAT=NEEDS-ACTION;RSVP=TRUE:mailto:%s\r\n", escapeString(details.Name), details.Email))
	b.WriteString("STATUS:CONFIRMED\r\n")
	b.WriteString("SEQUENCE:0\r\n")
	b.WriteString("END:VEVENT\r\n")
	b.WriteString("END:VCALENDAR\r\n")

	return b.String()
}
