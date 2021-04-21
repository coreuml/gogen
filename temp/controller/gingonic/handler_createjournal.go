package gingonic

import (
	"github.com/mirzaakhena/gogen2/temp/apperror"
	"github.com/mirzaakhena/gogen2/temp/infrastructure/log"
	"github.com/mirzaakhena/gogen2/temp/infrastructure/util"
	"github.com/mirzaakhena/gogen2/temp/usecase/createjournal"
	"net/http"

	"github.com/gin-gonic/gin"
)

// createJournalHandler ...
func (r *Controller) createJournalHandler(inputPort createjournal.Inport) gin.HandlerFunc {

	return func(c *gin.Context) {

		//
		//
		//
		//
		//
		//
		//

		ctx := log.ContextWithLogGroupID(c.Request.Context())

		var req createjournal.InportRequest
		if err := c.BindJSON(&req); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.ErrorResponse(ctx, err)
			c.JSON(http.StatusBadRequest, util.MustJSON(NewErrorResponse(newErr)))
			return
		}

		log.InfoRequest(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.ErrorResponse(ctx, err)
			c.JSON(http.StatusBadRequest, NewErrorResponse(err))
			return
		}

		log.InfoResponse(ctx, util.MustJSON(res))
		c.JSON(http.StatusOK, NewSuccessResponse(res))

	}
}
