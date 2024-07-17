package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	dbName := os.Getenv("MYSQL_DATABASE")
	dbUser := os.Getenv("MYSQL_USER")
	dbPasswd := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPasswd, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "../../internal/adapter/db/query",
		ModelPkgPath: "../../internal/domain/model",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(db)

	all := g.GenerateAllTable()
	invoice := g.GenerateModel("invoices",
		gen.FieldType("amount_due", "decimal.Decimal"),
		gen.FieldType("fee", "decimal.Decimal"),
		gen.FieldType("fee_rate", "decimal.Decimal"),
		gen.FieldType("consumption_tax", "decimal.Decimal"),
		gen.FieldType("tax_rate", "decimal.Decimal"),
		gen.FieldType("total_amount", "decimal.Decimal"),
	)
	companies := g.GenerateModel("companies")
	users := g.GenerateModel("users", gen.FieldRelate(field.HasOne, "Company", companies, &field.RelateConfig{}))
	g.ApplyBasic(invoice, companies, users)
	g.ApplyBasic(all...)

	g.Execute()
}
