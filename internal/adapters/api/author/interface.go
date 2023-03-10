package author

import (
	"clean_architecture_app/internal/domain/author"
	"context"
)

type Service interface {
	GetByUUID(ctx context.Context, uuid string) *author.Author
	GetAll(ctx context.Context, limit, offset int) []*author.Author
	Create(ctx context.Context, dto *author.CreateAuthorDTO) *author.Author
}
