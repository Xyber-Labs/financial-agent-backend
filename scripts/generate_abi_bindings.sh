#!/bin/bash

# This script generates Ethereum-compatible golang bindigns for the smart abis

usage() {
	echo "Usage: $0 --inputdir|-i <inputdir> [--pkg|-p <pkgname>] [--out|-o <outdir>]
  Generates contract bindings for each *.abi file in the provided directory

  Arguments:
    * --inputdir|-i: Directory containing *.abi files
    * --pkg|-p: Package name for the generated bindings (default: contracts)
    * --out|-o: Output directory for the generated bindings (default: contrib/contracts)
"
	exit 1
}

ABIGEN=abigen
inputdir="./core/abi/jsons"
pkgname="bindings"
outdir="./core/abi/bindings"

# Parse arguments
while [[ $# -gt 0 ]]; do
	case "$1" in
	--inputdir | -i)
		if [[ $# -lt 2 ]]; then
			echo "Error: Missing argument for $1"
			usage
		fi
		inputdir="$2"
		shift 2
		;;
	--pkg | -p)
		if [[ $# -lt 2 ]]; then
			echo "Error: Missing argument for $1"
			usage
		fi
		pkgname="$2"
		shift 2
		;;
	--out | -o)
		if [[ $# -lt 2 ]]; then
			echo "Error: Missing argument for $1"
			usage
		fi
		outdir="$2"
		shift 2
		;;
	*)
		echo "Error: Invalid argument: $1"
		usage
		;;
	esac
done

if [[ -z "$inputdir" ]]; then
  echo "Error: Missing input directory"
  usage
fi

echo "pkgname: $pkgname"
echo "outdir: $outdir"
echo "inputdir: $inputdir"

abiFiles=$(find "$inputdir" -name "*.abi")

echo "abiFiles: $abiFiles"

for abiFile in $abiFiles; do
  echo "Processing $abiFile"
  abiName="$(basename "$abiFile")"
  echo "abiName: $abiName"
  contractName="${abiName%.abi}"
  echo "contractName: $contractName"

  # Parse providerName from directory name
  dirName="$(dirname "$abiFile")"
  echo "dirName: $dirName"
  providerName=${dirName##*contracts/}
  echo "providerName: $providerName"

  bytecodeFile="$dirName/$contractName.bin"
  echo "bytecodeFile: $bytecodeFile"

  # Create the output directory if it doesn't exist
  outProviderPath="$outdir/$contractName"
  echo "outProviderPath: $outProviderPath"
  mkdir -p "$outProviderPath"

  if [ -r "$bytecodeFile" ]; then
    echo "Generating bindings with bytecode file: $bytecodeFile"
    echo $ABIGEN --abi "$abiFile" --bin "$bytecodeFile" --pkg "$contractName" --type "$contractName" --out "$outProviderPath/$contractName.go"
    $ABIGEN --abi "$abiFile" --bin "$bytecodeFile" --pkg "$contractName" --type "$contractName" --out "$outProviderPath/$contractName.go"
    continue
  else
    echo "Generating bindings without bytecode file"
    echo $ABIGEN --abi "$abiFile" --pkg "$contractName" --type "$contractName" --out "$outProviderPath/$contractName.go"

    # Generate golang contract bindings for the current contracts file
    $ABIGEN --abi "$abiFile" --pkg "$contractName" --type "$contractName" --out "$outProviderPath/$contractName.go"
  fi
done
