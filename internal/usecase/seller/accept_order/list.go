package accept_order

import (
	"context"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"log"
	"net/http"
)

func (p *acceptOrderUseCase) ListingOrderDetails(ctx context.Context, payload presentations.Paging, authId int) (presentations.ResponseListingOrderDetails, error) {
	payload.Limit = helper.LimitDefaultValue(payload.Limit)
	payload.Page = helper.PageDefaultValue(payload.Page)
	offset := helper.PageToOffset(payload.Limit, payload.Page)

	orderDetails, err := p.repoProduct.ListingOrderSeller(ctx, payload.Limit, offset, authId)
	if err != nil {
		log.Printf("failed get all order details %v", err)
		return presentations.ResponseListingOrderDetails{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	count, err := p.repoProduct.CountOrderSeller(ctx, authId)
	if err != nil {
		log.Printf("failed count transactions order seller %v", err)
		return presentations.ResponseListingOrderDetails{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	var dataOrderDetails []presentations.DataOrderDetails

	for _, od := range orderDetails {
		data := presentations.DataOrderDetails{
			Id: od.Id,
			SellerInformation: presentations.SellerInformation{
				Id:       od.SellerId,
				FullName: od.Seller,
				Address:  od.SellerAddress,
			},
			BuyerInformation: presentations.BuyerInformation{
				Id:       od.BuyerId,
				FullName: od.Buyer,
				Address:  od.BuyerAddress,
			},
			DataProductSeller: presentations.DataProductSeller{
				ProductName: od.ProductName,
				Price:       od.Price,
				Description: od.Description,
				Quantity:    od.Quantity,
			},
			Status:    od.Status,
			CreatedAt: od.CreatedAt.Format(consts.LayoutDateTimeFormat),
			UpdatedAt: od.UpdatedAt.Format(consts.LayoutDateTimeFormat),
		}

		dataOrderDetails = append(dataOrderDetails, data)
	}

	return presentations.ResponseListingOrderDetails{
		Code:    http.StatusOK,
		Message: "success",
		Data:    dataOrderDetails,
		MetaData: presentations.MetaData{
			Pagination: presentations.Pagination{
				Page:       payload.Page,
				Limit:      payload.Limit,
				TotalPage:  helper.PageCalculate(count, payload.Limit),
				TotalCount: count,
			}},
	}, nil
}
