// Package str 字符串辅助方法
package str

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// Plural user -> users
func Plural(word string) string {
	return pluralize.NewClient().Plural(word)
}

// Singular users -> user
func Singular(word string) string {
	return pluralize.NewClient().Singular(word)
}

// Snake eg: TopicComment -> topic_comment
func Snake(s string) string {
	return strcase.ToSnake(s)
}

// Camel eg: topic_comment -> TopicComment
func Camel(s string) string {
	return strcase.ToCamel(s)
}

// LowerCamel TopicComment -> topicComment
func LowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}
