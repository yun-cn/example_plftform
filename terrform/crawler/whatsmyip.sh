#!/bin/bash

# tries wireless and lan cable
ip=$(ipconfig getifaddr en0 || ipconfig getifaddr en1)

echo {\"ip\":\"$ip\"}
