#!/usr/bin/env bash

index_name="$1"

echo "=== DELETE INDEX === -> ${index_name}_listings"
curl -sS -XDELETE "http://localhost:9200/${index_name}_listings"

echo "=== DELETE INDEX === -> ${index_name}_search_requests"
curl -sS -XDELETE "http://localhost:9200/${index_name}_search_requests"
echo ''

echo "=== DELETE INDEX === -> ${index_name}_place_suggestions"
curl -sS -XDELETE "http://localhost:9200/${index_name}_place_suggestions"
echo ''

echo "=== DELETE INDEX === -> ${index_name}_users"
curl -sS -XDELETE "http://localhost:9200/${index_name}_users"
echo ''

echo "=== DELETE INDEX === -> ${index_name}_reviews"
curl -sS -XDELETE "http://localhost:9200/${index_name}_reviews"
echo ''
