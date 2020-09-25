package gopenapi

import (
	"fmt"
	"reflect"

	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/parser"
	"github.com/goccy/go-yaml/printer"
)

func yamlPathErrorMessage(source interface{}, errMsg string, pathStr string) (string, error) {
	key := ""
	if t := reflect.TypeOf(source); t.Kind() == reflect.Ptr {
		key = fmt.Sprintf("%s.%s", t.Elem().Name(), pathStr)
	} else {
		key = fmt.Sprintf("%s.%s", t.Name(), pathStr)
	}

	bytes, err := yaml.Marshal(source)
	if err != nil {
		panic(err)
	}
	file, err := parser.ParseBytes(bytes, 0)
	if err != nil {
		return "", err
	}
	path, err := yaml.PathString(fmt.Sprintf("%s%s", "$.", pathStr))
	if err != nil {
		return "", err
	}
	node, err := path.FilterFile(file)
	if err != nil {
		return "", err
	}

	var pp printer.Printer

	pos := fmt.Sprintf("[%d:%d] ", node.GetToken().Position.Line, node.GetToken().Position.Column)
	errKey := fmt.Sprintf("%s%s'%s', ", pos, "Key: ", key)
	msg := pp.PrintErrorMessage(fmt.Sprintf("%s%s", errKey, errMsg), true)
	msg += "\n" + pp.PrintErrorToken(node.GetToken(), true)

	return msg, nil

}
