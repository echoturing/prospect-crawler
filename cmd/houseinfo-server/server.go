package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/echoturing/prospect-crawler/db"
	"github.com/echoturing/prospect-crawler/etc"
	"github.com/echoturing/prospect-crawler/linkhome"
	"github.com/echoturing/prospect-crawler/logger"
	"github.com/gocraft/dbr"
	"go.uber.org/zap"
)

var (
	port       = flag.Int("port", 8000, "port")
	root       = flag.String("root", ".", "web root")
	configPath = flag.String("config", "./etc/config.yaml", "The config file path")
)

type server struct {
	http.ServeMux
	conn *dbr.Connection
}

func newServer(conn *dbr.Connection) http.Handler {
	srv := &server{conn: conn}
	// srv.HandleFunc("/", srv.index)
	srv.Handle("/", http.FileServer(http.Dir(*root)))
	srv.HandleFunc("/query", srv.query)
	return srv
}

func (s *server) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "TODO\n")
}

func (s *server) query(w http.ResponseWriter, r *http.Request) {
	const defaultLimit = 1000
	session := s.conn.NewSession(nil)
	query := session.Select(strings.Join(linkhome.TableHouseColumns, ",")).From(linkhome.TableHouseInfo).Limit(defaultLimit)
	var result []linkhome.HouseInfo
	if _, err := query.Load(&result); err != nil {
		http.Error(w, "failed to query", http.StatusInternalServerError)
		log := logger.GetLogger()
		log.Error("query", zap.String("err", err.Error()))
		return
	}
	setCORS(w)
	responseJSON(w, &result)
}

func responseJSON(w http.ResponseWriter, i interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w)
	e.SetIndent("", "  ")
	return e.Encode(i)
}

func setCORS(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
}

func main() {
	flag.Parse()
	log := logger.GetLogger()
	defer log.Sync()
	filePath, err := filepath.Abs(*configPath)
	if err != nil {
		log.Error("load config file path failed", zap.Error(err))
		os.Exit(1)
	}
	cfg, err := etc.LoadConfigFromFile(filePath)
	if err != nil {
		log.Error("load config failed", zap.Error(err))
		os.Exit(1)
	}
	conn, err := db.NewConn(cfg)
	if err != nil {
		log.Error("init db conn failed", zap.Error(err))
		os.Exit(1)
	}
	srv := newServer(conn)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *port), srv)
}
