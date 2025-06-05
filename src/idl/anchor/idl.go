package anchor

// https://github.com/solana-foundation/anchor/blob/8b0e965c65fb96b6865be53c478a16007984a566/ts/packages/anchor/src/idl.ts#L4
type Idl struct {
	Address      string          `json:"address"`
	Metadata     Metadata        `json:"metadata"`
	Docs         []byte          `json:"docs,omitempty"`
	Instructions []Instruction   `json:"instructions"`
	Accounts     []Account       `json:"accounts,omitempty"`
	Errors       []ErrorCode     `json:"errors,omitempty"`
	Events       []Event         `json:"events,omitempty"`
	Types        IdlTypeDefSlice `json:"types,omitempty"`
	Constants    []Constant      `json:"constants,omitempty"`
}

type Discriminator [8]byte

type Constant struct {
	Name  string
	Type  IdlType
	Value string
}

type IdlState struct {
	Struct  IdlTypeDef       `json:"struct"`
	Methods []IdlStateMethod `json:"methods"`
}

type IdlStateMethod = Instruction

type IdlField struct {
	Name string  `json:"name"`
	Type IdlType `json:"type"`
}

type IdlStructFieldSlice []IdlField

type IdlEnumVariantSlice []IdlEnumVariant

func (slice IdlEnumVariantSlice) IsAllUint8() bool {
	for _, elem := range slice {
		if !elem.IsUint8() {
			return false
		}
	}
	return true
}

func (slice IdlEnumVariantSlice) IsSimpleEnum() bool {
	return slice.IsAllUint8()
}

func (slice IdlEnumVariantSlice) GetEnumVariantTypeName() []string {
	var result []string
	for _, variant := range slice {
		result = append(result, variant.Name)

	}
	return result
}

type IdlTypeDefStruct = []IdlField

type IdlEnumVariant struct {
	Name   string         `json:"name"`
	Docs   []string       `json:"docs"` // @custom
	Fields *IdlEnumFields `json:"fields,omitempty"`
}

func (variant *IdlEnumVariant) IsUint8() bool {
	// it's a simple uint8 if there is no fields data
	return variant.Fields == nil
}

// TODO
// type IdlEnumFields = IdlEnumFieldsNamed | IdlEnumFieldsTuple;
type IdlEnumFields struct {
	IdlEnumFieldsNamed *IdlEnumFieldsNamed
	IdlEnumFieldsTuple *IdlEnumFieldsTuple
}

type IdlEnumFieldsNamed []IdlField

type IdlEnumFieldsTuple []IdlType

// TODO: verify with examples
//func (env *IdlEnumFields) UnmarshalJSON(data []byte) error {
//	var tmp interface{}
//	if err := json.Unmarshal(data, &tmp); err != nil {
//		return err
//	}
//
//	if tmp == nil {
//		return fmt.Errorf("envelope is nil: %v", env)
//	}
//
//	fields, ok := tmp.([]interface{})
//	if !ok {
//		return fmt.Errorf("fields must be a slice")
//	}
//
//	if len(fields) == 0 {
//		return nil
//	}
//	if m, ok := fields[0].(map[string]interface{}); ok && m["name"] != nil {
//		// If has `name` field, then it's most likely a IdlEnumFieldsNamed.
//		if err := TranscodeJSON(tmp, &env.IdlEnumFieldsNamed); err != nil {
//			return err
//		}
//	} else {
//		if err := TranscodeJSON(tmp, &env.IdlEnumFieldsTuple); err != nil {
//			return err
//		}
//	}
//
//	// panic(Sf("what is this?:\n%s", spew.Sdump(temp)))
//
//	return nil
//}
