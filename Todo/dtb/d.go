package dtb

func GetMigrationScripts() []MigrationScript {
	migrationScripts := []MigrationScript{
		{
			key: "initial",
			up: `
				CREATE TABLE todos (
					id uuid PRIMARY KEY,
					task text NOT NULL,
					done boolean DEFAULT false,
					datetimecreated timestamp with time zone,
					datetimeupdated timestamp with time zone,
				);
				CREATE TRIGGER set_todos_datetimecreated
				BEFORE INSERT ON todos
				FOR EACH ROW
				EXECUTE PROCEDURE trigger_set_datetimecreated();
				
				CREATE TRIGGER set_todos_datetimeupdated_in
				BEFORE INSERT ON todos
				FOR EACH ROW
				EXECUTE PROCEDURE trigger_set_datetimeupdated();

				CREATE TRIGGER set_todos_datetimeupdated_up
				BEFORE UPDATE ON todos
				FOR EACH ROW
				EXECUTE PROCEDURE trigger_set_datetimeupdated();

				CREATE TRIGGER todos_version_update
				BEFORE UPDATE ON todos
				FOR EACH ROW
				EXECUTE PROCEDURE version_update();`,
			down: `
				DROP TABLE todos;
				DROP TRIGGER todos_version_update on todos;
				DROP TRIGGER set_todos_datetimeupdated_up on todos;
				DROP TRIGGER set_todos_datetimeupdated_in on todos;
				DROP TRIGGER set_todos_datetimecreated on todos;`,
		},
	}
	return migrationScripts
}
