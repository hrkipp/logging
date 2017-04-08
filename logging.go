package logging

import (
	"context"
)

func Error(ctx context.Context, args ...interface{}) {
	if LevelFrom(ctx) <= ERROR {
		WriterFrom(ctx).Write(BuilderFrom(ctx)(ctx, ERROR, args...))
	}
}

func Warning(ctx context.Context, args ...interface{}) {
	if LevelFrom(ctx) <= WARNING {
		WriterFrom(ctx).Write(BuilderFrom(ctx)(ctx, WARNING, args...))
	}
}

func Info(ctx context.Context, args ...interface{}) {
	if LevelFrom(ctx) <= INFO {
		WriterFrom(ctx).Write(BuilderFrom(ctx)(ctx, INFO, args...))
	}
}
