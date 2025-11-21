# Vultisig System Components

## Services and Modules
![Vultisig ecosystem](general.png)

| Component | Role / Description                                                                                                                                                                                                                                                                 |
| :-- |:-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
| **Recipes** | Shared library and integration framework for standardized blockchain logic and protocol modules. Provides building blocks for apps and services to interact securely and consistently across supported blockchains.                                                                |
| **Fees** | Fee calculation and treasury logic; enforces fee-related rules and automates processing for transactions and treasury operations. Utilized by apps and Verifier policies.                                                                                                          |
| **Verifier** | Centralizes policy management. Receives transaction requests from apps, validates them against custom policies, and performs TSS-based signing for compliant transactions. Does NOT initiate or broadcast transactions; strictly enforces rules and maintains security boundaries. |
| **App A/B/...** | Applications that implement custom business logic or user-facing features. Compose and initiate transactions and handle broadcasting. Submit transactions to the Verifier for validation and signing, ensuring policy compliance.                                                  |
| **Blockchains** | Supported external networks such as Bitcoin, Ethereum, Solana, etc. Recipes standardizes their interfaces for system-wide compatibility.                                                                                                                                           |

## Key Interactions

- All applications interface with Recipes modules for standardized logic.
- Fee computations and treasury operations are funneled through the Fees module.
- Every transaction is routed to the Verifier for strict policy compliance and TSS signing.
- Blockchains are kept external, with all interaction managed through the Recipes abstraction layer.