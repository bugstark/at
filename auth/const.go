package auth

import (
	"context"
	"log"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	adapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

var CasbinServer *casbin.Enforcer

const casbinmodel = `
[request_definition]
r = role, urlpath, method

[policy_definition]
p = role, urlpath, method

[policy_effect]
e = some(where (p.eft == allow))

[role_definition]
g = _, _

[matchers]
m = (g(r.role, p.role) || r.role=="root") && (keyMatch2(r.urlpath, p.urlpath) || keyMatch(r.urlpath, p.urlpath) || p.urlpath == "*") && (r.method == p.method || p.method == "*" || regexMatch(r.method, p.method)) 
`

func Init(db *gorm.DB) error {
	adapter, err := adapter.NewAdapterByDBUseTableName(db.WithContext(context.Background()), "sys", "auth")
	if err != nil {
		return err
	}
	casbinmodle, err := model.NewModelFromString(casbinmodel)
	if err != nil {
		log.Println("casbinmodle error:", err.Error())
		return err
	}
	enforcer, err := casbin.NewEnforcer(casbinmodle, adapter, true)
	if err != nil {
		log.Println("enforcer error:", err.Error())
		return err
	}
	enforcer.EnableAutoSave(true)
	CasbinServer = enforcer
	return nil
}
