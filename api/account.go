package api

import (
	"errors"
	"net/http"

	db "com.wlq/simplebank/db/sqlc"
	"com.wlq/simplebank/token"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,currency"`
}

// 创建account
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	// 解析参数为createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 检查认证载荷中的用户名 登录的用户只能创建自己的账户
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	if authPayload.Username != req.Owner {
		err := errors.New("account's owner doesn't belong to the authorization user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		Owner:    authPayload.Username,
		Currency: req.Currency,
		Balance:  0,
	}

	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {

		errCode := db.ErrorCode(err)
		if errCode == db.ForeignKeyViolation || errCode == db.UniqueViolation {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

// 查询account
func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest

	// 解析参数为getAccountRequest,并从url中解析参数
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if errors.Is(err, db.ErrRecordNotFound) {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// 检查认证载荷中的用户名 登录的用户只能查询自己的账户
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)
	if account.Owner != authPayload.Username {
		err := errors.New("account doesn't belong to the authorization user")
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

// 分页查询
func (server *Server) listAccount(ctx *gin.Context) {
	var req listAccountRequest
	// 解析Params参数为listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// 检查认证载荷中的用户名 登录的用户只能查询自己的账户
	authPayload := ctx.MustGet(authorizationPayloadKey).(*token.Payload)

	arg := db.ListAccountsParams{
		Owner:  authPayload.Username,
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
