package paste

import (
	"net/http"
	"pasteProject/models"
	pastepb "pasteProject/paste/api/gen/go"
)

func P2MPaste(p *pastepb.PasteRequest) *models.Paste {
	return &models.Paste{
		Things: p.Things,
		Poster: p.Poster,
	}
}

func M2PPaste(m *models.Paste) *pastepb.PasteResponse {
	return &pastepb.PasteResponse{
		Things:     []string{m.Things},
		StatusCode: http.StatusOK,
	}
}
