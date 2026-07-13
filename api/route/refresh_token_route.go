package route

import (
	"time"

	"github.com/alfinmuzakkiiman/go-backend-clean-architecture/api/controller"
	"github.com/alfinmuzakkiiman/go-backend-clean-architecture/bootstrap"
	"github.com/alfinmuzakkiiman/go-backend-clean-architecture/domain"
	"github.com/alfinmuzakkiiman/go-backend-clean-architecture/mongo"
	"github.com/alfinmuzakkiiman/go-backend-clean-architecture/repository"
	"github.com/alfinmuzakkiiman/go-backend-clean-architecture/usecase"
	"github.com/gin-gonic/gin"
)

func NewRefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.NewRefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
