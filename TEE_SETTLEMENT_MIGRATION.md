# TEE Settlement System Migration Documentation

## Overview
This document describes the migration from ZK proof-based settlement to TEE signer-based settlement for the 0g-serving-broker inference system.

## Background
The original settlement system used ZK proofs to batch and verify multiple inference requests. This approach, while cryptographically secure, was computationally expensive and complex. The new system leverages TEE (Trusted Execution Environment) signatures for a more efficient settlement process.

## Architecture Changes

### 1. Smart Contract Modifications

#### Account Structure Enhancement (`InferenceAccount.sol`)
```solidity
struct Account {
    address user;
    address provider;
    uint nonce;
    uint balance;
    uint pendingRefund;
    uint[2] signer;           // User's BabyJub public key
    Refund[] refunds;
    string additionalInfo;
    uint[2] providerPubKey;    // Provider's BabyJub public key (for ZK)
    address teeSignerAddress;  // NEW: Provider's TEE ECDSA signer address
}
```

#### New Contract Functions (`InferenceServing.sol`)

1. **acknowledgeTEESigner**: Allows users to acknowledge a provider's TEE signer address
```solidity
function acknowledgeTEESigner(address provider, address teeSignerAddress) external
```

2. **settleFeesWithTEE**: New settlement function using TEE signatures
```solidity
function settleFeesWithTEE(TEESettlementData[] calldata settlements) external
```

### 2. Go Backend Implementation

#### New Settlement Controller (`settlement_tee.go`)
- **SettleFeesWithTEE**: Main settlement function that:
  - Aggregates fees by user
  - Signs settlement data with TEE private key
  - Submits signed data to smart contract

#### TEE Service Enhancement (`tee.go`)
- Added `Sign` method for message signing with TEE private key
```go
func (s *TeeService) Sign(messageHash []byte) ([]byte, error)
```

#### Contract Interface Update (`request.go`)
- Added `SettleFeesWithTEE` method to interact with the new contract function
- Defined `TEESettlementData` structure for settlement transactions

## Settlement Flow

### Old Flow (ZK Proof)
1. Collect user requests with signatures
2. Generate ZK proof for batch of requests
3. Submit proof to contract for verification
4. Contract verifies ZK proof and settles fees

### New Flow (TEE Signer)
1. Provider generates ECDSA keypair in TEE environment
2. Users acknowledge provider's TEE signer address via `acknowledgeTEESigner`
3. Provider collects requests and aggregates fees by user
4. TEE signs settlement data (user, totalFee, minNonce, maxNonce)
5. Provider submits signed settlements to contract
6. Contract verifies TEE signature against acknowledged address
7. Contract updates nonces and transfers fees

## Security Considerations

### Authentication
- TEE signer address must be acknowledged by users before settlement
- Each (user, provider) pair maintains its own TEE signer acknowledgment

### Replay Protection
- Nonce mechanism ensures each settlement can only be processed once
- minNonce must be greater than the last recorded nonce in the contract

### Trust Model
- Users trust the TEE environment to correctly calculate and sign fees
- Provider cannot forge signatures without access to TEE private key
- TEE attestation provides proof of secure execution environment

## Data Structures

### TEESettlementData
```solidity
struct TEESettlementData {
    address user;           // User's address
    address provider;       // Provider's address
    uint256 totalFee;       // Total fee to settle
    uint256 minNonce;       // Minimum nonce in this batch
    uint256 maxNonce;       // Maximum nonce in this batch
    uint256 timestamp;      // Settlement timestamp
    bytes signature;        // TEE signature
}
```

### Signature Format
The TEE signs a message containing:
```
keccak256(abi.encodePacked(
    user,
    provider,
    totalFee,
    minNonce,
    maxNonce,
    timestamp
))
```

## Migration Steps

### For Providers
1. Generate TEE signer keypair during initialization
2. Store TEE signer address for users to acknowledge
3. Update settlement logic to use `SettleFeesWithTEE`

### For Users
1. Call `acknowledgeTEESigner` with provider's TEE address
2. Continue using the service as normal

### For Contract Deployment
1. Deploy updated `InferenceAccount.sol` with new Account structure
2. Deploy updated `InferenceServing.sol` with TEE settlement functions
3. Existing accounts will have `teeSignerAddress` as zero address until acknowledged

## Benefits of TEE Settlement

1. **Performance**: Eliminates expensive ZK proof generation
2. **Simplicity**: Direct signature verification vs complex proof circuits
3. **Flexibility**: Easier to modify settlement logic without changing ZK circuits
4. **Cost**: Lower gas costs for settlement transactions
5. **Scalability**: Can process larger batches without proof size limitations

## Compatibility

- Original ZK settlement functions remain in place for backward compatibility
- Providers can choose between ZK and TEE settlement methods
- Users must explicitly acknowledge TEE signers to enable TEE settlement

## Testing Recommendations

1. Test TEE key generation and secure storage
2. Verify signature generation and validation
3. Test nonce progression and replay protection
4. Validate fee aggregation accuracy
5. Test edge cases (insufficient balance, invalid nonce, etc.)

## Future Enhancements

1. Multi-signature support for enhanced security
2. Batch acknowledgment of TEE signers
3. Automatic TEE signer rotation
4. Integration with remote attestation for TEE verification