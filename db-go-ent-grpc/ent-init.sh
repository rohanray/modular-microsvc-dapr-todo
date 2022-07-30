#!/bin/bash
echo "Initializing ent schema for $1"
go run entgo.io/ent/cmd/ent init $1
