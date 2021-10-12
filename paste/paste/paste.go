package paste

import (
	"context"
	"gorm.io/gorm"
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

func Init(db *gorm.DB) {
	paste = repositories.NewPasteRepo(db)
	errHandler = errs.NewErrs(db)
}

func init() {
	Init(errs.GetDB())
}

type Service struct {
	pastepb.UnimplementedPasteServiceServer
}

func (p *Service) Paste(ctx context.Context, pr *pastepb.PasteRequest) (*pastepb.PasteResponse, error) {
	errs.Refresh()
	Init(errs.GetDB())
	err := paste.CreateP(P2MPaste(pr))
	if err != nil {
		errHandler.ReciteErrors(errs.Err{Message: "error creating paste",
			When: time.Now(), Where: "paste service"})
		return &pastepb.PasteResponse{}, err
	}
	return M2PPaste(&models.Paste{Things: pr.Things}), nil
}

func (p *Service) GetAllPaste(ctx context.Context, pr *pastepb.PasteRequest) (*pastepb.PasteResponse, error) {
	errs.Refresh()
	Init(errs.GetDB())
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
	errs.Refresh()
	Init(errs.GetDB())
	err := paste.DeletePS()
	if err != nil {
		errHandler.ReciteErrors(errs.Err{Message: "error deleting paste",
			When: time.Now(), Where: "paste service"})
		return &pastepb.PasteResponse{}, err
	}
	return &pastepb.PasteResponse{StatusCode: http.StatusOK}, nil
}
