package work

import (
	"errors"
	"strings"
	"time"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
)

//
type Service interface {
	List(tags []string, order string, pageNum, pageSize int) ([]Sock, error) // GET /work
	Count(tags []string) (int, error)                                        // GET /work/size
	Get(id string) (Sock, error)                                             // GET /work/{id}
	Tags() ([]string, error)                                                 // GET /tags
	Health() []Health                                                        // GET /health
}

type Middleware func(Service) Service

// Work describes a work history item.
type Work struct {
	ID          string   `json:"id" db:"id"`
	Name        string   `json:"name" db:"name"`
	Description string   `json:"description" db:"description"`
	Price       float32  `json:"price" db:"price"`
	Count       int      `json:"count" db:"count"`
	Tags        []string `json:"tag" db:"-"`
}

// Health describes the health of a service
type Health struct {
	Service string `json:"service"`
	Status  string `json:"status"`
	Time    string `json:"time"`
}

// ErrNotFound is returned when there is no sock for a given ID.
var ErrNotFound = errors.New("not found")

// ErrDBConnection is returned when connection with the database fails.
var ErrDBConnection = errors.New("database connection error")

//
func NewWorkService(db *sqlx.DB, logger log.Logger) Service {
	return &workService{
		db:     db,
		logger: logger,
	}
}

type workService struct {
	db     *sqlx.DB
	logger log.Logger
}

func (s *workService) List(tags []string, order string, pageNum, pageSize int) ([]Work, error) {
	var workPlaces []Work
	query := baseQuery

	var args []interface{}

	for i, t := range tags {
		if i == 0 {
			query += " WHERE tag.name=?"
			args = append(args, t)
		} else {
			query += " OR tag.name=?"
			args = append(args, t)
		}
	}

	query += " GROUP BY id"

	if order != "" {
		query += " ORDER BY ?"
		args = append(args, order)
	}

	query += ";"

	err := s.db.Select(&workPlaces, query, args...)
	if err != nil {
		s.logger.Log("database error", err)
		return []Sock{}, ErrDBConnection
	}
	for i, s := range workPlaces {
		workPlaces[i].ImageURL = []string{s.ImageURL_1, s.ImageURL_2}
		workPlaces[i].Tags = strings.Split(s.TagString, ",")
	}

	// DEMO: Change 0 to 850
	time.Sleep(0 * time.Millisecond)

	workPlaces = cut(workPlaces, pageNum, pageSize)

	return workPlaces, nil
}

func (s *workService) Count(tags []string) (int, error) {
	query := "SELECT COUNT(DISTINCT sock.sock_id) FROM sock JOIN sock_tag ON sock.sock_id=sock_tag.sock_id JOIN tag ON sock_tag.tag_id=tag.tag_id"

	var args []interface{}

	for i, t := range tags {
		if i == 0 {
			query += " WHERE tag.name=?"
			args = append(args, t)
		} else {
			query += " OR tag.name=?"
			args = append(args, t)
		}
	}

	query += ";"

	sel, err := s.db.Prepare(query)

	if err != nil {
		s.logger.Log("database error", err)
		return 0, ErrDBConnection
	}
	defer sel.Close()

	var count int
	err = sel.QueryRow(args...).Scan(&count)

	if err != nil {
		s.logger.Log("database error", err)
		return 0, ErrDBConnection
	}

	return count, nil
}

func (s *workService) Get(id string) (Work, error) {
	query := baseQuery + " WHERE sock.sock_id =? GROUP BY sock.sock_id;"

	var sock Sock
	err := s.db.Get(&sock, query, id)
	if err != nil {
		s.logger.Log("database error", err)
		return Sock{}, ErrNotFound
	}

	sock.ImageURL = []string{sock.ImageURL_1, sock.ImageURL_2}
	sock.Tags = strings.Split(sock.TagString, ",")

	return sock, nil
}

func (s *workService) Health() []Health {
	var health []Health
	dbstatus := "OK"

	err := s.db.Ping()
	if err != nil {
		dbstatus = "err"
	}

	app := Health{"catalogue", "OK", time.Now().String()}
	db := Health{"catalogue-db", dbstatus, time.Now().String()}

	health = append(health, app)
	health = append(health, db)

	return health
}

func (s *workService) Tags() ([]string, error) {
	var tags []string
	query := "SELECT name FROM tag;"
	rows, err := s.db.Query(query)
	if err != nil {
		s.logger.Log("database error", err)
		return []string{}, ErrDBConnection
	}
	var tag string
	for rows.Next() {
		err = rows.Scan(&tag)
		if err != nil {
			s.logger.Log("database error", err)
			continue
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

func cut(workPlaces []Work, pageNum, pageSize int) []Work {
	if pageNum == 0 || pageSize == 0 {
		return []Work{} // pageNum is 1-indexed
	}
	start := (pageNum * pageSize) - pageSize
	if start > len(workPlaces) {
		return []Work{}
	}
	end := (pageNum * pageSize)
	if end > len(workPlaces) {
		end = len(workPlaces)
	}
	return workPlaces[start:end]
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
