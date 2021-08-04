package current_user

import (
	"encoding/json"
	"net/http"
	"time"

	"deltaboard-server/api/v1/response"
	"deltaboard-server/config/db"
	"deltaboard-server/config/session"
	"deltaboard-server/models"
)

const (
	UserSessionKey       = "user-token"
	UserVerificationInfo = "user-verification-info"
)

type CurrentUser struct {
	*models.User
	UserVerification  UserVerification  `json:"user_verification"`
	VaultVerification VaultVerification `json:"vault_verification"`
}

type UserVerification struct {
	UserId     string `json:"user_id"`
	VaultToken string `json:"vault_token"`
}

type VaultVerification struct {
	RandomString string    `json:"random_string"`
	Token        string    `json:"token"`
	Expire       time.Time `json:"expire"`
}

func NewCurrentUser() *CurrentUser {
	return &CurrentUser{
		User: &models.User{},
	}
}

func NewCurrentUserFromSession(request *http.Request) (*CurrentUser, error) {
	// 用户登录状态session获取
	userSession, err := session.GetSessionStore().Get(request, UserSessionKey)

	if err != nil {
		return nil, err
	}

	currentUser := NewCurrentUser()

	userVerificationBytes := userSession.Values[UserVerificationInfo]
	if userVerificationBytes == nil {
		return nil, response.NewAccessDeniedResponse()
	}

	userVerificationInfo := UserVerification{}
	if err := json.Unmarshal(userVerificationBytes.([]byte), &userVerificationInfo); err != nil {
		return nil, err
	}
	currentUser.UserVerification = userVerificationInfo

	currentUser.User.Id = currentUser.UserVerification.UserId
	if err := db.GetDB().First(currentUser.User, currentUser.User).Error; err != nil {
		return nil, err
	}

	// 保险箱登录状态session获取
	return currentUser, nil
}

func (currentUser *CurrentUser) RemoveUserLoginSession(request *http.Request, response http.ResponseWriter) error {
	store := session.GetSessionStore()
	userSession := session.NewSession(store, UserSessionKey) // 创建新session
	userSession.Options.MaxAge = 0
	return store.Save(request, response, userSession)
}

func (currentUser *CurrentUser) SaveUserLoginSession(request *http.Request, response http.ResponseWriter) error {

	store := session.GetSessionStore()
	userSession := session.NewSession(store, UserSessionKey) // 创建新session

	userVerificationBytes, err := json.Marshal(currentUser.UserVerification)
	if err != nil {
		return err
	}

	userSession.Values[UserVerificationInfo] = userVerificationBytes
	return store.Save(request, response, userSession)
}
