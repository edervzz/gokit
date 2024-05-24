package repository

import "context"

type ICreate[TEntity any] interface {
	Create(context.Context, TEntity) error
}

type IUpdate[TEntity any] interface {
	Update(context.Context, TEntity) error
}

type IReadSingle[TKey any, TEntity any] interface {
	Read(TKey, context.Context) (TEntity, error)
}

type IReadByExternalID[TKey any, TEntity any] interface {
	ReadByExternalID(context.Context, TKey) (TEntity, error)
}

type IReadAll[TEntity any] interface {
	ReadAll(context.Context) ([]TEntity, error)
}
