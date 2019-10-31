package database

type SQLHandler interface {
	Get(interface{}, string, ...interface{}) error
	Create(interface{}) error
	Update(interface{}, interface{}, string, ...interface{}) error
	Delete(interface{}, string, ...interface{}) error
}
