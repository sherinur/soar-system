package config

import (
	"time"

	"github.com/caarlos0/env/v7"
	"github.com/joho/godotenv"
	"github.com/sherinur/soar-system/backend/auth_service/pkg/postgrescon"
)

type (
	Config struct {
		Server    Server `envPrefix:"SERVER_"`
		ZapLogger ZapLogger
		Telemetry Telemetry
		Jwt       Jwt

		Postgres postgrescon.Postgres
	}

	Server struct {
		GRPCServer GRPCServer
		HTTPServer HTTPServer `envPrefix:"HTTP_"`
	}

	GRPCServer struct {
		Port                  int32         `env:"GRPC_PORT,notEmpty"`
		MaxRecvMsgSizeMiB     int           `env:"GRPC_MAX_MESSAGE_SIZE_MIB" envDefault:"12"`
		MaxConnectionAge      time.Duration `env:"GRPC_MAX_CONNECTION_AGE" envDefault:"30s"`
		MaxConnectionAgeGrace time.Duration `env:"GRPC_MAX_CONNECTION_AGE_GRACE" envDefault:"10s"`
	}

	HTTPServer struct {
		Port         int           `env:"HTTP_PORT" envDefault:"8080"`
		ReadTimeout  time.Duration `env:"HTTP_READ_TIMEOUT" envDefault:"30s"`
		WriteTimeout time.Duration `env:"HTTP_WRITE_TIMEOUT" envDefault:"30s"`
		Mode         string        `env:"GIN_MODE" envDefault:"release"` // release, debug, test
	}

	Nats struct{}

	ZapLogger struct {
		Directory string `env:"ZAP_LOGGING_DIRECTORY" envDefault:"./logs"`
		Mode      string `env:"ZAP_LOGGING_MODE" envDefault:"debug"` // release, debug, test
	}

	Jwt struct {
		JwtAccessSecret      string `env:"JWT_ACCESS_SECRET"`
		JwtRefreshSecret     string `env:"JWT_REFRESH_SECRET"`
		JwtAccessExpiration  int    `env:"JWT_ACCESS_EXPIRATION"`
		JwtRefreshExpiration int    `env:"JWT_REFRESH_EXPIRATION"`
	}

	Telemetry struct {
		Mode                 string `env:"OTEL_MODE" envDefault:"debug"` // release, debug, test
		ExporterOTLPEndpoint string `env:"OTEL_EXPORTER_OTLP_ENDPOINT" envDefault:"http://localhost:4318"`
		ExporterOTLPInsecure bool   `env:"OTEL_EXPORTER_OTLP_INSECURE" envDefault:"true"`
		ExporterPromPort     int    `env:"OTEL_EXPORTER_PROM_PORT" envDefault:"3003"`
	}
)

func New() (*Config, error) {
	var cfg Config
	opts := env.Options{Environment: nil, TagName: "env", Prefix: ""}

	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	err = env.Parse(&cfg, opts)
	if err != nil {
		return nil, err
	}

	return &cfg, err
}
