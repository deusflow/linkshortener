package domain

import "context"

type URLRepository interface {
	//Saver
	Save(ctx context.Context, url *URL) error

	//finder
	FindByShortCode(ctx context.Context, shortCode string) (*URL, error)

	//FindByID
	FindById(ctx context.Context, id int64) (*URL, error)

	// Delete
	Delete(ctx context.Context, id int64) error
}

// ClickRepository interface for click operations
type ClickRepository interface {
	//SaveClick saves a click record

	Save(ctx context.Context, click *Click) error

	// FindByURLID retrieves all clicks for a given URL ID
	FindByURLId(ctx context.Context, urlId int64) ([]*Click, error)

	//CountClicks returns the number of clicks for a given URL ID
	CountByURLId(ctx context.Context, urlId int64) (int64, error)
}
