# LMWN-GO-TEST
Exercise for Go developer test LMWN




37. Please create a simple go project that satisfy the following requirements*
You're asked to write a simple JSON API to summarize COVID-19 stats using this public API, https://static.wongnai.com/devinterview/covid-cases.json.

- Your project must use Go, Go module, and Gin framework
- You create a JSON API at this endpoint /covid/summary
- The JSON API must count number of cases by provinces and age group
- There are 3 age groups: 0-30, 31-60, and 60+ if the case has no age data, please count as "N/A" group
- We encourage you to write tests, which we will give you some extra score
- Please zip the whole project and upload to the form.

--- sample response --

```json
{
  "Province": {
    "Samut Sakhon": 3613,
    "Bangkok": 2774
  },
  "AgeGroup": {
    "0-30": 300,
    "31-60": 150,
    "61+": 250,
    "N/A": 4
  }
}
```
