package config

type Config struct {
    DBConnString string
}

func GetConfig() *Config {
    return &Config{
        DBConnString: "host=localhost port=5433 user=postgres password=pass111 dbname=laba4",
    }
}