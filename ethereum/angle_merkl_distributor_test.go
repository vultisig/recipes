package ethereum

import (
	"testing"
)

func TestAngleMerklDistributor_ValidateTransaction(t *testing.T) {
	validator := NewAngleMerklDistributor()

	tests := []struct {
		name          string
		functionName  string
		params        map[string]interface{}
		expectError   bool
		errorContains string
	}{
		// https://etherscan.io/tx/0x3f73880c0bc9305c33894f51f4924e10cb2945ff511186663245c74b1a007842
    //
		// 0	users	address[] 0x850CEC90edCA90981a3982bB6F06C16a59AB5F31
		// 1	tokens	address[] 0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984
		// 2	amounts	uint256[] 117876393971237861056
		// 3	proofs	bytes32[][]
		// 0x9ba5e970a488573eab9fb24a112f045892c89a4b826cd37421e661067ae8d3f1
		// 0x80fa8f86e4492d27373166c4e22b3600fcbdff653e879e4630c3a47f1feaa66f
		// 0xcabb1a7be16ca987d0bbfe4b8e5488259166aee4a07478a944bd9aabb3280c03
		// 0x772f0e91947e792b15d584e7806b47d0e4d5e7df4c0cd3a77eb870310b41973c
		// 0xa505ca405f289ef5d6a7ddf57b1dd9f4b14554716f954ebe20462eb642ad4147
		// 0xe14361fa0bfb37507337952f2c4e6cda16e5ecfc339a39a50a40b1766d4aa96f
		// 0x572521c14018ea070e2faff6037ac838a55b05337e6b359e265da6958ac92516
		// 0x1824179cada913f5ab6fd89a67637ff9d564d18aef38c1ce623a1df8f209ef6a
		// 0xdc066e78cf66ce967f8d08248d6a8aa4ef3f14b6d5ae3c74af39e9025d4b08f8
		// 0xee2bd03ae677967a0b6a5a74ba1760c907397f147ebaa05c5c93dadd75c9426f
		// 0x992a98acbd5b52f02a1225508f23ae84c2f5d566ad68bba17156cacf5727cedc
		// 0x8405374478954fe080eb9fa97c09ac903398d15a057043916e9444dda92350bb
		// 0xe0de761e9a5ab331fad1467b464a10df954b0b1fa4bc0a91d20ac3004a8a0361
		// 0xcf780902f3ddbbab9004f02f8221595e07047dd837fec043d0a3ff8ff82b6440
		// 0xad6fcd0b262be716f18b38f5eb7dd137e5ed7b388efe74518d963c7504bdb7d2
		// 0x2beb443c449573c8bf87731372c8de946a4b056fcc8111a9afd38f09a6e4fc6f
		// 0x1a662a6881b67db60bd4ab596cdf4c417e55953e56a239df60dcdeb1426c9f31
		// 0x7656b6e92c10b794bf8c934dc401c8efed3bfd38ea7b9f465b4a87fc0467d285
		{
			name:         "Valid claim",
			functionName: "claim",
			params: map[string]interface{}{
				"users": []string{"0x850CEC90edCA90981a3982bB6F06C16a59AB5F31"},
				"tokens": []string{"0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"},
				"amounts": []string{"117876393971237861056"},
				"proofs": []string{
					"0x9ba5e970a488573eab9fb24a112f045892c89a4b826cd37421e661067ae8d3f1",
					"0x80fa8f86e4492d27373166c4e22b3600fcbdff653e879e4630c3a47f1feaa66f",
					"0xcabb1a7be16ca987d0bbfe4b8e5488259166aee4a07478a944bd9aabb3280c03",
					"0x772f0e91947e792b15d584e7806b47d0e4d5e7df4c0cd3a77eb870310b41973c",
					"0xa505ca405f289ef5d6a7ddf57b1dd9f4b14554716f954ebe20462eb642ad4147",
					"0xe14361fa0bfb37507337952f2c4e6cda16e5ecfc339a39a50a40b1766d4aa96f",
					"0x572521c14018ea070e2faff6037ac838a55b05337e6b359e265da6958ac92516",
					"0x1824179cada913f5ab6fd89a67637ff9d564d18aef38c1ce623a1df8f209ef6a",
					"0xdc066e78cf66ce967f8d08248d6a8aa4ef3f14b6d5ae3c74af39e9025d4b08f8",
					"0xee2bd03ae677967a0b6a5a74ba1760c907397f147ebaa05c5c93dadd75c9426f",
					"0x992a98acbd5b52f02a1225508f23ae84c2f5d566ad68bba17156cacf5727cedc",
					"0x8405374478954fe080eb9fa97c09ac903398d15a057043916e9444dda92350bb",
					"0xe0de761e9a5ab331fad1467b464a10df954b0b1fa4bc0a91d20ac3004a8a0361",
					"0xcf780902f3ddbbab9004f02f8221595e07047dd837fec043d0a3ff8ff82b6440",
					"0xad6fcd0b262be716f18b38f5eb7dd137e5ed7b388efe74518d963c7504bdb7d2",
					"0x2beb443c449573c8bf87731372c8de946a4b056fcc8111a9afd38f09a6e4fc6f",
					"0x1a662a6881b67db60bd4ab596cdf4c417e55953e56a239df60dcdeb1426c9f31",
					"0x7656b6e92c10b794bf8c934dc401c8efed3bfd38ea7b9f465b4a87fc0467d285",
				},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator.ValidateTransaction(tt.functionName, tt.params)
			if tt.expectError {
				if err == nil {
					t.Errorf("Expected error but got none")
					return
				}
			}
			if err != nil {
				if tt.errorContains != "" && !contains(err.Error(), tt.errorContains) {
					t.Errorf("Expected error to contain '%s', but got: %s", tt.errorContains, err.Error())
				}
			}
		})
	}
}