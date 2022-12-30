package dtb

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/BetaLixT/tsqlx"
	// "go.uber.org/zap"
)

func RunMigrations(
	ctx context.Context,
	// lgr *zap.Logger,
	db *tsqlx.TracedDB,
	migrations []MigrationScript,
) error {

	tx := db.MustBegin()
	chck := ExistsEntity{}

	// - creating timestamp procedures if requried
	err := tx.Get(ctx, &chck, CheckTimestampProceduresExist)
	if err != nil {
		//	lgr.Error(
		//		"Failed fetching procedure info",
		//		zap.Error(err),
		//	)
		panic(fmt.Errorf("Failed fetching procedure info"))
	}

	if !chck.Exists {
		//	lgr.Info("Creating timestamp procedures")
		tx.MustExec(timestampProcedures.up)
	}

	// - creating migration table if required
	err = tx.Get(ctx, &chck, CheckMigrationExists)
	if err != nil {
		// lgr.Error(
		// 	"Failed fetching migration info",
		// 	zap.Error(err),
		// )
		panic(fmt.Errorf("Failed fetching migration info"))
	}
	var exMigrs []migrationEntity

	if !chck.Exists {
		// lgr.Info("Creating migration table")
		tx.MustExec(migrationTable.up)
		exMigrs = []migrationEntity{}
	} else {
		// lgr.Info("Fetching migration history")
		err = tx.Select(ctx, &exMigrs, GetAllMigrations)
		if err != nil {
			// lgr.Error(
			// 	"failed to fetch migrations",
			// 	zap.Error(err),
			// )
			panic(fmt.Errorf("Failed fetching migration"))
		}
	}
	sort.Slice(exMigrs, func(i, j int) bool {
		return exMigrs[i].Index < exMigrs[j].Index
	})

	exMigrsLen := len(exMigrs)

	for idx, migr := range migrations {
		if idx < exMigrsLen {
			if migr.key != exMigrs[idx].Key {
				panic(fmt.Errorf("migration key missmatch"))
			}
		} else {
			// lgr.Info("Running migration", zap.String("migration", migr.key))
			tx.MustExec(migr.up)
			tx.MustExec(AddMigration, migr.key)
		}
	}
	return tx.Commit()
}

type MigrationScript struct {
	key  string
	up   string
	down string
}

type migrationEntity struct {
	Index           int        `db:"index"`
	Key             string     `db:"key"`
	DateTimeCreated *time.Time `db:"datetimecreated"`
}

type ExistsEntity struct {
	Exists bool `db:"exists"`
}

var timestampProcedures = MigrationScript{
	up: `
		CREATE OR REPLACE FUNCTION trigger_set_datetimecreated()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.dateTimeCreated = NOW();
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
		
		CREATE OR REPLACE FUNCTION trigger_set_datetimeupdated()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.dateTimeUpdated = NOW();
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;

	  CREATE OR REPLACE FUNCTION version_update()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.version = OLD.version + 1;
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
		`,
	down: `
	  DROP FUNCTION version_update();
		DROP FUNCTION trigger_set_datetimeupdated();
		DROP FUNCTION trigger_set_datetimecreated();`,
}

var migrationTable = MigrationScript{
	up: `
		CREATE TABLE migrations (
			index SERIAL,
			key text PRIMARY KEY,
			datetimecreated timestamp with time zone NULL
		);
		
		CREATE TRIGGER set_migrations_datetimecreated
		BEFORE INSERT ON migrations
		FOR EACH ROW
		EXECUTE PROCEDURE trigger_set_datetimecreated();`,
	down: `
		DROP TRIGGER set_migrations_datetimecreated on migrations;
		DROP TABLE Migrations;`,
}

const (
	CheckTimestampProceduresExist = `
		SELECT EXISTS(
			SELECT * FROM (
				SELECT Count(p.proname) as count
				FROM pg_proc AS p
				JOIN pg_namespace n ON p.pronamespace = n.oid
				WHERE p.proname in (
					'trigger_set_datetimecreated', 
					'trigger_set_datetimeupdated'
					) 
					AND n.nspname = 'public'
			) as c
			WHERE c.count = 2
		) as exists`
	CheckMigrationExists = `
		SELECT EXISTS(
			SELECT * FROM pg_tables
			WHERE schemaname = 'public' AND tablename = 'migrations'
		) as exists`
	GetAllMigrations = `
		SELECT * FROM migrations`
	AddMigration = `
		INSERT INTO migrations (key) VALUES ($1)`
)
