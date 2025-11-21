# Fee System Documentation

## Overview

The Vultisig fee system is a comprehensive solution for collecting usage fees from app users and managing treasury operations. It supports multiple fee structures, automatic token conversions, and transparent revenue distribution while maintaining the security guarantees of the TSS architecture.

## Fee System Architecture

```
┌─────────────────┐    ┌─────────────────┐
│      Fees       │    │   Verifier      │
│                 │    │                 │
│ • Fee           │◄──►│ • Policy        │
│   calculation   │    │   validation    │
│ • Collection    │    │ • TSS signing   │
│ • Conversion    │    │                 │
│ • Distribution  │    └─────────────────┘
└─────────────────┘
        │
┌─────────────────┐
│  DEX Aggregator │
│                 │
│ • Token swaps   │
│ • Best prices   │
│ • Slippage      │
│   protection    │
└─────────────────┘
        │
┌─────────────────┐
│ Vultisig        │
│ Treasury        │
│                 │
│ • USDC storage  │
│ • Revenue       │
│   tracking      │
└─────────────────┘
```

## Fee Types and Structures

### 1. Per-Transaction Fees

Charged for each transaction executed by an app.

```json
{
  "fee_type": "per_transaction",
  "amount": "1000000",
  "denomination": "usdc",
  "collection_frequency": "immediate",
  "description": "Fee charged per executed transaction"
}
```

### 2. Subscription Fees

Fixed recurring fees charged at regular intervals.

```json
{
  "fee_type": "subscription",
  "amount": "5000000",
  "denomination": "usdc",
  "collection_frequency": "monthly",
  "billing_period": "30d",
  "description": "Monthly subscription fee"
}
```

### 3. Per installation

Fee for installation in app store

```json
{
  "fee_type": "per_transaction",
  "amount": "1000000",
  "denomination": "usdc",
  "collection_frequency": "immediate",
  "description": "Fee charged per installation"
}
```

## Fee Collection Mechanisms

- Immediate Collection (Fees are collected immediately when transactions are executed).
- Deferred Collection (Fees are accumulated and collected in batches)

## Token Conversion System

The fee system automatically converts various tokens to USDC for treasury management.
