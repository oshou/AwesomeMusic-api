package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.uber.org/zap"

	"github.com/oshou/AwesomeMusic-api/db"
	"github.com/oshou/AwesomeMusic-api/log"
	"github.com/pkg/errors"
)

var (
	maxconn  = 3
	dir      = "_db/seed"
	pattern  = "*.csv"
	filePath = "_config/config.yaml"
)

func main() {
	flag.StringVar(&dir, "dir", dir, "seeds directory")
	flag.StringVar(&pattern, "pattern", pattern, "seed files pattern")
	flag.IntVar(&maxconn, "maxconn", maxconn, "max db connection")
	flag.StringVar(&filePath, "filePath", filePath, "db config filePath")

	// Set Logger
	log.Init()
	defer log.Logger.Sync()
	flag.Parse()
	log.Init()

	// Load Environment
	if err := godotenv.Load(); err != nil {
		log.Logger.Fatal("failed to loading .env file", zap.Error(err))
	}

	// DB Connection
	err := db.Init()
	if err != nil {
		log.Logger.Fatal("failed to connect db", zap.Error(err))
	}
	defer func() {
		err := db.Pool.Close()
		if err != nil {
			log.Logger.Fatal("failed to release db", zap.Error(err))
		}
	}()

	run(pattern)
}

func run(pattern string) {
	fs, err := filepath.Glob(filepath.Join(dir, pattern))
	if err != nil {
		panic(errors.WithStack(err))
	}
	for _, f := range fs {
		fmt.Println(f)
		err := importCSV(f)
		if err != nil {
			panic(errors.WithStack(err))
		}
	}
}

func importCSV(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true

	for {
		// Columns
		cols, err := r.Read()
		if err != nil {
			return errors.WithStack(err)
		}
		fmt.Println("cols:", cols)

		// Rows
		rows, err := r.ReadAll()
		if err != nil {
			return errors.WithStack(err)
		}
		fmt.Println("rows:", rows)

		const sql = `INSERT INTO %s (%s)
									VALUES %s
									ON CONFLICT (%s)
									DO UPDATE SET %s`

	}
	return nil

	//table := strings.TrimSuffix(filepath.Base(path), ".csv")

	//var pkeys []string
	//for i, v := range cols {
	//	if strings.HasSuffix(v, "(PK)") {
	//		cols[i] = strings.Replace(v, "(PK)", "", -1)
	//		pkeys = append(pkeys, cols[i])
	//	}
	//}
	//pkey := strings.Join(pkeys, ",")

	//var setters []string
	//for _, c := range cols {
	//	setters = append(setters, fmt.Sprintf("%s = EXCLUDED.$s", c, c))
	//}

	//var vals []interface{}
	//for _, r := range rows {
	//	for _, v := range r {
	//		if v == "<NULL>" {
	//			vals = append(vals, nil)
	//		} else {
	//			vals = append(vals, v)
	//		}
	//	}
	//}

	//err = db.With(func(conn db.Conn) error {
	//	q := fmt.Sprintf(
	//		sql,
	//		strings.TrimSuffix(filepath.Base(path), ".csv"),
	//		strings.Join(cols, ","),
	//		db.BuilBulkedPlaceholders(len(cols), len(rows)),
	//		pkey,
	//		strings.Join(setters, ","),
	//	)
	//	_, err := conn.Exec(q, vals...)
	//})

	//if err != nil {
	//	return errors.WithStack(err)
	//}
	//return nil
}
