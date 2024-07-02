package service

import (
	"context"

	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/model"
)

type FactorizacionQrService interface {
	GetQrFromMatrix(ctx context.Context, model model.MatrixModel) model.MatrixResponse
}
