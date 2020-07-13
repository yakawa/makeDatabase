# AST YAML 解説

anywhereQL の AST を示すYAMLの解説

## SQL Section

```
SQL
|- [Statment]...
```

## Statement Section

```
Statement
|- ALTERStatement
|- CREATEStatement
|- DELETEStatement
|- DROPStatement
|- INSERTStatement
|- SELECTStatement
|- UPDATEStatement
|- VACUUMStatement
```

* 属性
| 属性 | 説明 |
| --- | --- |
| IsExplain | クエリプラン |


## SELECTStatement Section
```
SELECTStatement
 |- WITH
 |- SELECTClause
 |- ORDERBY
 |- LIMIT
```

## SELECTClause Section

```
SELECTClause
|- SELECT
|- VALUES
```
* 属性
| 属性 | 説明 |
| --- | --- |
| CompoundOperator | UNION / UNIONALL / INTERSECT / EXCEPT |

## SELECT Section
```
SELECT
 |- [ResultColumns]...
 |- FROM
 |- WHERE
 |- GROUPBY
 |- WINDOW
```

## VALUES Section
```
VALUES
|- [EXPRESSION]...
```

#### FROM Section
#### WHERE Section
#### GROUPBY Section
#### WINDOW Section


``` yaml
- SQL
 SELECTStatement:
  SELECT:
   ResultColumns:
    - Expression:
       Column: c1
    - Expression:
       Column: c2
    - Expression:
       Function:
        Name: SUM
        ARGS:
         - arg:
            Column: c1
  FROM:
   - ToS:
      Table: t1
  GROUPBy:
   Grouping:
    - Expression:
       Column: c1
    - Expression:
       Column: c2

```
