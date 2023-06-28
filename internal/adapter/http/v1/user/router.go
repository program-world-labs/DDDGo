package user

import (
	"github.com/gin-gonic/gin"
	"github.com/program-world-labs/pwlogger"

	application_user "github.com/program-world-labs/DDDGo/internal/application/user"
)

type userRoutes struct {
	u application_user.IService
	l pwlogger.Interface
}

func NewUserRoutes(handler *gin.RouterGroup, u application_user.IService, l pwlogger.Interface) {
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
// @Router      /user/getInfo [get].
func (r *userRoutes) getInfo(_ *gin.Context) {

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
// @Router      /user/register [post].
func (r *userRoutes) register(_ *gin.Context) {
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
