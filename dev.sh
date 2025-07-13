#!/bin/bash

trap "echo 'Stopping...'; kill 0" SIGINT

echo "Preparing backend environment"
cd backend
touch vouchers.db
go mod tidy

echo "Preparing frontend environment"
cd ../frontend
npm install 

cd ../
echo "Starting backend and frontend"
(cd backend && go run main.go) &
(cd frontend && npm run dev) &

wait