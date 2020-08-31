package database

// Config 構造体
type Config struct {
	DB struct {
		Production struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
		Test struct {
			Host     string
			Username string
			Password string
			DBName   string
		}
	}
	Routing struct {
		Port string
	}
}

// NewConfig 初期化関数
func NewConfig() *Config {
	c := new(Config)

	c.DB.Production.Host = "localhost"
	c.DB.Production.Username = "username"
	c.DB.Production.Password = "password"
	c.DB.Production.DBName = "db_name"

	c.DB.Test.Host = "localhost"
	c.DB.Test.Username = "username"
	c.DB.Test.Password = "password"
	c.DB.Test.DBName = "db_name"

	c.Routing.Port = "8080"

	return c
}
