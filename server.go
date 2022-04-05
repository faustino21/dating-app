package main

import (
	"dating_app_last/config"
	"dating_app_last/delivery/member"
	"dating_app_last/manager"
	"dating_app_last/util"
	"github.com/gin-gonic/gin"
)

type AppServer interface {
	Run()
}

type appServer struct {
	routerEngine   *gin.Engine
	cfg            config.Config
	InfraManager   manager.InfraManager
	RepoManager    manager.RepoManager
	UseCaseManager manager.UseCaseManager
}

func (a *appServer) initHandler() {
	a.v1()
}

func (a *appServer) v1() {
	datingGroup := a.routerEngine.Group("/dating")
	member.NewMemberApi(datingGroup, a.cfg.UseCaseManager.MemberSignUpUseCase(), a.cfg.UseCaseManager.MemberActivationUseCase())
}

func (a *appServer) Run() {
	a.initHandler()
	util.Log.Info().Msgf("Server run on %s", a.cfg.ApiConfig.Url)
	err := a.routerEngine.Run("localhost:8888")
	if err != nil {
		util.Log.Fatal().Msg("Server failed tp run")
	}
}

func Server() AppServer {
	r := gin.Default()
	c := config.NewConfig(".", "config")

	return &appServer{
		routerEngine: r,
		cfg:          c,
	}
}
