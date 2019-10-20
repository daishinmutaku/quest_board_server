package response

import "github.com/gin-gonic/gin"

type Response struct {
	Key   string
	Value interface{}
}

func (response *Response) FormatToJson() interface{} {
	return gin.H{
		response.Key: response.Value,
	}
}
