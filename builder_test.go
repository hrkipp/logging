package logging_test

import (
	"github.com/hrkipp/logging"
	. "gopkg.in/check.v1"
)

func (s *LoggingSuite) TestLoggingWithFormatString(c *C) {
	logging.Info(s.ctx, "%v : %v", "foo", 1)
	c.Assert(s.buf.String(), Equals, "INFO foo : 1")
}

func (s *LoggingSuite) TestLoggingJustANumber(c *C) {
	logging.Info(s.ctx, 1)
	c.Assert(s.buf.String(), Equals, "INFO 1")
}

func (s *LoggingSuite) TestLoggingAStruct(c *C) {
	type foo struct {
		Bar int
		Baz string
	}
	logging.Info(s.ctx, foo{1, "foo"})
	c.Assert(s.buf.String(), Equals, `INFO {"Bar":1,"Baz":"foo"}`)
}

func (s *LoggingSuite) TestLoggingMultipleStructs(c *C) {
	type foo struct {
		Bar int
		Baz string
	}
	logging.Info(s.ctx, foo{1, "foo"}, foo{2, "bar"})
	c.Assert(s.buf.String(), Equals, `INFO {"Bar":1,"Baz":"foo"} {"Bar":2,"Baz":"bar"}`)
}
