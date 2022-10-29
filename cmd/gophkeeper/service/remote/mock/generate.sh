#!/bin/bash

mockgen -source=../../interface/remote_use_interface.go -destination=external_interfaces.go -package mock
