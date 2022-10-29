#!/bin/bash

mockgen -source=../interface/data_interface.go -destination=external_interfaces.go -package mock
mockgen -source=../interface/remote_interface.go -destination=internal_interfaces.go -package mock
