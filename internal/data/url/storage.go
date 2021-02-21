// template version: 1.0.11

package url

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/davidyunus/shorty/internal/data"
)

// Storage ,
type Storage struct {
	q data.Queryable
}

//rowScanner represent rows object from sql
type rowScanner interface {
	Scan(dest ...interface{}) error
}

//rowScanner represent rows object from sql
type rowsResult interface {
	Next() bool
	Scan(dest ...interface{}) error
}

// NewStorage , Create new Storage.
func NewStorage(q data.Queryable) *Storage {
	return &Storage{
		q: q,
	}
}

// Single , find one Url record matching with condition specified by query and args.
func (s *Storage) Single(ctx context.Context, query string, args ...interface{}) (*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE %s LIMIT 2`, selectQuery(), query)
	rows, err := q.QueryContext(ctx, stmt, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	count := 0
	var result *Url
	for rows.Next() {
		if count > 1 {
			return nil, errors.New("found more than one record")
		}
		data := &Url{}
		err := scan(rows, data)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, nil
			}
			return nil, err
		}
		result = data
		count++
	}
	return result, nil
}

// First , find first Url record matching with condition specified by query and args.
func (s *Storage) First(ctx context.Context, query string, args ...interface{}) (*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE %s LIMIT 1`, selectQuery(), query)
	row := q.QueryRowContext(ctx, stmt, args...)

	data := &Url{}
	err := scan(row, data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

// FirstOrder , find first Url record matching with condition specified by query and args, and ordered.
func (s *Storage) FirstOrder(ctx context.Context, query, order string, args ...interface{}) (*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE %s ORDER BY %s LIMIT 1`, selectQuery(), query, order)
	row := q.QueryRowContext(ctx, stmt, args...)

	data := &Url{}
	err := scan(row, data)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return data, nil
}

// Where , find all Url records matching with condition specified by query and args.
func (s *Storage) Where(ctx context.Context, query string, args ...interface{}) ([]*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE (%s) AND %s`, selectQuery(), query, defaultFilter())
	rows, err := q.QueryContext(ctx, stmt, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

// WhereOrder , find all Url records matching with condition specified by query and args.
func (s *Storage) WhereOrder(ctx context.Context, query, order string, args ...interface{}) ([]*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE (%s) AND %s ORDER BY %s`, selectQuery(), query, defaultFilter(), order)
	rows, err := q.QueryContext(ctx, stmt, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

// WhereWithPaging , find all Url records matching with condition specified by query and args limiting the result specified by size
// when size has value less than 1, the function will use default value 20 for size.
// when page has value less than 1, the function will use default value 1 for page. page has base index 1
func (s *Storage) WhereWithPaging(ctx context.Context, page, size int, query, order string, args ...interface{}) ([]*Url, error) {
	q := s.pickQueryable(ctx)
	limit := size
	if limit < 1 {
		limit = 20
	}
	offset := page
	if offset < 1 {
		offset = 1
	}
	offset = (offset - 1) * limit
	stmt := fmt.Sprintf(`%s WHERE (%s) AND %s ORDER BY %s LIMIT %v OFFSET %v`, selectQuery(), query, defaultFilter(), order, limit, offset)
	rows, err := q.QueryContext(ctx, stmt, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

// WhereNoFilter , find all Url records matching with condition specified by query and args.
func (s *Storage) WhereNoFilter(ctx context.Context, query string, args ...interface{}) ([]*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE %s`, selectQuery(), query)
	rows, err := q.QueryContext(ctx, stmt, args...)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

// FindAll , find all Url records.
func (s *Storage) FindAll(ctx context.Context) ([]*Url, error) {
	q := s.pickQueryable(ctx)
	stmt := fmt.Sprintf(`%s WHERE %s`, selectQuery(), defaultFilter())
	rows, err := q.QueryContext(ctx, stmt)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return scanRows(rows)
}

// FindByKeys , find Url using it's primary key(s).
func (s *Storage) FindByKeys(ctx context.Context, id int) (*Url, error) {
	criteria := `"id" = $1`
	stmt := fmt.Sprintf(`(%s) AND %s`, criteria, defaultFilter())
	return s.Single(ctx, stmt, id)
}

// FindByKeysNoFilter , find Url using it's primary key(s) without filter.
func (s *Storage) FindByKeysNoFilter(ctx context.Context, id int) (*Url, error) {
	criteria := `"id" = $1`
	return s.Single(ctx, criteria, id)
}

// Create , create new Url record.
func (s *Storage) Create(ctx context.Context, p *Url) error {
	q := s.pickQueryable(ctx)

	stmt, args := InsertQuery(p)
	row := q.QueryRowContext(ctx, stmt, args...)
	return scan(row, p)
}

// Update , update Url record.
func (s *Storage) Update(ctx context.Context, p *Url) error {
	q := s.pickQueryable(ctx)
	record, err := s.FindByKeys(ctx,
		p.ID,
	)
	if err != nil {
		return err
	}
	if record == nil {
		return errors.New("record not found")
	}
	record.URL = p.URL
	record.Shortcode = p.Shortcode
	record.RedirectCount = p.RedirectCount
	record.StartDate = p.StartDate
	record.LastSeenDate = p.LastSeenDate

	stmt, args := UpdateQuery(record)
	row := q.QueryRowContext(ctx, stmt, args...)
	return scan(row, p)
}

// Save , create new Url if it doesn't exist or update if exists.
func (s *Storage) Save(ctx context.Context, p *Url) error {
	record, err := s.FindByKeys(ctx,
		p.ID,
	)
	if err != nil {
		return err
	}
	if record != nil {
		return s.Update(ctx, p)
	}
	return s.Create(ctx, p)
}

func (s *Storage) pickQueryable(ctx context.Context) data.Queryable {
	q, ok := data.QueryableFromContext(ctx)
	if !ok {
		q = s.q
	}
	return q
}
func fields() string {
	return `"id", "url", "shortcode", "redirect_count", "start_date", "last_seen_date"`
}

func selectQuery() string {
	return fmt.Sprintf(`SELECT %s FROM "url"`, fields())
}

// InsertQuery returns query statement and slice of arguments
func InsertQuery(data *Url) (string, []interface{}) {
	o := []string{
		"url",
		"shortcode",
		"redirect_count",
		"start_date",
		"last_seen_date",
	}
	m := map[string]interface{}{
		"url":            data.URL,
		"shortcode":      data.Shortcode,
		"redirect_count": data.RedirectCount,
		"start_date":     data.StartDate,
		"last_seen_date": data.LastSeenDate,
	}

	fs, ph := func(v map[string]interface{}) ([]string, []string) {
		fs := []string{}
		ph := []string{}
		for i, k := range o {
			fs = append(fs, fmt.Sprintf(`"%s"`, k))
			ph = append(ph, fmt.Sprintf(`$%d`, i+1))
		}
		return fs, ph
	}(m)
	args := func(v map[string]interface{}) []interface{} {
		args := []interface{}{}
		for _, k := range o {
			v := v[k]
			args = append(args, v)
		}
		return args
	}(m)

	return fmt.Sprintf(`
        INSERT INTO "url" (%s) 
        VALUES 
            (%s)
        RETURNING %s`, strings.Join(fs, ","), strings.Join(ph, ","), fields()), args
}

// UpdateQuery returns query statement and slice of arguments
func UpdateQuery(data *Url) (string, []interface{}) {
	return fmt.Sprintf(`
        UPDATE "url"
        SET   
            "url" = $1 ,   
            "shortcode" = $2 ,   
            "redirect_count" = $3 ,   
            "start_date" = $4 ,   
            "last_seen_date" = $5
        WHERE 
            "id" = $6
        RETURNING %s`, fields()),
		[]interface{}{data.URL, data.Shortcode, data.RedirectCount, data.StartDate, data.LastSeenDate, data.ID}
}

func defaultFilter() string {
	return `true`
}

func scan(scanner rowScanner, data *Url) error {
	var iID int
	var iURL string
	var iShortcode string
	var iRedirectCount int
	var iStartDate time.Time
	var iLastSeenDate time.Time

	err := scanner.Scan(&iID, &iURL, &iShortcode, &iRedirectCount, &iStartDate, &iLastSeenDate)
	if err != nil {
		return err
	}

	data.ID = iID
	data.URL = iURL
	data.Shortcode = iShortcode
	data.RedirectCount = iRedirectCount
	data.StartDate = iStartDate
	data.LastSeenDate = iLastSeenDate

	return nil
}

func scanRows(rows rowsResult) ([]*Url, error) {
	collection := []*Url{}
	for rows.Next() {
		data := &Url{}
		err := scan(rows, data)
		if err != nil {
			return nil, err
		}
		collection = append(collection, data)
	}
	return collection, nil
}
