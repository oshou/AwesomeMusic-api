package main

import (
	"encoding/csv"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
)

func importCSV(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.LazyQuotes = true

	cols, err := r.Read()
	if err != nil {
		return errors.WithStack(err)
	}

	const sql = `INSERT INTO %s (%s)
								VALUES %s
								ON CONFLICT (%s)
								DO UPDATE SET %s`

	table := strings.TrimSuffix(filepath.Base(path), ".csv")

	var pkeys []string
	for i, v := range cols {
		if strings.HasSuffix(v, "(PK)") {
			cols[i] = strings.Replace(v, "(PK)", "", -1)
			pkeys = append(pkeys, cols[i])
		}
	}
	pkey := strings.Join(pkeys, ",")

	var setters []string
	for _, c := range rows {
		for _, v := range r {
			if v == "<NULL>" {
				vals = append(vals, nil)
			} else {
				vals = append(vals, v)
			}
		}
	}

	err = db.With(func(conn db.Conn)) error {
		q := fmt.Sprintf(
			sql,
			strings.TrimSuffix(filepath.Base(path),".csv"),
			strings.Join(cols,","),
			db.BuilBulkedPlaceholders(len(cols),len(rows)),
			pkey,
			strings.Join(setters,",")
		)
		_,err := conn.Exec(q,vals...)
	})

	if err != nil{
		return errors.WithStack(err)
	}
	return nil
}
