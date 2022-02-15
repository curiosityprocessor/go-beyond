package main

import "fmt"

type User struct {
	Name string
	Id string
	Email string
}

type VipUser struct {
	UserInfo User
	VipLevel int
	Discount float32
}

type AdminUser struct {
	User //omit field name -> nested struct as included fields
	AdminLevel int
}

func NestedStruct() {
	user := User {"cp", "Cp", "cp@grr.la",}
	vipUser := VipUser {
		User {"John Doe", "jd", "jd@grr.la"},
		1,
		22,
	}
	fmt.Printf("\nuser: name: %s, id: %s, email: %s", user.Name, user.Id, user.Email)
	fmt.Printf("\nVIP: name: %s, id: %s, email: %s, vip level: %d, discount: %f", 
		vipUser.UserInfo.Name,
		vipUser.UserInfo.Id,
		vipUser.UserInfo.Email,
		vipUser.VipLevel,
		vipUser.Discount,
	)

	adminUser := AdminUser {
		User {"admin", "Administrator", "admin@grr.la"},
		3,
	}

	//nameless nested structs can be accessed with single . operator
	fmt.Printf("\nAdmin: name: %s, id: %s, email: %s, admin level: %d", 
		adminUser.Name,
		adminUser.Id,
		adminUser.Email,
		adminUser.AdminLevel,
	)


}