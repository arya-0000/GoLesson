package dao


import (
	"database/sql"
	"errors"
)
func GetCount() (int64,error){
	var count int64
	return count, errors.Unwrap(sql.ErrNoRows)
}
