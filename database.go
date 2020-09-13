package main

import (
	"log"

	migrate "github.com/rubenv/sql-migrate"
)

func migrateDatabase() {
	migrations := &migrate.MemoryMigrationSource{
		Migrations: []*migrate.Migration{
			{
				Id: "001",
				Up: []string{
					"CREATE TABLE `posts` (`ID`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,`post_title`	TEXT NOT NULL," +
						"`post_content`	TEXT NOT NULL);" +
						"insert into posts (post_title, post_content) values ('Hello World from Carlzberg', 'Hello World');"},
				Down: []string{
					"drop table posts;"},
			},
		},
	}
	_, err := migrate.Exec(db, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal(err)
	}
}
