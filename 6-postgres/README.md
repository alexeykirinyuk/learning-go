# PostgreSQL

<!-- TOC -->
- [PostgreSQL](#postgresql)
  - [Содеражние](#содеражние)
  - [PostgreSQL](#postgresql-1)
  - [Индексы](#индексы)
    - [Примеры создания индексов](#примеры-создания-индексов)
    - [Создание иднекса без блокировки записи в таблицу (postgres@>=8.2)](#создание-иднекса-без-блокировки-записи-в-таблицу-postgres82)
    - [Hash Index](#hash-index)
  - [Транзакционность](#транзакционность)
    - [Уровни изоляции](#уровни-изоляции)
      - [Аномалии](#аномалии)
      - [Уровни](#уровни)
  - [Драйверы в GO для подключения к SQL](#драйверы-в-go-для-подключения-к-sql)
    - [lib/pq](#libpq)
    - [go-pg/pg](#go-pgpg)
    - [pgx - PostgreSQL Driver and Toolkit](#pgx---postgresql-driver-and-toolkit)
  - [Подключение к Postgres](#подключение-к-postgres)
    - [pgx - GO](#pgx---go)
  - [Пул соединений](#пул-соединений)
    - [Настроки пула:](#настроки-пула)
  - [Миграции](#миграции)
  - [Фикстуры](#фикстуры)
  - [Выполнение запросов](#выполнение-запросов)
    - [Объект `sql.Rows`](#объект-sqlrows)
    - [Prepared Statements](#prepared-statements)
    - [Работа с соединениями](#работа-с-соединениями)
    - [Транзакции](#транзакции)
    - [NULL](#null)
    - [SQL Injection](#sql-injection)
    - [Плейсхолдеры](#плейсхолдеры)
  - [Проблемы database/sql](#проблемы-databasesql)
<!-- TOC -->

## Содеражние

- Общие сведения о postgres: индексы, транзакционность
- Подключение к PostgreSQL: драйверы, DSN, пул соединений
- Миграции
- Фикстуры
- Выполнение запросов
- NULL значения
- SQL инъекции
- Библиотека `jmoiron/sqlx` (расширение database/sql)
- Библиотека `Misterminds/squirrel` (query builder)

## PostgreSQL

PostgreSQL или Postgres - свободная, объектно-реляционная система управления базами данных (СУБД).

| Максимальный размер базы данных | Нет ограничений                        |
|---------------------------------|----------------------------------------|
| Максимальный размер таблицы     | 32ГБ                                   |
| Максимальный размер поля        | 1ГБ                                    |
| Максимальнум записей в таблице  | Ограничено размерами таблицы           |
| Максимальнум полей в записи     | 250-1600, в зависимости от типов полей |
| Максимум индексов в таблице     | Нет ограничений                        |

## Индексы
- B-tree - позволяет выбирать range, итерироваться вперед-назад, проверять nullable
- Hash - в новых вериях postgres-а работает лучше
- GiST
- SP-GiST
- GIN
- BRIN

По умолчанию команда `CREATE INDEX` создает индексы типа B-tree, эффективный в большинстве случаев.

### Примеры создания индексов
- Создание индекса (B-дерево) по столбцу title в таблице films
```postgresql
CREATE UNIQUE INDEX title_idx ON films(title);
```
- Создание индекса по выражению lower(title) для эффективного регистронезависимого поиска
```postgresql
CREATE INDEX ON films(lower(title));
```
(Имя индекса выберет система - films_lower_idx)

### Создание иднекса без блокировки записи в таблицу (postgres@>=8.2)
```postgresql
CREATE INDEX CONCURRENTLY my_table_index ON sales_table(quantity);
```
- без блокировки добавления/удаляения/изменения записей в таблице
- можно без проблем работать с базой пока индекс строится
- приведет к дополнительному сканированию таблицы
- выполняется значительно дольше чем обычное построение индекса
- нагружает систему, это может аффектить другие операции
```postgresql
REINDEX INDEX CONCURRENTLY my_table_index; -- C PG 12+
```
```postgresql
DROP INDEX CONCURRENTLY my_table_index;
```

### Hash Index
32-битный хеш-код.
Уникальные значения, которые нужно сравнивать по равенству.
Занимает меньше места чем b-tree. Алгоритмическая сложность проще

| Индекс | Сложность |
|--------|-----------|
| B-tree | O(log n)  |
| Hash   | O(1)      |

```postgresql
CREATE INDEX my_hash_idx ON my_table USING HASH (my_column);
```

## Транзакционность
Multiversion concurrency control, MVCC
Каждая транзакция работает со своим "снепшотом" данных

### Уровни изоляции
#### Аномалии
- dirty read - транзакция читает данные записанные незавершенной другой транзакцией
- nonrepeatable read - транзакция повторно читает те же данные, что и раньше, и обнаруживает что они были изменены другой транзакцией (которая завершилась после первого чтения)
- phantom read - изменилась выборка данных
- serialization anomaly

#### Уровни
- read uncommited - грязное чтение
- read commited - fix грязное чтение
- repeatable read - fix repeatable read
- serializable - fix serialization anomaly

## Драйверы в GO для подключения к SQL
```go
import "database/sql"
//...
sql.Open(driver, url string) (*DB, error)
db.ExecContext(...) (Result, error)
db.QueryRowContext(...) *Row
db.QueryContext(..) (*Rows, error)
// ... это лишь часть методов
```
- реализует базовый индерфейс для SQL-совместимых СУБД
- для работы нужен драйвер
- не привязан к какой либо конкретной СУБД

### lib/pq
```shell
go get github.com/lib/pq
```
maintenance mode

### go-pg/pg
```shell
go get github.com/go-pg/pg
```
ORM и Клиент для Postgres  
**maintenance mode**

### pgx - PostgreSQL Driver and Toolkit
```shell
go get github.com/jackc/pgx/v4
```
- работает только с Postgres
- популярная и активно развивается
- не использует cgo (pure go)
- поддержка фичей, специфичных для Postgres
- можно использовать свой интерфейс, отличный от database/sql

Интерфейс pgx можно выбрать если:
- используется только Postgres
- никакая другая библиотека в приложении не использует database/sql

## Подключение к Postgres
Open Database Connectivity (ODBC) - открытый стандарт API для доступа к БД.
DSN (data source name) - структура данных, содержащая информацию о конкретной БД, к которой ODBC драйвер может подключиться.

Пример DSN:
```
postgres://user:password@localhost:5000/db
```

или
```
user=user password=password host=host port=5432 database=db sslmode=disable
```

### pgx - GO

```go
package main

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v4"
)

func main() {
	// Open лишь валидирует аргументы, но не создает подключения
	db, err := sql.Open("pgx", "postgres://...")
	err = db.PingContext(context.Background())
}
```

## Пул соединений
`sql.DB` - это пул соединений с базой данных. Соединения будут открываться по мере необходимости.  
`sql.DB` - безопасен для конкурентного использования (также как http.Client).  

### Настроки пула:
```go
// Максимальное число открытых соединений от этого процесса
db.SetMaxOpenConns(n int)

// Максимальное число открытых неиспользуемых соединений
db.SetMaxIdleCons(n int)

// Максимальное время жизни одного подключения
db.SetConnMaxLifetime(d time.Duration)
```

## Миграции
Для миграции используется goose
```shell
go install github.com/pressly/goose/v3/cmd/goose
```

## Фикстуры
Для фикстур используем [testfixtures](https://github.com/go-testfixtures/testfixtures)
```shell
go get github.com/go-testfixtures/testfixtures/v3
```
```go
fixtures, err := testfixtures.New(
    testfixtures.Database(db),
	testfixtures.Dialect("postgres"),
	testfixtures.Paths(
	    "fixtures/products.yml"
	)
)
```
- можно использовать как CLI утилиту
- можно подключить как библиотеку
- есть встроенный шаблонизатор

Не используйте для тестов отличную СУБД от той, что у вас в prod!

## Выполнение запросов
### Объект `sql.Rows`

```go
// возвращает имена колонок в выборке
rows.Columns() ([]string, error)

// возвращает типы колонок в выборке
rows.ColumnTypes() ([]*ColumnType, error)

// переходит к следующей строке или возвращает false
rows.Next() bool

// заполняет переменные из текущей строки
rows.Scan(dest ...interface{}) error

// закрывает объект Rows
rows.Close()

// возвращает ошибку, встреченную при итерации
rows.Err() error
```

### Prepared Statements
PreparedStatement - это заранее разработанный запрос, который можно выполнять повторно.  
PreparedStatement - это временный объект, который создается в СУБД и живет в рамках сессии, или пока не будет закрыт

```go
// создаем подготовленный запрос
stmt, err := db.Prepare("delete from events where id = $1") // *sql.Stmt
if err != nil {
	log.Fatal(err)
}

// освобождаем ресурсы СУБД
defer stmt.Close()

// многократно выполняем запрос
for _, id := range ids {
	_, err = stmt.Exec(id)
	if err != nil {
        log.Fatal(err)
	}
}
```

### Работа с соединениями
`*sql.DB` - это пул соединений. Даже последовательные запросы могут использовать разные соединения с базой.  
Если нужно получить одно конкретное соединение, то
```go
conn, err := db.Conn(ctx) // *sql.Conn
// вернуть соединение в pool
defer conn.Close()

// далее обычная работа как с *sql.DB
err := conn.ExecContext(ctx, query1, arg1, arg2)
rows, err := conn.QueryContext(ctx, query2, arg1, arg2)
```

### Транзакции
Транзакция - группа запросов, которая либо выполняются, либо не выполняются вместе.
Внутри транзакция все запросы видят "согласованное" состояние.  
На уровне SQL для транзакций используются отдельные запросы: `BEGIN`, `COMMIT` и  `ROLLBACK`
```go
tx, err := db.BeginTx(ctx, nil) // *sql.Tx
if err != nil {
	log.Fatal(err)
}

// далее обычная работа как с *sql.DB
err := tx.ExecContext(ctx, query1, arg1, arg2)
rows, err := tx.QueryContext(ctx, query2, arg1, arg2)
err := tx.Commit() // или tx.Rollback()
if err != nil {
	// commit не прошел, данные не изменились
}
```

### NULL
В SQL базах любая колонка может быть объявлена как NULL / NOT NULL.
NULL - это не 0 и не пустая строка, это отсутствие значения.
```postgresql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    age INT NULL
)
```

Для обработки NULL в Go предлагается использовать специальные типы:
```go
var id, realAge int64
var name string
var age sql.NullInt64
err := db.
	QueryRowContext(ctx, "SELECT * FROM users WHERE id = 1").
	Scan(&id, &name, &age)

if age.Valid {
	realAge = age.Int64
} else {
	// обработка на ваше усмотрение
}
```

### SQL Injection
Опасно:
```go
query := "select * from users where name = '" + name + "'"
query := fmt.Sprintf("select * from users where name = '%s'", name)
```
Потому что в name может оказаться что-то вроде:
```postgresql
"jack'; truncate users; select 'pawned"
```
Итого
```postgresql
select * from users where name = 'jack'; truncate users; select 'pawned'
```

### Плейсхолдеры
Правильный подход - использовать placeholders для подстановки значений в SQL:
```go
row := db.QueryRowContext(ctx, "select * from users where name = $1;", name)
```
Однако, это не всегда возможно, это работать не будет:
```go
row1 := db.QueryRowContext(ctx, "select * from $1 where name = $2", table, name)
row2 := db.QueryRowContext(ctx, "select * from users order by $1 limit 3", column)
```

Проверить код на инъекции (и другие проблемы безопасности):
[securego/gosec](https://github.com/securego/gosec)

## Проблемы database/sql
- placeholder зависит от базы: $1 в Postgres, ? в MySQL и :name в Oracle
- есть только базовые типы, но нет например `sql.NullDate`
- `rows.Scan(arg1, arg2, arg3)` неудобен - легко ошибиться
- нет возможности `rows.StructScan(&event)`

