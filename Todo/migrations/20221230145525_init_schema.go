package migrations

import "database/sql"

func init() {
	migrator.AddMigration(&Migration{
		Version: "20221230145525",
		Up:      mig_20221230145525_init_schema_up,
		Down:    mig_20221230145525_init_schema_down,
	})
}

func mig_20221230145525_init_schema_up(tx *sql.Tx) error {
	_, err := tx.Exec("CREATE TABLE todos (task varchar(255));")
	if err != nil {
		return err
	}

	return nil
}

func mig_20221230145525_init_schema_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE todos")
	if err != nil {
		return err
	}

	return nil
}
