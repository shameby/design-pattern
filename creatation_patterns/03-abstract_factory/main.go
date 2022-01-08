package main

import "fmt"

func main() {
	f1, f2 := &RDBDaoFac{}, &XMLDaoFac{}
	getMainAndDetail(f1)
	getMainAndDetail(f2)
}

func getMainAndDetail(fac DaoFactory) {
	fac.CreateOrderMainDao().SaveOrderMain()
	fac.CreateOrderDetailDao().SaveOrderDetail()
}

type DaoFactory interface {
	CreateOrderMainDao() OrderMainDao
	CreateOrderDetailDao() OrderDetailDao
}

type OrderMainDao interface {
	SaveOrderMain()
}

type OrderDetailDao interface {
	SaveOrderDetail()
}

type RDBDaoFac struct{}

type RDBMainDao struct{}

func (this_ *RDBMainDao) SaveOrderMain() {
	fmt.Println("save main in RDB")
}

type RBDDetailDao struct{}

func (this_ *RBDDetailDao) SaveOrderDetail() {
	fmt.Println("save detail in RDB")
}

func (this_ RDBDaoFac) CreateOrderMainDao() OrderMainDao {
	return &RDBMainDao{}
}
func (this_ RDBDaoFac) CreateOrderDetailDao() OrderDetailDao {
	return &RBDDetailDao{}
}


type XMLDaoFac struct{}

type XMLMainDao struct{}

func (this_ *XMLMainDao) SaveOrderMain() {
	fmt.Println("save main in XML")
}

type XMLDetailDao struct{}

func (this_ *XMLDetailDao) SaveOrderDetail() {
	fmt.Println("save detail in XML")
}

func (this_ XMLDaoFac) CreateOrderMainDao() OrderMainDao {
	return &XMLMainDao{}
}
func (this_ XMLDaoFac) CreateOrderDetailDao() OrderDetailDao {
	return &XMLDetailDao{}
}

