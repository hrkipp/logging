package logging_test

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/hrkipp/logging"

	"context"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type LoggingSuite struct {
	ctx context.Context
	buf bytes.Buffer
}

func (s *LoggingSuite) SetUpTest(c *C) {
	s.ctx = context.Background()
	s.buf = bytes.Buffer{}
	s.ctx = logging.WithWriter(s.ctx, &s.buf)
}

var _ = Suite(&LoggingSuite{})

func (s *LoggingSuite) TestLoggingErrorsToABuffer(c *C) {
	logging.Error(s.ctx, "foo")
	c.Assert(s.buf.String(), Equals, "ERROR foo")
}

func (s *LoggingSuite) TestLoggingWarningsToABuffer(c *C) {
	logging.Warning(s.ctx, "foo")
	c.Assert(s.buf.String(), Equals, "WARNING foo")
}

func (s *LoggingSuite) TestLoggingInfo(c *C) {
	logging.Info(s.ctx, "foo")
	c.Assert(s.buf.String(), Equals, "INFO foo")
}

func (s *LoggingSuite) TestLoggingWithNoWriter(c *C) {
	logging.Info(context.Background(), "foo")
	logging.Warning(context.Background(), "foo")
	logging.Error(context.Background(), "foo")
}

func (s *LoggingSuite) TestLoggingWithLogLevel(c *C) {
	ctx := logging.WithLevel(s.ctx, logging.WARNING)
	logging.Info(ctx, "foo")
	c.Assert(s.buf.String(), Equals, "")
}

func (s *LoggingSuite) TestUsingNonStandardBuilder(c *C) {
	builder := func(ctx context.Context, level logging.Level, args ...interface{}) []byte {
		return []byte(fmt.Sprintf("custom %v %v", level.String(), args[0]))
	}
	ctx := logging.WithBuilder(s.ctx, builder)
	logging.Info(ctx, "foo")
	c.Assert(s.buf.String(), Equals, "custom INFO foo")

}
