package main

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/k0kubun/pp"
)

type User struct {
	UserID         int     `db:"user_id"`
	LoginID        string  `db:"login_id"`
	PasswordDigest string  `db:"password_digest"`
	UserName       string  `db:"user_name"`
	GroupID        int     `db:"group_id"`
	Email          *string `db:"email"`
	DeletedAt      *string `db:"deleted_at"`
	CreatedAt      *string `db:"created_at"`
	CreatedUserID  int     `db:"create_user_id"`
	UpdatedAt      *string `db:"updated_at"`
	UpdatedUserID  int     `db:"updated_user_id"`
	LockVersion    int     `db:"lock_version"`
}

type SQLDatabase struct {
	database *sql.DB
}

func NewDatabase(driver, dsn string) (*SQLDatabase, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return &SQLDatabase{database: db}, nil
}

func (d *SQLDatabase) Begin() (*sql.Tx, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	return d.database.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})
}

func (d *SQLDatabase) Close() error {
	return d.database.Close()
}

func (d *SQLDatabase) selectUserByName(db *sql.Tx) (*User, error) {
	//fmt.Println("start")
	//driver := "mysql"
	//dsn := "root:password@tcp(127.0.0.1:3306)/asagaku"
	//db, err := sql.Open(driver, dsn)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
	defer cancel()

	u := &User{}
	if err := db.QueryRowContext(ctx, "select * from t_user limit 1").Scan(
		&u.UserID,
		&u.LoginID,
		&u.PasswordDigest,
		&u.UserName,
		&u.GroupID,
		&u.Email,
		&u.DeletedAt,
		&u.CreatedAt,
		&u.CreatedUserID,
		&u.UpdatedAt,
		&u.UpdatedUserID,
		&u.LockVersion,
	); err != nil {
		log.Fatalln(err)
	}
	return u, nil
}

func main() {
	driver := "mysql"
	dsn := "root:password@tcp(127.0.0.1:3306)/asagaku"

	db, err := NewDatabase(driver, dsn)
	if err != nil {
		log.Fatalln(err)
	}

	tx, err := db.database.Begin()
	u, err := db.selectUserByName(tx)
	pp.Print(u)
	if err != nil {
		log.Fatalln(err)
	}
}
