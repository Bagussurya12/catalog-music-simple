package configs

type (
	Config struct {
		Service  Service
		Database Database
	}

	Service struct {
		Port      string
		SecretJWT string
	}

	Database struct {
		DataSourceName string
	}
)
