//go:build wireinject
// +build wireinject

// Copyright 2018 The Wire Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// The build tag makes sure the stub is not built in the final build.
package main

import (
	"app/conf"
	"app/domains/account"
	"app/domains/audio"
	"app/sql/db"
	"github.com/go-chi/jwtauth"
	"github.com/google/wire"
)

type reqControllers struct {
	AuthC  account.AuthController
	AccC   account.AccountController
	AudioC audio.AudioController
}

func controllers(config *conf.Config, jwtAuth *jwtauth.JWTAuth, queries *db.Queries) (
	reqControllers, error,
) {
	wire.Build(
		account.NewAuthController,
		account.NewAuthService,
		account.NewAccountController,
		account.NewAccountService,
		account.NewAccountRepo,
		audio.NewAudioController,
		wire.Bind(new(account.IAuthService), new(*account.AuthService)),
		wire.Bind(new(account.IAccountService), new(*account.AccountService)),
		wire.Bind(new(account.IAccountRepo), new(*account.AccountRepo)),
		newReqControllers,
	)
	return reqControllers{}, nil
}

func newReqControllers(
	authC account.AuthController,
	accC account.AccountController,
	audioC audio.AudioController,
) reqControllers {
	return reqControllers{
		AuthC:  authC,
		AccC:   accC,
		AudioC: audioC,
	}
}
