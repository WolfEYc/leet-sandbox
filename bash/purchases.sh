#!/bin/bash

purchases_url="https://purchase-json-974308738028.us-central1.run.app/purchase?page=${page}&purchase_number=${purchase_number}"
curl ${purchases_url}
