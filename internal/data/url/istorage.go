// template version: 1.0.11
package url

import (
	"context"
)

// IStorage interface that wraps methods for working with table url
type IStorage interface {
	// Single , find single Url record matching with condition specified by query and args.
	Single(ctx context.Context, query string, args ...interface{}) (*Url, error)
	// First , find first Url record matching with condition specified by query and args.
	First(ctx context.Context, query string, args ...interface{}) (*Url, error)
	// FirstOrder , find first Url record matching with condition specified by query and args, and ordered.
	FirstOrder(ctx context.Context, query, order string, args ...interface{}) (*Url, error)
	// Where , find all Url records matching with condition specified by query and args.
	Where(ctx context.Context, query string, args ...interface{}) ([]*Url, error)
	// WhereOrder , find all Url records matching with condition specified by query and args, and ordered.
	WhereOrder(ctx context.Context, query, order string, args ...interface{}) ([]*Url, error)
	// WhereWithPaging , find all Url records matching with condition specified by query and args limiting the result specified by size
	// when size has value less than 1, the function will use default value 20 for size.
	// when page has value less than 1, the function will use default value 1 for page. page has base index 1
	WhereWithPaging(ctx context.Context, page, size int, query, order string, args ...interface{}) ([]*Url, error)
	// WhereNoFilter , find all Url records matching with condition specified by query and args.
	WhereNoFilter(ctx context.Context, query string, args ...interface{}) ([]*Url, error)
	// FindAll , find all Url records.
	FindAll(ctx context.Context) ([]*Url, error)
	// FindByKeys , find Url using it's primary key(s).
	FindByKeys(ctx context.Context, id int) (*Url, error)
	// FindByKeysNoFilter , find Url using it's primary key(s).
	FindByKeysNoFilter(ctx context.Context, id int) (*Url, error)
	// Create , create new Url record.
	Create(ctx context.Context, pUrl *Url) error
	// Update , update Url record.
	Update(ctx context.Context, pUrl *Url) error
	// Save , create new Url if it doesn't exist or update if exists.
	Save(ctx context.Context, pUrl *Url) error
}
