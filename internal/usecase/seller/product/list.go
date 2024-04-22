package product

import (
	"context"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"log"
	"net/http"
)

func (p *sellerUseCase) ListProduct(ctx context.Context, payload presentations.Paging, authId int) (presentations.ResponseListingProduct, error) {

	payload.Limit = helper.LimitDefaultValue(payload.Limit)
	payload.Page = helper.PageDefaultValue(payload.Page)
	offset := helper.PageToOffset(payload.Limit, payload.Page)
	products, err := p.repoProduct.Listing(ctx, payload.Limit, offset, authId)
	if err != nil {
		log.Printf("failed get all products %v", err)
		return presentations.ResponseListingProduct{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	count, err := p.repoProduct.Count(ctx, authId)
	if err != nil {
		log.Printf("failed count transactions %v", err)
		return presentations.ResponseListingProduct{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	var dataProducts []presentations.DataProducts

	for _, ps := range products {
		data := presentations.DataProducts{
			Id: ps.Id,
			SellerInformation: presentations.SellerInformation{
				Id:       ps.UserId,
				FullName: ps.FullName,
				Address:  ps.Address,
			},
			ProductName: ps.ProductName,
			Price:       ps.Price,
			Description: ps.Description,
			CreatedAt:   ps.CreatedAt.Format(consts.LayoutDateTimeFormat),
			UpdatedAt:   ps.UpdatedAt.Format(consts.LayoutDateTimeFormat),
		}

		dataProducts = append(dataProducts, data)
	}

	return presentations.ResponseListingProduct{
		Code:    http.StatusOK,
		Message: "success",
		Data:    dataProducts,
		MetaData: presentations.MetaData{
			Pagination: presentations.Pagination{
				Page:       payload.Page,
				Limit:      payload.Limit,
				TotalPage:  helper.PageCalculate(count, payload.Limit),
				TotalCount: count,
			}},
	}, nil
}
