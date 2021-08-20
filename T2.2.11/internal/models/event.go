package models

//Order ...
type Event struct {
	UserId          int  `json:"user_id" validate:"required"`
	Date             string `json:"date" validate:"required"`
	Event 		string `json:"event" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateEventRequest struct{
	UserId          int  `json:"user_id" validate:"required"`
	Date             string `json:"date" validate:"required"`
	Event 		string `json:"event" validate:"required"`
	NewValue Event `json:"new_event"`
}

type DeleteEventRequest struct{
	UserId          int  `json:"user_id" validate:"required"`
	Date             string `json:"date" validate:"required"`
	Event 		string `json:"event" validate:"required"`
}

type EventsForDate struct{
	UserId          int  `json:"user_id" validate:"required"`
	Date             string `json:"date" validate:"required"`
}

type EventsForDateResponse struct{
	UserId          int  `json:"user_id" validate:"required"`
	Date             string `json:"date" validate:"required"`
	Event			[]Event `json:"events"`
}