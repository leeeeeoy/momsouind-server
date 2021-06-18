package models

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Age       string `json:"age"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// func (user *User) UpdateData() (string, []interface{}, error) {
// 	query := ""
// 	values := []interface{}{}

// 	if user.Name != "" {
// 		query += "name=?,"
// 		values = append(values, user.Name)
// 	}
// 	if user.Age != "" {
// 		query += "age=?,"
// 		values = append(values, user.Age)
// 	}
// 	if user.Password != "" {
// 		query += "password=?,"
// 		values = append(values, user.Password)
// 	}
// 	if user.BabyName != "" {
// 		query += "baby_name=?,"
// 		values = append(values, user.BabyName)
// 	}
// 	if user.BabyNickname != "" {
// 		query += "baby_nickname=?,"
// 		values = append(values, user.BabyNickname)
// 	}
// 	if user.BabyBirth != "" {
// 		query += "baby_birth=?,"
// 		values = append(values, user.BabyBirth)
// 	}
// 	if query == "" {
// 		return query, values, errors.New("업데이트 할 게 없어용")
// 	}

// 	query = query[:len(query)-1]

// 	return query, values, nil
// }
