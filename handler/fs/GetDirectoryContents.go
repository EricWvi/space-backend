package fs

import (
	"github.com/gin-gonic/gin"
)

func (b Base) GetDirectoryContents(c *gin.Context, req *GetDirectoryContentsRequest) *GetDirectoryContentsResponse {
	return &GetDirectoryContentsResponse{}
}

type GetDirectoryContentsRequest struct {
}

type GetDirectoryContentsResponse struct {
}
