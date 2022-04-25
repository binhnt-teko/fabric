OUTDIR=/Users/nguyenthanhbinh/Work/blockchain/nhs/network/dockers
build-peer:
	GO111MODULE=on GOOS=linux  go build -o peer cmd/peer/main.go
	mv peer ${OUTDIR}/node/bin/