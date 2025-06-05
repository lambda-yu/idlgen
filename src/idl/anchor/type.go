package anchor

import (
	"encoding/json"
	"fmt"
)

type TypeStr string

const (
	Bool      TypeStr = "bool"
	Uint8     TypeStr = "u8"
	Int8      TypeStr = "i8"
	Uint16    TypeStr = "u16"
	Int16     TypeStr = "i16"
	Uint32    TypeStr = "u32"
	Int32     TypeStr = "i32"
	Float32   TypeStr = "f32"
	Uint64    TypeStr = "u64"
	Int64     TypeStr = "i64"
	Float64   TypeStr = "f64"
	Uint128   TypeStr = "u128"
	Int128    TypeStr = "i128"
	Uint256   TypeStr = "u256"
	Int256    TypeStr = "i256"
	Bytes     TypeStr = "bytes"
	String    TypeStr = "string"
	PublicKey TypeStr = "pubkey"

	// Custom additions:
	UnixTimestamp TypeStr = "unixTimestamp"
	Hash          TypeStr = "hash"
	Duration      TypeStr = "duration"
)

type Serialization string

const (
	Borsh          Serialization = "borsh"
	ByteMuck       Serialization = "bytemuck"
	ByteMuckUnSafe Serialization = "bytemuckunsafe"
)

type IdlTypeDefTyKind string

func (t IdlTypeDefTyKind) Value() string {
	return string(t)
}

const (
	IdlTypeDefTyKindStruct IdlTypeDefTyKind = "struct"
	IdlTypeDefTyKindEnum   IdlTypeDefTyKind = "enum"
	IdlTypeDefTyKindEtType IdlTypeDefTyKind = "type"
)

type IdlTypeDef struct {
	Name          string              `json:"name"`
	Docs          []string            `json:"docs,omitempty"`
	Serialization Serialization       `json:"serialization,omitempty"`
	Repr          IdlRepr             `json:"repr"`
	Generics      []IdlTypeDefGeneric `json:"generics,omitempty"`
	Type          IdlTypeDefTy        `json:"type"`
}

type IdlTypeDefSlice []IdlTypeDef

func (named IdlTypeDefSlice) GetByName(name string) *IdlTypeDef {
	for i := range named {
		v := named[i]
		if v.Name == name {
			return &v
		}
	}
	return nil
}

type IdlTypeDefTy struct {
	TypeDefTyEnum   *TypeDefTyEnum
	TypeDefTyStruct *TypeDefTyStruct
	TypeDefTyType   *TypeDefTyType
}

func (t *IdlTypeDefTy) UnmarshalJSON(data []byte) error {
	// 解析为 map[string]any
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("invalid IdlTypeDefTy: %w", err)
	}
	kind, err := obj["kind"].MarshalJSON()
	if err != nil {
		return fmt.Errorf("invalid IdlTypeDefTy kind error: %w", err)
	}

	var kindStr string
	err = json.Unmarshal(kind, &kindStr)
	if err != nil {
		return fmt.Errorf("invalid Seed kind error: %w", err)
	}
	// 判断字段类型并递归解析
	switch kindStr {
	case IdlTypeDefTyKindStruct.Value():
		var enum TypeDefTyEnum
		if err := json.Unmarshal(data, &enum); err != nil {
			return err
		}
		t.TypeDefTyEnum = &enum
		return nil
	case IdlTypeDefTyKindEnum.Value():
		var typeStruct TypeDefTyStruct
		if err := json.Unmarshal(data, &typeStruct); err != nil {
			return err
		}
		t.TypeDefTyStruct = &typeStruct
		return nil
	case IdlTypeDefTyKindEtType.Value():
		var tt TypeDefTyType
		if err := json.Unmarshal(data, &tt); err != nil {
			return err
		}
		t.TypeDefTyType = &tt
		return nil
	default:
		return fmt.Errorf("unknown IdlTypeDefTy field")
	}

	return fmt.Errorf("unrecognized IdlTypeDefTy: %s", string(data))

}

type TypeDefTyEnum struct {
	Kind     string                 `json:"kind"`
	Variants []TypeDefTyEnumVariant `json:"variants"`
}

type TypeDefTyEnumVariant struct {
	Name   string            `json:"name"`
	Fields TypeDefinedFields `json:"fields,omitempty"`
}

type TypeDefTyStruct struct {
	Kind   string            `json:"kind"`
	Fields TypeDefinedFields `json:"fields,omitempty"`
}

type TypeDefinedFields struct {
	DefinedFieldsNamed []IdlField
	DefinedFieldsTuple []IdlType
}

func (t *TypeDefinedFields) UnmarshalJSON(data []byte) error {
	// 解析为 map[string]any
	var fieldsTuple []IdlType
	var fieldsNamed []IdlField
	if err := json.Unmarshal(data, &fieldsTuple); err == nil {
		t.DefinedFieldsTuple = fieldsTuple
		return nil
	} else if err := json.Unmarshal(data, &fieldsNamed); err == nil {
		t.DefinedFieldsNamed = fieldsNamed
		return nil
	} else {
		return fmt.Errorf("invalid IdlTypeDefTy: %w", err)
	}
}

type TypeDefTyType struct {
	Kind  string  `json:"kind"`
	Alias IdlType `json:"alias"`
}

type IdlRepr struct {
	Kind   string `json:"kind"`
	Packed bool   `json:"packed,omitempty"`
	Align  int    `json:"align,omitempty"`
}

type IdlTypeDefGeneric struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
	Type string `json:"type,omitempty"`
}

type IdlType struct {
	Primitive *string             `json:"-"`
	Option    *IdlType            `json:"option,omitempty"`
	COption   *IdlType            `json:"coption,omitempty"`
	Vec       *IdlType            `json:"vec,omitempty"`
	Array     *TypeArray[IdlType] `json:"array,omitempty"`
	Defined   *TypeDefinedItem    `json:"defined,omitempty"`
	Generic   *TypeGeneric        `json:"generic,omitempty"`
}

func (t *IdlType) UnmarshalJSON(data []byte) error {
	// 尝试解析为基本类型字符串（如 "u64"）
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		t.Primitive = &s
		return nil
	}

	// 解析为 map[string]any
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("invalid IdlType: %w", err)
	}

	// 判断字段类型并递归解析
	for key, value := range obj {
		switch key {
		case "option":
			var nested IdlType
			if err := json.Unmarshal(value, &nested); err != nil {
				return err
			}
			t.Option = &nested
			return nil
		case "coption":
			var nested IdlType
			if err := json.Unmarshal(value, &nested); err != nil {
				return err
			}
			t.COption = &nested
			return nil
		case "vec":
			var nested IdlType
			if err := json.Unmarshal(value, &nested); err != nil {
				return err
			}
			t.Vec = &nested
			return nil
		case "array":
			// 你可以定义专门的结构来解析固定长度数组等
			var arr TypeArray[IdlType]
			if err := json.Unmarshal(value, &arr); err != nil {
				return err
			}
			t.Array = &arr
			return nil
		case "defined":
			var def TypeDefinedItem
			if err := json.Unmarshal(value, &def); err != nil {
				return err
			}
			t.Defined = &def
			return nil
		case "generic":
			var g TypeGeneric
			if err := json.Unmarshal(value, &g); err != nil {
				return err
			}
			t.Generic = &g
			return nil
		default:
			return fmt.Errorf("unknown IdlType field: %s", key)
		}
	}

	return fmt.Errorf("unrecognized IdlType: %s", string(data))
}

type TypeGeneric struct {
	Generic string `json:"generic"`
}

type TypeDefinedItem struct {
	Defined    TypeDefined `json:"defined"`
	DefinedStr string      `json:"definedStr"`
}

func (t *TypeDefinedItem) UnmarshalJSON(data []byte) error {
	// 尝试解析为基本类型字符串
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		t.DefinedStr = s
		return nil
	} else {
		var s TypeDefined
		err := json.Unmarshal(data, &s)
		if err != nil {
			return fmt.Errorf("UnmarshalJSON TypeDefinedItem error: %s", err.Error())
		}

		t.Defined = s

		return nil
	}

	// 如果全失败，返回错误
	return fmt.Errorf("unsupported JSON type: %s", string(data))
}

type TypeDefined struct {
	Name     string        `json:"name"`
	Generics []TypeGeneric `json:"generics,omitempty"`
}

type TypeArrayLenValue uint

type TypeArrayGeneric struct {
	Generic string `json:"generic"`
}

type TypeArray[T any] struct {
	ArrayType T
	// and
	ArraySize TypeArrayLen
}

func (f *TypeArray[IdlType]) UnmarshalJSON(data []byte) error {
	// 尝试按常见类型依次解析
	var val []json.RawMessage
	var err error
	if err = json.Unmarshal(data, &val); err == nil {
		var d1 IdlType
		if err = json.Unmarshal(val[0], &d1); err == nil {
			f.ArrayType = d1
		}

		var d2 *TypeArrayLen
		if err = json.Unmarshal(val[1], &d2); err == nil {
			f.ArraySize = *d2
		}
	}

	// 如果全失败，返回错误
	if err != nil {
		return fmt.Errorf("unsupported JSON type: %s", string(data))
	} else {
		return nil
	}

}

type TypeArrayLen struct {
	ArraySize TypeArrayLenValue
	// or
	ArrayGenericSize TypeArrayGeneric
}

func (f *TypeArrayLen) UnmarshalJSON(data []byte) error {
	// 尝试按常见类型依次解析
	var val uint
	if err := json.Unmarshal(data, &val); err == nil {
		var lenValue = TypeArrayLenValue(val)
		f.ArraySize = lenValue
		return nil
	}

	var valGeneric TypeArrayGeneric
	if err := json.Unmarshal(data, &valGeneric); err == nil {
		f.ArrayGenericSize = valGeneric
		return nil
	}

	// 如果全失败，返回错误
	return fmt.Errorf("unsupported JSON type: %s", string(data))
}
