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
	_, err := tx.Exec("CREATE TABLE users ( id serial not null unique, name varchar(255) not null, username varchar(255) not null unique, password_hash varchar(255) not null );")
	if err != nil {
		return err
	}

	return nil
}

func mig_20221230145525_init_schema_down(tx *sql.Tx) error {
	_, err := tx.Exec("DROP TABLE users")
	if err != nil {
		return err
	}

	return nil
}

// "BEGIN;

// CREATE TABLE users
// (
//     id            serial       not null unique,
//     name          varchar(255) not null,
//     username      varchar(255) not null unique,
//     password_hash varchar(255) not null
// );

// CREATE TABLE todo_lists
// (
//     id          serial       not null unique,
//     title       varchar(255) not null,
//     description varchar(255)
// );

// CREATE TABLE users_lists
// (
//     id      serial                                           not null unique,
//     user_id int references users (id) on delete cascade      not null,
//     list_id int references todo_lists (id) on delete cascade not null
// );

// CREATE TABLE todo_items
// (
//     id          serial       not null unique,
//     title       varchar(255) not null,
//     description varchar(255),
//     done        boolean      not null default false
// );

// CREATE TABLE lists_items
// (
//     id      serial                                           not null unique,
//     item_id int references todo_items (id) on delete cascade not null,
//     list_id int references todo_lists (id) on delete cascade not null
// );

// COMMIT;
// "
