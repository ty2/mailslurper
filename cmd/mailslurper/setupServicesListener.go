// Copyright 2013-2018 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mailslurper/mailslurper/cmd/mailslurper/controllers"
	"github.com/mailslurper/mailslurper/pkg/auth/authfactory"
	"github.com/mailslurper/mailslurper/pkg/auth/authscheme"
	"github.com/mailslurper/mailslurper/pkg/auth/jwt"
	"github.com/mailslurper/mailslurper/pkg/mailslurper"
	"net/http"
)

func setupServicesListener() {
	middlewares := make([]echo.MiddlewareFunc, 0, 5)

	/*
	 * Start the services server
	 */
	serviceController := &controllers.ServiceController{
		AuthFactory: &authfactory.AuthFactory{
			Config: config,
		},
		CacheService: cacheService,
		Config:       config,
		Database:     database,
		JWTService: &jwt.JWTService{
			Config: config,
		},
		Logger:        mailslurper.GetLogger(*logLevel, *logFormat, "ServiceController"),
		ServerVersion: SERVER_VERSION,
	}

	service = echo.New()
	service.HideBanner = true

	if config.AuthenticationScheme != authscheme.NONE {
		middlewares = append(middlewares, serviceAuthorization)
	}

	service.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: []string{"*", config.ServicePublicURL},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	service.HEAD("/", serviceController.Head, middlewares...)
	service.HEAD("/mail", serviceController.Head, middlewares...)
	service.HEAD("/mail/:id", serviceController.Head, middlewares...)
	service.HEAD("/mail/:id/message", serviceController.Head, middlewares...)
	service.HEAD("/mail/:id/messageraw", serviceController.Head, middlewares...)
	service.GET("/mail/:id", serviceController.GetMail, middlewares...)
	service.GET("/mail/:id/message", serviceController.GetMailMessage, middlewares...)
	service.GET("/mail/:id/messageraw", serviceController.GetMailMessageRaw, middlewares...)
	service.DELETE("/mail", serviceController.DeleteMail, middlewares...)
	service.GET("/mail", serviceController.GetMailCollection, middlewares...)
	service.GET("/mailcount", serviceController.GetMailCount, middlewares...)
	service.GET("/mail/:mailID/attachment/:attachmentID", serviceController.DownloadAttachment)
	service.GET("/version", serviceController.Version, middlewares...)
	service.GET("/pruneoptions", serviceController.GetPruneOptions, middlewares...)

	if config.AuthenticationScheme != authscheme.NONE {
		service.POST("/login", serviceController.Login)
		service.DELETE("/logout", serviceController.Logout)
	}

	go func() {
		var err error

		if config.IsServiceSSL() {
			err = service.StartTLS(config.GetFullServiceAppAddress(), config.CertFile, config.KeyFile)
		} else {
			err = service.Start(config.GetFullServiceAppAddress())
		}

		if err != nil {
			logger.WithError(err).Info("Shutting down HTTP service listener")
		} else {
			logger.Infof("Service listener running on %s", config.GetFullServiceAppAddress())
		}
	}()
}
