package handler

import (
	"database/sql"
	"net/http"

	"github.com/basriyasin/sp-user/generated"
	"github.com/basriyasin/sp-user/repository"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

// [POST] /register
// create a new user profile
func (s Server) Register(ctx echo.Context) error {
	var req generated.RegisterRequest
	err := ctx.Bind(&req)
	if err != nil {
		return echo.ErrBadRequest
	}

	user := repository.User{
		Name:     req.Name,
		Phone:    req.Phone,
		Password: req.Password,
	}
	err = Validate(user)
	if err != nil {
		return err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	user.Password = string(bytes)
	userID, err := s.Repository.SaveProfile(ctx.Request().Context(), user)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, generated.Profile{
		Id:    &userID,
		Name:  req.Name,
		Phone: req.Phone,
	})
}

// [POST] /authenticate
//
//	authenticate the provided user and return a JWT token
func (s Server) Authenticate(ctx echo.Context) error {
	var req generated.AuthenticateRequest
	err := ctx.Bind(&req)
	if err != nil {
		return err
	}

	user, err := s.Repository.GetProfileByPhone(ctx.Request().Context(), req.Phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return echo.ErrBadRequest
		}
		return echo.ErrInternalServerError
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return echo.ErrBadRequest
	}

	token, err := generateToken(s.rsaPrivateKey, user)
	if err != nil {
		return err
	}

	err = s.Repository.UpdateLoginCount(ctx.Request().Context(), user.ID, user.LoginCount+1)
	if err != nil {
		return err
	}

	createdAt := user.CreatedAt.Format(DateTimeFormat)
	var updatedAt string
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time.Format(DateTimeFormat)
	}
	return ctx.JSON(http.StatusOK, generated.AuthenticateResponse{
		Id:        &user.ID,
		Name:      user.Name,
		Phone:     user.Phone,
		Token:     token,
		UpdateAt:  &updatedAt,
		CreatedAt: &createdAt,
	})
}

// [GET] /profile
// retrieve the latest profile information of the currently logged-in user
func (s Server) Profile(ctx echo.Context) error {
	user, err := verifyToken(ctx, s.rsaPrivateKey)
	if err != nil {
		return err
	}

	// get latest updated profile
	user, err = s.Repository.GetProfileByID(ctx.Request().Context(), user.ID)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, generated.Profile{
		Name:  user.Name,
		Phone: user.Phone,
	})
}

// [PUT] /profile
// update the current logged-in user's name or phone number if the provided information is valid.
func (s Server) UpdateProfile(ctx echo.Context) (err error) {
	user, err := verifyToken(ctx, s.rsaPrivateKey)
	if err != nil {
		return echo.ErrForbidden
	}

	var req generated.Profile
	err = ctx.Bind(&req)
	if err != nil {
		return err
	}

	user, err = s.Repository.GetProfileByID(ctx.Request().Context(), user.ID)
	if err != nil {
		return err
	}

	hasValidName := isValidName(req.Name)
	if hasValidName {
		user.Name = req.Name
	}

	hasValidPhone := isValidPhone(req.Phone)
	if hasValidPhone {
		user.Phone = req.Phone
	}

	if !hasValidName && !hasValidPhone {
		return ctx.JSON(http.StatusOK, generated.Profile{
			Name:  user.Name,
			Phone: user.Phone,
		})
	}

	err = s.Repository.UpdateUserByID(ctx.Request().Context(), user)
	if err != nil {
		return echo.ErrConflict
	}

	return ctx.JSON(http.StatusOK, generated.Profile{
		Name:  user.Name,
		Phone: user.Phone,
	})
}
