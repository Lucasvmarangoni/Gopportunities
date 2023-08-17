package handler

import (
	"net/http"

	"github.com/Lucasvmarangoni/go-api/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1/

// @Summary Create a new opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param Request body CreateOpeningRequest true "Request Body"
// @Success 200 {object} DefaultOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func CreateOpeningHandler(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	opening := schemas.Opening{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   request.Salary,
	}

	if err := db.Create(&opening).Error; err != nil {
		logger.Errorf("error creating opening: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error creating opening on database")
		return
	}

	sendSuccess(ctx, "create-opening", opening)
}
