package requestResponsemodels

type (
	ChangeEngineRequest struct {
		Query string `json:"query" binding:"required"`
	}
	ChangeDirPathRequest struct {
		Query string `json:"query" binding:"required"`
	}
	SearchRequest struct {
		Query string `json:"query" binding:"required"`
	}
	SignUpRequest struct {
		UserName string `json:"user_name" validate:"required"`
		Password string `json:"password" validate:"required"`
		Email    string `json:"email"`
	}
	SignInRequest struct {
		UserName string `json:"user_name"`
		Password string `json:"password"`
	}
	RemoveUserRequest struct {
		UserName string `json:"user_name"`
	}
)
type (
	SearchResponse struct {
		DocId         int `json:"Docid"`
		TermFrequency int `json:"TermFrequency"`
	}
)
