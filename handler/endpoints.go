package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetUser(ctx *fiber.Ctx) error {
	usr, err := s.UseCase.DisplayUser(ctx.Context(), ctx.FormValue("userid"))
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return ctx.JSON(usr)
}

func (s *Server) GetAllUser(ctx *fiber.Ctx) error {
	listUsr, err := s.UseCase.DisplayAllUsers(ctx.Context())
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).Send([]byte(err.Error()))
	}
	return ctx.JSON(listUsr)
}

func (s *Server) Test(ctx *fiber.Ctx) error {
	return ctx.JSON(http.StatusCreated, "success")
}
