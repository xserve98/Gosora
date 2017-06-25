/* WIP Under Construction */
package qgen

//import "fmt"
import "strings"
import "errors"

func init() {
	DB_Registry = append(DB_Registry,
		&Mysql_Adapter{Name:"mysql",Buffer:make(map[string]string)},
	)
}

type Mysql_Adapter struct
{
	Name string
	Buffer map[string]string
	BufferOrder []string // Map iteration order is random, so we need this to track the order, so we don't get huge diffs every commit
}

func (adapter *Mysql_Adapter) GetName() string {
	return adapter.Name
}

func (adapter *Mysql_Adapter) GetStmt(name string) string {
	return adapter.Buffer[name]
}

func (adapter *Mysql_Adapter) GetStmts() map[string]string {
	return adapter.Buffer
}

func (adapter *Mysql_Adapter) SimpleInsert(name string, table string, columns string, fields string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	if len(columns) == 0 {
		return "", errors.New("No columns found for SimpleInsert")
	}
	if len(fields) == 0 {
		return "", errors.New("No input data found for SimpleInsert")
	}
	
	var querystr string = "INSERT INTO `" + table + "`("
	
	// Escape the column names, just in case we've used a reserved keyword
	for _, column := range _process_columns(columns) {
		if column.Type == "function" {
			querystr += column.Left + ","
		} else {
			querystr += "`" + column.Left + "`,"
		}
	}
	
	// Remove the trailing comma
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += ") VALUES ("
	for _, field := range _process_fields(fields) {
		querystr += field.Name + ","
	}
	querystr = querystr[0:len(querystr) - 1]
	
	adapter.push_statement(name,querystr + ")")
	return querystr + ")", nil
}

func (adapter *Mysql_Adapter) SimpleReplace(name string, table string, columns string, fields string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	if len(columns) == 0 {
		return "", errors.New("No columns found for SimpleInsert")
	}
	if len(fields) == 0 {
		return "", errors.New("No input data found for SimpleInsert")
	}
	
	var querystr string = "REPLACE INTO `" + table + "`("
	
	// Escape the column names, just in case we've used a reserved keyword
	for _, column := range _process_columns(columns) {
		if column.Type == "function" {
			querystr += column.Left + ","
		} else {
			querystr += "`" + column.Left + "`,"
		}
	}
	// Remove the trailing comma
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += ") VALUES ("
	for _, field := range _process_fields(fields) {
		querystr += field.Name + ","
	}
	querystr = querystr[0:len(querystr) - 1]
	
	adapter.push_statement(name,querystr + ")")
	return querystr + ")", nil
}

func (adapter *Mysql_Adapter) SimpleUpdate(name string, table string, set string, where string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	if set == "" {
		return "", errors.New("You need to set data in this update statement")
	}
	
	var querystr string = "UPDATE `" + table + "` SET "
	for _, item := range _process_set(set) {
		querystr += "`" + item.Column + "` ="
		for _, token := range item.Expr {
			switch(token.Type) {
				case "function","operator","number","substitute":
					querystr += " " + token.Contents + ""
				case "column":
					querystr += " `" + token.Contents + "`"
				case "string":
					querystr += " '" + token.Contents + "'"
			}
		}
		querystr += ","
	}
	
	// Remove the trailing comma
	querystr = querystr[0:len(querystr) - 1]
	
	// Add support for BETWEEN x.x
	if len(where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						querystr += " `" + token.Contents + "`"
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleDelete(name string, table string, where string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	if where == "" {
		return "", errors.New("You need to specify what data you want to delete")
	}
	
	var querystr string = "DELETE FROM `" + table + "` WHERE"
	
	// Add support for BETWEEN x.x
	for _, loc := range _process_where(where) {
		for _, token := range loc.Expr {
			switch(token.Type) {
				case "function","operator","number","substitute":
					querystr += " " + token.Contents + ""
				case "column":
					querystr += " `" + token.Contents + "`"
				case "string":
					querystr += " '" + token.Contents + "'"
				default:
					panic("This token doesn't exist o_o")
			}
		}
		querystr += " AND"
	}
	
	querystr = strings.TrimSpace(querystr[0:len(querystr) - 4])
	adapter.push_statement(name,querystr)
	return querystr, nil
}

// We don't want to accidentally wipe tables, so we'll have a seperate method for purging tables instead
func (adapter *Mysql_Adapter) Purge(name string, table string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	adapter.push_statement(name,"DELETE FROM `" + table + "`")
	return "DELETE FROM `" + table + "`", nil
}

func (adapter *Mysql_Adapter) SimpleSelect(name string, table string, columns string, where string, orderby string, limit string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	if len(columns) == 0 {
		return "", errors.New("No columns found for SimpleSelect")
	}
	
	// Slice up the user friendly strings into something easier to process
	var colslice []string = strings.Split(strings.TrimSpace(columns),",")
	
	var querystr string = "SELECT "
	
	// Escape the column names, just in case we've used a reserved keyword
	for _, column := range colslice {
		querystr += "`" + strings.TrimSpace(column) + "`,"
	}
	// Remove the trailing comma
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += " FROM `" + table + "`"
	
	// Add support for BETWEEN x.x
	if len(where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						querystr += " `" + token.Contents + "`"
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if len(orderby) != 0 {
		querystr += " ORDER BY "
		for _, column := range _process_orderby(orderby) {
			querystr += column.Column + " " + strings.ToUpper(column.Order) + ","
		}
		querystr = querystr[0:len(querystr) - 1]
	}
	
	if limit != "" {
		querystr += " LIMIT " + limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleLeftJoin(name string, table1 string, table2 string, columns string, joiners string, where string, orderby string, limit string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table1 == "" {
		return "", errors.New("You need a name for the left table")
	}
	if table2 == "" {
		return "", errors.New("You need a name for the right table")
	}
	if len(columns) == 0 {
		return "", errors.New("No columns found for SimpleLeftJoin")
	}
	if len(joiners) == 0 {
		return "", errors.New("No joiners found for SimpleLeftJoin")
	}
	
	var querystr string = "SELECT "
	
	for _, column := range _process_columns(columns) {
		var source, alias string
		
		// Escape the column names, just in case we've used a reserved keyword
		if column.Table != "" {
			source = "`" + column.Table + "`.`" + column.Left + "`"
		} else if column.Type == "function" {
			source = column.Left
		} else {
			source = "`" + column.Left + "`"
		}
		
		if column.Alias != "" {
			alias = " AS `" + column.Alias + "`"
		}
		querystr += source + alias + ","
	}
	
	// Remove the trailing comma
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += " FROM `" + table1 + "` LEFT JOIN `" + table2 + "` ON "
	for _, joiner := range _process_joiner(joiners) {
		querystr += "`" + joiner.LeftTable + "`.`" + joiner.LeftColumn + "` " + joiner.Operator + " `" + joiner.RightTable + "`.`" + joiner.RightColumn + "` AND "
	}
	// Remove the trailing AND
	querystr = querystr[0:len(querystr) - 4]
	
	// Add support for BETWEEN x.x
	if len(where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						halves := strings.Split(token.Contents,".")
						if len(halves) == 2 {
							querystr += " `" + halves[0] + "`.`" + halves[1] + "`"
						} else {
							querystr += " `" + token.Contents + "`"
						}
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if len(orderby) != 0 {
		querystr += " ORDER BY "
		for _, column := range _process_orderby(orderby) {
			querystr += column.Column + " " + strings.ToUpper(column.Order) + ","
		}
		querystr = querystr[0:len(querystr) - 1]
	}
	
	if limit != "" {
		querystr += " LIMIT " + limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleInnerJoin(name string, table1 string, table2 string, columns string, joiners string, where string, orderby string, limit string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table1 == "" {
		return "", errors.New("You need a name for the left table")
	}
	if table2 == "" {
		return "", errors.New("You need a name for the right table")
	}
	if len(columns) == 0 {
		return "", errors.New("No columns found for SimpleInnerJoin")
	}
	if len(joiners) == 0 {
		return "", errors.New("No joiners found for SimpleInnerJoin")
	}
	
	var querystr string = "SELECT "
	
	for _, column := range _process_columns(columns) {
		var source, alias string
		
		// Escape the column names, just in case we've used a reserved keyword
		if column.Table != "" {
			source = "`" + column.Table + "`.`" + column.Left + "`"
		} else if column.Type == "function" {
			source = column.Left
		} else {
			source = "`" + column.Left + "`"
		}
		
		if column.Alias != "" {
			alias = " AS `" + column.Alias + "`"
		}
		querystr += source + alias + ","
	}
	
	// Remove the trailing comma
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += " FROM `" + table1 + "` INNER JOIN `" + table2 + "` ON "
	for _, joiner := range _process_joiner(joiners) {
		querystr += "`" + joiner.LeftTable + "`.`" + joiner.LeftColumn + "` " + joiner.Operator + " `" + joiner.RightTable + "`.`" + joiner.RightColumn + "` AND "
	}
	// Remove the trailing AND
	querystr = querystr[0:len(querystr) - 4]
	
	// Add support for BETWEEN x.x
	if len(where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						halves := strings.Split(token.Contents,".")
						if len(halves) == 2 {
							querystr += " `" + halves[0] + "`.`" + halves[1] + "`"
						} else {
							querystr += " `" + token.Contents + "`"
						}
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if len(orderby) != 0 {
		querystr += " ORDER BY "
		for _, column := range _process_orderby(orderby) {
			querystr += column.Column + " " + strings.ToUpper(column.Order) + ","
		}
		querystr = querystr[0:len(querystr) - 1]
	}
	
	if limit != "" {
		querystr += " LIMIT " + limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleInsertSelect(name string, ins DB_Insert, sel DB_Select) (string, error) {
	/* Insert Portion */
	
	var querystr string = "INSERT INTO `" + ins.Table + "`("
	
	// Escape the column names, just in case we've used a reserved keyword
	for _, column := range _process_columns(ins.Columns) {
		if column.Type == "function" {
			querystr += column.Left + ","
		} else {
			querystr += "`" + column.Left + "`,"
		}
	}
	querystr = querystr[0:len(querystr) - 1] + ") SELECT"
	
	/* Select Portion */
	
	for _, column := range _process_columns(sel.Columns) {
		var source, alias string
		
		// Escape the column names, just in case we've used a reserved keyword
		if column.Type == "function" || column.Type == "substitute" {
			source = column.Left
		} else {
			source = "`" + column.Left + "`"
		}
		
		if column.Alias != "" {
			alias = " AS `" + column.Alias + "`"
		}
		querystr += " " + source + alias + ","
	}
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += " FROM `" + sel.Table + "`"
	
	// Add support for BETWEEN x.x
	if len(sel.Where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(sel.Where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						querystr += " `" + token.Contents + "`"
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if len(sel.Orderby) != 0 {
		querystr += " ORDER BY "
		for _, column := range _process_orderby(sel.Orderby) {
			querystr += column.Column + " " + strings.ToUpper(column.Order) + ","
		}
		querystr = querystr[0:len(querystr) - 1]
	}
	
	if sel.Limit != "" {
		querystr += " LIMIT " + sel.Limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleInsertLeftJoin(name string, ins DB_Insert, sel DB_Join) (string, error) {
	/* Insert Portion */
	
	var querystr string = "INSERT INTO `" + ins.Table + "`("
	
	// Escape the column names, just in case we've used a reserved keyword
	for _, column := range _process_columns(ins.Columns) {
		if column.Type == "function" {
			querystr += column.Left + ","
		} else {
			querystr += "`" + column.Left + "`,"
		}
	}
	querystr = querystr[0:len(querystr) - 1] + ") SELECT"
	
	/* Select Portion */
	
	for _, column := range _process_columns(sel.Columns) {
		var source, alias string
		
		// Escape the column names, just in case we've used a reserved keyword
		if column.Table != "" {
			source = "`" + column.Table + "`.`" + column.Left + "`"
		} else if column.Type == "function" {
			source = column.Left
		} else {
			source = "`" + column.Left + "`"
		}
		
		if column.Alias != "" {
			alias = " AS `" + column.Alias + "`"
		}
		querystr += " " + source + alias + ","
	}
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += " FROM `" + sel.Table1 + "` LEFT JOIN `" + sel.Table2 + "` ON "
	for _, joiner := range _process_joiner(sel.Joiners) {
		querystr += "`" + joiner.LeftTable + "`.`" + joiner.LeftColumn + "` " + joiner.Operator + " `" + joiner.RightTable + "`.`" + joiner.RightColumn + "` AND "
	}
	querystr = querystr[0:len(querystr) - 4]
	
	// Add support for BETWEEN x.x
	if len(sel.Where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(sel.Where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						halves := strings.Split(token.Contents,".")
						if len(halves) == 2 {
							querystr += " `" + halves[0] + "`.`" + halves[1] + "`"
						} else {
							querystr += " `" + token.Contents + "`"
						}
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if len(sel.Orderby) != 0 {
		querystr += " ORDER BY "
		for _, column := range _process_orderby(sel.Orderby) {
			querystr += column.Column + " " + strings.ToUpper(column.Order) + ","
		}
		querystr = querystr[0:len(querystr) - 1]
	}
	
	if sel.Limit != "" {
		querystr += " LIMIT " + sel.Limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleInsertInnerJoin(name string, ins DB_Insert, sel DB_Join) (string, error) {
	/* Insert Portion */
	
	var querystr string = "INSERT INTO `" + ins.Table + "`("
	
	// Escape the column names, just in case we've used a reserved keyword
	for _, column := range _process_columns(ins.Columns) {
		if column.Type == "function" {
			querystr += column.Left + ","
		} else {
			querystr += "`" + column.Left + "`,"
		}
	}
	querystr = querystr[0:len(querystr) - 1] + ") SELECT"
	
	/* Select Portion */
	
	for _, column := range _process_columns(sel.Columns) {
		var source, alias string
		
		// Escape the column names, just in case we've used a reserved keyword
		if column.Table != "" {
			source = "`" + column.Table + "`.`" + column.Left + "`"
		} else if column.Type == "function" {
			source = column.Left
		} else {
			source = "`" + column.Left + "`"
		}
		
		if column.Alias != "" {
			alias = " AS `" + column.Alias + "`"
		}
		querystr += " " + source + alias + ","
	}
	querystr = querystr[0:len(querystr) - 1]
	
	querystr += " FROM `" + sel.Table1 + "` INNER JOIN `" + sel.Table2 + "` ON "
	for _, joiner := range _process_joiner(sel.Joiners) {
		querystr += "`" + joiner.LeftTable + "`.`" + joiner.LeftColumn + "` " + joiner.Operator + " `" + joiner.RightTable + "`.`" + joiner.RightColumn + "` AND "
	}
	querystr = querystr[0:len(querystr) - 4]
	
	// Add support for BETWEEN x.x
	if len(sel.Where) != 0 {
		querystr += " WHERE"
		for _, loc := range _process_where(sel.Where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						halves := strings.Split(token.Contents,".")
						if len(halves) == 2 {
							querystr += " `" + halves[0] + "`.`" + halves[1] + "`"
						} else {
							querystr += " `" + token.Contents + "`"
						}
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if len(sel.Orderby) != 0 {
		querystr += " ORDER BY "
		for _, column := range _process_orderby(sel.Orderby) {
			querystr += column.Column + " " + strings.ToUpper(column.Order) + ","
		}
		querystr = querystr[0:len(querystr) - 1]
	}
	
	if sel.Limit != "" {
		querystr += " LIMIT " + sel.Limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) SimpleCount(name string, table string, where string, limit string) (string, error) {
	if name == "" {
		return "", errors.New("You need a name for this statement")
	}
	if table == "" {
		return "", errors.New("You need a name for this table")
	}
	
	var querystr string = "SELECT COUNT(*) AS `count` FROM `" + table + "`"
	
	// Add support for BETWEEN x.x
	if len(where) != 0 {
		querystr += " WHERE"
		//fmt.Println("SimpleCount:",name)
		//fmt.Println("where:",where)
		//fmt.Println("_process_where:",_process_where(where))
		for _, loc := range _process_where(where) {
			for _, token := range loc.Expr {
				switch(token.Type) {
					case "function","operator","number","substitute":
						querystr += " " + token.Contents + ""
					case "column":
						querystr += " `" + token.Contents + "`"
					case "string":
						querystr += " '" + token.Contents + "'"
					default:
						panic("This token doesn't exist o_o")
				}
			}
			querystr += " AND"
		}
		querystr = querystr[0:len(querystr) - 4]
	}
	
	if limit != "" {
		querystr += " LIMIT " + limit
	}
	
	querystr = strings.TrimSpace(querystr)
	adapter.push_statement(name,querystr)
	return querystr, nil
}

func (adapter *Mysql_Adapter) Write() error {
	var stmts, body string
	
	for _, name := range adapter.BufferOrder {
		stmts += "var " + name + "_stmt *sql.Stmt\n"
		body += `	
	log.Print("Preparing ` + name + ` statement.")
	` + name + `_stmt, err = db.Prepare("` + adapter.Buffer[name] + `")
	if err != nil {
		return err
	}
	`
	}
	
	out := `// Code generated by Gosora. More below:
/* This file was generated by Gosora's Query Generator. Please try to avoid modifying this file, as it might change at any time. */
// +build !pgsql !sqlite !mssql
package main

import "log"
import "database/sql"

` + stmts + `
func gen_mysql() (err error) {
	if debug {
		log.Print("Building the generated statements")
	}
` + body + `
	return nil
}
`
	return write_file("./gen_mysql.go", out)
}

// Internal method, not exposed in the interface
func (adapter *Mysql_Adapter) push_statement(name string, querystr string ) {
	adapter.Buffer[name] = querystr
	adapter.BufferOrder = append(adapter.BufferOrder,name)
}
