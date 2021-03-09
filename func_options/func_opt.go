package func_options

import (
	"fmt"
)

type User1 struct {
	// 必要且必填
	ID 		int
	Name	string

	// 不必要不必填
	TelNum  string
	Age 	int

	//必要但非必填（有默认值
	Status 	string
	Sign	string
}


// go 不支持在结构里定义默认值以及重载函数，要创建多种类型的话就得写多个函数
func NewUser1(name, telNum, sign, status string, id,age int) *User1 {
	return &User1{
		ID:     id,
		Name:   name,
		TelNum: telNum,
		Age:    age,
		Status: status,
		Sign: sign,
	}
}

func NewDefaultUser1(id int, name string) *User1 {
	return &User1{
		ID:        id,
		Name:      name,
		TelNum:    "",
		Age:       0,
		Status:    "DEFAULT STATUS",
		Sign:      "DEFAULT SIGN",
	}
}

func Solution1() {
	user1 := NewDefaultUser1(1, "user1_name")
	user1.Sign = "USER1'S SIGN"
}


// Solution2

type User2 struct {
	// 必要且必填
	ID 		int
	Name	string

	// 不必填的
	Opt 	User2Opt
}

type User2Opt struct {
	// 不必要不必填
	TelNum  string
	Age 	int

	//必要但非必填（有默认值
	Status 	string
	Sign	string
}
func NewDefaultOpt() *User2Opt {
	return &User2Opt{
		TelNum: "",
		Age:    0,
		Status: "DEFAULT STATUS",
		Sign:   "DEFAULT SIGN",
	}
}
func NewUser2(id int, name string, opt *User2Opt) *User2 {
	if opt == nil {
		opt := NewDefaultOpt()
		return &User2{
			ID:   id,
			Name: name,
			Opt: *opt,
		}
	}

	// 里边判断再对应
	return &User2{
		ID:   id,
		Name: name,
		Opt: User2Opt{

		},
	}
}

func Solution2() {
	opt := &User2Opt{
		TelNum: "tel",
		Status: "new status string",
	}
	var _ = NewUser2(1, "user2_name", opt)
}



// Solution3: Functional Options(用闭包来初始化
type User3 struct {
	// 必要且必填
	ID 		int
	Name	string

	// 不必要不必填
	TelNum  string
	Age 	int

	//必要但非必填（有默认值
	Status 	string
	Sign	string
}

type Option func(*User3)

func NewUser3(id int, name string, opts ...Option) *User3 {
	user3 := User3{
		ID:     id,
		Name:   name,
		TelNum: "",
		Age:    0,
		Status: "DEFAULT STATUS",
		Sign:   "DEFAULT SIGN",
	}
	for _, opt := range opts {
		opt(&user3)
	}
	return &user3
}

func Status(status string) Option {
	return func(user3 *User3) {
		user3.Status = status
	}
}

func Sign(sign string) Option {
	return func(user3 *User3) {
		user3.Sign = sign
	}
}

func Solution3() {
	user3 := NewUser3(1, "user3_name", Status("USER3'S STATUS"))
	fmt.Println(user3.Status)
}