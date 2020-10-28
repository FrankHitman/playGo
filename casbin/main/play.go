package main

import (
	"fmt"

	"github.com/casbin/casbin"
	"siemens.com/wallbox/pkg/fastlog"
)

func main() {

	conf := "casbin/rbac_model.conf"
	csv := "casbin/rbac_policy.csv"
	e, err := casbin.NewEnforcer(conf, csv)
	if err != nil {
		fastlog.Errorln("Init casbin error: ", err)
	}
	res, err := e.Enforce("admin", "/api/v1.0/charger", "POST")
	fmt.Println(res)

	res, err = e.Enforce("zhangsan", "/api/v1.0/login", "POST")
	fmt.Println(res)

	res, err = e.Enforce("member_read", "/login", "POST")
	fmt.Println(res)

	res, err = e.Enforce("member_write", "/dddd", "POST")
	fmt.Println("member write is ", res)

	res, err = e.Enforce("member_write", "/dddd", "GET")
	fmt.Println("member write is ", res)

	res, err = e.Enforce("zhaoliu", "/charger/1", "POST")
	fmt.Println(res)

	allRoles := e.GetAllRoles()
	fmt.Println("allRoles is ", allRoles)
	allNamedRoles := e.GetAllNamedRoles("g")
	fmt.Println(allNamedRoles)

	fmt.Println("=====")
	policy := e.GetPolicy()
	fmt.Println(policy)
	// for _, v := range policy {
	// 	fmt.Println(v[0])
	// 	fmt.Printf("%t\n", v[0])
	// }

	filteredPolicy := e.GetFilteredPolicy(0, "admin")
	fmt.Println(filteredPolicy)

	namedPolicy := e.GetNamedPolicy("p")
	fmt.Println(namedPolicy)

	filteredNamedPolicy := e.GetFilteredNamedPolicy("p", 0, "admin")
	fmt.Println(filteredNamedPolicy)

	groupingPolicy := e.GetGroupingPolicy()
	fmt.Println(groupingPolicy)

	filteredGroupingPolicy := e.GetFilteredGroupingPolicy(0, "member_read")
	fmt.Println(filteredGroupingPolicy)

	hasPolicy := e.HasPolicy("admin", "/*", "^*$")
	fmt.Println(hasPolicy)

	role, err := e.GetRolesForUser("admin")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("role is ", role[0])

}


// -----output------
// true
// false
// false
// member write is  true
// member write is  false
// false
// allRoles is  [admin]
// [admin]
// =====
// [[admin /* ^*$] [anonymous */login POST] [member_read /* GET] [member_write /* (POST)|(PUT)|(DELETE)] [installer1 /charger/1 ^*$]]
// [[admin /* ^*$]]
// [[admin /* ^*$] [anonymous */login POST] [member_read /* GET] [member_write /* (POST)|(PUT)|(DELETE)] [installer1 /charger/1 ^*$]]
// [[admin /* ^*$]]
// [[admin admin]]
// []
// true
// role is  admin