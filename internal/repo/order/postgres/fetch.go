package pgorderrepo

import (
	"context"

	"github.com/RakhimovAns/L0/internal/model"
	slerr "github.com/RakhimovAns/logger/pkg/err"
)

func (r *Repo) FetchByID(ctx context.Context, id string) (model.Order, error) {
	var order model.Order

	orderQuery := `
		SELECT 
			order_uid,
			track_number,
			entry,
			locale,
			internal_signature,
			customer_id,
			delivery_service,
			shardkey,
			sm_id,
			date_created,
			oof_shard
		FROM orders
		WHERE order_uid = $1
	`
	err := r.db.QueryRow(ctx, orderQuery, id).Scan(
		&order.OrderUID,
		&order.TrackNumber,
		&order.Entry,
		&order.Locale,
		&order.InternalSignature,
		&order.CustomerID,
		&order.DeliveryService,
		&order.ShardKey,
		&order.SmID,
		&order.DateCreated,
		&order.OofShard,
	)
	if err != nil {
		return model.Order{}, slerr.WithSource(err)
	}

	deliveryQuery := `
		SELECT 
			name, phone, zip, city, address, region, email
		FROM delivery
		WHERE order_uid = $1
	`
	err = r.db.QueryRow(ctx, deliveryQuery, id).Scan(
		&order.Delivery.Name,
		&order.Delivery.Phone,
		&order.Delivery.Zip,
		&order.Delivery.City,
		&order.Delivery.Address,
		&order.Delivery.Region,
		&order.Delivery.Email,
	)
	if err != nil {
		return model.Order{}, slerr.WithSource(err)
	}

	paymentQuery := `
		SELECT 
			transaction, request_id, currency, provider, amount, payment_dt,
			bank, delivery_cost, goods_total, custom_fee
		FROM payments
		WHERE transaction = $1
	`
	err = r.db.QueryRow(ctx, paymentQuery, id).Scan(
		&order.Payment.Transaction,
		&order.Payment.RequestID,
		&order.Payment.Currency,
		&order.Payment.Provider,
		&order.Payment.Amount,
		&order.Payment.PaymentDt,
		&order.Payment.Bank,
		&order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
		&order.Payment.CustomFee,
	)
	if err != nil {
		return model.Order{}, slerr.WithSource(err)
	}

	itemsQuery := `
		SELECT 
			chrt_id, track_number, price, rid, name, sale, size, total_price,
			nm_id, brand, status
		FROM items
		WHERE order_uid = $1
	`
	rows, err := r.db.Query(ctx, itemsQuery, id)
	if err != nil {
		return model.Order{}, slerr.WithSource(err)
	}
	defer rows.Close()

	for rows.Next() {
		var item model.Item
		err = rows.Scan(
			&item.ChrtID,
			&item.TrackNumber,
			&item.Price,
			&item.RID,
			&item.Name,
			&item.Sale,
			&item.Size,
			&item.TotalPrice,
			&item.NmID,
			&item.Brand,
			&item.Status,
		)
		if err != nil {
			return model.Order{}, slerr.WithSource(err)
		}
		order.Items = append(order.Items, item)
	}

	return order, nil
}

func (r *Repo) FetchAllIDs(ctx context.Context) ([]string, error) {
	query := `SELECT order_uid FROM orders`

	rows, err := r.db.Query(ctx, query)
	if err != nil {
		return nil, slerr.WithSource(err)
	}
	defer rows.Close()

	var ids []string
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, slerr.WithSource(err)
		}
		ids = append(ids, id)
	}

	return ids, nil
}
