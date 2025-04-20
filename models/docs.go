package models

import "time"


type ErrorBadRequestResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"400"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Bad Request"`
}

type ErrorInternalServerErrorResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"500"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Internal Server Error"`
}

type ErrorNotFoundResponse struct {
	Status     string      `json:"status" example:"error"`
	StatusCode int         `json:"status_code" example:"404"`
	Success    bool        `json:"success" example:"false"`
	Message    string      `json:"message" example:"Not Found"`
}


type CreateResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"201"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Created successfully"`
}

type UpdateResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Updated successfully"`
}

type DeleteResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Deleted successfully"`
}

type CreateIngredientRequest struct {
	Name string `json:"name" example:"Chicken"`
	Unit string `json:"unit" example:"kg"`
	Quantity string `json:"quantity" example:"100"`
	Description string `json:"description" example:"onion"`
}

type CreateInstructionRequest struct {
	Step string `json:"step" example:"1"`
	ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
	Description string `json:"description" example:"Chop the chicken into small cubes or finely mince it (you can also use a food processor). This makes it easier to form patties later."`
}

type UpdateIngredientRequest struct {
	Name string `json:"name" example:"Chicken"`
	Unit string `json:"unit" example:"kg"`
	Quantity string `json:"quantity" example:"100"`
	Description string `json:"description" example:"onion"`
}

type UpdateInstructionRequest struct {
	Step string `json:"step" example:"1"`
	ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
	Description string `json:"description" example:"Chop the chicken into small cubes or finely mince it (you can also use a food processor). This makes it easier to form patties later."`
}

type UpdateRecipeRequest struct {
	Title string `json:"title" example:"Chicken Wings"`
	Description string `json:"description" example:"A feast for the senses"`
	EstimatedTime string `json:"estimated_time" example:"20 minutes"`
	ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
	Ingredients []struct {
		Name string `json:"name" example:"Chicken"`
		Unit string `json:"unit" example:"kg"`
		Quantity string `json:"quantity" example:"100"`
		Description string `json:"description" example:"onion"`
	} `json:"ingredients"`	
	Instructions []struct {
		Step string `json:"step" example:"1"`
		Description string `json:"description" example:"Chop the chicken into small cubes or finely mince it (you can also use a food processor). This makes it easier to form patties later."`
	} `json:"instructions"`
}

type DeleteRecipeRequest struct {
	Status string `json:"status" example:"success"`
	StatusCode int `json:"status_code" example:"200"`
	Success bool `json:"success" example:"true"`
	Message string `json:"message" example:"Deleted successfully"`
}
type GetAllRecipesResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved recipes"`
	Data       []struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		Description string          `json:"description" example:"A feast for the senses"`
		EstimatedTime string          `json:"estimated_time" example:"20 minutes"`
		ImageUrl    string          `json:"image_url" example:"https://example.com/image.png"`
		CreatedBy   string          `json:"created_by" example:"@facu.potti"`
		TotalComments  int     `json:"total_comments" example:"64"`
		TotalFavorites int     `json:"total_favorites" example:"4445"`
		TotalStars     float64 `json:"total_stars" example:"4"`
		Category       string  `json:"category" example:"Chicken"`
		SearchParam string          `json:"searchParam" example:"chicken"`
	} `json:"data"`
}


type GetRecipeByIdResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved recipe"`
	Data       []struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		Description string          `json:"description" example:"A feast for the senses"`
		EstimatedTime string          `json:"estimated_time" example:"20 minutes"`
		ImageUrl    string          `json:"image_url" example:"https://example.com/image.png"`
		CreatedBy   string          `json:"created_by" example:"@facu.potti"`
		TotalComments  int     `json:"total_comments" example:"64"`
		TotalFavorites int     `json:"total_favorites" example:"4445"`
		TotalStars     float64 `json:"total_stars" example:"4"`
		Ingredients []struct {
			Name string `json:"name" example:"Chicken"`
			Unit string `json:"unit" example:"kg"`
			Quantity string `json:"quantity" example:"100"`
			Description string `json:"description" example:"onion"`
		} `json:"ingredients"`
		Instructions []struct {
			Step string `json:"step" example:"1"`
			ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
			Description string `json:"description" example:"Chop the chicken into small cubes or finely mince it (you can also use a food processor). This makes it easier to form patties later."`
		} `json:"instructions"`
		Comments []struct {
			ID          int             `json:"id" example:"1"`
			Comment string `json:"comment" example:"This is a comment"`
			CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
			CreatedBy string `json:"created_by" example:"@facu.potti"`
		} `json:"comments"`
	} `json:"data"`
}


type GetAllCollectionsResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved collections"`
	Data       []struct {
		ID    int    `json:"id" example:"1"`
		Title string `json:"title" example:"Salty"`
		Favorites []struct {
			ID    int    `json:"id" example:"1"`
			Title string `json:"title" example:"Chicken Wings"`
			ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
			TotalStars     float64 `json:"total_stars" example:"4.5"`
			EstimatedTime string          `json:"estimated_time" example:"25 minutes"`
		} `json:"favorites"`
	} `json:"data"`
}

type GetCollectionByIdResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved collection"`
	Data       struct {
		ID    int    `json:"id" example:"1"`
		Title string `json:"title" example:"Salty"`
		Favorites []struct {
			ID    int    `json:"id" example:"1"`
			Title string `json:"title" example:"Chicken Wings"`
			ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
			TotalStars     float64 `json:"total_stars" example:"4.5"`
			EstimatedTime string          `json:"estimated_time" example:"25 minutes"`
		} `json:"favorites"`
	} `json:"data"`
}

type CreateCollectionRequest struct {
	Title string `json:"title" example:"Salty"`
}


type GetAllSearchesResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved searches"`
	Data       []struct {
		ID    int    `json:"id" example:"1"`
		SearchTerm string `json:"search_term" example:"Salty"`
		CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
		CreatedBy string `json:"created_by" example:"@facu.potti"`
	} `json:"data"`
}


type GetProfileResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved profile"`
	Data       struct {
		ID          int             `json:"id" example:"1"`
		Name        string          `json:"name" example:"Facu Potti"`
		Username    string          `json:"username" example:"@facu.potti"`
		Email       string          `json:"email" example:"facu.potti@gmail.com"`
		Description string          `json:"description" example:"I'm a software engineer"`
		TotalFollowers int     `json:"total_followers" example:"250"`
		TotalFollowing int     `json:"total_following" example:"120"`
		TotalPosts int     `json:"total_posts" example:"120"`
		ProfilePicture string          `json:"profile_picture" example:"https://example.com/image.png"`

	} `json:"data"`
}	

type GetFavoritesResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved favorites"`
	Data       []struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		ImageUrl    string          `json:"image_url" example:"https://example.com/image.png"`
		TotalStars     float64 `json:"total_stars" example:"4.5"`
		EstimatedTime string          `json:"estimated_time" example:"25 minutes"`
	} `json:"data"`
}


type GetChangeResponse struct{
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved change"`
	Data       []struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		ImageUrl    string          `json:"image_url" example:"https://example.com/image.png"`
		TotalStars     float64 `json:"total_stars" example:"4.5"`
		EstimatedTime string          `json:"estimated_time" example:"25 minutes"`
	} `json:"data"`
}


type GetUserByIdResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved user"`
	Data       struct {
		ID          int             `json:"id" example:"1"`
		Name        string          `json:"name" example:"Facu Potti"`
		Username    string          `json:"username" example:"@facu.potti"`
		Email       string          `json:"email" example:"facu.potti@gmail.com"`
		Description string          `json:"description" example:"I'm a software engineer"`
		ProfilePicture string          `json:"profile_picture" example:"https://example.com/image.png"`
		SocialMedia []struct {
			URL string `json:"url" example:"https://twitter.com/facu.potti"`
		} `json:"social_media"`
	} `json:"data"`
}

type UpdateUserRequest struct {
	Name string `json:"name" example:"Facu Potti"`
	Username string `json:"username" example:"@facu.potti"`
	Email string `json:"email" example:"facu.potti@gmail.com"`
	Description string `json:"description" example:"I'm a software engineer"`
	ProfilePicture string `json:"profile_picture" example:"https://example.com/image.png"`
	SocialMedia []struct {
		URL string `json:"url" example:"https://twitter.com/facu.potti"`
	} `json:"social_media"`
}


type GetInstructionByIdResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved instruction"`
	Data       struct {
		ID          int             `json:"id" example:"1"`
		Step string `json:"step" example:"1"`
		ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
		Description string `json:"description" example:"Chop the chicken into small cubes or finely mince it (you can also use a food processor). This makes it easier to form patties later."`
		Ingredients []struct {
			ID          int             `json:"id" example:"1"`
			Name string `json:"name" example:"Chicken"`
			Unit string `json:"unit" example:"kg"`
			Quantity string `json:"quantity" example:"100"`
			Description string `json:"description" example:"onion"`
		} `json:"ingredients"`
	} `json:"data"`
}


type GetCommentByIdResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved comment"`
	Data       struct {
		ID          int             `json:"id" example:"1"`
		Comment string `json:"comment" example:"This is a comment"`
		CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
		CreatedBy string `json:"created_by" example:"@facu.potti"`
		TotalLikes int `json:"total_likes" example:"100"`
		FatherComment struct {
			ID          int             `json:"id" example:"1"`
			Comment string `json:"comment" example:"This is a comment"`
			CreatedAt time.Time `json:"created_at" example:"2021-01-01T00:00:00Z"`
			CreatedBy string `json:"created_by" example:"@facu.potti"`
		} `json:"father_comment"`	
	} `json:"data"`
}

type UpdateCommentRequest struct {
	Comment string `json:"comment" example:"This is a comment"`
}


type CreateCommentRequest struct {
	Comment string `json:"comment" example:"This is a comment"`
	FatherCommentID int `json:"father_comment_id" example:"1"`
	RecipeID int `json:"recipe_id" example:"1"`
	
}

type CreateRecipeRequest struct {
	Title string `json:"title" example:"Chicken Wings"`
	Description string `json:"description" example:"A feast for the senses"`
	EstimatedTime string `json:"estimated_time" example:"20 minutes"`
	ImageUrl string `json:"image_url" example:"https://example.com/image.png"`
}


type GetAllRecipeAdminResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved recipe"`
	Data       []struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		Description string          `json:"description" example:"A feast for the senses"`
		EstimatedTime string          `json:"estimated_time" example:"20 minutes"`
		ImageUrl    string          `json:"image_url" example:"https://example.com/image.png"`
	} `json:"data"`	
}

type GetAllNotificationsResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved notifications"`
	Data       []struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		Description string          `json:"description" example:"A feast for the senses"`
		Type string `json:"type" example:"follow"`
		ActorID int `json:"actor_id" example:"1"`
		IsNew bool `json:"is_new" example:"true"`
		CreatedBy string `json:"created_by" example:"@facu.potti"`
	} `json:"data"`
}

type GetNotificationsByIdResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved notification"`
	Data       struct {
		ID          int             `json:"id" example:"1"`
		Title       string          `json:"title" example:"Chicken Wings"`
		Description string          `json:"description" example:"A feast for the senses"`
		Type string `json:"type" example:"follow"`
		ActorID int `json:"actor_id" example:"1"`
		IsNew bool `json:"is_new" example:"true"`
		CreatedBy string `json:"created_by" example:"@facu.potti"`
	} `json:"data"`
}

type CreateNotificationRequest struct {
	Title string `json:"title" example:"Chicken Wings"`
	Description string `json:"description" example:"A feast for the senses"`
	Type string `json:"type" example:"follow"`
	ActorID int `json:"actor_id" example:"1"`
	IsNew bool `json:"is_new" example:"true"`
	CreatedBy string `json:"created_by" example:"@facu.potti"`
}


type CreateFollowRequest struct {
	UserID int `json:"user_id" example:"1"`
}

type DeleteFollowRequest struct {
	UserID int `json:"user_id" example:"1"`
}	

type GetMultipleResponse struct {
	Status     string      `json:"status" example:"success"`
	StatusCode int         `json:"status_code" example:"200"`
	Success    bool        `json:"success" example:"true"`
	Message    string      `json:"message" example:"Successfully retrieved multiple"`
	Data       struct {
		ID          int             `json:"id" example:"1"`
		Half       bool          `json:"half" example:"true"`
		Duplicate 	bool          `json:"duplicate" example:"true"`
		RecipeID int `json:"recipe_id" example:"1"`
		UserID int `json:"user_id" example:"1"`
		SetServing struct {
			ID int `json:"id" example:"1"`
			Active bool `json:"active" example:"true"`
			Quantity int `json:"quantity" example:"1"`
		}
		AmountIngredient struct {
			ID int `json:"id" example:"1"`
			Quantity int `json:"quantity" example:"1"`
			Unit string `json:"unit" example:"kg"`
			IngredientID int `json:"ingredient_id" example:"1"`
		} `json:"set_serving"`
	} `json:"data"`
}

type CreateMultipleRequest struct {
	Half bool `json:"half" example:"true"`
}
