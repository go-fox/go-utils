package converter

import "github.com/jinzhu/copier"

type EnumTypeConverter[TO ~int32, FROM ~string] struct {
	nameMap  map[int32]string
	valueMap map[string]int32
}

func NewEnumTypeConverter[TO ~int32, FROM ~string](
	nameMap map[int32]string,
	valueMap map[string]int32,
) *EnumTypeConverter[TO, FROM] {
	return &EnumTypeConverter[TO, FROM]{
		valueMap: valueMap,
		nameMap:  nameMap,
	}
}

func (m *EnumTypeConverter[TO, FROM]) From(to *TO) *FROM {
	if to == nil {
		return nil
	}

	find, ok := m.nameMap[int32(*to)]
	if !ok {
		return nil
	}

	from := FROM(find)
	return &from
}

func (m *EnumTypeConverter[TO, FROM]) To(from *FROM) *TO {
	if from == nil {
		return nil
	}

	find, ok := m.valueMap[string(*from)]
	if !ok {
		return nil
	}

	to := TO(find)
	return &to
}

func (m *EnumTypeConverter[TO, FROM]) NewConverterPair() []copier.TypeConverter {
	srcType := FROM("")
	dstType := TO(0)

	fromFn := m.To
	toFn := m.From

	return NewGenericTypeConverterPair(&srcType, &dstType, fromFn, toFn)
}

func NewGenericTypeConverterPair[A interface{}, B interface{}](
	srcType A,
	dstType B,
	fromFn func(src A) B,
	toFn func(src B) A,
) []copier.TypeConverter {
	return []copier.TypeConverter{
		{
			SrcType: srcType,
			DstType: dstType,
			Fn: func(src interface{}) (interface{}, error) {
				return fromFn(src.(A)), nil
			},
		},
		{
			SrcType: dstType,
			DstType: srcType,
			Fn: func(src interface{}) (interface{}, error) {
				return toFn(src.(B)), nil
			},
		},
	}
}
