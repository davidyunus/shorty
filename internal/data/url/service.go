// template version: 1.0.11
package url    

import (
	"context" 
	"github.com/davidyunus/shorty/internal/data"
)
 
// Service ,
type Service struct { 
    q data.Queryable
    s *Storage
}
// Single , find single Url record matching with condition specified by query and args.
func (s *Service) Single (ctx context.Context, query string, args ...interface{}) (*Url, error) {
    return s.s.Single(ctx, query, args...)
}
// First , find first Url record matching with condition specified by query and args.
func (s *Service) First (ctx context.Context, query string, args ...interface{}) (*Url, error) {
    return s.s.First(ctx, query, args...)
}
// FirstOrder , find first Url record matching with condition specified by query and args, and ordered.
func (s *Service) FirstOrder (ctx context.Context, query, order string, args ...interface{}) (*Url, error) {
    return s.s.FirstOrder(ctx, query, order, args...)
}
// Where , find all Url records matching with condition specified by query and args.
func (s *Service) Where (ctx context.Context, query string, args ...interface{}) ([]*Url, error) {
    return s.s.Where(ctx, query, args...)
}
// WhereOrder , find all Url records matching with condition specified by query and args, and ordered.
func (s *Service) WhereOrder (ctx context.Context, query, order string, args ...interface{}) ([]*Url, error) {
    return s.s.WhereOrder(ctx, query, order, args...)
}
// WhereWithPaging , find all Url records matching with condition specified by query and args limiting the result specified by size
// when size has value less than 1, the function will use default value 20 for size.
// when page has value less than 1, the function will use default value 1 for page. page has base index 1
func (s *Service) WhereWithPaging (ctx context.Context, page, size int, query, order string, args ...interface{}) ([]*Url, error) {
    return s.s.WhereWithPaging(ctx, page, size, query, order, args...)
}
// WhereNoFilter , find all Url records matching with condition specified by query and args.
func (s *Service) WhereNoFilter (ctx context.Context, query string, args ...interface{}) ([]*Url, error) {
    return s.s.WhereNoFilter(ctx, query, args...)
}
// FindAll , find all Url records.
func (s *Service) FindAll(ctx context.Context) ([]*Url, error) { 
    return s.s.FindAll(ctx)
}
// FindByKeys , find Url using it's primary key(s).
func (s *Service) FindByKeys(ctx context.Context, id int) (*Url, error){
    return s.s.FindByKeys(ctx, id)
}
// Create , create new Url record.
func (s *Service) Create(ctx context.Context, pUrl *Url) error {
    return s.s.Create(ctx, pUrl)
}
// Update , update Url record.
func (s *Service) Update(ctx context.Context, pUrl *Url) error {
    return s.s.Update(ctx, pUrl)
}
// Save , create new Url if it doesn't exist or update if exists.
func (s *Service) Save(ctx context.Context, pUrl *Url) error{
    return s.s.Save(ctx, pUrl)
} 
// FindByKeysNoFilter , find Url using it's primary key(s).
func (s *Service) FindByKeysNoFilter(ctx context.Context, id int) (*Url, error){
    return s.s.FindByKeysNoFilter(ctx, id)
}
  
// NewService , returns new Service.
func NewService(q data.Queryable) *Service{
    s := NewStorage(q)
    service := &Service{
        q: q,
        s : s,
    }
    return service
}
 
type key int
const ctxKey key = 0

// NewContext , return new context with s.
func NewContext(ctx context.Context, s *Service) context.Context {
	return context.WithValue(ctx, ctxKey, s)
}

// FromContext , return a service from a context.
func FromContext(ctx context.Context) (*Service, bool) {
	service, ok := ctx.Value(ctxKey).(*Service)
	return service, ok
}
