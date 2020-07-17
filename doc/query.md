# anywhereQL
## Query
```
[<with_clause]
[(]<select_clause>[)]
[<compund_ope> [(]<select_clause>[)]]...
[ORDER BY <orderby>[, <orderby>]...][LIMIT <expr> [{OFFSET|,} <expr>]]
;

```

### Definition
```
<with_clause> ::= WITH [RECURSIVE] <cte_name> [(<col_name_list>)] AS (<select_clause>) [, <cte_name> [(<col_name_list)] AS (<select_clause>)]...
<select_clause> ::=
   SELECT [DISTINCT | ALL] <result_col_list> [FROM <from_clause>] [WHERE <expr>] [GROUP BY <expr>[, <expr>]... [HAVING <expr>]] [WINDOW <window_name> AS <window_spec> [, <window_name> AS <window_spec>]...][ORDER BY <orderby>[, <orderby>]...][LIMIT <expr> [{OFFSET|,} <expr>]]
 | VALUES (<expr>[,<expr>]...) [, (<expr> [, <expr>]...)]...
<compound_ope> ::= UNION | {UNION ALL} | INTERSECT | EXCEPT | DISTINCT
<orderby> ::= <expr> [COLLATE <collate_name>] [ASC | DESC] [NULLS {FIRST | LAST}]
<cte_name> ::= <IDENTIFIER>
<col_name_list> ::= <single_col_name>[, <single_col_name>]...
<single_col_name> ::= <IDENTIFIER>
<result_col_list> ::= <result_col>[, <result_col>]...
<result_col> ::=
   <expr> [[AS] <alias_name>]
 | *
 | <table_name> . *
<alias_name> ::= <IDENTIFIER>
<window_name> ::= <IDENTIFIER>
<table_name> ::= [<database_name> .] <IDENTIFIER>
<database_name> ::= [<schema_name> .]<IDENTIFIER>
<schema_name> ::= <IDENTIFIER>
<from_Clause> ::= <tos> [<joined_table>]...
<tos> ::=
   <table_name>
```


### Expression (<expr>)
