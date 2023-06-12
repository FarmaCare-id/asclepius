package dto

type (
	// EditUserProfileRequest EditUserProfileRequest
	EditUserProfileRequest struct {
		Fullname string `json:"fullname" validate:"omitempty"`
		Password string `json:"password" validate:"omitempty"`
		Weight   float64 `json:"weight" validate:"omitempty"`
		Height   float64 `json:"height" validate:"omitempty"`
		Age      int `json:"age" validate:"omitempty"`
		NoSip    string `json:"no_sip" validate:"omitempty"`
		NoSipa   string `json:"no_sipa" validate:"omitempty"`
		Specialist string `json:"specialist" validate:"omitempty"`
		Title    string `json:"title" validate:"omitempty"`
	}

	// EditUserProfileResponse EditUserProfileResponse
	EditUserProfileResponse struct {
		Fullname string `json:"fullname"`
		Weight   float64 `json:"weight"`
		Height   float64 `json:"height"`
		Age      int `json:"age"`
		NoSip    string `json:"no_sip"`
		NoSipa   string `json:"no_sipa"`
		Specialist string `json:"specialist"`
		Title    string `json:"title"`
	}
)