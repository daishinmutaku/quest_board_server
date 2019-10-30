package database

type SQLHandler interface {
	Get(interface{}, string, ...interface{}) error
	Create(interface{}) error
	Update(interface{}, int, interface{}) error
}
