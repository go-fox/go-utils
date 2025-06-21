package converter

import (
	"github.com/jinzhu/copier"
)

// Converter 转换器
type Converter[TO, FROM any] struct {
	copierOption copier.Option
}

// NewConverter 创建一个转换器
func NewConverter[TO, FROM any]() *Converter[TO, FROM] {
	return &Converter[TO, FROM]{
		copierOption: copier.Option{
			Converters: []copier.TypeConverter{},
			DeepCopy:   true,
		},
	}
}

// AddConverter 添加转换器
func (c *Converter[TO, FROM]) AddConverter(converters ...copier.TypeConverter) *Converter[TO, FROM] {
	c.copierOption.Converters = append(c.copierOption.Converters, converters...)
	return c
}

// To converts the given 'from' pointer of type FROM to a pointer of type TO using the configured copier options.
// If 'from' is nil, it returns nil. The conversion panics if an error occurs during copying.
func (c *Converter[TO, FROM]) To(from *FROM) *TO {
	if from == nil {
		return nil
	}

	var to TO
	if err := copier.CopyWithOption(&to, from, c.copierOption); err != nil {
		panic(err)
	}

	return &to
}

// From converts the given 'to' pointer of type TO to a pointer of type FROM using the configured copier options.
// If 'to' is nil, it returns nil. The conversion panics if an error occurs during copying.
func (c *Converter[TO, FROM]) From(to *TO) *FROM {
	if to == nil {
		return nil
	}

	var from FROM
	if err := copier.CopyWithOption(&from, to, c.copierOption); err != nil {
		panic(err)
	}
	return &from
}
