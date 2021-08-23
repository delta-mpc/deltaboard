/*
 * Copyright 2021 Seven Seals Technology
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package models

import (
	"crypto/sha256"
	"fmt"

	"deltaboard-server/internal/library/random"
)

type UserStatus int

const (
	USER_APPROV_STATUS_REGISTED = 1
	USER_APPROV_STATUS_APPROVED = 2
	USER_APPROV_STATUS_REJECTED = 3
)

const (
	ROLE_ADMIN = 1
)

type User struct {
	BaseModel
	Name        string `json:"name" gorm:"type:varchar(127);not null"`
	Password    string `json:"-" gorm:"type:varchar(64);not null"`
	Salt        string `json:"-" gorm:"type:varchar(32);not null"`
	Email       string `json:"email" gorm:"type:varchar(127);not null"`
	RealName    string `json:"real_name" gorm:"type:varchar(32);not null"`
	CardNo      string `json:"card_no" gorm:"type:varchar(32);not null"`
	Phone       string `json:"phonenumber" gorm:"type:varchar(11);not null"`
	Role        int64  `json:"role" gorm:"type:tinyint(8)"`
	DeltaToken  string `json:"delta_token" gorm:"type:varchar(191);" description:"auth token"`
	ApprvStatus int8   `json:"approve_status" gorm:"type:tinyint(8)"`
}

func NewUser(name, password, email, phonenumber string) *User {

	salt := random.GenerateRandomString(32)
	passwordHash := fmt.Sprintf("%x", sha256.Sum256([]byte(password)))
	password = fmt.Sprintf("%x", sha256.Sum256([]byte(passwordHash+salt)))

	return &User{
		Name:        name,
		Password:    password,
		Salt:        salt,
		Email:       email,
		Phone:       phonenumber,
		Role:        0,
		DeltaToken:  random.GenerateRandomString(32),
		ApprvStatus: USER_APPROV_STATUS_REGISTED,
	}
}
