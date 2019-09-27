#!/bin/bash
curl -X POST -d '{"state": "critical", "entity": "test", "name": "test", "data": "swag"}' http://localhost:8080/v1/triggers/default/91fb93bc-034e-443c-991d-969ff4be8ea5
