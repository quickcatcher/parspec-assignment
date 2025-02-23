package persistence

import (
	"parspec-assignment/core/domain"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(domain.Orders))
}

type OrderModel struct{}

func (a *OrderModel) AddOrder(order *domain.Orders) (orderid int, err error) {
	o := orm.NewOrm()
	id, err := o.Insert(order)
	orderid = int(id)
	return
}

func (a *OrderModel) GetOrderbyOrderId(orderId int) (order *domain.Orders, err error) {
	o := orm.NewOrm()
	order = &domain.Orders{}
	err = o.QueryTable(new(domain.Orders)).Filter("order_id", orderId).One(order)
	if err == orm.ErrNoRows {
		return nil, nil
	}
	return
}

func (a *OrderModel) UpdateOrderStatus(status string, processing_time float64, orderId int) (err error) {
	o := orm.NewOrm()
	query := `UPDATE orders SET status = ?, processing_time = ? WHERE order_id = ?`
	_, err = o.Raw(query, status, processing_time, orderId).Exec()
	return
}
