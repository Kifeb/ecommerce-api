package config_test

import (
	"ecommerce_api/config"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestDBConnection(t *testing.T) {
	db, err := config.NewDB()
	if err != nil {
		t.Errorf("Error pinging databse: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		t.Errorf("Error pinging databse: %v", err)
	}
	fmt.Println("Berhasil Terkoneksi")
}
