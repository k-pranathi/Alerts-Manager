# Alerts-Manager

Alert Manager maintains and updates the Alerts data using Alerts api built with in-memory database

# Steps to run the application

1. Clone the repoistory
2. Download all the dependencies
3. Open terminal and make sure you are in following directory,
    ```<local-path>/Alerts-Manager```
4. Run the following command in terminal to run the application,
    ```bash
    go run cmd/main.go
    ```

## API Spec
```
# To save the alert
POST /alerts
{
   "alert_id": "b950482e9911ec7e41f7ca5e5d9a424f",
   "service_id": "my_test_service_id",
   "service_name": "my_test_service",
   "model": "my_test_model",
   "alert_type": "anomaly",
   "alert_ts": "1695644160",
   "severity": "warning",
   "team_slack": "slack_ch"
}

# To get Alerts

GET /alerts
Supported filters as query params:
service_id
start_ts
end_ts
```