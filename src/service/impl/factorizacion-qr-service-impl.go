package impl

import (
	"context"

	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/model"
	"github.com/Jorge2824/api-factorizacion-qr-go-v1/src/service"
	"gonum.org/v1/gonum/mat"
)

func NewFactorizacionQrServiceImpl() service.FactorizacionQrService {
	return &factorizacionQrServiceImpl{}
}

type factorizacionQrServiceImpl struct {
}

// Create implements service.FactorizacionQrService.
func (p *factorizacionQrServiceImpl) GetQrFromMatrix(ctx context.Context, modeled model.MatrixModel) model.MatrixResponse {
	matrix := mat.NewDense(len(modeled.Matrix), len(modeled.Matrix[0]), nil)
	for i := range modeled.Matrix {
		for j := range modeled.Matrix[i] {
			matrix.Set(i, j, modeled.Matrix[i][j])
		}
	}
	var qr mat.QR
	qr.Factorize(matrix)

	q := mat.NewDense(0, 0, nil)
	r := mat.NewDense(0, 0, nil)
	qr.QTo(q)
	qr.RTo(r)

	qMatrix := make([][]float64, q.RawMatrix().Rows)
	for i := range qMatrix {
		qMatrix[i] = q.RawRowView(i)
	}

	rMatrix := make([][]float64, r.RawMatrix().Cols)
	for i := range rMatrix {
		rMatrix[i] = r.RawRowView(i)
	}

	return model.MatrixResponse{
		Q: qMatrix,
		R: rMatrix,
	}
}
