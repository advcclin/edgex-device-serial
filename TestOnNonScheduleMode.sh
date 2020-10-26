#!/bin/sh

if [ $# != 3 ];then
	echo "usage: $0 {edgex_ip} {times} {interval}"
	exit;
fi

EDGEX_IP=$1
TIMES=$2
INTERVAL=$3

for i in $(seq 1 $TIMES); do
    printf "Creating Event($i) : "
    curl http://${EDGEX_IP}:49990/api/v1/device/name/Serial-Integer-Generator01/GenerateSerialValue -X GET
    sleep ${INTERVAL};
done
