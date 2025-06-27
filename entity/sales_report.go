package entity

type SalesReport struct {
	EventID        uint    `json:"event_id"`
	EventName      string  `json:"event_name"`
	TicketsSold    int     `json:"tickets_sold"`
	TotalRevenue   float64 `json:"total_revenue"`
	RemainingQuota int     `json:"remaining_quota"`
	Status         string  `json:"status"`
}