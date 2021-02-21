// template version: 1.0.11
package url

import (
	"time"
)

// Url ,
type Url struct {
	// ID ,
	ID int `json:"id"`
	// URL ,
	URL string `json:"url"`
	// Shortcode ,
	Shortcode string `json:"shortcode"`
	// RedirectCount ,
	RedirectCount int `json:"redirectCount"`
	// StartDate ,
	StartDate time.Time `json:"startDate"`
	// LastSeenDate ,
	LastSeenDate time.Time `json:"lastSeenDate"`
}
