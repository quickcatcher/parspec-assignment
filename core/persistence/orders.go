package persistence

import (
	"parspec-assignment/core/domain"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(domain.Orders))
}

type OrderModel struct{}

func (a *OrderModel) AddOrder(order *domain.Orders) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(order)
	return
}
