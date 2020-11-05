package main

import (
	"net/http"
	"flag"
	trm "github.com/raphaelreyna/easi-trm-api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"math/rand"
	"net"
	"context"
)

type server struct {
	http.Server
	devKey string

	tlskey, tlscert string
	client *trm.APIClient
	logger *zap.Logger

	internalNet *net.IPNet
}

func serverFromFlags(client *http.Client, version *bool) (*server, error) {
	var err error
	s := &server{}

	// Allocate memory to read flag values into
	var (
		port string
		host string
		devKey string
		cidr string
	)

	// Declare the flags available to the user
	flag.StringVar(&port, "port", "80", `Port to bind to.`)
	flag.StringVar(&host, "host", "", `A reachable hostname.`)
	flag.StringVar(&s.tlskey, "tls-key", "", `Path to the key file to use for TLS/SSL/HTTP.`)
	flag.StringVar(&s.tlscert, "tls-cert", "", `Path to the key file to use for TLS/SSL/HTTPS.`)
	flag.StringVar(&devKey, "dev-key", "", `Path to file containing a valid dev key.
If set, the dev key in the file will be used as the default dev key for all requests from all clients.`)
	flag.StringVar(&cidr, "internal-net", "", `CIDR string describing the network allowed to request health checks.`)
	flag.BoolVar(version, "v", false, `Show version information.`)
	// Parse the arguments into flag values
	flag.Parse()

	// Return early if the user just wants the version
	if *version {
		return nil, nil
	}

	// Configure the server based on flag values
	s.Addr = host + ":" + port

	// Setup the HTTP client to the EASI server(s)
	if client == nil { client = http.DefaultClient }
	s.client = trm.NewAPIClient(devKey, client)

	// Parse internal network CIDR
	if cidr != "" {
		_, s.internalNet, err = net.ParseCIDR(cidr)
		if err != nil {
			return nil, err
		}
	}

	// Setup logging
	s.logger, err = zap.NewProduction()
	if err != nil { return nil, err }
	s.ErrorLog, err = zap.NewStdLogAt(s.logger, zapcore.ErrorLevel)
	if err != nil { return nil, err }

	// Every request gets an id in their context
	s.ConnContext = func(ctx context.Context, c net.Conn) context.Context {
		// Generate an id for this new request
		id := make([]byte, 8)
		_, err = rand.Read(id)
		if err != nil {
			panic(err)
			return context.Background()
		}

		return context.WithValue(ctx, key, id)
	}

	return s, s.routes()
}

func (s *server) ListenAndServe() error {
	if s.tlscert != "" && s.tlskey != "" {
		return s.Server.ListenAndServeTLS(s.tlscert, s.tlskey)
	}

	return s.Server.ListenAndServe()
}

func (s *server) routes() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/report",
		s.logRoute(
			s.jsonRoute(s.handleGetReport()),
		),
	)
	mux.HandleFunc("/health",
		s.logRoute(
			s.internalRoute(s.healthCheck),
		),
	)
	mux.HandleFunc("/health/full",
		s.logRoute(
			s.internalRoute(
				s.fullHealthCheck,
			),
		),
	)

	s.Handler = mux
	return nil
}

func (s *server) err(w http.ResponseWriter, r *http.Request, msg string, err error) {
	id := r.Context().Value(key).([]byte)
	s.logger.Error(msg,
		zap.String("request_ip", r.RemoteAddr),
		zap.Binary("request_id", id),
		zap.Error(err),
	)
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
