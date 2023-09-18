package domain

type CreatePollRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	StartsAt    string `json:"startsAt"`
	EndsAt      string `json:"endsAt"`
	Timezone    string `json:"timezone"`
}

type UpdatePollRequest struct {
	CreatePollRequest
	Id          int64 `json:"id"`
	OrganizerId int64 `json:"organizerId"`
}

type CreateUserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
