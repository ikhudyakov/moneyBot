package repository

import (
	"database/sql"
	m "moneybot/internal/model"
)

type ControlPosgres struct {
	DB *sql.DB
}

func NewControlPostgres(db *sql.DB) *ControlPosgres {
	return &ControlPosgres{DB: db}
}

func (c *ControlPosgres) GetUser(myMoneyId int64) (*m.User, error) {
	var id int
	rows, err := c.DB.Query("select ID from USERS where id_telegram=$1", myMoneyId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	if rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, nil
	}

	return &m.User{Id: id}, err
}

// // Получение списка всех категорий из базы данных
// func category() string {
// 	var result, name string
// 	var id int64
// 	c1 := make(chan string)

// 	go func() {
// 		dbConnect()
// 		defer db.Close()
// 		rows, err := db.Query("SELECT ID, NAME FROM CATEGORY WHERE type=$1;", 1)
// 		result += "Категории Расходов ➖💲\n"
// 		CheckError(err)
// 		defer rows.Close()

// 		for rows.Next() {
// 			err := rows.Scan(&id, &name)
// 			CheckError(err)
// 			result += fmt.Sprintf("id=%d. %s\n", id, name)
// 		}

// 		rows, err = db.Query("SELECT ID, NAME FROM CATEGORY WHERE type=$1;", 2)
// 		result += "Категории Доходов ➕💲\n"
// 		CheckError(err)
// 		defer rows.Close()

// 		for rows.Next() {
// 			err := rows.Scan(&id, &name)
// 			CheckError(err)
// 			result += fmt.Sprintf("id=%d. %s\n", id, name)
// 		}
// 		c1 <- result
// 	}()
// 	return <-c1
// }

// // Добавление записи
// func add(in []string, tableName string, userId int) int {
// 	if len(in) < 4 {
// 		return 0
// 	}

// 	category_type, err := strconv.ParseInt(in[1], 10, 64)
// 	if err != nil {
// 		CheckError(err)
// 		return 0
// 	}

// 	amount, err := strconv.ParseFloat(in[2], 64)
// 	if err != nil {
// 		CheckError(err)
// 		return 0
// 	}

// 	description := in[3]

// 	fmt.Printf("%d %f %s\n%d\n", category_type, amount, description, userId)

// 	c1 := make(chan int)
// 	go func() {
// 		dbConnect()
// 		defer db.Close()
// 		var lastID int
// 		err = db.QueryRow(
// 			"INSERT INTO "+tableName+" (amount, date, category_id, user_id, description) VALUES ($1, $2, $3, $4, $5) RETURNING id",
// 			amount,
// 			time.Now(),
// 			category_type,
// 			userId,
// 			description).Scan(&lastID)
// 		c1 <- lastID
// 		CheckError(err)
// 		fmt.Println(lastID)
// 	}()
// 	return <-c1
// }

// // Получение id пользователя
// func getUser(myMoneyId int) int {
// 	c1 := make(chan int)
// 	var userId int
// 	go func() {
// 		dbConnect()
// 		defer db.Close()
// 		rows, err := db.Query("select ID from USERS where id_telegram=$1", myMoneyId)
// 		CheckError(err)
// 		defer rows.Close()
// 		for rows.Next() {
// 			err := rows.Scan(&userId)
// 			CheckError(err)
// 		}
// 		c1 <- userId
// 	}()
// 	return <-c1
// }

// func CheckError(err error) {
// 	if err != nil {
// 		log.Printf(err.Error())
// 	}
// }
