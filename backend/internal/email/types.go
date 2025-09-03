package email

import "time"

// ContactMessage holds the data from the contact form.
type ContactMessage struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Message        string `json:"message"`
	SendCopyToSelf bool   `json:"sendCopyToSelf"`
}

// BookingConfirmationDetails holds all the necessary information
// to send a booking confirmation email.
type BookingConfirmationDetails struct {
	ToName          string
	ToEmail         string
	StartTime       time.Time
	EndTime         time.Time
	Timezone        string
	MeetLink        string
	CancellationURL string
	IcsUID          string
	IcsSummary      string
	IcsDescription  string
}

// GeneratedIdea holds the data for a single AI-generated idea.
type GeneratedIdea struct {
	Title   string
	Summary string
}
