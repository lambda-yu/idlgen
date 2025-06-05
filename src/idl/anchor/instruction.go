package anchor

import (
	"encoding/json"
	"fmt"
)

type Instruction struct {
	Name          string                    `json:"name"`
	Docs          []string                  `json:"docs,omitempty"`
	Discriminator Discriminator             `json:"discriminator"`
	Accounts      []*InstructionAccountItem `json:"accounts"`
	Args          []*IdlField               `json:"args"`
	Returns       IdlType                   `json:"returns,omitempty"`
}

type InstructionAccountItem struct {
	InstructionAccount  InstructionAccount  `json:"instructionAccount,omitempty"`
	InstructionAccounts InstructionAccounts `json:"instructionAccounts,omitempty"`
}

func (t *InstructionAccountItem) UnmarshalJSON(data []byte) error {
	// 解析为 map[string]any
	var insAccounts InstructionAccounts
	if err := json.Unmarshal(data, &insAccounts); err != nil {
		return fmt.Errorf("invalid InstructionAccounts: %w", err)
	}

	if insAccounts.Accounts == nil {
		var insAccount InstructionAccount
		if err := json.Unmarshal(data, &insAccount); err != nil {
			return fmt.Errorf("invalid InstructionAccount: %w", err)
		}
		t.InstructionAccount = insAccount
		return nil
	} else {
		t.InstructionAccounts = insAccounts
		return nil
	}

	return fmt.Errorf("unrecognized InstructionAccount: %s", string(data))
}

type InstructionAccount struct {
	Name      string   `json:"name"`
	Docs      []string `json:"docs,omitempty"`
	Writable  bool     `json:"writable,omitempty"`
	Signer    bool     `json:"signer,omitempty"`
	Optional  bool     `json:"optional,omitempty"`
	Address   string   `json:"address,omitempty"`
	PDA       PDA      `json:"pda,omitempty"`
	Relations []string `json:"relations,omitempty"`
}

type InstructionAccounts struct {
	Name     string               `json:"name"`
	Accounts []InstructionAccount `json:"accounts,omitempty"`
}

type PDA struct {
	Seeds   []Seed `json:"seeds"`
	Program Seed   `json:"program,omitempty"`
}

type Seed struct {
	SeedConst   *SeedConst   `json:"seed_const,omitempty"`
	SeedArg     *SeedArg     `json:"seed_arg,omitempty"`
	SeedAccount *SeedAccount `json:"seed_account,omitempty"`
}

func (t *Seed) UnmarshalJSON(data []byte) error {
	// 解析为 map[string]any
	var obj map[string]json.RawMessage
	if err := json.Unmarshal(data, &obj); err != nil {
		return fmt.Errorf("invalid Seed: %w", err)
	}
	kind, err := obj["kind"].MarshalJSON()
	if err != nil {
		return fmt.Errorf("invalid Seed kind: %w", err)
	}

	var kindStr string
	err = json.Unmarshal(kind, &kindStr)
	if err != nil {
		return fmt.Errorf("invalid Seed kind error: %w", err)
	}
	// 判断字段类型并递归解析
	switch kindStr {
	case "const":
		var seedConst SeedConst
		if err := json.Unmarshal(data, &seedConst); err != nil {
			return err
		}
		t.SeedConst = &seedConst
		return nil
	case "arg":
		var seedArg SeedArg
		if err := json.Unmarshal(data, &seedArg); err != nil {
			return err
		}
		t.SeedArg = &seedArg
		return nil
	case "account":
		var seedAccount SeedAccount
		if err := json.Unmarshal(data, &seedAccount); err != nil {
			return err
		}
		t.SeedAccount = &seedAccount
		return nil
	default:
		return fmt.Errorf("unknown Seed field")
	}

	return fmt.Errorf("unrecognized Seed: %s", string(data))
}

type SeedConst struct {
	Kind  string `json:"kind"`
	Value []int  `json:"value"`
}

type SeedArg struct {
	Kind string `json:"kind"`
	Path string `json:"path"`
}

type SeedAccount struct {
	Kind    string `json:"kind"`
	Path    string `json:"path"`
	Account string `json:"account,omitempty"`
}
