package pgorderrepo

import (
	"context"

	"github.com/RakhimovAns/L0/internal/model"
	slerr "github.com/RakhimovAns/logger/pkg/err"
)

func (r *Repo) Create(ctx context.Context, order model.Order) error {
	ib := r.qb.NewInsertBuilder()
	ib.InsertInto("orders").
		Cols(
			"order_uid",
			"track_number",
			"entry",
			"locale",
			"internal_signature",
			"customer_id",
			"delivery_service",
			"shardkey",
			"sm_id",
			"date_created",
			"oof_shard",
		).
		Values(
			order.OrderUID,
			order.TrackNumber,
			order.Entry,
			order.Locale,
			order.InternalSignature,
			order.CustomerID,
			order.DeliveryService,
			order.ShardKey,
			order.SmID,
			order.DateCreated,
			order.OofShard,
		)

	sql, args := ib.Build()
	if _, err := r.db.Exec(ctx, sql, args...); err != nil {
		return slerr.WithSource(err)
	}

	ib = r.qb.NewInsertBuilder()
	ib.InsertInto("delivery").
		Cols(
			"order_uid",
			"name",
			"phone",
			"zip",
			"city",
			"address",
			"region",
			"email",
		).
		Values(
			order.OrderUID,
			order.Delivery.Name,
			order.Delivery.Phone,
			order.Delivery.Zip,
			order.Delivery.City,
			order.Delivery.Address,
			order.Delivery.Region,
			order.Delivery.Email,
		)

	sql, args = ib.Build()
	if _, err := r.db.Exec(ctx, sql, args...); err != nil {
		return slerr.WithSource(err)
	}

	ib = r.qb.NewInsertBuilder()
	ib.InsertInto("payments").
		Cols(
			"transaction",
			"request_id",
			"currency",
			"provider",
			"amount",
			"payment_dt",
			"bank",
			"delivery_cost",
			"goods_total",
			"custom_fee",
		).
		Values(
			order.Payment.Transaction,
			order.Payment.RequestID,
			order.Payment.Currency,
			order.Payment.Provider,
			order.Payment.Amount,
			order.Payment.PaymentDt,
			order.Payment.Bank,
			order.Payment.DeliveryCost,
			order.Payment.GoodsTotal,
			order.Payment.CustomFee,
		)

	sql, args = ib.Build()
	if _, err := r.db.Exec(ctx, sql, args...); err != nil {
		return slerr.WithSource(err)
	}

	for _, item := range order.Items {
		ib = r.qb.NewInsertBuilder()
		ib.InsertInto("items").
			Cols(
				"order_uid",
				"chrt_id",
				"track_number",
				"price",
				"rid",
				"name",
				"sale",
				"size",
				"total_price",
				"nm_id",
				"brand",
				"status",
			).
			Values(
				order.OrderUID,
				item.ChrtID,
				item.TrackNumber,
				item.Price,
				item.RID,
				item.Name,
				item.Sale,
				item.Size,
				item.TotalPrice,
				item.NmID,
				item.Brand,
				item.Status,
			)

		sql, args = ib.Build()
		if _, err := r.db.Exec(ctx, sql, args...); err != nil {
			return slerr.WithSource(err)
		}
	}

	return nil
}
