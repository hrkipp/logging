package logging

import (
	"io"
	"os"

	"context"
)

type key int

const (
	writerKey key = iota
	builderKey
	levelKey
)

func WithWriter(ctx context.Context, writer io.Writer) context.Context {
	return context.WithValue(ctx, writerKey, writer)
}

func WriterFrom(ctx context.Context) io.Writer {
	if writer, ok := ctx.Value(writerKey).(io.Writer); ok {
		return writer
	}
	return os.Stdout
}

func WithBuilder(ctx context.Context, builder Builder) context.Context {
	return context.WithValue(ctx, builderKey, builder)
}

func BuilderFrom(ctx context.Context) Builder {
	if builder, ok := ctx.Value(builderKey).(Builder); ok {
		return builder
	}
	return DefaultBuilder
}

func WithLevel(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, levelKey, level)
}

func LevelFrom(ctx context.Context) Level {
	if level, ok := ctx.Value(levelKey).(Level); ok {
		return level
	}
	return INFO
}
