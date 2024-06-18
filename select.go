package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/bank")
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		viewSelect(w, db)
	})

	// сохранение отправленных значений через поля формы.
	http.HandleFunc("/postform", func(w http.ResponseWriter, r *http.Request) {
		earthpos := r.FormValue("earthpos")
		sunPosition := r.FormValue("sunPosition")
		moonPosition := r.FormValue("moonPosition")

		sQuery := "INSERT INTO position (earthpos, sunPosition, moonPosition) VALUES (?, ?, ?)"
		_, err := db.Exec(sQuery, earthpos, sunPosition, moonPosition)
		if err != nil {
			http.Error(w, "Ошибка при вставке данных: "+err.Error(), http.StatusInternalServerError)
			return
		}

		// Вызов хранимой процедуры
		err = callProcedure(db)
		if err != nil {
			http.Error(w, "Ошибка при вызове хранимой процедуры: "+err.Error(), http.StatusInternalServerError)
			return
		}

		viewSelect(w, db)
	})

	fmt.Println("Server is listening on http://localhost:8181/")
	http.ListenAndServe(":8181", nil)
}

func callProcedure(db *sql.DB) error {
	_, err := db.Exec("CALL NameProc()")
	return err
}

func viewHeadQuery(w http.ResponseWriter, db *sql.DB, sShow string) {
	type sHead struct {
		clnme string
	}
	rows, err := db.Query(sShow)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fmt.Fprintf(w, "<tr>")
	for rows.Next() {
		var p sHead
		err := rows.Scan(&p.clnme)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "<td>%s</td>", p.clnme)
	}
	fmt.Fprintf(w, "</tr>")

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewSelectQuery(w http.ResponseWriter, db *sql.DB, sSelect string) {
	type individual struct {
		id           int
		earthpos     sql.NullString
		sunPosition  sql.NullString
		moonPosition sql.NullString
	}
	position := []individual{}

	// получение значений в массив position из структуры типа individual.
	rows, err := db.Query(sSelect)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		p := individual{}
		err := rows.Scan(&p.id, &p.earthpos, &p.sunPosition, &p.moonPosition)
		if err != nil {
			fmt.Println(err)
			continue
		}
		position = append(position, p)
	}

	// перебор массива из БД.
	for _, p := range position {
		fmt.Fprintf(w, "<tr><td>%d</td><td>%s</td><td>%s</td><td>%s</td></tr>",
			p.id,
			nullStringToEmpty(p.earthpos),
			nullStringToEmpty(p.sunPosition),
			nullStringToEmpty(p.moonPosition))
	}
}

func nullStringToEmpty(ns sql.NullString) string {
	if ns.Valid && ns.String != "" {
		return ns.String
	}
	return "EMPTY"
}

func nullInt64ToEmpty(ni sql.NullInt64) string {
	if ni.Valid {
		return strconv.FormatInt(ni.Int64, 5)
	}
	return "EMPTY"
}

func viewSelectVerQuery(w http.ResponseWriter, db *sql.DB, sSelect string) {
	type sVer struct {
		ver string
	}
	rows, err := db.Query(sSelect)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var p sVer
		err := rows.Scan(&p.ver)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, p.ver)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewSelect(w http.ResponseWriter, db *sql.DB) {
	// чтение шаблона.
	file, err := os.Open("select.html")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//	кодовая фраза для вставки значений из БД.
		if scanner.Text() != "@tr" && scanner.Text() != "@ver" {
			fmt.Fprintf(w, scanner.Text())
		}
		if scanner.Text() == "@tr" {
			viewHeadQuery(w, db, "select COLUMN_NAME AS clnme from information_schema.COLUMNS where TABLE_NAME='position' ORDER BY ORDINAL_POSITION")
			viewSelectQuery(w, db, "SELECT * FROM position ORDER BY id ASC")
		}
		if scanner.Text() == "@ver" {
			viewSelectVerQuery(w, db, "SELECT VERSION() AS ver")
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
