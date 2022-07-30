package domain

import (
	"capi/errs"
	"capi/logger"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDB struct {
	db *sqlx.DB
}

func NewAccountRepositoryDB(db *sqlx.DB) *AccountRepositoryDB {
	return &AccountRepositoryDB{db}
}

func (d AccountRepositoryDB) InsertAccount(amount float64, customerId string, accountType string) (*Account, *errs.AppErr) {
	query := `INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5) RETURNING account_id`
	var time = time.Now().Format("2006-01-02T15:04:05")
	var a = Account{
		CustomerId:  customerId,
		OpeningDate: time,
		AccountType: accountType,
		Amount:      amount,
		Status:      "1",
	}
	// err := row.Scan(&c.ID, &c.Name, &c.DateOfBirth, &c.City, &c.ZipCode, &c.Status)
	// _, err := d.db.NamedExec(`INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) VALUES ($1, $2, $3, $4, $5)`,
	// 	map[string]interface{}{
	// 		"$1": a.CustomerId,
	// 		"$2": a.OpeningDate,
	// 		"$3": a.AccountType,
	// 		"$4": a.Amount,
	// 		"$5": a.Status,
	// 	})
	// if err != nil {
	// 	logger.Error("Error unexpected " + err.Error())
	// 	return nil, errs.NewUnExpectedError("Un expected error when insert rows")
	// }
	// rowsAffected, err := result.RowsAffected()
	// if err != nil {
	// 	return nil, errs.NewUnExpectedError("Unexpected error when insert rows")
	// }
	// if rowsAffected > 0 {
	// 	return &a, nil
	// } else {
	// 	return nil, errs.NewUnExpectedError("Unexpected error when insert data!")
	// }
	id := 0
	err := d.db.QueryRow(query, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status).Scan(&id)
	if err != nil {
		logger.Error("Error when insert " + err.Error())
	}
	fmt.Println("New record ID is:", id)
	return &a, nil
}
