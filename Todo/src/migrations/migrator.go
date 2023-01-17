package migrations

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"html/template"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
)

// to define Migration
type Migration struct {
	Version string
	Up      func(*sql.Tx) error
	Down    func(*sql.Tx) error
	done    bool
}

// to define Migrator
type Migrator struct {
	db         *sqlx.DB
	Versions   []string
	Migrations map[string]*Migration
}

var migrator = &Migrator{
	Versions:   []string{},
	Migrations: map[string]*Migration{},
}

//generating the migration file

// adding migration
func (m *Migrator) AddMigration(mg *Migration) {
	m.Migrations[mg.Version] = mg
	//	i = index
	i := 0

	for i < len(m.Versions) {
		if m.Versions[i] > mg.Version {
			break
		}
		i++
	}

	m.Versions = append(m.Versions, mg.Version)
	copy(m.Versions[i+1:], m.Versions[i:])
	m.Versions[i] = mg.Version

}

func Create(name string) error {
	version := time.Now().Format("20060102150405")

	in := struct {
		Version string
		Name    string
	}{
		Version: version,
		Name:    name,
	}

	var out bytes.Buffer

	t := template.Must(template.ParseFiles("./migrations/template.txt"))
	err := t.Execute(&out, in)
	if err != nil {
		return errors.New("Unable to execute template:" + err.Error())
	}

	fl, err := os.Create(fmt.Sprintf("./migrations/%s_%s.go", version, name))
	if err != nil {
		return errors.New("Create function giving error")
	}

	defer fl.Close()

	if _, err := fl.WriteString(out.String()); err != nil {
		return errors.New("Unable to write inside file" + err.Error())
	}

	fmt.Println("Generated new migration files...", fl.Name())
	return nil
}

func Init(db *sqlx.DB) (*Migrator, error) {
	migrator.db = db
	fmt.Println("91")
	// Create `schema_migrations` table to remember which migrations were executed.
	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS schema_migrations (
		version varchar(255)
	);`); err != nil {
		fmt.Println("Unable to create `schema_migrations` table", err)
		return migrator, err
	}
	fmt.Println("99")

	// Find out all the executed migrations
	rows, err := db.Query(`SELECT version FROM schema_migrations;`)
	fmt.Println("103")

	if err != nil {
		//return migrator, err
		panic(err)
	}
	fmt.Println("106")

	defer rows.Close()
	fmt.Println("109")

	// Mark the migrations as Done if it is already executed
	for rows.Next() {
		var version string
		err := rows.Scan(&version)
		if err != nil {
			return migrator, err
		}
		fmt.Println("118")

		if migrator.Migrations[version] != nil {
			migrator.Migrations[version].done = true
		}
	}
	fmt.Println("124")
	return migrator, err

}

func (m *Migrator) Up(step int) error {
	tx, err := m.db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	count := 0
	for _, v := range m.Versions {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[v]

		if mg.done {
			continue
		}

		fmt.Println("Running migration", mg.Version)
		fmt.Println("HEllllllll")
		if err := mg.Up(tx); err != nil {
			tx.Rollback()
			return err
		}
		fmt.Println("migrator.155")
		if _, err := tx.Exec(`INSERT INTO schema_migrations (version) VALUES ($1)`, mg.Version); err != nil {
			tx.Rollback()
			return err
		}
		// }`INSERT INTO schema_migrations (version) VALUES ?`
		fmt.Println("Finished running migration", mg.Version)
		fmt.Println("migrator.161")

		count++
	}

	tx.Commit()

	return nil
}

func (m *Migrator) Down(step int) error {
	tx, err := m.db.BeginTx(context.TODO(), &sql.TxOptions{})
	if err != nil {
		return err
	}

	count := 0
	for _, v := range reverse(m.Versions) {
		if step > 0 && count == step {
			break
		}

		mg := m.Migrations[v]

		if !mg.done {
			continue
		}

		fmt.Println("Reverting Migration", mg.Version)
		if err := mg.Down(tx); err != nil {
			tx.Rollback()
			panic(err)
		}

		if _, err := tx.Exec("DELETE FROM schema_migrations WHERE version = ($1)", mg.Version); err != nil {
			tx.Rollback()
			return err
		}
		fmt.Println("Finished reverting migration", mg.Version)

		count++
	}

	tx.Commit()

	return nil
}

func reverse(arr []string) []string {
	for i := 0; i < len(arr)/2; i++ {
		j := len(arr) - i - 1
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func (m *Migrator) MigrationStatus() error {
	for _, v := range m.Versions {
		mg := m.Migrations[v]

		if mg.done {
			fmt.Println(fmt.Sprintf("Migration %s... completed", v))
		} else {
			fmt.Println(fmt.Sprintf("Migration %s... pending", v))
		}
	}

	return nil
}
