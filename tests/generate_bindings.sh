#!/bin/bash

foundry_out="./contracts-source/out"

# Check that the contracts are compiled (out dir exists in foundry project)
if [ ! -d "$foundry_out" ]; then
  echo "Error: Contracts are not compiled. Compile contracts first"
  exit 1
fi

# Create temporary directory to store intermediate files
temp_dir="_out"
if [ ! -d "$temp_dir" ]; then
  mkdir -p "$temp_dir"
fi
# Cleanup temporary directory on exit
trap "rm -rf $temp_dir" EXIT

# Function to parse foundry output to output expected by abigen
parse_foundry_output() {
  local contract_name=$1
  
  if [ -z "$contract_name" ]; then
    echo "Error: Contract name not provided"
    exit 1
  fi
  
  local artifacts_file="$foundry_out/${contract_name}.sol/${contract_name}.json"
  
  if [ ! -f "$artifacts_file" ]; then
    echo "Error: Artifacts not found for $contract_name with path $artifacts_file"
    exit 1
  fi
  
  # Extract the ABI and bytecode from the contract JSON
  mkdir -p "$temp_dir"
  jq -r '.abi' "$artifacts_file" > "${temp_dir}/${contract_name}.abi"
  jq -r '.bytecode.object' "$artifacts_file" > "${temp_dir}/${contract_name}.bin"
}

parse_foundry_output "MockTEEWallet"
parse_foundry_output "MockAavePool"
parse_foundry_output "MockERC20"
parse_foundry_output "MockTrustManagementRouter"

ABIGEN=abigen

# Generate bindings with abigen
generate_bindings() {
  local contract_name=$1
  local pkg_name="contracts"
  local output_dir="./contracts"
  mkdir -p "$output_dir"
  
  if [ -z "$contract_name" ]; then
    echo "Error: Contract name name not provided"
    exit 1
  fi

  # Check that info is properly extracted
  if [ ! -f "${temp_dir}/${contract_name}.abi" ]; then
    echo "Error: ABI file not found for $contract_name"
    exit 1
  fi
  
  if [ ! -f "${temp_dir}/${contract_name}.bin" ]; then
    echo "Error: Bytecode file not found for $contract_name"
    exit 1
  fi
  
  echo "Generating bindings for $contract_name"
  $ABIGEN --abi "${temp_dir}/${contract_name}.abi" \
          --bin "${temp_dir}/${contract_name}.bin" \
          --pkg "$pkg_name" \
          --type "$contract_name" \
          --out "${output_dir}/${contract_name}.go"
}

generate_bindings "MockTEEWallet"
generate_bindings "MockERC20"
generate_bindings "MockAavePool"
generate_bindings "MockTrustManagementRouter"