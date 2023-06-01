package v1

import (
	"github.com/gin-gonic/gin"

	usecase "gitlab.com/demojira/template.git/internal/application/user"
	"gitlab.com/demojira/template.git/pkg/logger"
)

type userRoutes struct {
	u usecase.UserUseCase
	l logger.Interface
}

func newUserRoutes(handler *gin.RouterGroup, u usecase.UserUseCase, l logger.Interface) {
	r := &userRoutes{u, l}

	h := handler.Group("/user")
	{
		h.GET("/getInfo", r.getInfo)
		h.POST("/register", r.register)
	}
}

// @Summary     Show user info
// @Description Show user info
// @ID          GetUser
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Success     200 {object} historyResponse
// @Failure     500 {object} response
// @Router      /user/getInfo [get]
func (r *userRoutes) getInfo(c *gin.Context) {
	// translations, err := r.u.GetByID(c.Request.Context())
	// if err != nil {
	// 	r.l.Error(err, "http - v1 - history")
	// 	errorResponse(c, http.StatusInternalServerError, "database problems")

	// 	return
	// }

	// c.JSON(http.StatusOK, historyResponse{translations})
}

type doTranslateRequest struct {
	Source      string `json:"source"       binding:"required"  example:"auto"`
	Destination string `json:"destination"  binding:"required"  example:"en"`
	Original    string `json:"original"     binding:"required"  example:"текст для перевода"`
}

// @Summary     Register User
// @Description Register User
// @ID          RegisterUser
// @Tags  	    user
// @Accept      json
// @Produce     json
// @Param       request body doTranslateRequest true "Set up translation"
// @Success     200 {object} entity.Translation
// @Failure     400 {object} response
// @Failure     500 {object} response
// @Router      /user/register [post]
func (r *userRoutes) register(c *gin.Context) {
	// var request doTranslateRequest
	// if err := c.ShouldBindJSON(&request); err != nil {
	// 	r.l.Error(err, "http - v1 - doTranslate")
	// 	errorResponse(c, http.StatusBadRequest, "invalid request body")

	// 	return
	// }

	// translation, err := r.t.Translate(
	// 	c.Request.Context(),
	// 	entity.Translation{
	// 		Source:      request.Source,
	// 		Destination: request.Destination,
	// 		Original:    request.Original,
	// 	},
	// )
	// if err != nil {
	// 	r.l.Error(err, "http - v1 - doTranslate")
	// 	errorResponse(c, http.StatusInternalServerError, "translation service problems")

	// 	return
	// }

	// c.JSON(http.StatusOK, translation)
}
