package order_product

import (
	"context"
	"gokomodo_test/internal/consts"
	"gokomodo_test/internal/helper"
	"gokomodo_test/internal/presentations"
	"log"
	"net/http"
)

func (o *orderBuyerUseCase) ListOrder(ctx context.Context, payload presentations.Paging, authId int) (presentations.ResponseListingOrderBuyer, error) {
	payload.Limit = helper.LimitDefaultValue(payload.Limit)
	payload.Page = helper.PageDefaultValue(payload.Page)
	offset := helper.PageToOffset(payload.Limit, payload.Page)

	viewOrders, err := o.repoOrder.Listing(ctx, payload.Limit, offset, authId)
	if err != nil {
		log.Printf("failed get all orders %v", err)
		return presentations.ResponseListingOrderBuyer{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	count, err := o.repoOrder.Count(ctx, authId)
	if err != nil {
		return presentations.ResponseListingOrderBuyer{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error",
		}, err
	}

	var dataProducts []presentations.DataOrderBuyer

	for _, ps := range viewOrders {
		data := presentations.DataOrderBuyer{
			Id: ps.Id,
			BuyerInformation: presentations.BuyerInformation{
				Id:       authId,
				FullName: ps.Buyer,
				Address:  ps.DeliveryDestinationAddress,
				Email:    ps.Email,
			},
			DataProductBuyer: presentations.DataProductBuyer{
				InformationSeller: presentations.InformationSeller{
					FullName: ps.Seller,
					Address:  ps.DeliverySourceAddress,
				},
				ProductName: ps.ProductName,
				TotalPrice:  ps.TotalPrice,
				Description: ps.Description,
				Quantity:    ps.Quantity,
			},
			Status:    ps.Status,
			CreatedAt: ps.CreatedAt.Format(consts.LayoutDateTimeFormat),
		}

		dataProducts = append(dataProducts, data)
	}

	return presentations.ResponseListingOrderBuyer{
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
