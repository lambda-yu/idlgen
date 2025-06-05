package anchor

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"testing"
)

func Test_UnmarshalJSON(t *testing.T) {
	//data1 := "{\"type\": \"123\"}"
	data2 := `"types": [
    {
      "name": "SetCollectionSizeArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "size",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "CreateMasterEditionArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "maxSupply",
            "type": {
              "option": "u64"
            }
          }
        ]
      }
    },
    {
      "name": "MintNewEditionFromMasterEditionViaTokenArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "edition",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "TransferOutOfEscrowArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "amount",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "CreateMetadataAccountArgsV3",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "data",
            "type": {
              "defined": "DataV2"
            }
          },
          {
            "name": "isMutable",
            "type": "bool"
          },
          {
            "name": "collectionDetails",
            "type": {
              "option": {
                "defined": "CollectionDetails"
              }
            }
          }
        ]
      }
    },
    {
      "name": "UpdateMetadataAccountArgsV2",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "data",
            "type": {
              "option": {
                "defined": "DataV2"
              }
            }
          },
          {
            "name": "updateAuthority",
            "type": {
              "option": "publicKey"
            }
          },
          {
            "name": "primarySaleHappened",
            "type": {
              "option": "bool"
            }
          },
          {
            "name": "isMutable",
            "type": {
              "option": "bool"
            }
          }
        ]
      }
    },
    {
      "name": "ApproveUseAuthorityArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "numberOfUses",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "UtilizeArgs",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "numberOfUses",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "AuthorizationData",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "payload",
            "type": {
              "defined": "Payload"
            }
          }
        ]
      }
    },
    {
      "name": "AssetData",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "symbol",
            "type": "string"
          },
          {
            "name": "uri",
            "type": "string"
          },
          {
            "name": "sellerFeeBasisPoints",
            "type": "u16"
          },
          {
            "name": "creators",
            "type": {
              "option": {
                "vec": {
                  "defined": "Creator"
                }
              }
            }
          },
          {
            "name": "primarySaleHappened",
            "type": "bool"
          },
          {
            "name": "isMutable",
            "type": "bool"
          },
          {
            "name": "tokenStandard",
            "type": {
              "defined": "TokenStandard"
            }
          },
          {
            "name": "collection",
            "type": {
              "option": {
                "defined": "Collection"
              }
            }
          },
          {
            "name": "uses",
            "type": {
              "option": {
                "defined": "Uses"
              }
            }
          },
          {
            "name": "collectionDetails",
            "type": {
              "option": {
                "defined": "CollectionDetails"
              }
            }
          },
          {
            "name": "ruleSet",
            "type": {
              "option": "publicKey"
            }
          }
        ]
      }
    },
    {
      "name": "Collection",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "verified",
            "type": "bool"
          },
          {
            "name": "key",
            "type": "publicKey"
          }
        ]
      }
    },
    {
      "name": "Creator",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "address",
            "type": "publicKey"
          },
          {
            "name": "verified",
            "type": "bool"
          },
          {
            "name": "share",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "Data",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "symbol",
            "type": "string"
          },
          {
            "name": "uri",
            "type": "string"
          },
          {
            "name": "sellerFeeBasisPoints",
            "type": "u16"
          },
          {
            "name": "creators",
            "type": {
              "option": {
                "vec": {
                  "defined": "Creator"
                }
              }
            }
          }
        ]
      }
    },
    {
      "name": "DataV2",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "name",
            "type": "string"
          },
          {
            "name": "symbol",
            "type": "string"
          },
          {
            "name": "uri",
            "type": "string"
          },
          {
            "name": "sellerFeeBasisPoints",
            "type": "u16"
          },
          {
            "name": "creators",
            "type": {
              "option": {
                "vec": {
                  "defined": "Creator"
                }
              }
            }
          },
          {
            "name": "collection",
            "type": {
              "option": {
                "defined": "Collection"
              }
            }
          },
          {
            "name": "uses",
            "type": {
              "option": {
                "defined": "Uses"
              }
            }
          }
        ]
      }
    },
    {
      "name": "Reservation",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "address",
            "type": "publicKey"
          },
          {
            "name": "spotsRemaining",
            "type": "u64"
          },
          {
            "name": "totalSpots",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "ReservationV1",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "address",
            "type": "publicKey"
          },
          {
            "name": "spotsRemaining",
            "type": "u8"
          },
          {
            "name": "totalSpots",
            "type": "u8"
          }
        ]
      }
    },
    {
      "name": "SeedsVec",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "seeds",
            "type": {
              "vec": "bytes"
            }
          }
        ]
      }
    },
    {
      "name": "ProofInfo",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "proof",
            "type": {
              "vec": {
                "array": [
                  "u8",
                  32
                ]
              }
            }
          }
        ]
      }
    },
    {
      "name": "Payload",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "map",
            "type": {
              "hashMap": [
                "string",
                {
                  "defined": "PayloadType"
                }
              ]
            }
          }
        ]
      }
    },
    {
      "name": "Uses",
      "type": {
        "kind": "struct",
        "fields": [
          {
            "name": "useMethod",
            "type": {
              "defined": "UseMethod"
            }
          },
          {
            "name": "remaining",
            "type": "u64"
          },
          {
            "name": "total",
            "type": "u64"
          }
        ]
      }
    },
    {
      "name": "BurnArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              }
            ]
          }
        ]
      }
    },
    {
      "name": "DelegateArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "CollectionV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "SaleV1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "TransferV1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "DataV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "UtilityV1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "StakingV1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "StandardV1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              }
            ]
          },
          {
            "name": "LockedTransferV1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "locked_address",
                "type": "publicKey"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "ProgrammableConfigV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AuthorityItemV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "DataItemV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "CollectionItemV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "ProgrammableConfigItemV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "PrintDelegateV1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "RevokeArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "CollectionV1"
          },
          {
            "name": "SaleV1"
          },
          {
            "name": "TransferV1"
          },
          {
            "name": "DataV1"
          },
          {
            "name": "UtilityV1"
          },
          {
            "name": "StakingV1"
          },
          {
            "name": "StandardV1"
          },
          {
            "name": "LockedTransferV1"
          },
          {
            "name": "ProgrammableConfigV1"
          },
          {
            "name": "MigrationV1"
          },
          {
            "name": "AuthorityItemV1"
          },
          {
            "name": "DataItemV1"
          },
          {
            "name": "CollectionItemV1"
          },
          {
            "name": "ProgrammableConfigItemV1"
          },
          {
            "name": "PrintDelegateV1"
          }
        ]
      }
    },
    {
      "name": "MetadataDelegateRole",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "AuthorityItem"
          },
          {
            "name": "Collection"
          },
          {
            "name": "Use"
          },
          {
            "name": "Data"
          },
          {
            "name": "ProgrammableConfig"
          },
          {
            "name": "DataItem"
          },
          {
            "name": "CollectionItem"
          },
          {
            "name": "ProgrammableConfigItem"
          }
        ]
      }
    },
    {
      "name": "HolderDelegateRole",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "PrintDelegate"
          }
        ]
      }
    },
    {
      "name": "CreateArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "asset_data",
                "type": {
                  "defined": "AssetData"
                }
              },
              {
                "name": "decimals",
                "type": {
                  "option": "u8"
                }
              },
              {
                "name": "print_supply",
                "type": {
                  "option": {
                    "defined": "PrintSupply"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "MintArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "TransferArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "amount",
                "type": "u64"
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "UpdateArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "new_update_authority",
                "type": {
                  "option": "publicKey"
                }
              },
              {
                "name": "data",
                "type": {
                  "option": {
                    "defined": "Data"
                  }
                }
              },
              {
                "name": "primary_sale_happened",
                "type": {
                  "option": "bool"
                }
              },
              {
                "name": "is_mutable",
                "type": {
                  "option": "bool"
                }
              },
              {
                "name": "collection",
                "type": {
                  "defined": "CollectionToggle"
                }
              },
              {
                "name": "collection_details",
                "type": {
                  "defined": "CollectionDetailsToggle"
                }
              },
              {
                "name": "uses",
                "type": {
                  "defined": "UsesToggle"
                }
              },
              {
                "name": "rule_set",
                "type": {
                  "defined": "RuleSetToggle"
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsUpdateAuthorityV2",
            "fields": [
              {
                "name": "new_update_authority",
                "type": {
                  "option": "publicKey"
                }
              },
              {
                "name": "data",
                "type": {
                  "option": {
                    "defined": "Data"
                  }
                }
              },
              {
                "name": "primary_sale_happened",
                "type": {
                  "option": "bool"
                }
              },
              {
                "name": "is_mutable",
                "type": {
                  "option": "bool"
                }
              },
              {
                "name": "collection",
                "type": {
                  "defined": "CollectionToggle"
                }
              },
              {
                "name": "collection_details",
                "type": {
                  "defined": "CollectionDetailsToggle"
                }
              },
              {
                "name": "uses",
                "type": {
                  "defined": "UsesToggle"
                }
              },
              {
                "name": "rule_set",
                "type": {
                  "defined": "RuleSetToggle"
                }
              },
              {
                "name": "token_standard",
                "type": {
                  "option": {
                    "defined": "TokenStandard"
                  }
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsAuthorityItemDelegateV2",
            "fields": [
              {
                "name": "new_update_authority",
                "type": {
                  "option": "publicKey"
                }
              },
              {
                "name": "primary_sale_happened",
                "type": {
                  "option": "bool"
                }
              },
              {
                "name": "is_mutable",
                "type": {
                  "option": "bool"
                }
              },
              {
                "name": "token_standard",
                "type": {
                  "option": {
                    "defined": "TokenStandard"
                  }
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsCollectionDelegateV2",
            "fields": [
              {
                "name": "collection",
                "type": {
                  "defined": "CollectionToggle"
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsDataDelegateV2",
            "fields": [
              {
                "name": "data",
                "type": {
                  "option": {
                    "defined": "Data"
                  }
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsProgrammableConfigDelegateV2",
            "fields": [
              {
                "name": "rule_set",
                "type": {
                  "defined": "RuleSetToggle"
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsDataItemDelegateV2",
            "fields": [
              {
                "name": "data",
                "type": {
                  "option": {
                    "defined": "Data"
                  }
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsCollectionItemDelegateV2",
            "fields": [
              {
                "name": "collection",
                "type": {
                  "defined": "CollectionToggle"
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          },
          {
            "name": "AsProgrammableConfigItemDelegateV2",
            "fields": [
              {
                "name": "rule_set",
                "type": {
                  "defined": "RuleSetToggle"
                }
              },
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "CollectionToggle",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Clear"
          },
          {
            "name": "Set",
            "fields": [
              {
                "defined": "Collection"
              }
            ]
          }
        ]
      }
    },
    {
      "name": "UsesToggle",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Clear"
          },
          {
            "name": "Set",
            "fields": [
              {
                "defined": "Uses"
              }
            ]
          }
        ]
      }
    },
    {
      "name": "CollectionDetailsToggle",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Clear"
          },
          {
            "name": "Set",
            "fields": [
              {
                "defined": "CollectionDetails"
              }
            ]
          }
        ]
      }
    },
    {
      "name": "RuleSetToggle",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Clear"
          },
          {
            "name": "Set",
            "fields": [
              "publicKey"
            ]
          }
        ]
      }
    },
    {
      "name": "PrintArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "edition",
                "type": "u64"
              }
            ]
          },
          {
            "name": "V2",
            "fields": [
              {
                "name": "edition",
                "type": "u64"
              }
            ]
          }
        ]
      }
    },
    {
      "name": "LockArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "UnlockArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "UseArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "authorization_data",
                "type": {
                  "option": {
                    "defined": "AuthorizationData"
                  }
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "VerificationArgs",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "CreatorV1"
          },
          {
            "name": "CollectionV1"
          }
        ]
      }
    },
    {
      "name": "TokenStandard",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "NonFungible"
          },
          {
            "name": "FungibleAsset"
          },
          {
            "name": "Fungible"
          },
          {
            "name": "NonFungibleEdition"
          },
          {
            "name": "ProgrammableNonFungible"
          },
          {
            "name": "ProgrammableNonFungibleEdition"
          }
        ]
      }
    },
    {
      "name": "Key",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Uninitialized"
          },
          {
            "name": "EditionV1"
          },
          {
            "name": "MasterEditionV1"
          },
          {
            "name": "ReservationListV1"
          },
          {
            "name": "MetadataV1"
          },
          {
            "name": "ReservationListV2"
          },
          {
            "name": "MasterEditionV2"
          },
          {
            "name": "EditionMarker"
          },
          {
            "name": "UseAuthorityRecord"
          },
          {
            "name": "CollectionAuthorityRecord"
          },
          {
            "name": "TokenOwnedEscrow"
          },
          {
            "name": "TokenRecord"
          },
          {
            "name": "MetadataDelegate"
          },
          {
            "name": "EditionMarkerV2"
          },
          {
            "name": "HolderDelegate"
          }
        ]
      }
    },
    {
      "name": "CollectionDetails",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "size",
                "type": "u64"
              }
            ]
          },
          {
            "name": "V2",
            "fields": [
              {
                "name": "padding",
                "type": {
                  "array": [
                    "u8",
                    8
                  ]
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "EscrowAuthority",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "TokenOwner"
          },
          {
            "name": "Creator",
            "fields": [
              "publicKey"
            ]
          }
        ]
      }
    },
    {
      "name": "PrintSupply",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Zero"
          },
          {
            "name": "Limited",
            "fields": [
              "u64"
            ]
          },
          {
            "name": "Unlimited"
          }
        ]
      }
    },
    {
      "name": "ProgrammableConfig",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "V1",
            "fields": [
              {
                "name": "rule_set",
                "type": {
                  "option": "publicKey"
                }
              }
            ]
          }
        ]
      }
    },
    {
      "name": "MigrationType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "CollectionV1"
          },
          {
            "name": "ProgrammableV1"
          }
        ]
      }
    },
    {
      "name": "TokenState",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Unlocked"
          },
          {
            "name": "Locked"
          },
          {
            "name": "Listed"
          }
        ]
      }
    },
    {
      "name": "TokenDelegateRole",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Sale"
          },
          {
            "name": "Transfer"
          },
          {
            "name": "Utility"
          },
          {
            "name": "Staking"
          },
          {
            "name": "Standard"
          },
          {
            "name": "LockedTransfer"
          },
          {
            "name": "Migration"
          }
        ]
      }
    },
    {
      "name": "AuthorityType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "None"
          },
          {
            "name": "Metadata"
          },
          {
            "name": "Holder"
          },
          {
            "name": "MetadataDelegate"
          },
          {
            "name": "TokenDelegate"
          }
        ]
      }
    },
    {
      "name": "PayloadKey",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Amount"
          },
          {
            "name": "Authority"
          },
          {
            "name": "AuthoritySeeds"
          },
          {
            "name": "Delegate"
          },
          {
            "name": "DelegateSeeds"
          },
          {
            "name": "Destination"
          },
          {
            "name": "DestinationSeeds"
          },
          {
            "name": "Holder"
          },
          {
            "name": "Source"
          },
          {
            "name": "SourceSeeds"
          }
        ]
      }
    },
    {
      "name": "PayloadType",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Pubkey",
            "fields": [
              "publicKey"
            ]
          },
          {
            "name": "Seeds",
            "fields": [
              {
                "defined": "SeedsVec"
              }
            ]
          },
          {
            "name": "MerkleProof",
            "fields": [
              {
                "defined": "ProofInfo"
              }
            ]
          },
          {
            "name": "Number",
            "fields": [
              "u64"
            ]
          }
        ]
      }
    },
    {
      "name": "UseMethod",
      "type": {
        "kind": "enum",
        "variants": [
          {
            "name": "Burn"
          },
          {
            "name": "Multiple"
          },
          {
            "name": "Single"
          }
        ]
      }
    }
  ]`

	type TestData struct {
		Data IdlType `json:"type"`
	}

	var raw TestData
	if err := json.Unmarshal([]byte(data2), &raw); err == nil {
		spew.Dump(raw)
	} else {
		spew.Dump(err)
	}

}
