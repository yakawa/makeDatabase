## MySQL 8.0
### Query
```
<query> ::=
  [<with_clause>]
  <select_statement>
  [UNION [ALL | DISTINCT] <select_statement>]...
```

#### Definition
```
<alias_name> ::= <IDENTIFIER>
<bit_expr> ::=
    <bit_expr> | <bit_expr>
  | <bit_expr> & <bit_expr>
  | <bit_expr> << <bit_expr>
  | <bit_expr> >> <bit_expr>
  | <bit_expr> + <bit_expr>
  | <bit_expr> - <bit_expr>
  | <bit_expr> * <bit_expr>
  | <bit_expr> / <bit_expr>
  | <bit_expr> DIV <bit_expr>
  | <bit_expr> MOD <bit_expr>
  | <bit_expr> % <bit_expr>
  | <bit_expr> ^ <bit_expr>
  | <bit_expr> + <interval_expr>
  | <bit_expr> - <interval_expr>
  | <simple_expr>
<boolean_primary> ::=
    <boolean_primary> IS [NOT] NULL
  | <boolean_primary> <comparison_operator> <predicate>
  | <boolean_primary> <comparison_operator> {ALL | ANY} (<subquery>)
  | <predicate>
<case_expr> ::=
  | IF(<expr>, <expr>, <expr>)
  | IFNULL(<expr>, <expr>)
  | NULL(<expr>, <expr>)
  | {CASE [<case_value>] WHEN <when_value> THEN <statement_list> [WHEN <when_value> THEN <statement_list>]... [ELSE <statement_list>] END CASE}
  | {CASE WHEN <search_condition> THEN <statement_list> [WHEN <search_condition> THEN <statement_list>]... [ELSE <statement_list>] END CASE}
<charset_name> ::= <IDENTIFIER>
<col_name> ::= [<tbl_name>.]<IDENTIFIER>
<collation_name> ::= <IDENTIFIER>
<comparison_operator> ::= = | >= | > | <= | < | <> | != | <=>
<database_name> ::= [<schema_name>.]<IDENTIFIER>
<escaped_tbl_reference> ::=
    <tbl_reference>
  | OJ <tbl_reference>
<export_options> ::= [FIELDS TERMINATED BY '<STRING>'][[OPTIONALY] ENCLOSED BY '<CHAR>'][ESCAPED BY '<CHAR>'][LINES STARTED BY '<STRING>'][LINES TERMINATED BY '<STRING>']
<expr> ::=
    <expr> OR <expr>
  | <expr> || <expr>
  | <expr> XOP <expr>
  | <expr> AND <expr>
  | <expr> && <expr>
  | NOT <expr>
  | ! <expr>
  | <boolean_primary> IS [NOT] {TRUE | FALSE | UNKNOWN}
  | <boolean_primary>
  | CONVERT(<expr> USING <transcoding_name>) [COLLATE <collation_name>]
  | CONVERT(<expr>, <type> [CHARACTER SET <charset_name>]) [COLLATE <collation_name>]
  | CAST(<expr> AS <type> [CHARACTER SET <chaset_name>]) [COLLATE <collation_name>]
<function_call> ::= <IDENTIFIER> ([<expr> [, <expr>]...])
<index_hint> ::=
    USE {INDEX | KEY} [FOR {JOIN | {ORDER BY} | {GROUP BY}}] ([<index_list>])
  | {IGNORE | FORCE}[FOR {JOIN | {ORDER BY} | {GROUP BY}}] (<index_list>)
<index_hint_list> ::= <index_hint>[, <index_hint>]...
<index_list> ::= <index_name>[, <index_name>]...
<index_name> ::= <IDENTIFIER>
<interval_expr> ::= INTERVAL <expr> <unit>
<into_option> ::= {INTO OUTFILE '<outfile_name>' [CHARACTER SET <charset_name>] <export_options>} | {INTO DUMPFILE '<outfile_name>'} | {INTO <var_name> [, <var_name>]...}
<join_col_list> ::= <col_name>[, <col_name>]...
<join_spec> ::=
    ON <search_condition>
  | USING (<join_col_list>)
<joined_table> ::=
    <tbl_reference> {{[INNER | CROSS] JOIN} | STRAIGHT_JOIN} table_factor [<join_spec>]
  | <tbl_reference> {LEFT | RIGHT} [OUTER] JOIN <tbl_reference> <join_spec>
  | <tbl_reference> NATURAL [INNER | {{LEFT | RIGHT} [OUTER]}] JOIN <tbl_factor>
<literal> ::=
    [_char_name]'<STRING>' [COLLATE <collation_name>]
  | <NUMERIC>
  | DATE '<STRING>'
  | TIME '<STRING>'
  | TIMESTAMP '<STRING>'
  | {X|x}'<0-9A-Fa-f>...'
  | 0x<0-9A-Fa-f>...
  | {B|b}'<01>...'
  | 0b<01>...
  | TRUE # case insensitive
  | FALSE # case insensitive
  | NULL # case insensitive
<match_expr> ::= MATCH (<col_name> [, <col_name>]...) AGAINST (<expr> [<search_modifier>])
<offset> ::= <NUMBER>
<outfile_name> := <IDENTIFIER>
<parameter_marker> ::= ?
<partition_list> ::= <IDENTIFIER> [, <IDENTIFIER>]...
<position> ::= <NUMBER>
<predicate> ::=
    <bit_expr> [NOT] IN (<subquery>)
  | <bit_expr> [NOT] IN (<expr> [, <expr>]...)
  | <bit_expr> [NOT] BETWEEN <bit_expr> AND <predicate>
  | <bit_expr> SOUNDS LIKE <bit_expr>
  | <bit_expr> [NOT] LIKE <simple_expr> [ESCAPE <simple_expr>]
  | <bit_expr> [NOT] REGEXP <bit_expr>
  | <bit_expr>
<row_constractor_list> ::= ROW(<value_list>)[, ROW(<value_list>)]...
<row_count> ::= <NUMBER>
<schema_name> ::= <IDENTIFIER>
<search_modifier> ::=
    IN NATURAL LANGUAGE MODE
  | IN NATURAL LANGUAGE MODE WITH QUERY EXPANSION
  | IN BOOLEAN MODE
  | WITH QUERY EXPANSION
<select_clause> ::=
  SELECT
   [ALL | DISTINCT | DISTINCTROW]
   [HIGH_PRIORITY]
   [STRAIGHT_JOIN]
   [SQL_SMALL_RESULT] [SQL_BIG_RESULT] [SQL_BUFFER_RESULT]
   [SQL_NO_CACHE] [SQL_CALC_FOUND_ROWS]
   <select_expr> [, <select_expr>]...
   <into_option>
   [FROM <tbl_references> [PARTITION <partition_list>]]
   [WHERE <where_condition>]
   [GROUP BY {<col_name> | <expr> | <position>}[, {<col_name> | <expr> | <position>}]... [WITH ROLLUP]]
   [HAVING <where_condition>]
   [WINDOW <window_name> AS (<window_spec>) [, <window_name> AS (<window_spec>)]...]
   [ORDER BY {<col_name> | <expr> | <position>} [ASC | DESC][, {<col_name> | <expr> | <position>} [ASC | DESC]]...[WITH ROLLUP] ]]
   [LIMIT {{[<offset>,] <row_count>} | {<row_count> OFFSET <offset>}}]
   [<into_option>]
   [FOR {UPDATE | SHARE} [OF <tbl_name>[, <tbl_name>]...] [NOWAIT | {SKIP LOCKED}] | [LOCK IN SHARE MODE]]
   [<into_option>]
<select_expr> ::= <expr> [[AS] <alias_name>]
<select_statement> ::=
    <select_clause>
  | <values_clause>
<simple_expr> ::=
    <literal>
  | <identifier>
  | <function_call>
  | <simple_expr> COLLATE <collation_name>
  | <param_marker>
  | <variable>
  | <simple_expr> || <simple_expr>
  | + <simple_expr>
  | - <simple_expr>
  | ~ <simple_expr>
  | ! <simple_expr>
  | BINARY <simple_expr>
  | (<expr>[, <expr>]...)
  | ROW (<expr>, <expr>[, <expr>]...)
  | (<subquery>)
  | EXISTS (<subquery>)
  | {<identifier> <expr>}
  | <match_expr>
  | <case_expr>
  | <interval_expr>
<subquery> ::=
    <select_statement>
  | TABLE <table_name>
<tbl_factor> ::=
    <tbl_name> [PARTITION (partition_list)] [[AS] <alias_name>][<index_hint_list>]
  | [LATERAL] <tbl_subquery> [AS] <alias_name> [(<col_list>)]
  | (<tbl_references>)
<tbl_name> ::= [<database_name> .]<IDENTIFIER>
<tbl_references> ::= <escaped_tbl_reference>[, <escaped_tbl_reference>]...
<tbl_reference> ::=
    <tbl_factor>
  | <joined_tbl>
<transcoding_name> ::= <IDENTIFIER>
<type> ::=
<unit> ::=
    MICROSECOND
  | SECOND
  | MINUTE
  | HOUR
  | DAY
  | WEEK
  | MONTH
  | QUATER
  | YEAR
  | SECOND_MICROSECOND
  | MINUTE_MICROSECOND
  | MINUTE_SECOND
  | HOUR_MICROSECOND
  | HOUR_SECOND
  | HOUR_MINUTE
  | DAY_MICROSECOND
  | DAY_SECOND
  | DAY_MINUTE
  | DAY_HOUR
  | YEAR_MONTH
<values_clause> ::= VALUES <row_constructor_list> [ORDER BY <column_designator>][LIMIT BY <NUMBER>]
<value_list> ::= <expr> [, <expr>]...
<variable> ::= @[@]<IDENTIFIER>
<where_condition> ::= <expr>
<window_name> ::= <IDENTIFIER>
<window_spec> ::=
<with_clause> ::= WITH [RECURSIVE] <cte_name> [(<col_list>)] AS (<subquery>) [, <cte_name> [(<col_list>)] AS (<subquery>)]...
```


#### KEYWORD
- AGAINST
- ALL
- AND
- ANY
- AS
- ASC
- B
- BETWEEN
- BINARY
- BOOLEAN
- BY
- CASE
- CAST
- CHARACTER
- COLLATE
- CONVERT
- CROSS
- DATE
- DAY
- DAY_HOUR
- DAY_MICROSECOND
- DAY_MINUTE
- DAY_SECOND
- DESC
- DISTINCT
- DISTINCTROW
- DIV
- DUMPFILE
- ELSE
- ENCLOSED
- END
- ESCAPED
- EXISTS
- EXPANSION
- FALSE
- FIELDS
- FOR
- FORCE
- FROM
- GROUP
- HAVING
- HIGH_PRIORITY
- HOUR
- HOUR_MICROSECOND
- HOUR_MINUTE
- HOUR_SECOND
- IF
- IFNULL
- IGNORE
- IN
- INDEX
- INNER
- INTERVAL
- INTO
- IS
- ISNULL
- JOIN
- KEY
- LANGUAGE
- LATERAL
- LEFT
- LIKE
- LINES
- LOCK
- LOCKED
- MATCH
- MICROSECOND
- MINUTE
- MINUTE_MICROSECOND
- MINUTE_SECOND
- MODE
- MONTH
- NATURAL
- NOT
- NOWAIT
- NULL
- OF
- OJ
- ON
- OPTIONALY
- OR
- ORDER
- OUTER
- OUTFILE
- PARTITION
- QUARTER
- QUERY
- RECURSIVE
- REGEXP
- RIGHT
- ROLLUP
- SECOND
- SECOND_MICROSECOND
- SELECT
- SET
- SHARE
- SKIP
- SOUNDS
- SQL_BIG_RESULT
- SQL_BUFFER_RESULT
- SQL_CALC_FOUND_ROWS
- SQL_NO_CACHE
- SQL_SMALL_RESULT
- STARTED
- STRAIGHT_JOIN
- TABLE
- TERMINATED
- THEN
- TIME
- TIMESTAMP
- TRUE
- UNION
- UNKNOWN
- UPDATE
- USE
- USING
- VALUES
- WEEK
- WHEN
- WHERE
- WINDOW
- WITH
- X
- XOR
- YEAR
- YEAR_MONTH

#### SYMBOL
```
| & << >> + - * / % ^ ( , ) . = >= > <= < <> != <=> || && ! ' ? ~
```
