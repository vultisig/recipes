
# App Submission Process

## Overview

To launch your app in the Vultisig ecosystem, you must submit your app/agent for review and approval. After approval, developers are responsible for deploying and maintaining their own applications, ensuring uptime, scaling, and operational security outside of Vultisig core infrastructure.

**Note:** Verifier and Fees are managed exclusively by Vultisig and run as core infrastructure services. All other apps are deployed, operated, and updated directly by their respective developers.
***

## Official App Examples

Below are typical app types found in Vultisig’s ecosystem. Refer to them when designing your own submission:


| App ID            | Title         | Description | Category | Endpoint                                         |
|:------------------|:--------------| :-- |:---------|:-------------------------------------------------|
| vultisig-dca-0000 | Recurring buy | Dollar Cost Averaging automation | app      | https://apps.vultisig.com/apps/vultisig-dca-0000 |
TBA

The marketplace supports both standard apps with business logic or automation, and agent-based apps for advanced trading and verification tasks.

***

## Submission Requirements

Create a detailed `plugin-config.yaml` describing:

- ID, title, description, and endpoint (if applicable)
- Category (app or ai-agent)
- Supported blockchains
- Payment and resource requirements

Follow this example:
```
plugins:
  - id: vultisig-dca-0000
    title: DCA Plugin
    description: Dollar Cost Averaging automation plugin
    server_endpoint: https://dca.vultisigplugin.app
    category: plugin

  - id: vultisig-payroll-0000
    title: Payroll Plugin
    description: Automated payroll distribution plugin
    server_endpoint: https://plugin.vultisigplugin.app/payroll
    category: plugin

  - id: vultisig-fees-feee
    title: Fees Plugin
    description: Transaction fees management plugin
    server_endpoint: ""
    category: plugin

  - id: vultisig-copytrader-0000
    title: Copy Trader Plugin
    description: Copy trading automation plugin
    server_endpoint: http://149.28.158.120:8081
    category: ai-agent

  - id: nbits-labs-merkle-e93d
    title: Merkle Plugin
    description: Merkle tree verification plugin
    server_endpoint: ""
    category: plugin
```

***

## Submission Process

1. Prepare your submission package: YAML config and documentation.
2. Complete the security checklist and performance requirements for APIs and resource use.
3. Submit via GitHub and email, including all required files and a summary of your app’s purpose.

***

## Review and Approval

The multi-phase review process includes documentation checks, security and code audits, policy validation, performance testing, and final deployment steps. Use the example apps above as references for best practices and compliance.

***

## Publication and Marketplace Listing

Once approved, your app is provisioned and listed in the Vultisig marketplace alongside standard entries such as DCA and Fees apps. Those examples illustrate category standards, endpoint conventions, and API scope for production apps.

Revenue sharing models are available for fee, subscription, and premium features, negotiated during approval.

***

## Support

Questions? Reach out via email, documentation portal, or the developer Discord for guidance on submission and best practices.

***