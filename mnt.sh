#!/bin/bash

sudo mkdir drive
sudo mount -w -t ufs -o ufstype=ufs2 $1 drive
exit
