package controller

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	coreResponse "github.com/datsukan/datsukan-blog-comment-core/response"
	"github.com/datsukan/datsukan-blog-comment-core/usecase"
	"github.com/datsukan/datsukan-blog-comment-ref/request"
	"github.com/datsukan/datsukan-blog-comment-ref/response"
)

func Ref(r events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	req := request.GetRequest(r.QueryStringParameters)
	if err := req.Validate(); err != nil {
		return coreResponse.ResponseBadRequestError(err)
	}

	cs, err := usecase.Ref(req.ArticleID)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}

	res := []response.ResponseItem{}
	for _, c := range cs {
		reply := []response.ReplyComment{}
		ri := response.ResponseItem{
			ID:            c.ID,
			ArticleID:     c.ArticleID,
			ParentID:      c.ParentID,
			UserName:      c.UserName,
			Content:       c.Content,
			CreatedAt:     c.CreatedAt,
			ReplyComments: reply,
		}
		if len(c.ReplyComments) > 0 {
			var rcs []response.ReplyComment
			for _, brc := range c.ReplyComments {
				rc := response.ReplyComment{
					ID:        brc.ID,
					ArticleID: brc.ArticleID,
					ParentID:  brc.ParentID,
					UserName:  brc.UserName,
					Content:   brc.Content,
					CreatedAt: brc.CreatedAt,
				}
				rcs = append(rcs, rc)
			}
			ri.ReplyComments = rcs
		}
		res = append(res, ri)
	}

	j, err := json.Marshal(res)
	if err != nil {
		return coreResponse.ResponseInternalServerError(err)
	}
	js := string(j)

	return coreResponse.ResponseSuccess(js)
}
