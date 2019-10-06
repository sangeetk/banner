package datastore

const (
	// define DB related errors here
	ErrorCannotConnect = "Cannot connect"
)

type Mysql struct {
	// Define mysql realted fields
}

func NewMysqlStorage(host, user, pass string) *Mysql {
	// connect to MySQL server and return handler
	return nil
}

func (m *Mysql) Put(key string, data interface{}) {
	// insert or update into database
}

func (m *Mysql) Get(key string) (interface{}, error) {
	// read from database
	return nil, nil
}

func (m *Mysql) Del(key string) error {
	// delete from database
	return nil
}
