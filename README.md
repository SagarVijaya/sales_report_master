# sales_report_master

## Tech stack

- Language : Go 1.2+
- CSV : Encoding/csv
- Schedular : Fix the time to run the data refresh go routines

## External Package

- GIN
- Toml
- UUID

### clone URL

### **1️⃣ Install Dependencies**

```sh
go mod tidy
```

## sample Request

```sh
curl http://localhost:3011/topProduct/overallData?start_date=2016-01-01&end_date=2025-05-06

curl http://localhost:3011/topProduct/categoryData?start_date=2016-01-01&end_date=2025-05-06

curl http://localhost:3011/topProduct/regionData?start_date=2016-01-01&end_date=2025-05-06

curl http://localhost:3011/trigger


```

## sample Response 1
- Endpoint :  /topProduct/overallData?start_date=2016-01-01&end_date=2025-05-06

```json
{"data":{"end":"2025-05-06","overall_sold_quantity":10,"start":"2016-01-01"},"message":"Overall Sold Product","status":"S"}

```

### sample Response 2
- Endpoint :  /topProduct/categoryData?start_date=2016-01-01&end_date=2025-05-06

```json
{"data":{"end":"2025-05-06","overall_details":[{"quantity":3,"Category":"Clothing"},{"quantity":3,"Category":"Shoes"},{"quantity":4,"Category":"Electronics"}],"start":"2016-01-01"},"message":"Top Product For Category","status":"S"}

```

### sample Response 3
- Endpoint :  /topProduct/regionData?start_date=2016-01-01&end_date=2025-05-06

```json
{"data":{"end":"2025-05-06","overall_details":[{"quantity":5,"Region":"Asia"},{"quantity":3,"Region":"North America"},{"quantity":1,"Region":"Europe"},{"quantity":1,"Region":"South America"}],"start":"2016-01-01"},"message":"Top Product For Region","status":"S"}

```

### sample Response 4
- Endpoint :  /topProduct/trigger

```json
{"data":{"end":"2025-05-06","revenues":[{"ProductName":"iPhone 15 Pro","Revenue":11301.300000000001},{"ProductName":"UltraBoost Running Shoes","Revenue":1512},{"ProductName":"Sony WH-1000XM5 Headphones","Revenue":892.4744999999999},{"ProductName":"Levi's 501 Jeans","Revenue":431.928}],"start":"2012-01-01"},"message":"Revenue by product","status":"S"}

```



