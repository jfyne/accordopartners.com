#!/usr/bin/env bash
SECRETS="sa.json
"

for f in $SECRETS;
do
    gcloud kms encrypt --ciphertext-file="$f".enc --plaintext-file="$f" --location=global --keyring=deployment --key=accordo;
done
