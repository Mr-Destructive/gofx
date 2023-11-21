package gofx

import (
	"fmt"
	"reflect"
)

type Map[K comparable, V any] map[K]V

func (m *Map[K, V]) Get(key K) (V, bool) {
	value, ok := (*m)[key]
	return value, ok
}

func (m *Map[K, V]) GetSet(key K, defaultValue V) (V, bool) {
	value, ok := (*m)[key]
	if !ok {
		(*m)[key] = defaultValue
		return defaultValue, ok
	}
	(*m)[key] = value
	return value, ok
}

func (m *Map[K, V]) Delete(key K) (V, bool) {
	value, ok := (*m)[key]
	delete(*m, key)
	return value, ok
}

func (m *Map[K, V]) Set(key K, value V) {
	(*m)[key] = value
}

func (m *Map[K, V]) Keys() []K {
	keys := make([]K, 0, len(*m))
	for key := range *m {
		keys = append(keys, key)
	}
	return keys
}

func (m *Map[K, V]) Values() []V {
	values := make([]V, 0, len(*m))
	for _, value := range *m {
		values = append(values, value)
	}
	return values
}

func (m *Map[K, V]) Len() int {
	return len(*m)
}

func Dict(obj interface{}) map[string]interface{} {
	result := make(map[string]interface{})

	val := reflect.ValueOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	switch val.Kind() {
	case reflect.Struct:
		typ := val.Type()
		for i := 0; i < val.NumField(); i++ {
			field := val.Field(i)
			fieldName := typ.Field(i).Name
			result[fieldName] = field.Interface()
		}
	case reflect.Map:
		iter := val.MapRange()
		for iter.Next() {
			key := iter.Key()
			value := iter.Value()
			result[fmt.Sprintf("%v", key.Interface())] = value.Interface()
		}
	default:
		result["value"] = val.Interface()
	}

	return result
}
