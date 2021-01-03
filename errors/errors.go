package errors

import (
	"bytes"
	"fmt"
	"strings"
)

type ErrorType string

const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	ERROR_TYPE_SCHEDULER ErrorType = "scheduler error"
)

type CrawlerError interface {
	Type() ErrorType //string 别名
	Error() string
}
// myCrawlerError 代表爬虫错误的实现类型。
type myCrawlerError struct {
	// errType 代表错误的类型。
	errType ErrorType
	// errMsg 代表错误的提示信息。
	errMsg string
	// fullErrMsg 代表完整的错误提示信息。
	fullErrMsg string
}

// NewCrawlerError 用于创建一个新的爬虫错误值。
func NewCrawlerError(errType ErrorType, errMsg string) CrawlerError {
	return &myCrawlerError{
		errType: errType,
		errMsg:  strings.TrimSpace(errMsg),
	}
}

// NewCrawlerErrorBy 用于根据给定的错误值创建一个新的爬虫错误值。
func NewCrawlerErrorBy(errType ErrorType, err error) CrawlerError {
	return NewCrawlerError(errType, err.Error())
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return  ce.fullErrMsg
}

func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}