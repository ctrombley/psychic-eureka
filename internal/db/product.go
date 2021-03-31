package db

import (
	"context"
	"os"

	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type Store struct {
	ID   int
	Name string
}

type Product struct {
	ID    int
	Name  string
	Store Store
}

func LoadProducts(ctx context.Context, ids ...int) ([]Product, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	defer conn.Close(ctx)

	products := []Product{}
	pgIDs := &pgtype.Int8Array{}
	pgIDs.Set(ids)

	var productID, storeID int
	var productName, storeName string
	_, err = conn.QueryFunc(
		ctx,
		`SELECT
		     products.id,
		     products.name,
		     stores.id,
		     stores.name
		   FROM stock_bindings
		   INNER JOIN products
		     on products.id = stock_bindings.product_id
		   JOIN stores
		     ON stores.id = stock_bindings.store_id
		   WHERE products.id = ANY($1)`,
		[]interface{}{pgIDs},
		[]interface{}{
			&productID,
			&productName,
			&storeID,
			&storeName,
		},
		func(pgx.QueryFuncRow) error {
			products = append(products, Product{
				ID:   productID,
				Name: productName,
				Store: Store{
					ID:   storeID,
					Name: storeName,
				},
			})
			return nil
		},
	)

	if err != nil {
		return nil, err
	}

	return products, nil
}

func FindLocationCountByCoords(ctx context.Context, productID int, lat float32, lng float32) (int, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return 0, err
	}
	defer conn.Close(ctx)

	var count int
	err = conn.QueryRow(
		ctx,
		`SELECT COUNT(DISTINCT stores.id)
		 FROM "stores"
		 INNER JOIN "stock_bindings"
		   ON "stores"."id" = "stock_bindings"."store_id"
		 WHERE "stock_bindings"."product_id" = $1
		   AND (stock_bindings.disabled_at IS NULL AND stock_bindings.price IS NOT NULL AND stock_bindings.price > 0)
		   AND "stores"."discarded_at" IS NULL
		   AND (ST_Distance(lon_lat, ST_MakePoint($2,$3)::geography) < 32186)`,
		productID, lat, lng).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}

func FindLocationCountByZip(ctx context.Context, productID int, zip string) (int, error) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		return 0, err
	}
	defer conn.Close(ctx)

	var count int
	err = conn.QueryRow(
		ctx,
		`SELECT COUNT(DISTINCT stores.id)
		 FROM "stores"
		 INNER JOIN "zip_codes"
		   ON "zip_codes"."store_id" = "stores"."id"
		 INNER JOIN "stock_bindings"
		   ON "stores"."id" = "stock_bindings"."store_id"
		 WHERE "stock_bindings"."product_id" = $1
		   AND (stock_bindings.disabled_at IS NULL AND stock_bindings.price IS NOT NULL AND stock_bindings.price > 0)
		   AND "stores"."discarded_at" IS NULL
		   AND "zip_codes"."code" = $2`,
		productID, zip).Scan(&count)

	if err != nil {
		return 0, err
	}

	return count, nil
}
