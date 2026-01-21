package handles

import (
	"net/http"

	"github.com/OpenListTeam/OpenList/v4/internal/ai"
	"github.com/OpenListTeam/OpenList/v4/internal/conf"
	"github.com/OpenListTeam/OpenList/v4/internal/op"
	"github.com/OpenListTeam/OpenList/v4/server/common"
	"github.com/gin-gonic/gin"
)

type AiRequest struct {
	Type    string `json:"type"` // translate, summarize, vision
	Content string `json:"content"`
	Image   string `json:"image"` // base64
	Target  string `json:"target"`
}

func AiProcess(c *gin.Context) {
	var req AiRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		common.ErrorResp(c, err, http.StatusBadRequest)
		return
	}

	featureSetting := ""
	switch req.Type {
	case "translate":
		featureSetting = conf.AiFeatureTranslation
	case "summary":
		featureSetting = conf.AiFeatureSummary
	case "vision":
		featureSetting = conf.AiFeatureOcrPlus
	}

	item, _ := op.GetSettingItemByKey(featureSetting)
	mode := ""
	if item != nil {
		mode = item.Value
	}
	var resp ai.Response

	if mode == "Cooperation" {
		resp = ai.ProcessWithCooperation(c.Request.Context(), req.Type, req.Content)
	} else {
		provider := ai.GetProvider(mode)
		if provider == nil {
			common.ErrorMsg(c, "AI provider not configured", http.StatusInternalServerError)
			return
		}

		ctx := c.Request.Context()
		switch req.Type {
		case "translate":
			resp = provider.Translate(ctx, req.Content, req.Target)
		case "summary":
			resp = provider.Summarize(ctx, req.Content)
		case "vision":
			// Handle vision separately if image is provided
			aiReq := ai.Request{Prompt: req.Content}
			if req.Image != "" {
				// decode or pass as is
				aiReq.Image = []byte(req.Image) // Simplified for now
			}
			resp = provider.Chat(ctx, aiReq)
		}
	}

	if resp.Error != nil {
		common.ErrorResp(c, resp.Error, http.StatusInternalServerError)
		return
	}

	common.SuccessResp(c, gin.H{
		"content": resp.Content,
	})
}
