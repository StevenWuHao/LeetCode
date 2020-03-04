package factory

import "testing"

func TestIndex(t *testing.T) {

	var logger Logger
	logger = FileFactory{}.new("safsaf")
	logger = StdoutFactory{}.new()
	logger.echo()
}
