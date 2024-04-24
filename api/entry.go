package api

import (
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/AbdelilahOu/GoThingy/db/sqlc"
	"github.com/AbdelilahOu/GoThingy/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type createEntryRequest struct {
	AccountID uuid.UUID `json:"entry_id" binding:"required,uuid"`
	Amount    int64     `json:"amount" binding:"required"`
}

func (server *Server) createEntry(ctx *gin.Context) {
	var req createEntryRequest
	// validate the request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		server.logger.Error(fmt.Sprintf("invalid request: %s", err))
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	// create entry
	entry, err := server.store.CreateEntry(ctx, db.CreateEntryParams{
		AccountID: req.AccountID,
		Amount:    req.Amount,
	})
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				server.logger.Error(fmt.Sprintf("create entry db error foreign_key_violation or unique_violation: %s", err))
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		server.logger.Error(fmt.Sprintf("create entry error: %s", err))
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// return res
	ctx.JSON(http.StatusOK, entry)
}

type getEntryRequest struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid"`
}

func (server *Server) getEntry(ctx *gin.Context) {
	var req getEntryRequest
	// validate the request
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.logger.Error(fmt.Sprintf("invalid request: %s", err))
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	// get entry
	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			server.logger.Error(fmt.Sprintf("get entry db error no row found: %s", err))
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		server.logger.Error(fmt.Sprintf("get entry error: %s", err))
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// return res
	ctx.JSON(http.StatusOK, entry)
}

type listEntriesRequest struct {
	PageID   int32 `form:"page" binding:"required,min=1"`
	PageSize int32 `form:"size" binding:"required,min=5,max=10"`
}

func (server *Server) listEntries(ctx *gin.Context) {
	var req listEntriesRequest
	// validate the request
	if err := ctx.ShouldBindQuery(&req); err != nil {
		server.logger.Error(fmt.Sprintf("invalid request: %s", err))

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	// get entries
	entries, err := server.store.ListEntries(ctx, db.ListEntriesParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	})
	if err != nil {
		server.logger.Error(fmt.Sprintf("list entries error: %s", err))

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// return res
	ctx.JSON(http.StatusOK, entries)
}

type updateEntryRequest struct {
	ID     uuid.UUID `uri:"id" binding:"required,uuid"`
	Amount int64     `json:"balance" binding:"required"`
}

func (server *Server) updateEntry(ctx *gin.Context) {
	var req updateEntryRequest
	// validate the request
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.logger.Error(fmt.Sprintf("invalid request: %s", err))

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	// get entry
	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			server.logger.Error(fmt.Sprintf("get entry for update db error no row found: %s", err))

			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// update entry
	entry, err = server.store.UpdateEntry(ctx, db.UpdateEntryParams{
		ID:     entry.ID,
		Amount: req.Amount,
	})
	if err != nil {
		server.logger.Error(fmt.Sprintf("update entry db error: %s", err))

		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// return res
	ctx.JSON(http.StatusOK, entry)
}

type deleteEntryRequest struct {
	ID uuid.UUID `uri:"id" binding:"required,uuid"`
}

func (server *Server) deleteEntry(ctx *gin.Context) {
	var req deleteEntryRequest
	// validate the request
	if err := ctx.ShouldBindUri(&req); err != nil {
		server.logger.Error(fmt.Sprintf("invalid request: %s", err))

		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	// get entry
	entry, err := server.store.GetEntry(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			server.logger.Error(fmt.Sprintf("get entry for delete db error no row found: %s", err))
			ctx.JSON(http.StatusNotFound, utils.ErrorResponse(err))
			return
		}
		server.logger.Error(fmt.Sprintf("get entry for delete error no row found: %s", err))
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// delete entry
	err = server.store.DeleteEntry(ctx, req.ID)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				server.logger.Error(fmt.Sprintf("delete entry db error foreign_key_violation or unique_violation: %s", err))
				ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
				return
			}
		}
		server.logger.Error(fmt.Sprintf("delete entry error: %s", err))
		ctx.JSON(http.StatusInternalServerError, utils.ErrorResponse(err))
		return
	}
	// return res
	ctx.JSON(http.StatusOK, entry)
}
