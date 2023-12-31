// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type CreateTaskInput struct {
	OwnerID                string    `json:"ownerId" bson:"ownerId"`
	Description            string    `json:"description" bson:"description"`
	Category               *string   `json:"category,omitempty" bson:"category"`
	TaskRequirements       *string   `json:"taskRequirements,omitempty" bson:"taskRequirements"`
	Location               *string   `json:"location,omitempty" bson:"location"`
	Budget                 *float64  `json:"budget,omitempty" bson:"budget"`
	SpecificSkillsRequired []*string `json:"specificSkillsRequired,omitempty" bson:"specificSkillsRequired"`
	Urgency                *string   `json:"urgency,omitempty" bson:"urgency"`
	Priority               *string   `json:"priority,omitempty" bson:"priority"`
	Status                 *string   `json:"status,omitempty" bson:"status"`
}

type CreateUserInput struct {
	Username string  `json:"username" bson:"username"`
	Email    string  `json:"email" bson:"email"`
	Phone    string  `json:"phone" bson:"phone"`
	Password string  `json:"password" bson:"password"`
	Sub      *string `json:"sub,omitempty" bson:"sub"`
}

type DeleteTaskResponse struct {
	DeletedTaskID string `json:"deletedTaskId" bson:"deletedTaskId"`
}

type DeleteUserResponse struct {
	DeletedUserID string `json:"deletedUserId" bson:"deletedUserId"`
}

type LogoutRequest struct {
	AccessToken string `json:"access_token" bson:"access_token"`
}

type LogoutResponse struct {
	Success bool `json:"success" bson:"success"`
}

type SignInRequest struct {
	Username string `json:"username" bson:"username"`
	Password string `json:"password" bson:"password"`
}

type SignInResponse struct {
	AccessToken  *string `json:"access_token,omitempty" bson:"access_token"`
	ExpiresIn    *string `json:"expires_in,omitempty" bson:"expires_in"`
	IDToken      *string `json:"id_token,omitempty" bson:"id_token"`
	RefreshToken *string `json:"refresh_token,omitempty" bson:"refresh_token"`
	TokenType    *string `json:"token_type,omitempty" bson:"token_type"`
}

type Task struct {
	ID                     string    `json:"_id" bson:"_id"`
	Owner                  *User     `json:"owner" bson:"owner"`
	OwnerID                *string   `json:"ownerId,omitempty" bson:"ownerId"`
	Description            string    `json:"description" bson:"description"`
	Category               *string   `json:"category,omitempty" bson:"category"`
	TaskRequirements       *string   `json:"taskRequirements,omitempty" bson:"taskRequirements"`
	Location               *string   `json:"location,omitempty" bson:"location"`
	Budget                 *float64  `json:"budget,omitempty" bson:"budget"`
	SpecificSkillsRequired []*string `json:"specificSkillsRequired,omitempty" bson:"specificSkillsRequired"`
	Urgency                *string   `json:"urgency,omitempty" bson:"urgency"`
	Priority               *string   `json:"priority,omitempty" bson:"priority"`
	Status                 *string   `json:"status,omitempty" bson:"status"`
}

type UpdateTaskInput struct {
	Description            *string   `json:"description,omitempty" bson:"description"`
	Category               *string   `json:"category,omitempty" bson:"category"`
	TaskRequirements       *string   `json:"taskRequirements,omitempty" bson:"taskRequirements"`
	Location               *string   `json:"location,omitempty" bson:"location"`
	Budget                 *float64  `json:"budget,omitempty" bson:"budget"`
	SpecificSkillsRequired []*string `json:"specificSkillsRequired,omitempty" bson:"specificSkillsRequired"`
	Urgency                *string   `json:"urgency,omitempty" bson:"urgency"`
	Priority               *string   `json:"priority,omitempty" bson:"priority"`
	Status                 *string   `json:"status,omitempty" bson:"status"`
}

type UpdateUserInput struct {
	Username *string `json:"username,omitempty" bson:"username"`
	Email    *string `json:"email,omitempty" bson:"email"`
	Password *string `json:"password,omitempty" bson:"password"`
}

type User struct {
	ID                 string    `json:"_id" bson:"_id"`
	Username           string    `json:"username" bson:"username"`
	Email              string    `json:"email" bson:"email"`
	Password           string    `json:"password" bson:"password"`
	Phone              *string   `json:"phone,omitempty" bson:"phone"`
	ContactInformation *string   `json:"contactInformation,omitempty" bson:"contactInformation"`
	ProfilePicture     *string   `json:"profilePicture,omitempty" bson:"profilePicture"`
	TaskPreferences    []*string `json:"taskPreferences,omitempty" bson:"taskPreferences"`
	VerificationStatus *bool     `json:"verificationStatus,omitempty" bson:"verificationStatus"`
	Tasks              []*Task   `json:"tasks,omitempty" bson:"tasks"`
}
