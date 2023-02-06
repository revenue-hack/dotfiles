package helper

import (
	"fmt"
	"os"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	gormysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbName string
)

func init() {
	dbName = fmt.Sprintf("test_%d", os.Getpid())
	// テスト用のDBに向き先を変更
	os.Setenv("DB_DATABASE", dbName)
}

func OpenDb(t *testing.T) *gorm.DB {
	db, err := openDb(dbName)
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func CloseDb(t *testing.T, db *gorm.DB) {
	err := closeDB(db)
	if err != nil {
		t.Fatal(err)
	}
}

func TruncateAll(t *testing.T) {
	if err := truncateAllTable(); err != nil {
		t.Fatal(err.Error())
	}
}

// マイグレーション用のDB接続
func openDb(dbName string) (*gorm.DB, error) {
	// multiStatements=trueがないと1ファイルに複数クエリ記述されている場合にエラーになる（migration用）
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true&loc=%s&multiStatements=true",
		os.Getenv("DB_USER_NAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST_NAME"),
		dbName,
		"Asia%2FTokyo",
	)

	db, err := gorm.Open(gormysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func closeDB(db *gorm.DB) error {
	d, err := db.DB()
	if err != nil {
		return err
	}
	return d.Close()
}

// テスト用のデータベースを構築
func createTestDB(dbName string) error {
	db, err := openDb("")
	if err != nil {
		return err
	}
	defer closeDB(db)
	if err = db.Exec(fmt.Sprintf("CREATE DATABASE `%s` DEFAULT CHARACTER SET utf8mb4;", dbName)).Error; err != nil {
		return err
	}
	if err = migrateDB(); err != nil {
		return err
	}
	// 一応全データ消しておく
	return truncateAllTable()
}

// テスト用のデータベースを破棄
func cleanTestDB() error {
	db, err := openDb("")
	if err != nil {
		return err
	}
	defer closeDB(db)
	return db.Exec(fmt.Sprintf("DROP DATABASE IF EXISTS `%s`;", dbName)).Error
}

// データベースのマイグレーションを実施
func migrateDB() error {
	gdb, err := openDb(dbName)
	if err != nil {
		return err
	}
	defer closeDB(gdb)

	db, err := gdb.DB()
	if err != nil {
		return err
	}

	dd, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		fmt.Sprintf("file://%s", os.Getenv("MIGRATE_DIR")),
		dbName,
		dd,
	)
	if err != nil {
		return err
	}
	return m.Up()
}

// 全テーブルのTruncateを行う
func truncateAllTable() error {
	db, err := openDb(dbName)
	if err != nil {
		return err
	}
	defer closeDB(db)

	rows, err := db.Raw("SHOW TABLES;").Rows()
	if err != nil {
		return err
	}

	if err := db.Exec("SET FOREIGN_KEY_CHECKS=0;").Error; err != nil {
		return err
	}

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return err
		}
		if table == "schema_migrations" {
			// migration管理用のテーブルはスキップ
			continue
		}
		if err := db.Exec("TRUNCATE " + table + ";").Error; err != nil {
			return err
		}
	}

	return db.Exec("SET FOREIGN_KEY_CHECKS=1;").Error
}
