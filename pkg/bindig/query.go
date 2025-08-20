package binding

import (
	"errors"
	"reflect"
	"strings"

	"github.com/gofiber/fiber/v3"
)

type QueryBinder struct{}

func NewQueryBinder() *QueryBinder {
	return &QueryBinder{}
}

func (q *QueryBinder) Name() string {
	return "query"
}

func (q *QueryBinder) MIMETypes() []string {
	return []string{}
}

func (q *QueryBinder) Parse(c fiber.Ctx, out any) error {
	if err := c.Bind().Query(out); err != nil {
		return err
	}

	queries := c.Queries()

	v := reflect.ValueOf(out)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("out must be a non-nil pointer")
	}

	v = v.Elem()
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		val := v.Field(i)

		if !val.CanSet() {
			continue
		}

		tag := field.Tag.Get("query")
		if tag == "" {
			tag = strings.ToLower(field.Name)
		}

		rawValue, ok := queries[tag]
		if !ok {
			continue
		}

		if strings.Contains(rawValue, ",") && field.Type.Kind() == reflect.Slice && field.Type.Elem().Kind() == reflect.String {
			parts := strings.Split(rawValue, ",")
			val.Set(reflect.ValueOf(parts))
		}
	}

	return nil
}
