package models

import "gorm.io/gorm"

type Schema struct {
}

type Columns struct {
	Type            string `json:"type"`
	Name            string `json:"name"`
	IsAutoIncrement bool   `json:"is_autoincrement"`
	IsPrimaryKey    bool   `json:"is_primary_key"`
}

func GetTableSchema(db *gorm.DB) ([]Columns, error) {
	var columns []Columns
	SQL := `
	SELECT
		cols.column_name as name,
		cols.data_type as type,
		cols.is_nullable,
		cols.column_default LIKE 'nextval%' AS is_autoincrement,
		tc.is_primary_key = 'YES' as is_primary_key
	FROM
    	information_schema.columns AS cols
	LEFT JOIN (
    	SELECT 
			kcu.table_name,
			kcu.table_schema,
			kcu.column_name,
			'YES' AS is_primary_key
    	FROM
        	information_schema.key_column_usage AS kcu
    	JOIN
			information_schema.table_constraints AS tc
		ON  kcu.table_schema = tc.table_schema
		AND kcu.table_name = tc.table_name
		AND kcu.constraint_name = tc.constraint_name
    	WHERE
      		tc.constraint_type = 'PRIMARY KEY'
	) AS tc
	ON  cols.table_name = tc.table_name
	AND cols.table_schema = tc.table_schema
	AND cols.column_name = tc.column_name
	WHERE
		cols.table_name = 'articles'
		AND cols.table_schema = 'public';
	`
	err := db.Raw(SQL).Scan(&columns).Error
	println(err)
	return columns, err
}

func GetTableData(db *gorm.DB) {
	db.Exec("select * from articles")
}

func ConnectDataBase() {

}
