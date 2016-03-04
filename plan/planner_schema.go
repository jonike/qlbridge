package plan

import (
	u "github.com/araddon/gou"
)

var (
	_ = u.EMPTY
)

func (m *PlannerDefault) WalkCommand(p *Command) error {
	u.Debugf("VisitCommand %+v", p.Stmt)
	return ErrNotImplemented
}

func (m *PlannerDefault) WalkShow(p *Show) error {

	//u.Debugf("exec.VisitShow create?%v  identity=%q  raw=%s", stmt.Create, stmt.Identity, stmt.Raw)
	/*

		taskName := "show"
		stmt := p.Stmt
		var source schema.Scanner
		proj := rel.NewProjection()

		raw := strings.ToLower(stmt.Raw)
		switch {
		case stmt.Create && strings.ToLower(stmt.CreateWhat) == "table":
			// SHOW CREATE TABLE
			tbl, _ := m.Ctx.Schema.Table(stmt.Identity)
			if tbl == nil {
				u.Warnf("no table? %q", stmt.Identity)
				return fmt.Errorf("No table found for %q", stmt.Identity)
			}
			source = membtree.NewStaticDataSource("tables", 0, [][]driver.Value{{stmt.Identity, tbl}}, []string{"Table", "Create Table"})
			proj.AddColumnShort("Table", value.StringType)
			proj.AddColumnShort("Create Table", value.StringType)
			taskName = "show-create-table"

		case stmt.ShowType == "columns":
			// SHOW FULL COLUMNS FROM `user` FROM `mysql` LIKE '%';
			// SHOW COLUMNS FROM `user` FROM `mysql` LIKE '%';
			tbl, _ := m.Ctx.Schema.Table(stmt.Identity)
			if tbl == nil {
				u.Warnf("no table? %q", stmt.Identity)
				return fmt.Errorf("No table found for %q", stmt.Identity)
			}
			source, proj = DescribeTable(tbl, stmt.Full)
			taskName = "show columns"
		case strings.ToLower(stmt.Identity) == "variables":
			// SHOW variables;
			vals := make([][]driver.Value, 2)
			vals[0] = []driver.Value{"auto_increment_increment", "1"}
			vals[1] = []driver.Value{"collation", "utf8"}
			source = membtree.NewStaticDataSource("variables", 0, vals, []string{"Variable_name", "Value"})
			proj.AddColumnShort("Variable_name", value.StringType)
			proj.AddColumnShort("Value", value.StringType)

		case strings.ToLower(stmt.Identity) == "databases":
			// SHOW databases;
			source = membtree.NewStaticDataSource("databases", 0, [][]driver.Value{{m.Ctx.Schema.Name}}, []string{"Database"})
			proj.AddColumnShort("Database", value.StringType)
			taskName = "show-databases"
		case strings.ToLower(stmt.Identity) == "collation":
			// SHOW collation;
			vals := make([][]driver.Value, 1)
			// utf8_general_ci          | utf8     |  33 | Yes     | Yes      |       1 |
			vals[0] = []driver.Value{"utf8_general_ci", "utf8", 33, "Yes", "Yes", 1}
			cols := []string{"Collation", "Charset", "Id", "Default", "Compiled", "Sortlen"}
			source = membtree.NewStaticDataSource("collation", 0, vals, cols)
			proj.AddColumnShort("Collation", value.StringType)
			proj.AddColumnShort("Charset", value.StringType)
			proj.AddColumnShort("Id", value.IntType)
			proj.AddColumnShort("Default", value.StringType)
			proj.AddColumnShort("Compiled", value.StringType)
			proj.AddColumnShort("Sortlen", value.IntType)
			taskName = "show-collation"
		case strings.HasPrefix(raw, "show session"):
			//SHOW SESSION VARIABLES LIKE 'lower_case_table_names';
			source, proj = ShowVariables("lower_case_table_names", 0)
			taskName = "show-session-vars"
		case strings.ToLower(stmt.ShowType) == "tables" || strings.ToLower(stmt.Identity) == m.Ctx.Schema.Name:
			if stmt.Full {
				// SHOW FULL TABLES;
				u.Debugf("show tables: %+v", m.Ctx)
				tables := m.Ctx.Schema.Tables()
				vals := make([][]driver.Value, len(tables))
				row := 0
				for _, tbl := range tables {
					vals[row] = []driver.Value{tbl, "BASE TABLE"}
					row++
				}
				source = membtree.NewStaticDataSource("tables", 0, vals, []string{"Tables", "Table_type"})
				proj.AddColumnShort("Tables", value.StringType)
				proj.AddColumnShort("Table_type", value.StringType)
				taskName = "show-full-tables"
			} else {
				// SHOW TABLES;
				source, proj = ShowTables(m.Ctx)
				taskName = "show tables"

			}
		case strings.ToLower(stmt.Identity) == "procedure":
			// SHOW PROCEDURE STATUS WHERE Db='mydb'
			return m.emptyTask(p, "Procedures")
		case strings.ToLower(stmt.Identity) == "function":
			// SHOW FUNCTION STATUS WHERE Db='mydb'
			return m.emptyTask(p, "Function")
		default:
			// SHOW FULL TABLES FROM `auths`
			desc := rel.SqlDescribe{}
			desc.Identity = stmt.Identity
			return m.WalkDescribe(&Describe{Stmt: &desc})
		}

		//tasks := m.TaskMaker.Sequential(taskName)

		sourcePlan := NewSourceStaticPlan(m.Ctx)
		sourceTask := NewSource(sourcePlan, source)

		p.Add(sourceTask)
		m.Ctx.Projection = NewProjectionStatic(proj)

		if stmt.Like != nil {
			// TODO:  this Like needs to be written to support which column we are filtering on
			// where := NewWhereFinal(m.Ctx, stmt)
			// tasks.Add(where)
		}
	*/
	return nil
}

// DESCRIBE statements
//
func (m *PlannerDefault) WalkDescribe(p *Describe) error {
	u.Debugf("VisitDescribe %+v", p.Stmt)
	/*
		if m.Ctx == nil || m.Ctx.Schema == nil {
			return ErrNoSchemaSelected
		}
		tbl, err := m.Ctx.Schema.Table(strings.ToLower(p.Stmt.Identity))
		if err != nil {
			u.Errorf("could not get table: %v", err)
			return err
		}
		source, proj := DescribeTable(tbl, false)
		m.Ctx.Projection = NewProjectionStatic(proj)

		sourcePlan := NewSourceStaticPlan(m.Ctx)
		sourceTask := NewSource(sourcePlan, source)

		//u.Infof("source:  %#v", source)
		p.Add(sourceTask)
	*/
	return nil
}

/*
func (m *PlannerDefault) emptyTask(p Task, name string) error {
	source := membtree.NewStaticDataSource(name, 0, nil, []string{name})
	proj := rel.NewProjection()
	proj.AddColumnShort(name, value.StringType)
	m.Ctx.Projection = NewProjectionStatic(proj)
	sourcePlan := NewSourceStaticPlan(m.Ctx)
	sourceTask := NewSource(sourcePlan, source)
	p.Add(sourceTask)
	return nil
}

func FieldsAsMessages(tbl *schema.Table) []schema.Message {
	msgs := make([]schema.Message, len(tbl.Fields))
	for i, f := range tbl.Fields {
		msgs[i] = f
	}
	return msgs
}

// Describe A table
//
func DescribeTable(tbl *schema.Table, full bool) (schema.Scanner, *rel.Projection) {
	if len(tbl.Fields) == 0 {
		u.Warnf("NO Fields!!!!! for %s p=%p", tbl.Name, tbl)
	}
	proj := rel.NewProjection()
	if full {
		for _, f := range schema.DescribeFullHeaders {
			proj.AddColumnShort(string(f.Name), f.Type)
		}
	} else {
		for _, f := range schema.DescribeHeaders {
			proj.AddColumnShort(string(f.Name), f.Type)
		}
	}
	//u.Warnf("columns?  %#v   \nvals?%v", proj, tbl.FieldMap)
	//tbl.FieldsAsMessages()
	ss := datasource.NewStaticSource(tbl.Name, tbl.Columns(), FieldsAsMessages(tbl))
	//tableVals := membtree.NewStaticDataSource("describetable", 0, tbl.DescribeValues, schema.DescribeFullCols)
	return ss, proj
}

func ShowTables(ctx *Context) (*membtree.StaticDataSource, *rel.Projection) {
*/
/*
	mysql> show full tables from temp like '%';
	+--------------------+------------+
	| Tables_in_temp (%) | Table_type |
	+--------------------+------------+
	| emails             | BASE TABLE |
	| events             | BASE TABLE |
	| evtnames           | BASE TABLE |
	| username           | BASE TABLE |
	+--------------------+------------+
	5 rows in set (0.00 sec)

	mysql> show tables;
	+----------------+
	| Tables_in_temp |
	+----------------+
	| emails         |
	| events         |
	| evtnames       |
	| username       |
	+----------------+
	5 rows in set (0.00 sec)

	mysql> show tables from temp like '%';
	+--------------------+
	| Tables_in_temp (%) |
	+--------------------+
	| emails             |
	| events             |
	| evtnames           |
	| username           |
	+--------------------+
	5 rows in set (0.00 sec)

*/
/*
	tables := ctx.Schema.Tables()
	vals := make([][]driver.Value, len(tables))
	idx := 0
	if len(tables) == 0 {
		u.Warnf("NO TABLES!!!!! for %+v", ctx.Schema)
	}
	for _, tbl := range tables {
		vals[idx] = []driver.Value{tbl}
		//u.Infof("found table: %v   vals=%v", tbl, vals[idx])
		idx++
	}
	showTableVals := membtree.NewStaticDataSource("schematables", 0, vals, []string{"Table"})
	proj := rel.NewProjection()
	proj.AddColumnShort("Table", value.StringType)
	//u.Infof("showtables:  %v", m.showTableVals)
	return showTableVals, proj
}

func ShowVariables(name string, val driver.Value) (*membtree.StaticDataSource, *rel.Projection) {
	/*
	/*
	   MariaDB [(none)]> SHOW SESSION VARIABLES LIKE 'lower_case_table_names';
	   +------------------------+-------+
	   | Variable_name          | Value |
	   +------------------------+-------+
	   | lower_case_table_names | 0     |
	   +------------------------+-------+
*/
/*
	vals := make([][]driver.Value, 1)
	vals[0] = []driver.Value{name, val}
	dataSource := membtree.NewStaticDataSource("schematables", 0, vals, []string{"Variable_name", "Value"})
	p := rel.NewProjection()
	p.AddColumnShort("Variable_name", value.StringType)
	p.AddColumnShort("Value", value.StringType)
	return dataSource, p
}
*/