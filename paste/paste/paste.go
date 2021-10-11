package paste

import (
	"context"
	"net/http"
	"pasteProject/errs"
	"pasteProject/models"
	pastepb "pasteProject/paste/api/gen/go"
	"pasteProject/repositories"
	"time"
)

var (
	paste      repositories.IPasteRepo
	errHandler errs.IErr
)

func init() {
	paste = repositories.NewPasteRepo(errs.GetDB())
	errHandler = errs.NewErrs(errs.GetDB())
}

type Service struct {
	pastepb.UnimplementedPasteServiceServer
}

func (p *Service) Paste(ctx context.Context, pr *pastepb.PasteRequest) (*pastepb.PasteResponse, error) {
	err := paste.CreateP(P2MPaste(pr))
	if err != nil {
		errHandler.ReciteErrors(errs.Err{Message: "error creating paste",
			When: time.Now(), Where: "paste service"})
		return &pastepb.PasteResponse{}, err
	}
	return M2PPaste(&models.Paste{Things: pr.Things}), nil
}

func (p *Service) GetAllPaste(ctx context.Context, pr *pastepb.PasteRequest) (*pastepb.PasteResponse, error) {
	var pb pastepb.PasteResponse
	ps := paste.SelectPS()
	if len(ps) == 0 {
		return &pb, nil
	}
	for _, pn := range ps {
		pb.Things = append(pb.Things, pn.Things)
	}
	return &pb, nil
}

func (p *Service) DeleteAllPaste(ctx context.Context, pr *pastepb.PasteRequest) (*pastepb.PasteResponse, error) {
	err := paste.DeletePS()
	if err != nil {
		errHandler.ReciteErrors(errs.Err{Message: "error deleting paste",
			When: time.Now(), Where: "paste service"})
		return &pastepb.PasteResponse{}, err
	}
	return &pastepb.PasteResponse{StatusCode: http.StatusOK}, nil
}
