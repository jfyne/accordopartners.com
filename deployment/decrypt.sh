#!/usr/bin/env bash

cd "$(dirname "$0")"

for f in $(find -type f | grep \\.enc);
do
    gcloud kms decrypt --ciphertext-file="$f" --plaintext-file="${f/%.enc/}" --location=global --keyring=deployment --key=accordo;
done
