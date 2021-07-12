package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"jike-goLearning/week2/dao"
)


func main(){
	count,err:=dao2.GetCount()
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("sql error: %T %v\n", errors.Cause(err), errors.Cause(err))
		fmt.Printf("stack trace: %+v\n", err)
		fmt.Println(errors.Unwrap(err))
	} else {
		fmt.Println(count)
	}
}
